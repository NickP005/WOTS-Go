package main

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"errors"
)

type Keychain struct {
	Seed  [32]byte
	Index uint64
}

type Keypair struct {
	PublicKey  [2144]byte
	PrivateKey [32]byte
	Components Components
}

type Components struct {
	PrivateSeed [32]byte
	PublicSeed  [32]byte
	AddrSeed    [32]byte
}

func mochimoHash(data []byte) [32]byte {
	hash := sha256.Sum256(data)
	return hash
}

/*
 * ComponentsGenerator derives three different seeds from an initial WOTS seed
 *
 * Parameters:
 * - wotsSeed: byte array of 32 bytes used as the initial seed
 *
 * Returns:
 * - Components: struct containing:
 *   1. PrivateSeed: 32 bytes used for WOTS secret key generation
 *   2. PublicSeed: 32 bytes used for WOTS public key generation
 *   3. AddrSeed: 32 bytes used for MCM address generation
 *
 * The function appends different strings to the seed ("seed", "publ", "addr")
 * and hashes each combination to generate the components
 */
func componentsGenerator(wotsSeed [32]byte) Components {
	seedAscii := string(wotsSeed[:])
	privateSeed := mochimoHash([]byte(seedAscii + "seed"))
	publicSeed := mochimoHash([]byte(seedAscii + "publ"))
	addrSeed := mochimoHash([]byte(seedAscii + "addr"))

	return Components{
		PrivateSeed: privateSeed,
		PublicSeed:  publicSeed,
		AddrSeed:    addrSeed,
	}
}

func keygenFromSeed(private_key [32]byte) Keypair {
	components := componentsGenerator(private_key)
	pk := wotsPkgen(components.PrivateSeed, components.PublicSeed, components.AddrSeed)

	return Keypair{
		PublicKey:  pk,
		PrivateKey: private_key,
		Components: components,
	}
}

/*
 * Keygen generates a new WOTS keypair
 *
 * Parameters:
 * - private_key: Optional 32-byte array for the private key seed
 *   If not provided, a cryptographically secure random key will be generated
 *
 * Returns:
 * - Keypair: containing the generated public key, private key and components
 * - error: returned if random generation fails when no private key is provided
 */
func Keygen(private_key ...[32]byte) (Keypair, error) {
	if len(private_key) == 0 {
		var randomKey [32]byte
		_, err := rand.Read(randomKey[:])
		if err != nil {
			return Keypair{}, errors.New("failed to generate random private key")
		}
		return keygenFromSeed(randomKey), nil
	}
	return keygenFromSeed(private_key[0]), nil
}

/*
 * Sign generates a WOTS signature for a given message
 *
 * Parameters:
 * - message: 32-byte array containing the message to sign
 *
 * Returns:
 * - [2144]byte: The generated signature
 */
func (keypair *Keypair) Sign(message [32]byte) [2144]byte {
	var sig [2144]byte
	sig = wotsSign(message, keypair.Components.PrivateSeed, keypair.Components.PublicSeed, keypair.Components.AddrSeed)
	return sig
}

/*
 * Verify checks if a signature is valid for a given message
 *
 * Parameters:
 * - message: 32-byte array containing the message to verify
 * - signature: 2144-byte array containing the signature to verify
 *
 * Returns:
 * - bool: true if signature is valid, false otherwise
 */
func (keypair *Keypair) Verify(message [32]byte, signature [2144]byte) bool {
	pk := wotsPkFromSig(signature, message, keypair.Components.PublicSeed, keypair.Components.AddrSeed)
	return pk == keypair.PublicKey
}

func newKeychainFromSeed(seed [32]byte) Keychain {
	return Keychain{
		Seed:  seed,
		Index: 0,
	}
}

/*
 * NewKeychain creates a new keychain
 *
 * Parameters:
 * - seed: Optional 32-byte array used as the master seed
 *   If not provided, a cryptographically secure random seed will be generated
 *
 * Returns:
 * - Keychain: initialized with the provided or random seed and index 0
 * - error: returned if random generation fails when no seed is provided
 */
func NewKeychain(seed ...[32]byte) (Keychain, error) {
	if len(seed) == 0 {
		var randomKey [32]byte
		_, err := rand.Read(randomKey[:])
		if err != nil {
			return Keychain{}, errors.New("failed to generate random private key")
		}
		return newKeychainFromSeed(randomKey), nil
	}
	return newKeychainFromSeed(seed[0]), nil
}

/*
 * Next generates the next keypair in the keychain sequence
 *
 * Returns:
 * - Keypair: the next keypair in the sequence, incrementing the internal index
 */
func (keychain *Keychain) Next() Keypair {
	secret, _ := DeriveSeed(keychain.Seed[:], keychain.Index)
	var privateKey [32]byte
	copy(privateKey[:], secret)

	keypair, _ := Keygen(privateKey)
	keychain.Index++
	return keypair
}

func DeriveSeed(deterministicSeed []byte, id uint64) ([]byte, *DigestRandomGenerator) {
	// Convert id to bytes
	idBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(idBytes, id)

	// Combine seed and id
	input := append(deterministicSeed, idBytes...)

	// Create SHA-512 hash
	h := sha512.New()
	h.Write(input)
	localSeed := h.Sum(nil)

	// Initialize PRNG
	prng := NewDigestRandomGenerator()
	prng.AddSeedMaterial(localSeed)

	// Generate secret (first 32 bytes)
	secret := prng.NextBytes(32)

	return secret, prng
}

func main() {

}

package wots

import (
	"crypto/sha512"
	"encoding/binary"
)

/*
 * DigestRandomGenerator implements a deterministic random number generator
 * using SHA-512 for generating sequences of random bytes.
 *
 * Fields:
 * - seed: accumulated seed material
 * - digest: current SHA-512 digest buffer
 * - count: counter for digest generation
 */
type DigestRandomGenerator struct {
	seed   []byte
	digest []byte
	count  uint32
}

/*
 * NewDigestRandomGenerator creates a new PRNG instance
 *
 * Returns:
 * - *DigestRandomGenerator: Initialized with empty seed and digest
 */
func NewDigestRandomGenerator() *DigestRandomGenerator {
	return &DigestRandomGenerator{
		seed:   make([]byte, 0),
		digest: make([]byte, 0),
		count:  0,
	}
}

/*
 * AddSeedMaterial adds entropy to the generator's seed
 * After adding seed material, a new digest is generated
 *
 * Parameters:
 * - seed: Byte slice containing entropy to add
 */
func (g *DigestRandomGenerator) AddSeedMaterial(seed []byte) {
	g.seed = append(g.seed, seed...)
	g.generateNextDigest()
}

/*
 * generateNextDigest creates a new SHA-512 digest
 * Combines the accumulated seed with a counter for uniqueness
 */
func (g *DigestRandomGenerator) generateNextDigest() {
	h := sha512.New()
	countBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(countBytes, g.count)

	h.Write(g.seed)
	h.Write(countBytes)
	g.digest = h.Sum(nil)
	g.count++
}

/*
 * NextBytes generates the requested number of random bytes
 * Will generate new digests as needed to fulfill the request
 *
 * Parameters:
 * - length: Number of random bytes to generate
 *
 * Returns:
 * - []byte: Slice containing the generated random bytes
 */
func (g *DigestRandomGenerator) NextBytes(length int) []byte {
	result := make([]byte, length)
	var pos int

	for pos < length {
		if len(g.digest) == 0 {
			g.generateNextDigest()
		}

		remaining := length - pos
		digestRemaining := len(g.digest)
		copyLen := remaining
		if digestRemaining < remaining {
			copyLen = digestRemaining
		}

		copy(result[pos:], g.digest[:copyLen])
		g.digest = g.digest[copyLen:]
		pos += copyLen
	}

	return result
}

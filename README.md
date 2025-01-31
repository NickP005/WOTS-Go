# WOTS-Go

A Go wrapper of the Winternitz One-Time Signature (WOTS) scheme of the Mochimo cryptocurrency.

## Installation

To install the package, run:

```sh
go get github.com/NickP005/WOTS-Go
```

## Usage Example

Here's a complete example showing all features with detailed explanations:

```go
package main

import (
	"fmt"
	"crypto/rand"
	"github.com/NickP005/WOTS-Go"
)

func main() {
	var seed [32]byte
	rand.Read(seed[:])

	keychain, _ := NewKeychain(seed)
	keypair := keychain.Next()

	fmt.Printf("Public Key: %x\n", keypair.PublicKey)
	fmt.Printf("Private Key: %x\n", keypair.PrivateKey)

	var message [32]byte
	copy(message[:], []byte("Hello, world!"))
	// Initialize message with some value
	signature := keypair.Sign(message)

	fmt.Println("Signature: ", len(signature))

	isValid := keypair.Verify(message, signature)
	fmt.Printf("Signature valid: %v\n", isValid)

	// Tampering with the signature
	signature[0] ^= 0xFF
	isValid = keypair.Verify(message, signature)
	fmt.Printf("Signature valid after tampering: %v\n", isValid)

}
```

# Support & Community

Join our communities for support and discussions:

<div align="center">

[![NickP005 Development Server](https://img.shields.io/discord/709417966881472572?color=7289da&label=NickP005%20Development%20Server&logo=discord&logoColor=white)](https://discord.gg/Q5jM8HJhNT)   
[![Mochimo Official](https://img.shields.io/discord/460867662977695765?color=7289da&label=Mochimo%20Official&logo=discord&logoColor=white)](https://discord.gg/SvdXdr2j3Y)

</div>

- **NickP005 Development Server**: Technical support and development discussions
- **Mochimo Official**: General Mochimo blockchain discussions and community

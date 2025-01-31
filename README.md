# WOTS-Go

A Go implementation of the Winternitz One-Time Signature (WOTS) scheme.

## Installation

To install the package, run:

```sh
go get github.com/NickP005/WOTS-Go
```

## Usage

### Key Generation

There are two ways to generate a new keypair:

```go
package main

import (
	"fmt"
	"crypto/rand"
	"github.com/NickP005/WOTS-Go"
)

func main() {
	var seed [32]byte
	_, err := rand.Read(seed[:])
	if err != nil {
		panic(err)
	}

	keychain := wots.NewKeychain(seed)
	keypair := keychain.Next()

	fmt.Printf("Public Key: %x\n", keypair.PublicKey)
	fmt.Printf("Private Key: %x\n", keypair.PrivateKey)
}
```

### Signing a Message

To sign a message:

```go
package main

import (
	"fmt"
	"github.com/NickP005/WOTS-Go"
)

func main() {
	var seed [32]byte
	// Initialize seed with some value
	keychain := wots.NewKeychain(seed)
	keypair := keychain.Next()

	var message [32]byte
	// Initialize message with some value
	signature := keypair.Sign(message)

	fmt.Printf("Signature: %x\n", signature)
}
```

### Verifying a Signature

To verify a signature:

```go
package main

import (
	"fmt"
	"github.com/NickP005/WOTS-Go"
)

func main() {
	var seed [32]byte
	// Initialize seed with some value
	keychain := wots.NewKeychain(seed)
	keypair := keychain.Next()

	var message [32]byte
	// Initialize message with some value
	signature := keypair.Sign(message)

	isValid := keypair.Verify(message, signature)
	fmt.Printf("Signature valid: %v\n", isValid)
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
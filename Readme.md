# Cryptographic Hash Functions

Golang implementation of the Cryptographic Hash Functions (CHF) used in CryptoPlasm.

First CHF implemented is Blake3, a pure Golang implementation.
Forked from: `https://github.com/lukechampine/blake3`

Added CustomOutput size option.

## Example for custom Output:
```
package main

import (
    blake3 "Cryptographic-Hash-Functions/Blake3"
    "encoding/hex"
    "fmt"
)

MyString = "Hello, playground"
MyStringToByteSlice := []byte(MyString)

S3 := blake3.SumCustom(MyStringToByteSlice,1024)

var S3ByteArray []byte
for i := 0; i < len(S3); i++ {
    S3ByteArray = append(S3ByteArray, S3[i])
    }
    Hex3 := hex.EncodeToString(S3ByteArray)

```
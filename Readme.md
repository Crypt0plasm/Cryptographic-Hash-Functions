# Cryptographic Hash Functions

Golang implementation of the Cryptographic Hash Functions (CHF) used in CryptoPlasm.

`go get github.com/Crypt0plasm/Cryptographic-Hash-Functions`

First CHF implemented is Blake3, a pure Golang implementation.
Forked from: `https://github.com/lukechampine/blake3`

Added CustomOutput size option for unkeyed hash.

## Example for custom Output:
```Go
package main

import (
    "encoding/hex"
    "fmt"

    blake3 "github.com/Crypt0plasm/Cryptographic-Hash-Functions/Blake3"
)

func main() {
    var (
	S3ByteArray []byte
	MyString = "Hello, playground"
    )

    //Converting the String to a slice of bytes
    MyStringToByteSlice := []byte(MyString)
    //Hashing the slice of bytes with 1024 byte Output size.
    S := blake3.SumCustom(MyStringToByteSlice, 1024)

    //Converting the resulting hash which is a slice of bytes, to hex (byte to hex)
    for i := 0; i < len(S); i++ {
	S3ByteArray = append(S3ByteArray, S[i])
    }
    Hex := hex.EncodeToString(S3ByteArray)

    fmt.Println("Unkeyed [", MyString, "], hashed in blake3 with 1024 bit output as byte is:", S)
    fmt.Println("Unkeyed [", MyString, "], hashed in blake3 with 1024 bit output as hex  is:", Hex)
}
```

To verify Output, go to
https://connor4312.github.io/blake3/index.html

For Input select "utf-8", and write the string: "Hello, playground" (without quotes). For Output select "hex" and 1024 bytes.
Output matches.
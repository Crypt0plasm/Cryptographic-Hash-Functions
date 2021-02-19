package main

import (
    blake3 "Cryptographic-Hash-Functions/Blake3"
    "encoding/hex"
    "fmt"
)

func main() {
    var MyString string

    MyString = "Hello, playground"
    MyStringToByteSlice := []byte(MyString)

    //start := time.Now()
    S1 := blake3.Sum256(MyStringToByteSlice)
    S2 := blake3.Sum512(MyStringToByteSlice)
    S3 := blake3.SumCustom(MyStringToByteSlice,1024)
    //elapsed := time.Since(start)
    //fmt.Println("Hashing took", elapsed)

    fmt.Println("My string is:",MyString)
    fmt.Println("")
    fmt.Println("Unkeyed [",MyString,"], hashed in blake3 with   32 bit output as byte is",S1)
    fmt.Println("Unkeyed [",MyString,"], hashed in blake3 with   64 bit output as byte is",S2)
    fmt.Println("Unkeyed [",MyString,"], hashed in blake3 with 1024 bit output as byte is",S3)
    fmt.Println("")

    var S1ByteArray []byte
    for i := 0; i < len(S1); i++ {
        S1ByteArray = append(S1ByteArray, S1[i])
    }
    Hex1 := hex.EncodeToString(S1ByteArray)
    fmt.Println("Unkeyed [",MyString,"], hashed in blake3 with   32 bit output as hex is",Hex1)

    var S2ByteArray []byte
    for i := 0; i < len(S2); i++ {
        S2ByteArray = append(S2ByteArray, S2[i])
    }
    Hex2 := hex.EncodeToString(S2ByteArray)
    fmt.Println("Unkeyed [",MyString,"], hashed in blake3 with   64 bit output as hex is",Hex2)

    var S3ByteArray []byte
    for i := 0; i < len(S3); i++ {
        S3ByteArray = append(S3ByteArray, S3[i])
    }
    Hex3 := hex.EncodeToString(S3ByteArray)
    fmt.Println("Unkeyed [",MyString,"], hashed in blake3 with 1024 bit output as hex is",Hex3)

    }
package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
)

func main() {
	res := someFunc(100)
	fmt.Println(res)
}
func someFunc(n int) string {
	v := createHugeString(n)
	return v
}
func createHugeString(n int) string {
	bytes := make([]byte, n)
	rand.Read(bytes)
	return base64.RawURLEncoding.EncodeToString(bytes)
}

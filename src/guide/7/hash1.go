package main

import (
	"fmt"
	"crypto/sha1"
	"crypto/md5"
	"encoding/hex"
)

func main() {
	TestString := []byte("abc")

	Md5Inst := md5.New()
	Md5Inst.Write(TestString)
	// Result := Md5Inst.Sum(nil)
	// r := md5.Sum(nil)
	fmt.Printf("md5:%x %x\n", Md5Inst.Sum(nil), md5.Sum(TestString))

	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(TestString))
	// Result = Sha1Inst.Sum(nil)
	// // r := sha1.Sum([]byte(TestString))
	fmt.Printf("sha1:%s\n", hex.EncodeToString(Sha1Inst.Sum(nil)))
}

package transform

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Hash(ppwd *[]byte) {
	if ppwd == nil { return }
	hash, err := bcrypt.GenerateFromPassword(*ppwd, bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error: failed to hash.")
		return
	}
	
	err = bcrypt.CompareHashAndPassword(hash, *ppwd)
	if err != nil {
		fmt.Println("Error: hash and password don't match")
		return
	}

	fmt.Println("Password hashed")
	for i := range *ppwd { (*ppwd)[i] = 0 }
	for i := range hash { 
		*ppwd = append(*ppwd, hash[i])
		hash[i] = 0
	}
	fmt.Println("Hash erased")
}
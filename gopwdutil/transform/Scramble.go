package transform

import (
	"fmt"
	"math/rand"
)

func Scramble(ppwd *[]byte) {
	if ppwd == nil { return }

	fmt.Println("Scrambling ...")
	length := len(*ppwd)
	scrambled := []byte{}
	for i := 0; i < length; i++ {
		j := rand.Intn(length - i)
		scrambled = append(scrambled, (*ppwd)[j])
		*ppwd = append((*ppwd)[:j],(*ppwd)[j + 1:]...)
	}

	for i := 0; i < length; i++ {
		*ppwd = append(*ppwd, scrambled[i])
		scrambled[i] = 0
	}
	fmt.Println("Scrambled password:", string(*ppwd))
	fmt.Print("Press Enter to continue ...")
	fmt.Scanln()
}
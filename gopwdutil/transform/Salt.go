package transform

import (
	"crypto/rand"
	"fmt"
	"gopwdutil/tools"
	"math/big"
)

func Salt(psalt *[]byte) error {
	if psalt == nil {
		return tools.Errors.NilError
	}

	fmt.Println("Generating salt...")
	pdLength, err := rand.Int(rand.Reader, big.NewInt(int64(tools.MaxSalt - tools.MinSalt)))
	if err != nil {
		fmt.Println("Error: randomizing error.")
		return err
	}

	length := tools.MinSalt + int(pdLength.Int64())

	var characters = append(tools.LowerChar, tools.UpperChar...)
	characters = append(characters, tools.NumericChar...)
	characters = append(characters, tools.SpecialChar...)

	bigLenCharacters := big.NewInt(int64(len(characters)))

	tools.Reset(psalt)
	for i := 0; i < length; i++ {
		j, err := rand.Int(rand.Reader, bigLenCharacters)
		if err != nil {
			fmt.Println("Error: randomizing error")
			return err
		}
		*psalt = append(*psalt, characters[j.Int64()])
	}
	fmt.Println("...")
	fmt.Println("Salt generated.")

	return nil
}

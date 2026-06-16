package transform

import (
	"encoding/base64"
	"fmt"
	"gopwdutil/tools"

	"golang.org/x/crypto/argon2"
)

func Hash(ppwd *tools.Password) error {
	if ppwd == nil {
		return tools.Errors.NilError
	}

	fmt.Println("Hashing password...")
	if len(ppwd.Salt) == 0 {
		fmt.Println("Warning: no salt")
	}

	if len(ppwd.Pepper) == 0 {
		fmt.Println("Warning: no pepper")
	}

	fmt.Println("...")
	key := argon2.IDKey(
		append(ppwd.Password, ppwd.Pepper...),
		ppwd.Salt,
		3,
		62 * 1024,
		1,
		32,
	)

	ppwd.HashedKey = []byte(base64.StdEncoding.EncodeToString(key))

	fmt.Println("Password hashed")

	return nil
}

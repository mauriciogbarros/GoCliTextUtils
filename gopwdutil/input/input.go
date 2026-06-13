package input

import (
	"bufio"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

var characters = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz .,;:!?'\"()[]{}<>-_/\\|+*=%^@#$&~`")

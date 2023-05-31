package trithemius_cipher

import (
	"fmt"
)

func GetEncryptedLetterCode(alphabetPos int, textPos int, alphabetSize int) int {
	return (alphabetPos + GetOffsetStep(textPos)) % alphabetSize
}

func Encrypt(text string, alphabet map[string]int) string {
	alphabetSize := len(alphabet)
	alphabetReversed := ReverseMap(alphabet)
	var encrypted string
	for i, ch := range text {
		code := GetEncryptedLetterCode(alphabet[string(ch)], i, alphabetSize)
		encrypted += alphabetReversed[code]
	}

	return encrypted
}

func EncryptMethod() string {
	text := "sourcetext"

	fmt.Println("Text:", text)

	encrypted := Encrypt(text, Alphabet)
	fmt.Println("Encrypted:", encrypted)
	return encrypted
}

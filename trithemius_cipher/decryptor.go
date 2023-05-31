package trithemius_cipher

import (
	"fmt"
)

func GetDecryptedLetterCode(alphabetPos int, textPos int, alphabetSize int) int {
	return (alphabetPos - GetOffsetStep(textPos)) % alphabetSize
}

func Decrypt(text string, alphabet map[string]int) string {
	alphabetSize := len(alphabet)
	alphabetReversed := ReverseMap(alphabet)
	var decrypted string
	for i, ch := range text {
		code := GetDecryptedLetterCode(alphabet[string(ch)], i, alphabetSize)
		decrypted += alphabetReversed[code]
	}

	return decrypted
}

func DecryptMethod1(encryptedText string) {
	decrypted := Decrypt(encryptedText, Alphabet)
	fmt.Println("decrypted:", decrypted)
}

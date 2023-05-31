package caesar_cipher

// Алфавит
const lowerCaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
const upperCaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Шифрует строку открытого текста, сдвигая каждый
// символ с помощью предоставленного ключа.
func EncryptPlaintext(plaintext string, key int) string {
	return rotateText(plaintext, key)
}

// Расшифровывает строку зашифрованного текста путем обратного сдвига каждого
// символа с предоставленным ключом.
func DecryptCiphertext(ciphertext string, key int) string {
	return rotateText(ciphertext, -key)
}

// Берет строку и поворачивает каждый
// символ на указанную величину.
func rotateText(inputText string, rot int) string {
	rot %= 26
	rotatedText := []byte(inputText)

	for index, byteValue := range rotatedText {
		if byteValue >= 'a' && byteValue <= 'z' {
			rotatedText[index] = lowerCaseAlphabet[(int((26+(byteValue-'a')))+rot)%26]
		} else if byteValue >= 'A' && byteValue <= 'Z' {
			rotatedText[index] = upperCaseAlphabet[(int((26+(byteValue-'A')))+rot)%26]
		}
	}
	return string(rotatedText)
}

// func main() {

// 	plaint_text := "SOMETEXTHERE"
// 	key := 13
// 	encrypted := caesar_cipher.EncryptPlaintext(plaint_text, key)
// 	fmt.Println(encrypted)

// 	plaintext := caesar_cipher.DecryptCiphertext(encrypted, key)
// 	fmt.Println(plaintext)

// }

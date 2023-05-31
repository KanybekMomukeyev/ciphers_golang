package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	caesar_cipher "github.com/golang_ciphers/caesar_cipher" // _1_
	// _2_
	matrix_vector "github.com/golang_ciphers/matrix_vector"         // _2_
	playfair_cipher "github.com/golang_ciphers/playfair_cipher"     // _4_
	trithemius_cipher "github.com/golang_ciphers/trithemius_cipher" // _3_
	vernamotp_cipher "github.com/golang_ciphers/vernamotp_cipher"   // _2_
	vigenere_cipher "github.com/golang_ciphers/vigenere_cipher"     // _2_
)

func main2222() {

	key := "Typewriter"
	message := "Now is the time for all good men to come to the aid of their country"

	encoded := vigenere_cipher.Encipher(message, key)
	fmt.Println(encoded)

	decoded := vigenere_cipher.Decipher(encoded, key)
	fmt.Println(decoded)

}

// encoded ==> GMLMOKPXXZFCUSNRTEKFHBBIJKWVSDXRDXDVIBHFYRWIEIKHYEMPN
// decoded ==> NOWISTHETIMEFORALLGOODMENTOCOMETOTHEAIDOFTHEIRCOUNTRY

func main() {
	// if false {
	// 	frequency_analysis.MainCrackCall()
	// }

	if false {
		text := trithemius_cipher.EncryptMethod()
		trithemius_cipher.DecryptMethod1(text)
	}

	if false {
		plaintext := caesar_cipher.DecryptCiphertext("Ebznaf, tb ubzr.", 13)
		fmt.Println(plaintext)

		key := "Typewriter"
		message := "Now is the time for all good men to come to the aid of their country"

		encoded := vigenere_cipher.Encipher(message, key)
		fmt.Println(encoded)

		decoded := vigenere_cipher.Decipher(encoded, key)
		fmt.Println(decoded)
		// -----

		text := "kanbek"
		fmt.Println("Source:", text)

		encrypted := trithemius_cipher.Encrypt(text, trithemius_cipher.Alphabet)
		fmt.Println("Encrypted:", encrypted)

		decrypted := trithemius_cipher.Decrypt(encrypted, trithemius_cipher.Alphabet)
		fmt.Println("decrypted:", decrypted)
	}

	if false {
		// str := "kanybek"
		// str = strings.ToUpper(str)
		// cipher := playfair_cipher.SplitLetters(data, str[:len(str)-1])
		// fmt.Println(cipher)
		playfair_cipher.DemoPlayfairCipher()

		// data := playfair_cipher.InitializeMatrix()
		// for i := 0; i < 5; i++ {
		// 	for j := 0; j < 5; j++ {
		// 		fmt.Print(data[i][j], " ")
		// 	}
		// 	fmt.Println()
		// }
	}

	if false {
		alp := matrix_vector.NewAlphabet("ABCDEFGHIJKLMNÑOPQRSTUVWXYZ")
		key := "FORTALEZA"
		msg := "CONSUL"
		// POH
		fmt.Println("message :", msg)
		fmt.Println("key :", key)

		cip, err := matrix_vector.NewCipher(alp)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Encrypt and decrypt messages

		cipherText, err := cip.Encrypt(msg, key)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Encrypted text :", cipherText)

		plainText, err := cip.Decrypt(cipherText, key)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Decrypted text :", plainText)
	}

	if false {
		m := make([]byte, 16*2048)
		rand.Read(m)

		// Создаем новый блокнот с 16-байтовыми страницами.
		//  Установите указатель на страницу 1.
		pad, err := vernamotp_cipher.NewPad(m, 16, 1)
		if err != nil {
			fmt.Printf("%s", err)
			return
		}

		original_message := "this is a test"
		fmt.Printf("message: %s\n\n", original_message)
		encrypted, _ := pad.Encrypt([]byte(original_message))
		fmt.Printf("encrypted: %s\n\n", base64.StdEncoding.EncodeToString(encrypted))
		decrypted, _ := pad.Decrypt(encrypted)
		fmt.Printf("decrypted: %s\n\n", decrypted)

		pad.NextPage()

		encrypted, _ = pad.Encrypt([]byte("this is a test"))
		fmt.Println(base64.StdEncoding.EncodeToString(encrypted))
		decrypted, _ = pad.Decrypt(encrypted)
		fmt.Printf("%s\n\n", decrypted)

		fmt.Printf("Total pages: %d\n", pad.TotalPages())
		fmt.Printf("Page pointer: %d\n", pad.CurrentPage())
		fmt.Printf("Remaining pages: %d\n", pad.RemainingPages())
	}

	// if false {
	// 	h := gost28147.NewCipher()
	// 	h.Write([]byte("hello world"))
	// 	fmt.Println(hex.EncodeToString(h.Sum(nil)))
	// }
}

package matrix_vector

import (
	"fmt"
)

func getKeyMatrix(key string, keyMatrix [][]int) {
	k := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			keyMatrix[i][j] = int(key[k]) % 65
			k++
		}
	}
}

func encrypt(cipherMatrix [][]int, keyMatrix [][]int, messageVector [][]int) {
	var x, i, j int
	for i = 0; i < 3; i++ {
		for j = 0; j < 1; j++ {
			cipherMatrix[i][j] = 0
			for x = 0; x < 3; x++ {
				cipherMatrix[i][j] += keyMatrix[i][x] * messageVector[x][j]
			}
			cipherMatrix[i][j] = cipherMatrix[i][j] % 26
		}
	}
}

func HillCipher(message string, key string) {
	keyMatrix := make([][]int, 3)
	for i := range keyMatrix {
		keyMatrix[i] = make([]int, 3)
	}
	getKeyMatrix(key, keyMatrix)
	messageVector := make([][]int, 3)
	for i := range messageVector {
		messageVector[i] = make([]int, 1)
		messageVector[i][0] = int(message[i]) % 65
	}
	cipherMatrix := make([][]int, 3)
	for i := range cipherMatrix {
		cipherMatrix[i] = make([]int, 1)
	}
	encrypt(cipherMatrix, keyMatrix, messageVector)
	var CipherText string
	for i := 0; i < 3; i++ {
		CipherText += string(cipherMatrix[i][0] + 65)
	}
	fmt.Println(CipherText)
}

func MainCallMatrix() {
	message := "ACT"
	// key := "GYBNQKURP"
	key := "KANYBESFTO"
	fmt.Println("message :", message)
	fmt.Println("key :", key)

	HillCipher(message, key)
}

package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

// шифрование использует шифр Цезаря для шифрования сообщений путем сдвига каждой буквы shift
// места вокруг алфавита. Неалфавитные символы оставляются в покое.
// Тут же происходит расшифрока при shift отрицательном.
func Encrypt(msg []byte, shift int) {
	// Make sure the shift value is positive.
	if shift < 0 {
		shift %= 26
		shift += 26
	}
	n := byte(shift)
	for i, c := range msg {
		if c >= 'A' && c <= 'Z' {
			msg[i] = 'A' + (c-'A'+n)%26
		} else if c >= 'a' && c <= 'z' {
			msg[i] = 'a' + (c-'a'+n)%26
		}
	}
}

// englishFreqs — относительная частота букв в английском языке.
// значения взяты из <http://en.wikipedia.org/wiki/Letter_frequency>. Они
// выражаются десятичными знаками и в сумме дают 1,0.
var EnglishFreqs = []float64{
	0.08167, 0.01492, 0.02782, 0.04253, 0.12702, 0.02228, 0.02015, 0.06094,
	0.06966, 0.00153, 0.00772, 0.04025, 0.02406, 0.06749, 0.07507, 0.01929,
	0.00095, 0.05987, 0.06327, 0.09056, 0.02758, 0.00978, 0.02360, 0.00150,
	0.01974, 0.00074,
}

// частоты вычисляет относительную частоту букв в тексте.
func Frequencies(text []byte) []float64 {
	freqs := make([]float64, 26)
	nLetters := 0
	for _, c := range text {
		if c >= 'A' && c <= 'Z' {
			freqs[c-'A']++
			nLetters++
		} else if c >= 'a' && c <= 'z' {
			freqs[c-'a']++
			nLetters++
		}
	}
	if nLetters != 0 {
		for i := range freqs {
			freqs[i] /= float64(nLetters)
		}
	}
	return freqs
}

// chiSqr вычисляет кумулятивную статистику критерия хи-квадрат Пирсона для
// заданные наблюдаемые частоты поворачивали позиции вращения, используя englishFreqs в качестве
// ожидаемые (теоретические) частоты.
func Chisqr(freqs []float64, rot int) float64 {
	// Make sure rot is positive.
	if rot < 0 {
		rot %= 26
		rot += 26
	}
	sum := 0.0
	for i, e := range EnglishFreqs {
		o := freqs[(i+rot)%26]
		sum += ((o - e) * (o - e)) / e
	}
	return sum
}

// Crack использует частотный анализ для определения значения сдвига Цезаря, которое было
// скорее всего используется для шифрования msg. Чем длиннее сообщение, тем выше шанс
// что он будет успешно взломан.
func Crack(msg []byte) int {
	freqs := Frequencies(msg)
	best := math.MaxFloat64
	shift := 0
	for i := 0; i < 26; i++ {
		chi := Chisqr(freqs, i)
		if chi < best {
			best = chi
			shift = i
		}
	}
	return shift
}

func main() {
	programName := filepath.Base(os.Args[0])
	args := os.Args[1:]
	usage := "usage: " + programName + " -e shift | -f | -c"

	var fn func([]byte)
	switch {
	case len(args) == 2 && args[0] == "-e":
		n, err := strconv.ParseInt(args[1], 0, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s: %s\n",
				programName, args[1], err.(*strconv.NumError).Err)
			os.Exit(1)
		}
		fn = func(msg []byte) {
			Encrypt(msg, int(n))
			os.Stdout.Write(msg)
		}
	case len(args) == 1 && args[0] == "-f":
		fn = func(text []byte) {
			fmt.Println(Frequencies(text))
		}
	case len(args) == 1 && args[0] == "-c":
		fn = func(msg []byte) {
			Encrypt(msg, -Crack(msg))
			os.Stdout.Write(msg)
		}
	case len(args) == 1 && (args[0] == "-h" || args[0] == "--help"):
		fmt.Println(usage)
		return
	default:
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	if input, err := ioutil.ReadAll(os.Stdin); err == nil {
		fn(input)
	}
}

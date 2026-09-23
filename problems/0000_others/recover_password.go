package _000_others

import (
	"crypto/md5"
	"fmt"
)

var alphabet = []rune{'a', 'b', 'c', 'd', '1', '2', '3'}

func RecoverPassword(h []byte) string {
	length := 8
	for i := 0; i < length; i++ {
		fmt.Println(i)
		if p := generateCombinations(length, h); p != "" {
			return p
		}
	}

	return ""
}

func generateCombinations(length int, Hash []byte) string {
	indices := make([]int, length)
	fmt.Println(indices)

	for {
		pass := make([]rune, length)

		for i, idx := range indices {
			pass[i] = alphabet[idx]
		}

		if string(HashPassword(string(pass))) == string(Hash) {
			return string(pass)
		}

		if incrIndec(indices, len(alphabet)) {
			continue
		}
	}

	//return ""
}

func incrIndec(indices []int, max int) bool {
	for i := len(indices) - 1; i >= 0; i-- {
		indices[i]++
		if indices[i] < max {
			return true
		}
		indices[i] = 0
	}
	return false
}

func HashPassword(in string) []byte {
	h := md5.Sum([]byte(in))
	return h[:]
}

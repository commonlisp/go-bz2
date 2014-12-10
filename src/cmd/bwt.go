package main

import (
	"fmt"
	"sort"
)

// Burrows-Wheeler Transform (aka block-sorting compression)

func bwt(buf []byte) (sorted []byte) {
	buflen := len(buf)
	concated := make([]byte, 2*buflen)
	copy(concated, buf)
	copy(concated[buflen:], buf)
	rotations := make([]string, buflen)
	for i := 0; i < buflen; i++ {
		rotations[i] = string(concated[i : i+buflen])
	}
	sort.Strings(rotations)
	sorted = make([]byte, buflen)
	for i := 0; i < buflen; i++ {
		sorted[i] = rotations[i][buflen-1]
	}

	return sorted
}

func bwtInv(buf []byte) (restored []byte) {
	buflen := len(buf)
	restored = make([]byte, buflen)
	table := make([][]byte, buflen)
	for i := 0; i < buflen; i++ {
		table[i] = make([]byte, buflen)
	}
	for i := 0; i < buflen; i++ {
		for j := 0; j < buflen; j++ {
			table[j][buflen-i-1] = buf[j]
		}
		tmp := make([]string, buflen)
		for i := 0; i < buflen; i++ {
			tmp[i] = string(table[i])
		}
		sort.Strings(tmp)
		for i := 0; i < buflen; i++ {
			table[i] = []byte(tmp[i])
		}
	}
	result := make([]byte, buflen)
	for i := range table {
		if table[i][buflen-1] == '.' {
			result = table[i]
			break
		}
	}
	return result
}

func main() {
	compressed := bwt([]byte("BANANA."))
	fmt.Printf("compressed: %s\n", compressed)
	fmt.Printf("uncompressed: %s\n", bwtInv(compressed))
}

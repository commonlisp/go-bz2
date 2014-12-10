package main

import (
	"fmt"
	"priorityqueue"
)

// WIP
func huffmanEncode(buf []byte) []byte {
	buflen := len(buf)
	freq := make(map[byte]int)
	for _, b := range buf {
		freq[b] += 1
	}
	pq := make(priorityqueue.PriorityQueue, len(buf))
	for b, count := range freq {
		pq.Push(&priorityqueue.Item{Value: b, Priority: count / buflen})
	}
	// TODO
	return []byte{}
}

func main() {
	fmt.Printf("%s", huffmanEncode([]byte("aaabbbcceeeddwwss")))
}

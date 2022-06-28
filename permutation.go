package main

const NUM_CHARS = 26 // a-z

func printAllKLengthRec(prefix []byte, k int, ch chan []byte) {
	if k == 0 {
		ch <- prefix
		return
	}

	for i := 0; i < NUM_CHARS; i++ {
		newPrefix := append(prefix, byte('a'+i))
		printAllKLengthRec(newPrefix, k-1, ch)
	}
}

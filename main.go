package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const Length = 4

func main() {
	ch := make(chan []byte)
	go func() {
		printAllKLengthRec(nil, Length, ch)
	}()

	slots := map[uint16]string{}
	count := 0
	for key := range ch {
		slot := crc16(key) % 16384
		if _, ok := slots[slot]; !ok {
			slots[slot] = string(key)
			count++
		}
		if count == 16384 {
			break
		}
	}
	fmt.Printf("Count is %d", count)

	data, err := json.MarshalIndent(slots, "", "  ")
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("slots.json", data, 0o644)
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

func main() {
	var b []byte
	randomizedSliceSize := 10000000
	//b = randomNumberInserts(randomizedSliceSize)
	b = arrayScrambler(randomizedSliceSize)

	err := ioutil.WriteFile("randomslice.txt", b, 0644)
	if err != nil {
		fmt.Errorf("unable to write to file: %v", err)
	}
}

func arrayScrambler(randomizedSliceSize int) []byte {
	scrambleXTimes := 5

	randomizedSlice := make([]int64, randomizedSliceSize)
	for index := range randomizedSlice {
		randomizedSlice[index] = int64(index)
	}

	var placeholder int64
	var randIndex int64
	for scrambleTimes := 0; scrambleTimes < scrambleXTimes; scrambleTimes++ {
		for index := range randomizedSlice {
			randIndex = rand.Int63n(int64(randomizedSliceSize))
			placeholder = randomizedSlice[randIndex]

			randomizedSlice[randIndex] = randomizedSlice[index]
			randomizedSlice[index] = placeholder
		}
	}

	b, err := json.Marshal(&randomizedSlice)
	if err != nil {
		fmt.Errorf("unable to marshal json: %v", err)
	}
	return b
}

func randomNumberInserts(randomizedSliceSize int) []byte {
	randomizedSlice := make([]int64, randomizedSliceSize)

	for index := range randomizedSlice {
		randomizedSlice[index] = rand.Int63n(int64(randomizedSliceSize))
	}

	b, err := json.Marshal(&randomizedSlice)
	if err != nil {
		fmt.Printf("unable to marshal json: %v\n", err)
	}
	return b
}

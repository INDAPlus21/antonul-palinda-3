package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"
)

const DataFile = "loremipsum.txt"

// Return the word frequencies of the text argument.
func WordCount(text string) map[string]int {
	words := strings.Fields(text)
	freqs := make(map[string]int)

	amount := len(words)

	//Approx 1 ms per goroutine
	routineSize := 1

	freq_chan := make(chan map[string]int, amount/routineSize+1)

	wg := new(sync.WaitGroup)

	for i, j := 0, routineSize; j < amount+1; i, j = j, j+routineSize {

		if j > amount {
			j = amount
		}

		wg.Add(1)
		go func(i, j int) {
			freq := make(map[string]int)
			for k := i; k < j; k++ {
				word := strings.ToLower(words[k])
				word = strings.ReplaceAll(word, ".", "")
				word = strings.ReplaceAll(word, ",", "")
				freq[word] += 1
			}
			freq_chan <- freq
			wg.Done()
		}(i, j)
	}
	wg.Wait()
	close(freq_chan)
	for something := range freq_chan {
		for k, v := range something {
			freqs[k] += v
		}
	}
	return freqs
}

// Benchmark how long it takes to count word frequencies in text numRuns times.
//
// Return the total time elapsed.
func benchmark(text string, numRuns int) int64 {
	start := time.Now()
	for i := 0; i < numRuns; i++ {
		WordCount(text)
	}
	runtimeMillis := time.Since(start).Nanoseconds() / 1e6

	return runtimeMillis
}

// Print the results of a benchmark
func printResults(runtimeMillis int64, numRuns int) {
	fmt.Printf("amount of runs: %d\n", numRuns)
	fmt.Printf("total time: %d ms\n", runtimeMillis)
	average := float64(runtimeMillis) / float64(numRuns)
	fmt.Printf("average time/run: %.2f ms\n", average)
}

func main() {
	// read in DataFile as a string called data
	data, err := ioutil.ReadFile("loremipsum.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", WordCount(string(data)))

	numRuns := 1
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}

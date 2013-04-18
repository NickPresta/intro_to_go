package main

import ("flag";"fmt";"io/ioutil";"runtime";"strings";"time")

func contains(needle string, haystack []string, result chan int) int {
	var out []string
	for _, word := range haystack {
		if strings.Contains(word, needle) {
			out = append(out, word)
		}
	}
	result <- len(out)
	return len(out)
}

func getSource(path string) []string {
	file, _ := ioutil.ReadFile(path)
	return strings.Split(string(file), "\n")
}

var word = flag.String("word", "", "Specify the word")
var path = flag.String("path", "", "Specify the path")
var workers = flag.Int("workers", runtime.NumCPU(), "Specify the workers")

func main() {
	flag.Parse()
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)

	haystack := getSource(*path)
	haystackSize := len(haystack)
	size := haystackSize / *workers
	partitions := make([][]string, haystackSize / size)
	numPartitions := cap(partitions)

	startTime := time.Now()
	startIndex := 0
	endIndex := size
	done := make(chan int, numPartitions)
	for i := 0; i < numPartitions; i++ {
		go contains(*word, haystack[startIndex:endIndex], done)
		startIndex = startIndex + size
		endIndex = startIndex + size
		if endIndex >= haystackSize {
			endIndex = haystackSize - 1
		}
	}

	total := 0
	for i := 0; i < numPartitions; i++ {
		total = total + <-done
	}
	fmt.Printf("Matched: %d\n", total)
	fmt.Printf("Critical section: %f seconds\n", (time.Now().Sub(startTime)).Seconds())
}

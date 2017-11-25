package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	sc.Split(bufio.ScanWords)

	buckets := make([][]string, 12)

	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}

	for sc.Scan() {
		word := sc.Text()
		n := HashBucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}

	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}

	// fmt.Println(buckets[6])
}

func HashBucket(word string, buckets int) int {
	var sum int

	for _, v := range word {
		sum += int(v)
	}

	return sum % buckets
}

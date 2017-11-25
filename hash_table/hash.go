package hash

import (
	"bufio"
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"
	//"os"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// bs, _ := ioutil.ReadAll(res.Body)
	// str := string(bs)
	// fmt.Println(str)

	// words := make(map[string]string)

	sc := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	sc.Split(bufio.ScanWords)

	buckets := make([]int, 10000)

	for sc.Scan() {
		n := HashBucket(sc.Text())
		buckets[n]++
	}

	fmt.Println(buckets)
	// for sc.Scan() {
	//  words[sc.Text()] = ""
	// }

	// if err := sc.Err(); err != nil {
	//  fmt.Println(os.Stderr, "reading input:", err)
	// }

	// i := 0
	// for k := range words {
	//  fmt.Println(k)
	//  if i == 200 {
	//      break
	//  }
	//  i++
	// }
}

func HashBucket(word string) uint64 {
	// letter := int(word[0])
	// bucket := letter % buckets
	// return bucket
	var hash uint64 = 5381

	for i := 0; i < len(word); i++ {
		hash = ((hash << 5) + hash) + uint64(word[i])
	}

	return hash % 10000
}

package main

import (
	"encoding/json"
	"fmt"
	//"os"
	"strings"
)

type Person struct {
	First       string
	Last        string // `json:"-"`
	Age         int    `json:"wisdom score"`
	notExported int
}

func main() {
	// p1 := Person{"James", "Bond", 20, 007}
	// bs, _ := json.Marshal(p1)
	// fmt.Println(bs)
	// fmt.Printf("%T \n", bs)
	// fmt.Println(string(bs)) // will not show notExported field, becasue it's unexported (starts with small n)

	// p2 := Person{"James", "Bond", 20}
	// fmt.Println(p1)
	// bs2, _ := json.Marshal(p2)
	// fmt.Println(string(bs2))

	var p3 Person
	// fmt.Println("----------------")
	// fmt.Println(p3.First)
	// fmt.Println(p3.Last)
	// fmt.Println(p3.Age)

	// bs3 := []byte(`{"First":"James", "Last":"Bond", "wisdom score":20}`)
	// json.Unmarshal(bs3, &p3)

	// fmt.Println("----------------")
	// fmt.Println(p3.First)
	// fmt.Println(p3.Last)
	// fmt.Println(p3.Age)
	// fmt.Printf("%T \n", p3)

	// json.NewEncoder(os.Stdout).Encode(p1)
	rdr := strings.NewReader(`{"First":"James", "Last":"Bond", "wisdom score":20}`)
	json.NewDecoder(rdr).Decode(&p3)

	fmt.Println(p3.First)
	fmt.Println(p3.Last)
	fmt.Println(p3.Age)
}

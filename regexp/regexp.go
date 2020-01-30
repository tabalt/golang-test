package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Match
	matched1, err := regexp.Match(`foo.*`, []byte(`seafood`))
	fmt.Println(matched1, err)

	// MatchString
	matched2, err := regexp.MatchString(`foo.*`, "seafood")
	fmt.Println(matched2, err)

	// Find
	re1 := regexp.MustCompile(`foo.?`)
	fmt.Printf("%q\n", re1.Find([]byte(`seafood fool`)))

	// FindAll
	re2 := regexp.MustCompile(`foo.?`)
	fmt.Printf("%q\n", re2.FindAll([]byte(`seafood fool`), -1))

	// ReplaceAll
	re3 := regexp.MustCompile(`a(x*)b`)
	fmt.Printf("%q\n", re3.ReplaceAll([]byte("-ab-axxb-"), []byte("T")))

	// ReplaceAllString
	re4 := regexp.MustCompile(`a(x*)b`)
	fmt.Printf("%q\n", re4.ReplaceAllString("-ab-axxb-", "T"))

	// Split
	a := regexp.MustCompile(`a`)
	fmt.Println(a.Split("banana", -1))
	fmt.Println(a.Split("banana", 0))
	fmt.Println(a.Split("banana", 1))

}

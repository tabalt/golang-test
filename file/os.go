package main

import (
	"log"
	"os"
)

func main() {

	file := "./_os.txt"

	// write
	f1, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f1.Write([]byte("appended some data\n")); err != nil {
		log.Fatal("read file error", err)
	}

	if err := f1.Close(); err != nil {
		log.Fatal(err)
	}

	// read
	f2, err := os.OpenFile(file, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	count, err := f2.Read(data)
	if err != nil {
		log.Fatal("read file error: ", err)
	}

	log.Printf("read %d bytes: %s\n", count, data[:count])

	if err := f2.Close(); err != nil {
		log.Fatal(err)
	}
}

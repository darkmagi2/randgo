package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func input(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var chars []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		chars = append(chars, scanner.Text())
	}
	return chars, scanner.Err()
}

func main() {
	chars, err := input("chars.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println("Random Choice:")
	rand.Seed(time.Now().Unix())
	message := fmt.Sprint(chars[rand.Intn(len(chars))])
	messagebyte := []byte(message + "\n")
	fmt.Println(message)

	f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(messagebyte)); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)

	}
	fmt.Print("Press 'Enter' to continue")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

}

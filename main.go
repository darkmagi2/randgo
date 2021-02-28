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

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func main() {
	used, err := input("output.txt")
	chars, err := input("chars.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	fmt.Println("Choosing a random character:")

	if len(used) == len(chars) {
		time.Sleep(2 * time.Second)
		fmt.Println("You played all your characters this week, nerd.")
		os.Exit(1)
	}

	rand.Seed(time.Now().Unix())
	message := fmt.Sprint(chars[rand.Intn(len(chars))])
	_, found := Find(used, message)
	for found {
		fmt.Println("I chose", message, "but you played him already this week. Trying again..")
		rand.Seed(time.Now().Unix())
		rand.Shuffle(len(chars), func(i, j int) {
			chars[i], chars[j] = chars[j], chars[i]
		})
		message = (chars[1])
		_, found = Find(used, message)

	}
	if !found {
		fmt.Println("I choose", message)
		messagebyte := []byte(message + "\n")
		//fmt.Println(message)

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
		fmt.Println("Press 'Enter' to exit")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}

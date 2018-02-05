package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var autoincrement = 0

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func init() {
	if _, err := os.Stat("AtID.txt"); os.IsNotExist(err) {
		ID := []byte("1")
		ioutil.WriteFile("AtID.txt", ID, 0644)
	}
}

func insert(key, value string) {
	file, err := os.OpenFile("database.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	check(err)
	insertion := "" + key + ":" + value + ""

	defer file.Close()

	file.WriteString(insertion)
	file.WriteString("\r\n")
	fmt.Println("Inserted key: " + key)
	fmt.Println("Inserted value: " + value)
	fmt.Println("==> Done Writing to file...")
}

func read(keyArg string) {
	readFile, err := os.Open("database.txt")
	check(err)
	defer readFile.Close()

	reader := bufio.NewReader(readFile)

	m := make(map[string]string)
	for {
		line, err := reader.ReadString('\n')
		if len(line) != 0 {

			fmt.Println(line)
			s := strings.Split(line, ":")
			key, value := s[0], s[1]
			//fmt.Println("Key: " + key)
			//fmt.Println("Value: " + value)
			m[key] = value
			check(err)
		}
		if err != nil {
			break
		}
	}

	fmt.Printf("Map: %#v\n", m[keyArg])
	/*
		if err := json.Unmarshal(read, &m); err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		fmt.Println(read)
		fmt.Printf("Map: %#v\n", m)
	*/
}

func main() {
	// Insert CLI
	insertCmd := flag.NewFlagSet("insert", flag.ExitOnError)
	key := insertCmd.String("key", "", "The key you want to add")
	value := insertCmd.String("value", "", "The value you want to pass to the key")

	// Update CLI
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)

	// Delete CLI
	delCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	// Read CLI
	readCmd := flag.NewFlagSet("read", flag.ExitOnError)
	get := readCmd.String("get", "", "Used to retrieve data from the database file")

	/*
		readFile, err := ioutil.ReadFile("AtID.txt")

		convert := string(readFile[:])
		fmt.Println(convert)
		number, err := strconv.ParseInt(convert, 10, 0)
		check(err)

		if number == 1 {
			fmt.Println(true)
		}
	*/
	switch os.Args[1] {
	case "insert":
		err := insertCmd.Parse(os.Args[2:])
		check(err)
	case "update":
		err := updateCmd.Parse(os.Args[2:])
		check(err)
	case "delete":
		//delete(entries, os.Args[2])

	case "read":
		err := readCmd.Parse(os.Args[2:])
		check(err)
	default:
		fmt.Printf("%s", "Error")
	}

	if insertCmd.Parsed() {
		fmt.Println("Parsed INSERT")
		insert(*key, *value)
		//ioutil.WriteFile("AtID.txt", number, 0644)
	}

	if updateCmd.Parsed() {
		fmt.Println("Parsed UPDATE")

	}
	if delCmd.Parsed() {
		fmt.Println("Parsed DELETE")
	}
	if readCmd.Parsed() {
		fmt.Println("Parsed READ")
		read(*get)
	}
}

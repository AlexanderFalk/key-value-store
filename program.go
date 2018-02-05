package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"strconv"
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

func index(key string, offset int64, length int) {
	file, err := os.OpenFile("index.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	check(err)
	fmt.Println("Offset: ", offset)
	fmt.Println("Length: ", length)
	insertion := "" + key + ":" + strconv.Itoa(int(offset)) + ":" + strconv.Itoa(int(length)) + ""
	defer file.Close()

	file.WriteString(insertion)
	file.WriteString("\r\n")
	fmt.Println("==> Done writing to index file...")
}


func insert(key, value string) {
	file, err := os.OpenFile("database.db", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	check(err)
	defer file.Close()

	insertion := "" + key + ":" + value + ","
	convertToByte := []byte(insertion)
	
	// Offset is how many bytes to move
    // Offset can be positive or negative
    var offset int64 = 0

    // Whence is the point of reference for offset
    // 0 = Beginning of file
    // 1 = Current position
    // 2 = End of file
    var whence int = 2

    currentPosition, err := file.Seek(offset, whence)
    check(err)
    fmt.Println("Position: ", currentPosition)
    // Add to index file
    index(key, currentPosition, len(convertToByte))

	fmt.Println("Bytes: ", convertToByte)
	file.Write(convertToByte)
	fmt.Println("==> Done writing to database file...")
}

func read(keyArg string) {
	readFile, err := os.Open("index.txt")
	check(err)
	defer readFile.Close()

	reader := bufio.NewReader(readFile)

	m := make(map[string][]string)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimRight(line, "\r\n")
		if len(line) != 0 {
			s := strings.Split(line, ":")
			key, index, length := s[0], s[1], s[2]
			m[key] = []string{string(index), string(length)}
			check(err)
		}
		if err != nil {
			break
		}
	}

	fmt.Println("Key Argument: " + keyArg)
	//fmt.Printf("Map: %#v\n", m[keyArg])

	// Gets value from key and converts it to an integer, so it can be passed to the database function
	i := m[keyArg][0]
	l := m[keyArg][1]
	//fmt.Println("Data Index: ", i)
	//fmt.Println("Data Length: ", l)
	indexConvert, err := strconv.Atoi(i)
	check(err)
	lengthConvert, err := strconv.Atoi(l)
	check(err)
	readDatabase(indexConvert, lengthConvert)
}

func readDatabase(index, length int) {
	file, err := os.Open("database.db")
	check(err)
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	check(err)
	//fmt.Println("Data Length: ", len(data))
	//fmt.Println("Data Sliced Length: ", len(data[index:]))
	// Calculating the slice of the bytes needed to return the key/value pair
	result := (len(data) - len(data[index:])) + length
	s := string(data[index:result])
	finalResult := strings.Split(strings.TrimRight(s, ","), ":")
	fmt.Println("Result: ", finalResult[1])
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

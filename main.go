package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	var name, address string
	personMap := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Your name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	personMap["name"] = name

	fmt.Print("Your address: ")
	address, _ = reader.ReadString('\n')
	address = strings.TrimSpace(address)
	personMap["address"] = address

	jsonData, err := json.Marshal(personMap)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("JSON: ", string(jsonData))
}

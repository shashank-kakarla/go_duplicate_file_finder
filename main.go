package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
)

const path = "/home/shanks/go/src/github.com/shashank-kakarla/Test_Folder"

func printMap(duplicatesMap map[string][]string) {
	for k, v := range duplicatesMap {
		fmt.Print(k + " ")
		for _, each := range v {
			fmt.Print(each + " ")
		}
		fmt.Println()
	}
}

func main() {
	fileMap := make(map[[32]byte]string)
	duplicatesMap := make(map[string][]string)
	fmt.Println("Duplicate Finder")

	// Reading all the files in directory
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {

		fileRead, err := os.ReadFile(path + "/" + file.Name())

		if err != nil {
			panic(err)
		}

		fileHash := sha256.Sum256(fileRead)

		_, ok := fileMap[fileHash]
		if ok {
			duplicatesMap[fileMap[fileHash]] = append(duplicatesMap[fileMap[fileHash]], file.Name())
		} else {
			fileMap[fileHash] = file.Name()
		}
	}
	printMap(duplicatesMap)
}

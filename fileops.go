package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
)

func ChooseFile() string {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	textFiles := []string{}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".txt") {
			textFiles = append(textFiles, file.Name())
		}
	}

	idx, err := fuzzyfinder.Find(
		textFiles,
		func(i int) string {
			return textFiles[i]
		},
	)

	if err != nil {
		log.Fatalf("File selection cancelled or no files found: %v", err)
	}

	return textFiles[idx]
}

func ReadFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	fmt.Println(string(data))
}

func CountLines(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	content := string(data)
	lines := strings.Split(content, "\n")
	fmt.Printf("Number of lines: %d\n", len(lines))
}

func FindInFile(filePath, textToFind string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), textToFind) {
			fmt.Printf("Found '%s' at line: %d\n", textToFind, lineNumber)
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %s", err)
	}
}

func WriteToFile(filePath string, content string) {
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		log.Fatalf("Failed to write to file: %s", err)
	}
}

func ReplaceInFile(filePath, oldText, newText string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	newData := strings.ReplaceAll(string(data), oldText, newText)
	WriteToFile(filePath, newData)
}

func CountWords(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}

	re := regexp.MustCompile(`\w+`)
	words := re.FindAllString(string(data), -1)

	fmt.Printf("Number of words: %d\n", len(words))
}

func CountCharacters(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	content := string(data)
	fmt.Printf("Number of characters: %d\n", len(content))
}

func SortLines(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	content := string(data)
	lines := strings.Split(content, "\n")
	sort.Strings(lines)
	newData := strings.Join(lines, "\n")
	WriteToFile(filePath, newData)
}

func ConcatenateFiles(filePath1, newFilePath string) {
	fmt.Println("Select a file to concatenate with the current file:")
	filePath2 := ChooseFile()

	data1, err := ioutil.ReadFile(filePath1)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}

	data2, err := ioutil.ReadFile(filePath2)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}

	err = ioutil.WriteFile(newFilePath, append(data1, data2...), 0644)
	if err != nil {
		log.Fatalf("Failed to write to file: %s", err)
	}

	fmt.Println("Files concatenated successfully.")
}

func SplitFile(filePath string, numFiles int) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	content := string(data)
	lines := strings.Split(content, "\n")
	numLines := len(lines)
	linesPerFile := numLines / numFiles
	for i := 0; i < numFiles; i++ {
		startIndex := i * linesPerFile
		endIndex := startIndex + linesPerFile
		if i == numFiles-1 {
			endIndex = numLines
		}
		newData := strings.Join(lines[startIndex:endIndex], "\n")
		newFilePath := fmt.Sprintf("%s_part%d.txt", filePath, i+1)
		WriteToFile(newFilePath, newData)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var FilePath string

func main() {
	reader := bufio.NewReader(os.Stdin)

	FilePath = ChooseFile()
	PrintBanner()

	for {
		fmt.Printf("Enter command for file '%s' (type 'help' for available commands): ", FilePath)
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
		case "read":
			ReadFile(FilePath)
		case "write":
			fmt.Print("Enter text to write: ")
			textToWrite, _ := reader.ReadString('\n')
			textToWrite = strings.TrimSpace(textToWrite)
			WriteToFile(FilePath, textToWrite)
		case "lines":
			CountLines(FilePath)
		case "find":
			fmt.Print("Enter text to find: ")
			textToFind, _ := reader.ReadString('\n')
			textToFind = strings.TrimSpace(textToFind)
			FindInFile(FilePath, textToFind)
		case "replace":
			fmt.Print("Enter old text: ")
			oldText, _ := reader.ReadString('\n')
			oldText = strings.TrimSpace(oldText)
			fmt.Print("Enter new text: ")
			newText, _ := reader.ReadString('\n')
			newText = strings.TrimSpace(newText)
			ReplaceInFile(FilePath, oldText, newText)
		case "words":
			CountWords(FilePath)
		case "chars":
			CountCharacters(FilePath)
		case "sort":
			SortLines(FilePath)
		case "concat":
			fmt.Print("Enter new file path: ")
			newFilePath, _ := reader.ReadString('\n')
			newFilePath = strings.TrimSpace(newFilePath)
			ConcatenateFiles(FilePath, newFilePath)

		case "split":
			fmt.Print("Enter number of files to split into: ")
			numFilesStr, _ := reader.ReadString('\n')
			numFilesStr = strings.TrimSpace(numFilesStr)
			numFiles, err := strconv.Atoi(numFilesStr)
			if err != nil {
				fmt.Println("Invalid number")
				continue
			}
			SplitFile(FilePath, numFiles)
		case "change_file":
			FilePath = ChooseFile()
			PrintBanner()
		case "help":
			PrintHelp()
		case "quit":
			return
		default:
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}
	}
}

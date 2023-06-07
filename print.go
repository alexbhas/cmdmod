package main

import "fmt"

func PrintBanner() {
	fmt.Println(`
	 $$$$$$\  $$\      $$\ $$$$$$$\  $$\      $$\  $$$$$$\  $$$$$$$\  
	$$  __$$\ $$$\    $$$ |$$  __$$\ $$$\    $$$ |$$  __$$\ $$  __$$\ 
	$$ /  \__|$$$$\  $$$$ |$$ |  $$ |$$$$\  $$$$ |$$ /  $$ |$$ |  $$ |
	$$ |      $$\$$\$$ $$ |$$ |  $$ |$$\$$\$$ $$ |$$ |  $$ |$$ |  $$ |
	$$ |      $$ \$$$  $$ |$$ |  $$ |$$ \$$$  $$ |$$ |  $$ |$$ |  $$ |
	$$ |  $$\ $$ |\$  /$$ |$$ |  $$ |$$ |\$  /$$ |$$ |  $$ |$$ |  $$ |
	\$$$$$$  |$$ | \_/ $$ |$$$$$$$  |$$ | \_/ $$ | $$$$$$  |$$$$$$$  |
	 \______/ \__|     \__|\_______/ \__|     \__| \______/ \_______/
	
	 			A Text File Tool
	`)
}

func PrintHelp() {
	fmt.Println(`
Available commands:
    read        - Reads and prints the entire file.
    write       - Prompts for text and writes it to the file.
    lines       - Counts and prints the number of lines in the file.
    find        - Prompts for a string and searches for that string in the file.
    replace     - Prompts for old and new text and replaces all occurrences of old text with new text.
    words       - Counts and prints the number of words in the file.
    chars       - Counts and prints the number of characters in the file.
    sort        - Sorts the lines in the file alphabetically.
    concat      - Prompts for a new file name then a second file and concatenates the two files.
    split       - Prompts for a number and splits the file into that many files.
    change_file - Prompts to select a new file to operate on.
    help        - Prints this help message.
    quit        - Exits the program.
	`)
}

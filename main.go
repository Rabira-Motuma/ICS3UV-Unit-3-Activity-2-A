/*
 * Author: Rabira Motuma
 * Version: 1.0.0
 * Date: 2025-11-22
 * Fileoverview: This program asks the user for their name and then age and then returns that information.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// variables
	var userName string
	var ageAsString string
	var ageAsNumber int
	var ageFiveYearsAgo int

	reader := bufio.NewReader(os.Stdin)

	// input
	fmt.Print("What is your name? ")
	userName, _ = reader.ReadString('\n')
	userName = strings.TrimSpace(userName)

	fmt.Print("How old are you? ")
	ageAsString, _ = reader.ReadString('\n')
	ageAsString = strings.TrimSpace(ageAsString)

	// process
	ageAsNumber, _ = strconv.Atoi(ageAsString)
	ageFiveYearsAgo = ageAsNumber - 5

	// output
	fmt.Println()
	fmt.Printf("%s", "Hello, " + userName + "!\n")
	fmt.Printf("%s", "You are currently " + strconv.Itoa(ageAsNumber) + " years old!\n")
	fmt.Printf("%s", "Five years ago you were " + strconv.Itoa(ageFiveYearsAgo) + " years old.")

	fmt.Println("\nDone.")
}

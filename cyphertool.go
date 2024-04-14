package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var resStr string

	fmt.Printf("\nWelcome to the Cypher Tool!\n\n")
	toEncrypt, encoding, message := getInput() // here you declare variables and get the input from user what they want

	switch {
	case toEncrypt && encoding == "1":
		{
			resStr = encrypt_rot13(message)
			encoding = "rot13"
		}
	case toEncrypt && encoding == "2":
		{
			resStr = encrypt_reverse(message)
			encoding = "reverse"
		}
	case toEncrypt && encoding == "3":
		{
			resStr = encrypt_reverseAlphabetValue(message)
			encoding = "reverse arphabet value"
		}
	case !toEncrypt && encoding == "1":
		{
			resStr = decrypt_rot13(message)
			encoding = "rot13"
		}
	case !toEncrypt && encoding == "2":
		{
			resStr = decrypt_reverse(message)
			encoding = "reverse"
		}
	case !toEncrypt && encoding == "3":
		{
			resStr = decrypt_reverseAlphabetValue(message)
			encoding = "reverse alphabet value"
		}
	}

	action := "Decrypted"

	if toEncrypt {
		action = "Encrypted"
	}

	fmt.Printf("\n%v message using %v:\n", action, encoding)
	fmt.Println(resStr)
}

func getInput() (toEncrypt bool, encoding string, message string) {

	var operation string
	var validOperation bool

	for !validOperation {
		fmt.Println("Select operation (1/2): \n", "1. Encrypt.\n", "2. Decrypt.")
		fmt.Scan(&operation)
		if operation == "1" || operation == "2" {
			validOperation = true
		} else {
			fmt.Println("\nYou have entered unsuitable selection for operation. Please select again.\n")
		}
	}

	if operation == "1" {
		toEncrypt = true
	}
	if operation == "2" {
		toEncrypt = false
	}

	var validEncoding bool

	for !validEncoding {

		fmt.Println("\nSelect cypher (1/2/3): \n", "1. ROT13.\n", "2. Reverse.\n", "3. Reverse Alphabet Value.")
		fmt.Scan(&encoding)
		if encoding == "1" || encoding == "2" || encoding == "3" {
			validEncoding = true
		} else {
			fmt.Println("\nYou have entered unsuitable selection for operation. Please select again.")
		}
	}

	in := bufio.NewReader(os.Stdin)
	fmt.Println("\nEnter the message: ")
	text, _ := in.ReadString('\n')
	text = strings.TrimSpace(text)
	message += text

	return
}

// Encrypt the message with rot13
func encrypt_rot13(s string) string {
	var encryptedStr string

	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			if char+13 > 'Z' {
				encryptedStr += string(char - 13)
			} else {
				encryptedStr += string(char + 13)
			}
		} else if char >= 'a' && char <= 'z' {
			if char+13 > 'z' {
				encryptedStr += string(char - 13)
			} else {
				encryptedStr += string(char + 13)
			}
		} else {
			encryptedStr += string(char) // Add non-letter characters as they are
		}
	}

	return encryptedStr
}

// Encrypt the message with reverse
func encrypt_reverse(s string) string {

	var encryptedStr string

	for x := 0; x < len(s); x++ {

		if s[x] >= 'A' && s[x] <= 'Z' {

			encryptedStr += string(('Z' - s[x]) + 'A')

		} else if s[x] >= 'a' && s[x] <= 'z' {

			encryptedStr += string(('z' - s[x]) + 'a')

		} else {

			encryptedStr += string(s[x])

		}
	}

	return encryptedStr

}

func encrypt_reverseAlphabetValue(s string) string {
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result

}

// Decrypt the message with rot13
func decrypt_rot13(s string) string {

	var encryptedStr string

	for x := 0; x < len(s); x++ {

		if s[x] >= 'A' && s[x] <= 'Z' {

			if s[x]+13 > 'Z' {

				encryptedStr += string(s[x] - 13)

			} else {

				encryptedStr += string(s[x] + 13)

			}

		} else if s[x] >= 'a' && s[x] <= 'z' {

			if s[x]+13 > 'z' {

				encryptedStr += string(s[x] - 13)

			} else {

				encryptedStr += string(s[x] + 13)

			}

		} else {

			encryptedStr += string(s[x])

		}
	}

	return encryptedStr

}

// Decrypt the message with reverse
func decrypt_reverse(s string) string {

	var encryptedStr string

	for x := 0; x < len(s); x++ {

		if s[x] >= 'A' && s[x] <= 'Z' {

			encryptedStr += string(('Z' - s[x]) + 'A')

		} else if s[x] >= 'a' && s[x] <= 'z' {

			encryptedStr += string(('z' - s[x]) + 'a')

		} else {

			encryptedStr += string(s[x])

		}
	}

	return encryptedStr

}

func decrypt_reverseAlphabetValue(s string) string {
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}

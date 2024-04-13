package main

import (
	"fmt"
)

// Main logic, envoking other functions
func main() {

	fmt.Println("\nWelcome to the Cypher Tool!\n")

	var toEncrypt bool
	var encoding, message, resStr, action string

	toEncrypt, encoding, message = getInput()

	if encoding == "rot13" {

		if toEncrypt {

			action = "Encrypted"

			resStr = encrypt_rot13(message)

		} else {

			action = "Decrypted"

			resStr = decrypt_rot13(message)

		}
	}

	if encoding == "reverse" {

		if toEncrypt {

			action = "Encrypted"

			resStr = encrypt_reverse(message)

		} else {

			action = "Decrypted"

			resStr = decrypt_reverse(message)

		}
	}

	fmt.Printf("\n%v message using %v:\n", action, encoding)

	fmt.Println(resStr)

}

// Get the input data required for the operation
func getInput() (toEncrypt bool, encoding string, message string) {

	var actionChoise int
	var cypherChoise int

	fmt.Println("Select operation (1/2):\n1. Encrypt.\n2. Decrypt.")

	for actionChoise != 1 && actionChoise != 2 {

		fmt.Scan(&actionChoise)

	}

	fmt.Println("\nSelect cypher (1/2):\n1. ROT13.\n2. Reverse.")

	for cypherChoise != 1 && cypherChoise != 2 {

		fmt.Scan(&cypherChoise)

	}

	fmt.Println("\nEnter the message:")

	fmt.Scan(&message)

	if actionChoise == 1 {

		toEncrypt = true

	}

	if cypherChoise == 1 {

		encoding = "rot13"

	} else {

		encoding = "reverse"

	}

	return

}

// Encrypt the message with rot13
func encrypt_rot13(s string) string {

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

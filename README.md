___
# Cypher Tool
---
---
## Best Team Ever:
---
## Teamleader

***Dmytro Mykhailenko***

## Members

***Merly Teder***

***Viktoria Tsetskaja***

---
---

>***"If something works - don't touch it"***
>
>*1st rule of programming*

---
---

# Introduction

### "What does the tool do?"

The tool prompts the user to select an operation (encrypt or decrypt) and an encryption method. In our case, Rot13, Reverse and ReverseAlphabetValue.

After the user enters a message, the program applies the selected encryption method to the message.

### The usage of the tool, with example

If encryption is selected and the Rot13 method is selected, the program applies Rot13 to the message.

If decryption and the Rot13 method are selected, the program performs the reverse Rot13 operation on the message.

Similar actions are performed for the Reverse and ReverseAlphabetValue methods.

### The cyphers used

Rot13: Replaces every letter in the message with a letter 13 positions ahead in the alphabet. Applies to the Latin alphabet.

Reverse: The letter changes to the corresponding letter, which is the same distance from the beginning of the alphabet, but in reverse order. For example, given the letter "i". This is the 9th letter from the beginning of the alphabet. After encryption, the letter changes to 9th from the end.

ReverseAlphabetValue: Reverses the order of characters in the message. For example, "hello" will become "olleh".

![Picture](https://media.geeksforgeeks.org/wp-content/cdn-uploads/string-reverse.jpg)


# Goals and objectives

### The program has to do the following:

- [X] Greet the user.

- [X] Allow the user to select the operation (encrypt or decrypt).

- [X] Allow the user to select the encryption type.

- [X] Allow the user to input the message to encrypt/decrypt.

- [X] Output the result of the operation.

- [X] Rot13, similar to ShiftBy you created for a single rune.

- [X] Reverse, similar to ReverseAlphabetValue you created for a single rune.

- [X] One more encryption of your choice.


# Nantes Hangman-Web Project

## Description

This project is a web server implementation of the Hangman game. The goal is to guess a randomly chosen word by suggesting letters before José (our hangman character) is fully hanged. You have 10 attempts to guess the word in the default difficulty mode. The game ends when all the letters of the word are revealed, the word is guessed, or the number of attempts reaches zero.

## Features

### Configuration
1. By default, the game will be configured in English, without any additional extensions.
2. Once the website is launched, you will have access to three pages: index.html (general presentation of the game), game.html (the Hangman game itself), and configuration.html (game configuration with various extensions).

### Game Settings
In the game configuration panel (configuration.html), you have access to the following settings:
- Language: This parameter allows you to change the language of the game. You can choose "en" for English or "fr" for French.
- VictoryCounter: Enables the win/loss counter.
- EnableDifficulty: Activates different difficulty levels (veryeasy, easy, middle, hard, expert, hacker). For example:
    - In veryeasy mode, you get more attempts. 
    - In hacker mode, no hints are given to guess the word, but the program reveals a random number of letters equal to **(word length / 2) - 1**. Additionally, error tolerances will cause you to lose one attempt.
- AddWordAfterGame: Allows you to add an additional word to the game based on the selected difficulty mode.
- EnableJokers: Lets you activate the use of jokers; to earn a joker, you need a victory.
- Compatibility: All these features can be enabled simultaneously to enrich the gaming experience.

### Basic Parameters
The Hangman program operates as follows:

1. Word File: The program uses a .txt file. These files will automatically be created in the /files/ directory if they don't already exist. If a file is empty, it will be populated with default values. However, if the file already contains words, the game will not modify it, allowing you to add words manually as needed. Each word is separated by a newline (\n).

2. Starting the game:
    - The program randomly selects a word from the file. If the game detects an invalid word, it will automatically move on to another word; if all words are invalid, the game will inform you automatically.
    - A word is valid if it consists of these characters only:
        - Uppercase: A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z
        - Lowercase: a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z
        - Accented characters: À, Á, Â, Ã, Ä, Å, Ç, È, É, Ê, Ë, Ì, Í, Î, Ò, Ó, Ô, Õ, Ö, Ù, Ú, Û, Ü, Ý, à, á, â, ã, ä, å, è, é, ê, ë, ò, ó, ô, õ, ö, ù, ú, û, ü
    - It reveals a certain number of random letters from the word, where n = (length of the word / 2) - 1 (except in hacker mode if it is enabled).
    - You have 10 attempts to guess the word (except in super easy mode if it is enabled).
    - You propose a letter by entering input from the keyboard. If the letter is not present, you lose an attempt and an error message is displayed.
    - You have the option to enter a word; however, if the answer is incorrect, you lose two attempts, and an error message is displayed.
    - If the letter is correct, all its occurrences in the word are revealed.
    - The game continues until the word is fully guessed or the number of attempts is exhausted.
    - Additionally, messages will inform you if a word or letter has already been proposed.

### José's Graphics

A unique aspect of this project is José, the character who will be hanged if you fail to guess the word. With each mistake, José's position gradually changes until he is fully hanged.

## Prerequisites

Before running the project, ensure that the following are installed on your machine:

- ![Go Version](https://img.shields.io/badge/Go-1.23-blue)
- **Git** to clone the repository

## Installation

To install and run the project, follow these steps:

1. Clone the project repository using Git:
   ```bash
   git clone https://ytrack.learn.ynov.com/git/gremy/hangman-web.git
   ```
2. Access the project directory:
   ```bash
   cd hangman-web
   ```
3. Ensure that all necessary files are present, including main.go, before launching the game.

4. Start the game by executing the command:
   ```bash
   go run main.go
   ```
## How to Test the Project?

To test the project, simply run the following command:

```bash
    go run main.go
```
This command should be executed at the root of the project.

## Project Developers

This project was developed by:

- **[Guibert Rémy](https://github.com/thedevrems)** - Project Manager - Developer
- **Emma Lainé** - Developer

If you wish to contribute or discuss ideas, feel free to contact us!

### Acknowledgments

A big thank you to everyone who contributed to the project and to the open-source community for their support and resources.


Thank you for testing our project!

![Go Version](https://img.shields.io/badge/Go-1.23-blue)
![Licence](https://img.shields.io/badge/License-MIT-green)
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

    "github.com/gobuffalo/packr"
)

const (
	lives = 7
	easy = 1
	medium = 2
	hard = 3
)

func printTheMan(missedCount int) {
	fmt.Println("=====================")
	fmt.Println("        \\          ")
	switch {
	case missedCount == 0:
		fmt.Println("                   ")
		fmt.Println("                   ")
		fmt.Println("                   ")
		fmt.Println("                   ")
	case missedCount == 1:
		fmt.Println("         |         ")
		fmt.Println("                   ")
		fmt.Println("                   ")
		fmt.Println("                   ")
	case missedCount == 2:
		fmt.Println("         |         ")
		fmt.Println("         O         ")
		fmt.Println("                   ")
		fmt.Println("                   ")
	case missedCount == 3:
		fmt.Println("         |         ")
		fmt.Println("         O         ")
		fmt.Println("        /           ")
		fmt.Println("                   ")
	case missedCount == 4:
		fmt.Println("         |         ")
		fmt.Println("         O         ")
		fmt.Println("        /|          ")
		fmt.Println("                   ")
	case missedCount == 5:
		fmt.Println("         |         ")
		fmt.Println("         O         ")
		fmt.Println("        /|\\         ")
		fmt.Println("                   ")
	case missedCount == 6:
		fmt.Println("         |         ")
		fmt.Println("         O         ")
		fmt.Println("        /|\\         ")
		fmt.Println("        /          ")
	case missedCount == 7:
		fmt.Println("         |         ")
		fmt.Println("         O         ")
		fmt.Println("        /|\\         ")
		fmt.Println("        / \\        ")
	}
	fmt.Println()
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func readLines(filename string) []string {
    box := packr.NewBox("./assets")
    s, err := box.FindString(filename)
    if err != nil {
        log.Fatal(err)
    }

	lines := make([]string, 0)
    for _, line := range strings.Split(s, "\n") {
		lines = append(lines, line)
	}
	return lines
}

func printWord(word string) {
	for idx, l := range word {
		if idx != 0 {
			fmt.Print(" ")
		}
		fmt.Print(string(l))
	}
	fmt.Println()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	clearScreen()
	fmt.Println("Welcome to Hangman")

	words := readLines("words-and-phrases.txt")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Choose your difficulty level(%d easy, %d medium(default), %d hard): ", easy, medium, hard)
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	maxMissCount := lives
	missedCount := 1
	switch num {
	case easy:
		missedCount = 0
	case hard:
		missedCount = 2
	}

	guessWord := words[rand.Intn(len(words))]
	tmpRevealWord := make([]rune, 0)
	for _, l := range guessWord {
		if l == ' ' {
			tmpRevealWord = append(tmpRevealWord, ' ')
		} else {
			tmpRevealWord = append(tmpRevealWord, '_')
		}
	}
	revealWord := string(tmpRevealWord)
	guesses := ""

	clearScreen()
	printTheMan(missedCount)
	printWord(revealWord)
	fmt.Println("\nYou've guessed:", guesses)

	for fmt.Print("\n=> "); scanner.Scan(); fmt.Print("\n=> ") {
		guess := scanner.Text()
		if len(guess) == 0 || guess[0] < 'a' || guess[0] > 'z' {
			fmt.Println("Invalid guess")
			continue;
		} else if strings.Contains(guesses, string(guess[0])) {
			fmt.Println("Already guessed")
			continue
		} else {
			clearScreen()
		}
		if guesses == "" {
			guesses = string(guess[0])
		} else {
			guesses = guesses + "," + string(guess[0])
		}

		tmpRevealWord := []byte(revealWord)
		revealed := false

		for idx, letter := range guessWord {
			if uint8(letter) == guess[0] {
				revealed = true
				tmpRevealWord[idx] = byte(letter)
			}
		}
		if !revealed {
			missedCount++
		}
		printTheMan(missedCount)
		revealWord = string(tmpRevealWord)
		printWord(revealWord)
		if !revealed && missedCount == maxMissCount {
			printWord(guessWord)
			fmt.Println("\nYou lose!")
			return
		}
		if revealWord == guessWord {
			fmt.Println("\nYou win!")
			return
		}
		fmt.Println("\nYou've guessed:", guesses)
	}
}

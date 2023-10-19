package main

import (
	"fmt"
	"strings"
)

const (
	LettersMode = "Letters"
	NumbersMode = "Numbers"
)

type Watch struct {
	Mode string
}

type Button interface {
	Press() error
}

type LetterButton struct {
	Watch  *Watch
	Letter string
}

type NumberButton struct {
	Watch  *Watch
	Number int
}

func (lb *LetterButton) Press() error {
	if lb.Watch.Mode != LettersMode {
		return fmt.Errorf("Button cannot be pressed in %s mode", lb.Watch.Mode)
	}
	fmt.Println(lb.Letter)
	return nil
}

func (nb *NumberButton) Press() error {
	if nb.Watch.Mode != NumbersMode {
		return fmt.Errorf("Button cannot be pressed in %s mode", nb.Watch.Mode)
	}
	fmt.Println(nb.Number)
	return nil
}

func (w *Watch) ChangeMode(newMode string) {
	w.Mode = newMode
	fmt.Printf("Watch mode changed to %s\n", newMode)
}

func click(watch *Watch, digit string) error {
	buttons := map[string]Button{}
	if watch.Mode == LettersMode {
		buttons = map[string]Button{
			"1": &LetterButton{Watch: watch, Letter: "A"},
			"2": &LetterButton{Watch: watch, Letter: "B"},
			"3": &LetterButton{Watch: watch, Letter: "C"},
		}
	} else if watch.Mode == NumbersMode {
		buttons = map[string]Button{
			"1": &NumberButton{Watch: watch, Number: 1},
			"2": &NumberButton{Watch: watch, Number: 2},
			"3": &NumberButton{Watch: watch, Number: 3},
		}
	}
	button, exists := buttons[digit]
	if !exists {
		return fmt.Errorf("invalid button: %s", digit)
	}

	return button.Press()
}

func main() {
	watch := &Watch{Mode: LettersMode} // Create a pointer to Watch

	for {
		fmt.Print("Enter button (1, 2, 3, 4): ")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		switch strings.TrimSpace(input) {
		case "1", "2", "3":
			err := click(watch, input)
			if err != nil {
				fmt.Println(err)
			}
		case "4":
			if watch.Mode == "Letters" {
				watch.ChangeMode("Numbers")
			} else if watch.Mode == "Numbers" {
				watch.ChangeMode("Letters")
			}
		default:
			fmt.Println("Invalid button. Please try again.")
		}
	}
}

package main

import (
	"testing"
)

func TestClick(t *testing.T) {
	watchLetters := &Watch{Mode: LettersMode}
	watchNumbers := &Watch{Mode: NumbersMode}

	err := click(watchLetters, "1")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = click(watchNumbers, "1")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = click(watchLetters, "2")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = click(watchNumbers, "2")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = click(watchLetters, "3")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = click(watchNumbers, "3")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = click(watchNumbers, "1")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = click(watchLetters, "5")
	expectedErrMsg := "invalid button: 5"
	if err == nil || err.Error() != expectedErrMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedErrMsg, err)
	}
}

func TestChangeMode(t *testing.T) {
	watch := &Watch{Mode: LettersMode}
	watch.ChangeMode(NumbersMode)
	if watch.Mode != NumbersMode {
		t.Errorf("Expected mode to be NumbersMode, got %s", watch.Mode)
	}

	watch.ChangeMode(LettersMode)
	if watch.Mode != LettersMode {
		t.Errorf("Expected mode to be LettersMode, got %s", watch.Mode)
	}
}

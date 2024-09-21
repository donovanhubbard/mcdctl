package tui

import (
	"fmt"
	"testing"
)

func TestOneUpReturnsText(t *testing.T) {
	cmd1 := "get foo"
	ch := CommandHistory{}
	ch.CommandTexts = []CommandText{
		{UserText: cmd1},
		{ResponseText: "Error"},
	}

	targetText := ch.GetLastCommand()

	if targetText.UserText != cmd1 {
		t.Fatal(fmt.Sprintf("Expected: '%s', got '%s'", cmd1, targetText.UserText))
	}
}

func TestTwoUpReturnsFirstText(t *testing.T) {
	cmd1 := "get foo"
	ch := CommandHistory{}
	ch.CommandTexts = []CommandText{
		{UserText: cmd1},
		{ResponseText: "ERROR"},
	}

	ch.GetLastCommand()
	targetText := ch.GetLastCommand()

	if targetText.UserText != cmd1 {
		t.Fatal(fmt.Sprintf("Expected: '%s', got '%s'", cmd1, targetText.UserText))
	}
}

func TestOneUpWithTwoCommandsReturnsText(t *testing.T) {
	cmd1 := "get foo"
	cmd2 := "get bar"
	ch := CommandHistory{}
	ch.CommandTexts = []CommandText{
		{UserText: cmd1},
		{ResponseText: "Failure"},
		{UserText: cmd2},
		{ResponseText: "Success"},
	}

	targetText := ch.GetLastCommand()

	if targetText.UserText != cmd2 {
		t.Fatal(fmt.Sprintf("Expected: '%s', got '%s'", cmd2, targetText.UserText))
	}
}

func TestTwoUpWithTwoTextReturnsText(t *testing.T) {
	cmd1 := "get foo"
	cmd2 := "get bar"
	ch := CommandHistory{}
	ch.CommandTexts = []CommandText{
		{UserText: cmd1},
		{ResponseText: "Failure"},
		{UserText: cmd2},
		{ResponseText: "Success"},
	}

	ch.GetLastCommand()
	targetText := ch.GetLastCommand()

	if targetText.UserText != cmd1 {
		t.Fatal(fmt.Sprintf("Expected: '%s', got '%s'", cmd1, targetText.UserText))
	}
}

func TestUpPastStartReturnsFirstText(t *testing.T) {
	cmd1 := "get foo"
	cmd2 := "get bar"
	ch := CommandHistory{}
	ch.CommandTexts = []CommandText{
		{UserText: cmd1},
		{ResponseText: "Failure"},
		{UserText: cmd2},
		{ResponseText: "Success"},
	}

	ch.GetLastCommand()
	ch.GetLastCommand()
	targetText := ch.GetLastCommand()

	if targetText.UserText != cmd1 {
		t.Fatal(fmt.Sprintf("Expected: '%s', got '%s'", cmd1, targetText.UserText))
	}
}

func TestUpNoCommandsReturnsEmptyResponse(t *testing.T) {
	ch := CommandHistory{}
	targetText := ch.GetLastCommand()

	if targetText.UserText != "" {
		t.Fatal(fmt.Sprintf("Expected empty response got: '%s'", targetText.UserText))
	}
}

func TestTwoUpOneDownReturnsText(t *testing.T) {
	cmd1 := "get foo"
	cmd2 := "get bar"
	ch := CommandHistory{}
	ch.CommandTexts = []CommandText{
		{UserText: cmd1},
		{ResponseText: "Failure"},
		{UserText: cmd2},
		{ResponseText: "Success"},
	}

	ch.GetLastCommand()
	ch.GetLastCommand()
	targetText := ch.GetNextCommand()

	if targetText.UserText != cmd2 {
		t.Fatal(fmt.Sprintf("Expected: '%s', got '%s'", cmd2, targetText.UserText))
	}
}

func TestOneUpTwoDownPassedBeginningReturnsFirstText(t *testing.T) {
	cmd1 := "get foo"
	ch := CommandHistory{}
	ch.CommandTexts = []CommandText{
		{UserText: cmd1},
		{ResponseText: "Failure"},
	}

	ch.GetLastCommand()
	targetText := ch.GetNextCommand()

	if targetText.UserText != cmd1 {
		t.Fatal(fmt.Sprintf("Expected: '%s', got '%s'", cmd1, targetText.UserText))
	}
}

func TestThreeUpTwoDownReturnsText(t *testing.T) {
	cmd1 := "get foo"
	cmd2 := "get bar"
	cmd3 := "get baz"
	ch := CommandHistory{}
	ch.CommandTexts = []CommandText{
		{UserText: cmd1},
		{ResponseText: "Failure"},
		{UserText: cmd2},
		{ResponseText: "Success"},
		{UserText: cmd3},
		{ResponseText: "Success"},
	}

	ch.GetLastCommand()
	ch.GetLastCommand()
	ch.GetLastCommand()
	ch.GetNextCommand()
	targetText := ch.GetNextCommand()

	if targetText.UserText != cmd3 {
		t.Fatal(fmt.Sprintf("Expected: '%s', got '%s'", cmd3, targetText.UserText))
	}
}

func TestThreeUpBeyondStartOneDownReturnsText(t *testing.T) {
	cmd1 := "get foo"
	cmd2 := "get bar"
	ch := CommandHistory{}
	ch.CommandTexts = []CommandText{
		{ResponseText: "Connected successfully"},
		{UserText: cmd1},
		{ResponseText: "Failure"},
		{UserText: cmd2},
		{ResponseText: "Success"},
	}

	ch.GetLastCommand()
	ch.GetLastCommand()
	ch.GetLastCommand()
	targetText := ch.GetNextCommand()

	if targetText.UserText != cmd2 {
		t.Fatal(fmt.Sprintf("Expected: '%s', got '%s'", cmd2, targetText.UserText))
	}
}

func TestDownNoResponsesReturnsBlankStruct(t *testing.T) {
	ch := CommandHistory{}

	targetText := ch.GetNextCommand()

	if targetText.UserText != "" {
		t.Fatal(fmt.Sprintf("Expected empty response got: '%s'", targetText.UserText))
	}
}

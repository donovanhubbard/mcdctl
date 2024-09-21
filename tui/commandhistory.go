package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/donovanhubbard/mcdctl/logging"
)

const (
	MAX_HISTORY = 6
)

type CommandHistory struct {
	CommandTexts []CommandText
	Height       int
	Width        int
	currentIndex *int
}

func (ch CommandHistory) View() string {
	var borderStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Align(lipgloss.Left).
		Width(ch.Width)

	var sb strings.Builder

	if len(ch.CommandTexts) > MAX_HISTORY {
		for i := len(ch.CommandTexts) - MAX_HISTORY; i < len(ch.CommandTexts); i++ {
			sb.WriteString(ch.CommandTexts[i].View())
		}

	} else {
		for _, v := range ch.CommandTexts {
			sb.WriteString(v.View())
		}
	}

	return borderStyle.Render(sb.String())
}

func (ch *CommandHistory) GetLastCommand() *CommandText {
	logging.Debug("Starting GetLastCommand")

	if len(ch.CommandTexts) == 0 {
		return &CommandText{}
	}

	var targetIndex int

	if ch.currentIndex == nil {
		targetIndex = len(ch.CommandTexts) - 1
		ch.currentIndex = &targetIndex
	} else {
		targetIndex = *ch.currentIndex - 1
	}

	logging.Debug(fmt.Sprintf("Starting loop at %d", *ch.currentIndex))
	for i := targetIndex; i >= 0; i-- {
		logging.Debug(fmt.Sprintf("i=%d", i))
		if ch.CommandTexts[i].UserText != "" {
			ch.currentIndex = &i
			return &ch.CommandTexts[i]
		}
	}
	logging.Debug("No command history was found")
	return &ch.CommandTexts[*ch.currentIndex]
}

func (ch *CommandHistory) GetNextCommand() *CommandText {
	logging.Debug("Starting GetNextCommand")
	if ch.currentIndex == nil || len(ch.CommandTexts) == 0 {
		return &CommandText{}
	}

	logging.Debug(fmt.Sprintf("currentIndex=%d", *ch.currentIndex))

	for j, v := range ch.CommandTexts {
		logging.Debug(fmt.Sprintf("j=%d", j))
		logging.Debug(fmt.Sprintf("UserText=%s", v.UserText))
		logging.Debug(fmt.Sprintf("ResponseText=%s", v.ResponseText))
	}

	for i := *ch.currentIndex + 1; i < len(ch.CommandTexts); i++ {
		logging.Debug(fmt.Sprintf("i=%d", i))
		logging.Debug(fmt.Sprintf("UserText=%s", ch.CommandTexts[i].UserText))
		logging.Debug(fmt.Sprintf("ResponseText=%s", ch.CommandTexts[i].ResponseText))
		if ch.CommandTexts[i].UserText != "" {
			logging.Debug("Found next command")
			ch.currentIndex = &i
			return &ch.CommandTexts[i]
		}
	}

	if ch.currentIndex == nil {
		logging.Debug("Couldn't find any commands and didn't start with any. Returning blank struct")
		return &CommandText{}
	}

	logging.Debug("Couldn't find any commands. Returning the one we started with")
	return &ch.CommandTexts[*ch.currentIndex]
}

func (ch *CommandHistory) ResetCurrentIndex() {
	ch.currentIndex = nil
}

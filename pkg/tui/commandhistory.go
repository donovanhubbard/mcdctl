package tui

import (
  "strings"
  "github.com/charmbracelet/lipgloss"
)

const (
  MAX_HISTORY = 3
)

type CommandHistory struct {
  CommandTexts []CommandText
  Height int
  Width int
}

func (ch CommandHistory) View() string {
  var borderStyle = lipgloss.NewStyle().
    BorderStyle(lipgloss.RoundedBorder()).
    BorderForeground(lipgloss.Color("63")).
    Align(lipgloss.Left).
    Width(ch.Width)

  var sb strings.Builder

  if len(ch.CommandTexts) > MAX_HISTORY {
    for i:=len(ch.CommandTexts)-MAX_HISTORY;i<len(ch.CommandTexts);i++ {
      sb.WriteString(ch.CommandTexts[i].View())
    }

  } else {
    for _, v := range ch.CommandTexts {
      sb.WriteString(v.View())
    }
  }

  return borderStyle.Render(sb.String())
}

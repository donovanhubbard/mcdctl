package tui

import (
  "strings"
  "github.com/charmbracelet/lipgloss"
)

type CommandText struct {
  UserText string
  ResponseText string
  Success bool
}

const (
  RED = "1"
  GREEN = "2"
  YELLOW = "3"
)

func (ct CommandText) View() string {
  var responseStyle lipgloss.Style
  if ct.Success {
    responseStyle = lipgloss.
      NewStyle().
      Foreground(lipgloss.Color(GREEN))
  } else {
    responseStyle = lipgloss.
      NewStyle().
      Foreground(lipgloss.Color(RED))
  }
  userStyle := lipgloss.
    NewStyle().
    Foreground(lipgloss.Color(YELLOW))

  var sb strings.Builder

  if ct.UserText != "" {
    sb.WriteString("> ")
    sb.WriteString(userStyle.Render(ct.UserText))
    sb.WriteString("\n")
  }

  if ct.ResponseText != "" {
    lines := strings.Split(ct.ResponseText, "\r\n")
    for _, line := range lines {
      sb.WriteString(responseStyle.Render(line))
      sb.WriteString("\n")
    }
  }

  return sb.String()
}

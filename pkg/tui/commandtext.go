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

func (ct CommandText) View() string {
  userStyle := lipgloss.
    NewStyle().
    Foreground(lipgloss.Color("2"))
  responseStyle := lipgloss.
    NewStyle().
    Foreground(lipgloss.Color("1"))
  var sb strings.Builder

  if ct.UserText != "" {
    sb.WriteString("> ")
    sb.WriteString(userStyle.Render(ct.UserText))
    sb.WriteString("\n")
  }

  if ct.ResponseText != "" {
    sb.WriteString(responseStyle.Render(ct.ResponseText))
    sb.WriteString("\n")
  }

  return sb.String()
}

package tui

import (
  "strings"
)

type CommandText struct {
  UserText string
  ResponseText string
  Success bool
}

func (ct CommandText) View() string {
  var sb strings.Builder
  sb.WriteString(">")
  sb.WriteString(ct.UserText)
  sb.WriteString("\n")
  sb.WriteString(ct.ResponseText)
  sb.WriteString("\n")
  return sb.String()
}

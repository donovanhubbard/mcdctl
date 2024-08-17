package tui

import (
  "strings"
)

const (
  MAX_HISTORY = 3
)

type CommandHistory struct {
  CommandTexts []CommandText
}

func (ch CommandHistory) View() string {
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

  return sb.String()
}

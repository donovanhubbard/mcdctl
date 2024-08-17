package tui

import (
  "strings"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/donovanhubbard/mcdctl/pkg/memcachedclient"
)

type Model struct {
  socketAddress memcachedclient.SocketAddress
  textInput textinput.Model
  commandHistory CommandHistory
}

func NewModel(socketAddress memcachedclient.SocketAddress) Model {
  ti := textinput.New()
  ti.Placeholder = "memcached command"
  ti.Focus()
  ti.CharLimit = 125
  ti.Width = 75

  return Model{
    socketAddress: socketAddress,
    textInput: ti,
  }
}

func (m Model) Init() tea.Cmd {
  return textinput.Blink
}

func (m Model) View() string {
  var sb strings.Builder
  sb.WriteString(m.commandHistory.View())
  sb.WriteString("Enter memcached commands\n")
  sb.WriteString(m.textInput.View())

  return sb.String()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
    case tea.KeyEnter:
      text := m.textInput.Value()
      m.textInput.Reset()
      commandText := CommandText {
        UserText: text,
      }
      m.commandHistory.CommandTexts = append(m.commandHistory.CommandTexts, commandText)
		}
  }

  m.textInput, cmd = m.textInput.Update(msg)
  return m, cmd
}

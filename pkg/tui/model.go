package tui

import (
  "strings"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/donovanhubbard/mcdctl/pkg/memcachedclient"
  "github.com/charmbracelet/lipgloss"
)

type Model struct {
  height int
  width int
  client memcachedclient.Client
  textInput textinput.Model
  commandHistory CommandHistory
}

func NewModel(socketAddress memcachedclient.SocketAddress) Model {
  ti := textinput.New()
  ti.Placeholder = "memcached command"
  ti.Focus()
  ti.CharLimit = 125
  ti.Width = 75

  client := memcachedclient.Client{
    SocketAddress: socketAddress,
  }

  return Model{
    client: client,
    textInput: ti,
  }
}

func generateConnectionCmd(c *memcachedclient.Client) tea.Cmd {
  return func () tea.Msg {
    err := c.Dial()
    connectMsg := ConnectMsg {
      Error: err,
    }
    return connectMsg
  }
}

func (m *Model) SetSize(height, width int) {
  m.height = height
  m.width = width

  m.textInput.Width = width - 13

  m.commandHistory.Width = width - 10
}

func (m Model) Init() tea.Cmd {
  connectCmd := generateConnectionCmd(&m.client)
  return tea.Batch(textinput.Blink,connectCmd)
}

func (m Model) View() string {
  var sb strings.Builder

  var textInputStyle = lipgloss.NewStyle().
    BorderStyle(lipgloss.RoundedBorder()).
    BorderForeground(lipgloss.Color("63"))

  borderedCommandHistory := m.commandHistory.View()
  borderedTextInput := textInputStyle.Render(m.textInput.View())

  joinedText := lipgloss.JoinVertical(lipgloss.Center, borderedCommandHistory, borderedTextInput)

  sb.WriteString(joinedText)

  return sb.String()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd

	switch msg := msg.(type) {
  case tea.WindowSizeMsg:
    m.SetSize(msg.Height, msg.Width)
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
  case ConnectMsg:
    var commandText CommandText
    if msg.Error == nil {
      commandText.ResponseText = "Connected successfully"
    } else {
      commandText.ResponseText = msg.Error.Error()
    }
    m.commandHistory.CommandTexts = append(m.commandHistory.CommandTexts, commandText)
  }

  m.textInput, cmd = m.textInput.Update(msg)
  return m, cmd
}

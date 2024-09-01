package tui

import (
  "fmt"
  "strings"
  "reflect"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

  "github.com/charmbracelet/lipgloss"
	"github.com/donovanhubbard/mcdctl/pkg/memcachedclient"
	"github.com/donovanhubbard/mcdctl/pkg/utils"
  "github.com/charmbracelet/bubbles/cursor"
)

type Model struct {
  height int
  width int
  client *memcachedclient.Client
  textInput textinput.Model
  commandHistory CommandHistory
}

func NewModel(socketAddress memcachedclient.SocketAddress) Model {
  ti := textinput.New()
  ti.Placeholder = "memcached command"
  ti.Focus()
  ti.CharLimit = 125
  ti.Width = 75

  client := &memcachedclient.Client{
    SocketAddress: socketAddress,
  }

  utils.Sugar.Debug(fmt.Sprintf("Memcached server is at %s", client.SocketAddress.String()))

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

    if c.IsConnected() {
      utils.Sugar.Debug("Is connected inside of cmd")
    }else{
      utils.Sugar.Error("Is NOT connected inside of cmd.")
    }

    return connectMsg
  }
}

func sendMemcachedCmd(c *memcachedclient.Client, text string) tea.Cmd {
  return func() tea.Msg{
    response, err := c.SendCommand(text)
    if err != nil {
      return MemcachedResponseMsg{
        Error: err,
      }
    }
    return MemcachedResponseMsg{
      Response: response,
    }
  }
}

func (m *Model) SetSize(height, width int) {
  m.height = height
  m.width = width

  m.textInput.Width = width - 13

  m.commandHistory.Width = width - 10
}

func (m Model) Init() tea.Cmd {
  connectCmd := generateConnectionCmd(m.client)
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

  if reflect.TypeOf(msg) != reflect.TypeOf(cursor.BlinkMsg{}) {
    utils.Sugar.Debug(fmt.Sprintf("model.Update received msg of type %T", msg))
  }

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
      if text == "quit" {
        return m, tea.Quit
      }
      commandText := CommandText {
        UserText: text,
      }
      m.commandHistory.CommandTexts = append(m.commandHistory.CommandTexts, commandText)
      return m, sendMemcachedCmd(m.client, text)
		}
  case ConnectMsg:
    var commandText CommandText
    if msg.Error == nil {
      commandText.ResponseText = "Connected successfully"
      commandText.Success = true
    } else {
      commandText.ResponseText = msg.Error.Error()
      commandText.Success = false
    }
    m.commandHistory.CommandTexts = append(m.commandHistory.CommandTexts, commandText)

  case MemcachedResponseMsg:
    var commandText CommandText
    if msg.Error != nil {
      utils.Sugar.Debug(fmt.Sprintf("Received msg with error '%s'", msg.Error.Error()))
      commandText.ResponseText = msg.Error.Error()
      commandText.Success = false
    } else {
      utils.Sugar.Debug(fmt.Sprintf("Received msg with response '%s'", msg.Response))
      commandText.ResponseText = msg.Response
      commandText.Success = true
    }
    m.commandHistory.CommandTexts = append(m.commandHistory.CommandTexts, commandText)
  }

  m.textInput, cmd = m.textInput.Update(msg)
  return m, cmd
}

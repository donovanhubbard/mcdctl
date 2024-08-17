package tui

import (
  tea "github.com/charmbracelet/bubbletea"
  "github.com/donovanhubbard/mcdctl/pkg/memcachedclient"
)

type Model struct {
  socketAddress memcachedclient.SocketAddress
}

func NewModel(socketAddress memcachedclient.SocketAddress) Model {
  return Model{
    socketAddress: socketAddress,
  }
}

func (m Model) Init() tea.Cmd {
  return nil
}

func (m Model) View() string {
  return "hi bob"
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  // var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
  }
  return m, nil
}

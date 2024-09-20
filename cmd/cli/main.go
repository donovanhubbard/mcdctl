package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/donovanhubbard/mcdctl/pkg/logging"
	"github.com/donovanhubbard/mcdctl/pkg/memcachedclient"
	"github.com/donovanhubbard/mcdctl/pkg/tui"
)

func main() {
	logging.InitializeLogger()
	logging.Info("Starting program")
	socketAddress, err := getSocketAddress(os.Args[1:])

	if err != nil {
		usageStatement := generateUsageStatement(os.Args[0])
		fmt.Println(usageStatement)
		fmt.Println(err)
		os.Exit(1)
	}

	p := tea.NewProgram(tui.NewModel(socketAddress), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func generateUsageStatement(programName string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("usage: %s <hostname>[:port]", programName))
	return sb.String()
}

func getSocketAddress(args []string) (memcachedclient.SocketAddress, error) {
	s := memcachedclient.SocketAddress{}
	if len(args) != 1 {
		return s, errors.New("Must specify the host and port")
	}

	split := strings.Split(args[0], ":")

	if len(split) == 1 {
		s.Host = split[0]
		s.Port = memcachedclient.DEFAULT_PORT
		return s, nil
	} else if len(split) == 2 {
		s.Host = split[0]
		port, err := strconv.Atoi(split[1])
		if err != nil {
			return s, errors.New("Must specify numeric port")
		}
		s.Port = port
		return s, nil
	} else {
		return s, errors.New("Must specify the host and port in 'host:port' format")
	}
}

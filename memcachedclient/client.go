package memcachedclient

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/donovanhubbard/mcdctl/logging"
)

const (
	DEFAULT_PORT = 11211
)

type SocketAddress struct {
	Host string
	Port int
}

func (s SocketAddress) String() string {
	return fmt.Sprintf("%v:%d", s.Host, s.Port)
}

type Client struct {
	SocketAddress SocketAddress
	connection    net.Conn
}

func (c *Client) IsConnected() bool {
	if c.connection == nil {
		return false
	} else {
		return true
	}
}

func (c *Client) Dial() error {
	logging.Info(fmt.Sprintf("Attempting to connect to %s", c.SocketAddress.String()))
	conn, err := net.Dial("tcp", c.SocketAddress.String())
	if err == nil {
		logging.Debug("connection was successful")
		c.connection = conn
		return nil
	} else {
		logging.Error(fmt.Sprintf("Failed to connect to memcached server. %s", err.Error()))
		return err
	}
}

func (c *Client) SendCommand(text string) (string, error) {
	logging.Debug(fmt.Sprintf("Sending command to server: '%s'", text))
	if c.connection == nil {
		logging.Error("Trying to send command to server while disconnected")
		return "", errors.New("Tried to send command to server while disconnected")
	}
	fmt.Fprintf(c.connection, "%s\r\n", text)

	var sb strings.Builder
	var line string
	var err error
	reader := bufio.NewReader(c.connection)

	// flood the buffer
	reader.Peek(1)

	for reader.Buffered() > 0 {
		line, err = reader.ReadString('\n')
		logging.Debug(fmt.Sprintf("Received non-error from memcached: '%s'", line))
		sb.WriteString(line)
		if err != nil {
			logging.Error(fmt.Sprintf("Received connection error from memcached: '%s'", err.Error()))
			return "", err
		}
	}

	if strings.HasPrefix(line, "ERROR") || strings.HasPrefix(line, "NOT_FOUND") || strings.HasPrefix(line, "NOT_STORED") {
		errorText := strings.TrimSpace(sb.String())
		logging.Error(fmt.Sprintf("Received error message from memcached: '%s'", errorText))
		return "", errors.New(errorText)
	} else {
		text := sb.String()
		responseText := strings.TrimSpace(text)
		logging.Debug(fmt.Sprintf("Received message from memcached: '%s'", responseText))
		return responseText, nil
	}
}

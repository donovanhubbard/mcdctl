package memcachedclient

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/donovanhubbard/mcdctl/pkg/utils"
	// "github.com/charmbracelet/log"
)

const (
  DEFAULT_PORT = 11211
)

type SocketAddress struct {
  Host string
  Port int
}

func (s SocketAddress) String() string{
  return fmt.Sprintf("%v:%d",s.Host,s.Port)
}

type Client struct {
  SocketAddress SocketAddress
  connection net.Conn
}

func (c *Client)IsConnected()bool{
  if c.connection == nil {
    return false
  }else{
    return true
  }
}

func (c *Client) Dial() error{
  utils.Sugar.Info(fmt.Sprintf("Attempting to connect to %s",c.SocketAddress.String()))
  conn, err := net.Dial("tcp",c.SocketAddress.String())
  if err == nil {
    utils.Sugar.Debug("connection was successful")
    c.connection = conn
    return nil
  } else {
    utils.Sugar.Error(fmt.Sprintf("Failed to connect to memcached server. %s", err.Error()))
    return err
  }
}

func (c *Client) SendCommand(text string) (string,error) {
  utils.Sugar.Debug(fmt.Sprintf("Sending command to server: '%s'",text))
  if c.connection == nil {
    utils.Sugar.Error("Trying to send command to server while disconnected")
    return "", errors.New("Tried to send command to server while disconnected")
  }
  fmt.Fprintf(c.connection, "%s\r\n", text)

  var sb strings.Builder
  reader := bufio.NewReader(c.connection)

  for {
    utils.Sugar.Debug("Waiting on reader.ReadString()")
    line, err := reader.ReadString('\n')
    utils.Sugar.Debug(fmt.Sprintf("Received non-error from memcached: '%s'",line))
    sb.WriteString(line)
    if err != nil {
      utils.Sugar.Error(fmt.Sprintf("Received connection error from memcached: '%s'",err.Error()))
      return "", err
    }
    if strings.HasPrefix(line, "ERROR") || strings.HasPrefix(line,"NOT_FOUND") || strings.HasPrefix(line,"NOT_STORED") {
      errorText := strings.TrimSpace(sb.String())
      utils.Sugar.Error(fmt.Sprintf("Received error from memcached: '%s'",errorText))
      return "", errors.New(errorText)
    }
    if strings.HasPrefix(line, "END") || strings.HasPrefix(line, "STORED") || strings.HasPrefix(line,"DELETED") || strings.HasPrefix(line,"OK"){
      text := sb.String()
      utils.Sugar.Debug(fmt.Sprintf("before trim %s",text))
      responseText := strings.TrimSpace(text)
      utils.Sugar.Debug(fmt.Sprintf("Received error from memcached: '%s'",responseText))
      return responseText, nil
    }
  }

}

package memcachedclient

import (
  "net"
  "fmt"
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

func (c *Client) Dial() error{
  conn, err := net.Dial("tcp",c.SocketAddress.String())
  if err == nil {
    c.connection = conn
  }
  return err
}

package memcachedclient

type SocketAddress struct {
  Host string
  Port int
}

const (
  DEFAULT_PORT = 11211
)

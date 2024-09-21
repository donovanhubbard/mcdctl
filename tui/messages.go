package tui

type ConnectMsg struct {
  Error error
}

type MemcachedResponseMsg struct {
  Error error
  Response string
}

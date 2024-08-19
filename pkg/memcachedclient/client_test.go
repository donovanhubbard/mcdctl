package memcachedclient

import (
  "testing"
)

func TestSocketAddressToStringSucceeds(t *testing.T){
  s := SocketAddress{
    Host: "memcached.foo.com",
    Port: 11211,
  }

  str := s.String()

  if str != "memcached.foo.com:11211" {
    t.Fatal("Should convert to <host>:<port>")
  }
}

package main

import (
	"fmt"
	"testing"

	"github.com/donovanhubbard/mcdctl/memcachedclient"
)

func TestNoProgramArgumentsReturnsError(t *testing.T) {
	var args []string
	_, err := getSocketAddress(args)

	if err == nil {
		t.Fatal("Should return error when no arguments are supplied")
	}
}

func TestTooManyProgramArgumentsReturnsError(t *testing.T) {
	args := []string{"memcache.aws.com", "11211"}
	_, err := getSocketAddress(args)

	if err == nil {
		t.Fatal("Should return error when too many arguments are supplied")
	}
}

func TestNoPortUsesDefaultPort(t *testing.T) {
	host := "memcached.aws.com"
	args := []string{host}
	s, err := getSocketAddress(args)

	if err != nil {
		t.Fatal("Should not return error when supplied with default port")
	}

	if s.Host != host {
		t.Fatal("Should set the hostname correctly")
	}

	if s.Port != memcachedclient.DEFAULT_PORT {
		t.Fatal("Should infer the default port")
	}
}

func TestHostAndPortReturnsNoError(t *testing.T) {
	host := "memcached.aws.com"
	port := 42069
	args := []string{fmt.Sprintf("%s:%d", host, port)}
	s, err := getSocketAddress(args)

	if err != nil {
		t.Fatal("Should not return error when supplied with default port")
	}

	if s.Host != host {
		t.Fatal("Should set the hostname correctly")
	}

	if s.Port != port {
		t.Fatal("Should set the correct port")
	}
}

package handler

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"sf035/proverb"
	"strings"
	"testing"
)

func TestHandle(t *testing.T) {
	message := "With great power comes great responsibility"
	want := fmt.Sprintf("%v\n", message)

	proverb, err := proverb.NewProverb(io.NopCloser(strings.NewReader(message)))
	if err != nil {
		t.Fatal(err)
	}

	connServer, connClient := net.Pipe()

	go func() {
		Handle(connServer, proverb)
		connServer.Close()
	}()

	readerClient := bufio.NewReader(connClient)
	got, err := readerClient.ReadString('\n')
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Errorf("got %#v; want %#v", got, want)
	}
}

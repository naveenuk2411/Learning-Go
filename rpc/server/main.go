package main

import (
	"fmt"
	"net"
	"net/rpc"
	shared_types "rpc/shared_types"
	"strings"
)

type WordCountServer struct {
	address string
}

func (wordCountServer *WordCountServer) Compute(args *shared_types.WordCountRequest, reply *shared_types.WordCountReply) error {
	content := args.Content
	counts := make(map[string]int)

	words := strings.Fields(content)
	for _, word := range words {
		counts[word]++
	}

	reply.Counts = counts
	return nil
}

func (wordCountServer *WordCountServer) Listen() error {
	rpc.Register(wordCountServer)
	listener, err := net.Listen("tcp", wordCountServer.address)

	if err != nil {
		return err
	}
	go func() {
		fmt.Println("Server is up at", wordCountServer.address)
		rpc.Accept(listener)
		fmt.Println("Server shutdown")
	}()

	return nil
}

func main() {
	wordCountServer := new(WordCountServer)
	wordCountServer.address = "localhost:5001"

	err := wordCountServer.Listen()

	if err != nil {
		fmt.Printf("Unable to spin up the word count server %s", err.Error())
	}
}

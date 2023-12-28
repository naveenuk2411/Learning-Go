package main

import (
	"fmt"
	"net/rpc"
	shared_types "rpc/shared_types"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:5001")

	if err != nil {
		fmt.Printf("Unable to spin up a client %s", err.Error())
	}

	args := new(shared_types.WordCountRequest)
	args.Content = " Hello there, this is a good day. But what matters is how good you are feeling! LOL just a random dumb quote!"

	reply := new(shared_types.WordCountReply)

	done := client.Go("WordCountServer.Compute", args, reply, nil)
	<-done.Done
	fmt.Println("Word count received", reply.Counts)
}

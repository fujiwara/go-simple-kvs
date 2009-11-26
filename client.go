package main

import (
	"kvs";
	"log";
	"rpc";
	"os";
	"fmt";
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1975");
	if err != nil {
		log.Exit("dialing:", err);
	}
	reply := new(kvs.Reply);
	args  := new(kvs.Args);
	args.Key = os.Args[1];

	var command string;
	switch len(os.Args) {
	case 2: command = "Server.Get";
	case 3: command = "Server.Set"; args.Value = os.Args[2];
	}
	err = client.Call(command, args, reply);
	if err != nil {
		log.Exit("kvs error:", err);
	}
	fmt.Printf("%s(%s) = %s\n", command, args.Key, reply.Value);
}

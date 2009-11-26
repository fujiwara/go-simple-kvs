package main

import (
	"kvs";
	"log";
	"os";
	"fmt";
)

func main() {
	client, err := kvs.NewClient("localhost:1975");

	var command string;
	var value string;
	switch len(os.Args) {
	case 2: command = "get"; value, err = client.Get(os.Args[1])
	case 3: command = "set"; value, err = client.Set(os.Args[1], os.Args[2])
	}
	if err != nil {
		log.Exit("kvs error:", err);
	}
	fmt.Printf("%s(%s) = %s\n", command, os.Args[1], value);
}


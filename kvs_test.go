package kvs

import (
	"testing";
	"fmt";
	"time";
	"net";
	"log";
)
var tested int = 1;

func empty_port () int {
	for port := 10000; port < 20000; port++ {
		addr   := fmt.Sprintf("localhost:%d", port);
		l, err := net.Listen("tcp", addr);
		if (err == nil) {
			defer l.Close();
			return port;
		}
	}
	panic("can't listen empty port");
}

func ok(t *testing.T, ok bool, name string) {
	if !ok {
		t.Errorf("not ok %d - %s", tested, name);
	} else {
		t.Logf("ok %d - %s", tested, name);
	}
	tested++;
}

func prepare (key string, value string) (*Args, *Reply) {
	args := new(Args);
	args.Key   = key;
	args.Value = value;
	return args, new(Reply);
}

func TestServerStandalone (t *testing.T) {
	server := new(Server).Init();
	var args  *Args;
	var reply *Reply;

	args, reply = prepare("foo", "bar");
	server.Set(args, reply);
	ok(t, reply.Value == "ok", "set foo => bar");

	args, reply = prepare("foo", "");
	server.Get(args, reply);
	ok(t, reply.Value == "bar", "get foo == bar");

	args, reply = prepare("bar", "");
	server.Get(args, reply);
	ok(t, reply.Value == "", "get bar == \"\"");

	args, reply = prepare("bar", "baz");
	server.Set(args, reply);
	ok(t, reply.Value == "ok", "set bar => baz");

	args, reply = prepare("bar", "");
	server.Get(args, reply);
	ok(t, reply.Value == "baz", "get bar == baz");

	args, reply = prepare("bar", "BAZ");
	server.Set(args, reply);
	ok(t, reply.Value == "ok", "set bar => BAZ");

	args, reply = prepare("bar", "");
	server.Get(args, reply);
	ok(t, reply.Value == "BAZ", "get bar == BAZ");
}

func TestServerClient (t *testing.T) {
	port := empty_port();
	ok(t, port != 0, "empty_port");
	log.Stderrf("got empty port: %d", port);

	addr := fmt.Sprintf("localhost:%d", port);
	go RunServer(addr);

	time.Sleep(int64(1) * 1e8); // wait for Server started

	client, err := NewClient(addr);
	ok(t, err == nil, "NewClient()");

	var value string;
	value, err = client.Get("xxx");
	ok(t, err == nil, "");
	ok(t, value == "", "client.Get(xxx) == \"\"");

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key_%d", i);
		val := fmt.Sprintf("val_%d", i);

		_, err = client.Set(key, val);
		ok(t, err == nil, fmt.Sprintf("client.Set(%s,%s)", key, val));

		value, err = client.Get(key);
		ok(t, err == nil, "");
		ok(t, value == val, fmt.Sprintf("client.Get(%s)", key));
	}
}


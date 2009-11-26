package kvs

import (
	"testing";
)
var tested int = 1;

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
	ok(t, reply.Value == "", "get bar == ''");

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
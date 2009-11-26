package kvs

import(
	"rpc";
	"http";
	"log";
	"net";
	"os";
)

type Server struct {
	storage map[string] string;
}

type Args struct {
	Key   string;
	Value string;
}

type Reply struct {
	Value string;
}

func RunServer(addr string) {
	self := new(Server).Init();
	rpc.Register(self);
	rpc.HandleHTTP();
	l, e := net.Listen("tcp", addr);
	if e != nil {
		log.Exit("listen error: ", e);
	}
	http.Serve(l, nil);
}

func (self *Server) Init() *Server {
	self.storage = make(map[string] string);
	return self;
}

func (self *Server) Get(args *Args, reply *Reply) os.Error {
	if value, ok := self.storage[args.Key]; ok {
		reply.Value = value;
	}
	return nil;
}

func (self *Server) Set(args *Args, reply *Reply) os.Error {
	self.storage[args.Key] = args.Value;
	reply.Value = "ok";
	return nil;
}

type Client struct {
	client *rpc.Client;
}

func NewClient(addr string) (*Client, os.Error) {
	self := new(Client);
	client, err := rpc.DialHTTP("tcp", addr);
	self.client = client;
	return self, err;
}

func (c *Client) Get(key string) (string, os.Error) {
	reply := new(Reply);
	args  := new(Args);
	args.Key = key;
	err := c.client.Call("Server.Get", args, reply);
	return reply.Value, err;
}

func (c *Client) Set(key string, value string) (string, os.Error) {
	reply := new(Reply);
	args  := new(Args);
	args.Key   = key;
	args.Value = value;
	err := c.client.Call("Server.Set", args, reply);
	return reply.Value, err;
}

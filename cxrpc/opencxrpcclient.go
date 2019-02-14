package cxrpc

import (
	"fmt"
	"net/rpc"
)

// OpencxRPCClient is a RPC client for the opencx server
type OpencxRPCClient struct {
	Conn *rpc.Client
}

// Call calls the servicemethod with name stirng, args args, and reply reply
func (cl *OpencxRPCClient) Call(serviceMethod string, args interface{}, reply interface{}) error {
	return cl.Conn.Call(serviceMethod, args, reply)
}

// SetupConnection creates a new RPC client
func (cl *OpencxRPCClient) SetupConnection(server string, port int) error {
	var err error

	cl.Conn, err = rpc.Dial("tcp", server+":"+fmt.Sprintf("%d", port))
	if err != nil {
		return err
	}

	// logging.Infof("Connected to exchange at %s\n", server+":"+fmt.Sprintf("%d", port))
	return nil
}

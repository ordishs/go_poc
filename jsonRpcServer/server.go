package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Handler comment
type Handler struct{}

// Request comment
type Request interface{}

// Response comment
type Response struct {
	Message interface{} `json:"message"`
}

// Hello is a comment
func (h *Handler) Hello(req Request, res *Response) (err error) {
	// fmt.Printf("%+v\n", req)

	// req is a map with keys of string and values of interface{}
	a := req
	b := a.(map[string]interface{}) // Cast to map[string]interface{}
	c := b["firstName"]             // Get the value for the string key of "firstName" - returns an interface{}
	d := c.(string)                 // Cast to a string
	fmt.Printf("a is %T\nb is %T\nb is %T\nd is %T\n", a, b, c, d)

	firstName := req.(map[string]interface{})["firstName"].(string)
	lastName := req.(map[string]interface{})["lastName"].(string)

	fmt.Printf("%+v - %+v\n", firstName, lastName)

	if firstName == "" || lastName == "" {
		err = errors.New("A name must be specified")
		return err
	}

	res.Message = "Hello " + firstName + " " + lastName
	return nil
}

func startServer() {
	handler := new(Handler)

	server := rpc.NewServer()
	server.Register(handler)

	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

func main() {
	startServer()
}

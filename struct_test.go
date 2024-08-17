package flag

import (
	"fmt"
	"testing"
)

// name, shorthand string, value string, usage string
type User struct {
	Name  string `json:"name"`
	Age   int    `name:"age" short:"a" value:"1233"`
	Hello string `name:"hello" short:"h" value:"hee" usage:"usage"`
}

func TestBind(t *testing.T) {
	u := &User{}
	Scan(u)
	_ = CommandLine.Parse([]string{"--age=13"})
	fmt.Println(">>>", u.Hello, u.Age)
}

package flag

import (
	"fmt"
	"testing"
)

// name, shorthand string, value string, usage string
type User struct {
	Name  string `json:"name"`
	Age   int    `short:"a" value:"1233"`
	Hello string `name:"hello" short:"h" value:"hee" usage:"usage"`
	Sex   bool   `name:"sex" short:"s" def:"false"`
	Sex2  bool   `name:"sex2" short:"x"`
}

func TestBind(t *testing.T) {
	u := &User{}
	Scan(u)
	var b bool
	BoolVarP(&b, "bb", "b", false, "ddd")
	// _ = CommandLine.Parse([]string{"--age=13 -s"})
	_ = CommandLine.Parse([]string{"--sex=true", "-x", "-b"})

	sex, _ := CommandLine.GetBool("sex")
	fmt.Println(">>>", u.Hello, u.Age, u.Sex, ">>", b, "sex get", sex)

}

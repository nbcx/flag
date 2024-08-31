package flag

import (
	"net"
	"reflect"
	"testing"
)

// name, shorthand string, value string, usage string
type Target struct {
	Name  string   `name:"name"`
	Names []string `name:"names"`
	Yes   bool     `name:"yes" short:"y"`
	Ok    byte     `name:"ok" short:"o" value:"1" def:"2"`
	Sexes []bool   `name:"sexes" short:"s"`
	Count int      `name:"count" short:"c" value:"1" def:"2" type:"count"`
	Ip    net.IP   `name:"ip" short:"i" value:"127.0.0.1" def:"0.0.0.0"`
}

func TestScan(t *testing.T) {
	target := &Target{}
	tests := []struct {
		name  string
		input []string
		v1    any
		v2    any
	}{
		{"string", []string{"--name=hello"}, &target.Name, "hello"},
		{"bool", []string{"--yes=true"}, &target.Yes, true},
		{"byte", []string{"-o"}, &target.Ok, uint8(2)},
		{"count", []string{"-c"}, &target.Count, int(2)},
		{"count", []string{"-c=3"}, &target.Count, int(3)},
		// {"count", []string{"--ip=192.168.0.1"}, &target.Count, net.IP("192.168.0.1")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := NewFlagSet("Example", ContinueOnError)
			fs.Scan(target)
			err := fs.Parse(tt.input)
			if err != nil {
				t.Errorf("parse err: %v", err)
				return
			}
			v1 := reflect.ValueOf(tt.v1)
			v2 := reflect.ValueOf(tt.v2)
			if v1.Kind() == reflect.Pointer {
				v1 = v1.Elem()
			}
			if v2.Kind() == reflect.Pointer {
				v2 = v2.Elem()
			}
			if !v1.Equal(v2) {
				t.Errorf("v1 != v2: %v", tt.v2)
				return
			}
		})
	}
}

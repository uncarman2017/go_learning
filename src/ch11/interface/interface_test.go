package interface_test

import (
	"fmt"
	"testing"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

type Architect interface {
	DesignDB(req string) bool
}

type ArchitectImpl struct {

}


func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func (a *ArchitectImpl) DesignDB(req string) bool {
	fmt.Printf("test %s\n",req)
	return true
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())

	var a Architect
	a = new(ArchitectImpl)
	a.DesignDB("数据库设计要求")
}

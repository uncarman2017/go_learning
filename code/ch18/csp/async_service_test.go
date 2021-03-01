package concurrency

import (
	"fmt"
	"testing"
	"time"
)

type Person struct {
	SName string
	Age   int
}

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	var (
		p1 *Person
	 	p2 Person
	)
	fmt.Println(service())
	otherTask()
	p1 = new(Person)
	p1.Age = 1
	p2 = Person{}
	p2.Age = 2
}

/*
此函数返回一个channel
*/
func AsyncService() chan string {
	retCh := make(chan string, 1)
	//retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret //将处理结果放到channel中
		fmt.Println("service exited.")
	}()
	return retCh
}

//
func TestAsynService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh) //从channel中取得处理结果
	time.Sleep(time.Second * 1)
}

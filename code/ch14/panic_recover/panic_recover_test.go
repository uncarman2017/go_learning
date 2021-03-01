package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {

	defer func() {
		//panic的返回值，通过recover函数来获取;recover函数只有在defer代码块中才会有效果
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("Start")
	//panic存在的意义，不仅可以控制异常处理流程，还可以用来返回异常原因。
	panic(errors.New("Something wrong!")) //调用panic后，调用方函数执行从当前调用点退出

	//os.Exit(-1)
	// fmt.Println("End")
}

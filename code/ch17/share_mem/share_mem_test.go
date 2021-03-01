package share_mem

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {

	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)

}

func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex   // 协锁
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)  // 等待所有协程完成任务后继续执行后续代码
	t.Logf("counter = %d", counter)

}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()  //通知等待组当前协程已结束任务
		}()
	}
	wg.Wait() //等待的协程任务全部完成后，后续代码继续执行
	t.Logf("counter = %d", counter)

}

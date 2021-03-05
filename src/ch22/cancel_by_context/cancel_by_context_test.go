package cancel

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done(): // 取消通知发到Channel
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	// context.Background()创建根节点上下文对象并返回
	// context.WithCancel创建子节点上下文对象并返回,方法参数中传递上级Context对象.
	// 当某个Context被取消后，子Context都会被取消
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}

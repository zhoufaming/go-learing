package awesomeProjecr

import (
	"fmt"
	"testing"
)


func TestHuman(t *testing.T) {
	fmt.Println("are you ok")
	t.Log("I am  ok")
}

func Benchmark_Add(b *testing.B) {
	// 重置计时器
	b.ResetTimer()
	// 停止计时器
	b.StopTimer()
	// 开始计时器
	var n int
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n++
		fmt.Sprintf("%d", n)
	}

}

func TestMa(t *testing.T) {
	fmt.Println("hello world")
}

func TestFailNow(t *testing.T) {
	fmt.Println("before fail")
	t.FailNow()
	fmt.Println("after fail")
}

func TestFail(t *testing.T) {
	fmt.Println("before fail")
	t.Fail()
	fmt.Println("after fail")
}


package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
    a, b := 0, 1
    n := 0
    return func() int {
        if n++; n==1 {
            return 0
        }
        a, b = b, a + b
        return a
    }
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}


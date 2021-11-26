package lib

// 斐波那契数列
// 求出第n个数的值
//Fibonacci(0) = 0
//Fibonacci(1) = 1
//Fibonacci(2) = Fib(1) - Fib(0) = 1- 0 =1

func Fibonacci(n int64) int64 {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

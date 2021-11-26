package lib

import "testing"

func TestFibonacci(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{0}, want: 0},
		{name: "test2", args: args{1}, want: 1},
		{name: "test3", args: args{2}, want: 1},
		{"test4", args{3}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.args.n); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(int64(i))
	}
}

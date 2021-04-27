package diffsq

func SquareSum(n int) int {
	nums := []int{}
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	return square(sum(nums))
}

func SumSquare(n int) int {
	sqs := []int{}
	for i := 1; i <= n; i++ {
		sqs = append(sqs, square(i))
	}
	return sum(sqs)
}

func DiffSquare(n int) int {
	return SquareSum(n) - SumSquare(n)
}

func square(n int) int {
	return n * n
}

func sum(ns []int) int {
	s := 0
	for _, n := range ns {
		s += n
	}
	return s
}

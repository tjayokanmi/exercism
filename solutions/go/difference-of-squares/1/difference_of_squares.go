package differenceofsquares


func SquareOfSum(n int) int {
	var num int
	for i := 1; i <= n; i++ {
		num += i
	}
	return num * num

	// panic("Please implement the SquareOfSum function")
}

func SumOfSquares(n int) int {
	//panic("Please implement the SumOfSquares function")
	var result int
	for i := 1; i <= n; i++ {
		result += i * i
	}
	return result
}

func Difference(n int) int {
    return SquareOfSum(n) - SumOfSquares(n)
	panic("Please implement the Difference function")
}

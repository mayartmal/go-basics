package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	doubled := transformNums(&nums, double)
	tripled := transformNums(&nums, triple)
	doubledAnonomous := transformNums(&nums, func(number int) int { return number * 2 })
	transformer := createTransformer(4)
	squred := transformNums(&nums, transformer)

	fmt.Println(nums)
	fmt.Println(doubled)
	fmt.Println(tripled)
	fmt.Println(doubledAnonomous)
	fmt.Println(squred)
	fmt.Println()
	fmt.Println()

	res := factorial(1)
	fmt.Println(res)

	fmt.Println()
	nnnn := []int{1, 5, 4, 11}
	fmt.Println(sumup(1, 2, 7))
	fmt.Println(sumup(1, 2, 7))
	fmt.Println(sumup(1, nnnn...))

}

func transformNums(numbers *[]int, transform func(int) int) []int {
	tNumbrs := []int{}
	for _, val := range *numbers {
		tNumbrs = append(tNumbrs, transform(val))
	}
	return tNumbrs
}

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

// factorial 5 => 5*4*3*2*1 => 120
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return (num * factorial(num-1))
}

//variatic func
func sumup(startNum int, nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}

	return sum
}

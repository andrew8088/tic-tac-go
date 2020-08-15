package main

type predicate func(int) bool

func GetIsNum(num1 int) func(int) bool {
	return func(num2 int) bool {
		return num1 == num2
	}
}

func Any(arr []int, predicate predicate) bool {
	for _, m := range arr {
		if predicate(m) {
			return true
		}
	}
	return false
}

func All(arr []int, predicate predicate) bool {
	for _, m := range arr {
		if !predicate(m) {
			return false
		}
	}
	return true
}

func Contains(arr []int, item int) bool {
	return Any(arr, GetIsNum(item))
}

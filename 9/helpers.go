package main

func derivative(numbers []int) []int {
    d := make([]int, len(numbers)-1)
    for i, n := range(numbers[1:]) {
        d[i] = n - numbers[i]
    }
    return d
}

func allZeros(numbers []int) bool {
    for _, n := range(numbers) {
        if n != 0 {
            return false
        }
    }
    return true
}

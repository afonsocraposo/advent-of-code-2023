package common

func Histogram(cards string) map[rune]int {
	hist := make(map[rune]int)
	for _, card := range cards {
		if _, exists := hist[card]; !exists {
			hist[card] = 0
		}
		hist[card]++
	}
	return hist
}

package maps

import "fmt"

func Run() {
	colors := make(map[string]string)
	colors["white"] = "whitewhitewhite"
	colors["yellow"] = "yellowyellowyellow"
	printMap(colors)
}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

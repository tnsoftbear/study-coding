package main

import "fmt"

// const INPUT = "Ayowe awxewr nwaalfw die tiy rgw fklf ua xgixiklrw! Tiy lew qwkxinw."
const INPUT = "SYNTPrfneVfPbbyOhgAbgFrpher"

func main() {
	// var results []string
	for i := 0; i < 256; i++ {
		var result []byte
		for _, ch := range INPUT {
			result = append(result, byte(ch)+byte(i))
		}
		fmt.Printf("%d] %s\n", i, result)
	}
}
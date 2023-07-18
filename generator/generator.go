package generator

import "fmt"

func generateResults() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			for key, val := range Operations {
				res := val(i, j)
				fmt.Printf("\"%v%v%v\": \"%v\",\n", i, key, j, res)

				if res > 0 {
					ir := ToRoman(i)
					jr := ToRoman(j)
					resr := ToRoman(res)
					fmt.Printf("\"%v%v%v\": \"%v\",\n", ir, key, jr, resr)
				}
			}
		}
	}
}

var Operations = map[string]func(a int, b int) int{
	"+": func(a int, b int) int {
		return a + b
	},
	"/": func(a int, b int) int {
		return a / b
	},
	"*": func(a int, b int) int {
		return a * b
	},
	"-": func(a int, b int) int {
		return a - b
	},
}

var ones = [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
var tens = [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
var hunds = [2]string{"", "C"}

func ToRoman(num int) string {
	h := num / 100 % 10
	t := num / 10 % 10
	o := num % 10

	return hunds[h] + tens[t] + ones[o]
}

package week01

import "fmt"

func RangeSlice() {
	me := []string{"I", "am", "stupid", "and", "weak"}

	for index, value := range me {
		switch value {
		case "stupid":
			me[index] = "smart"
		case "weak":
			me[index] = "strong"
		}
	}

	fmt.Println(me)
}

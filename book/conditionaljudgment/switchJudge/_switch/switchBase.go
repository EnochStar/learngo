package _switch

import "fmt"

func SwitchBase() {
	var grade string = "B"
	var marks int = 1
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 70:
		grade = "C"
	default:
		grade = "E"
	}
	switch {
	case grade == "A":
		fmt.Println("优秀\n")
	case grade == "B":
		fmt.Println("良好\n")
	case grade == "C":
		fmt.Println("及格\n")
	default:
		fmt.Println("差\n")
	}
	fmt.Printf("你的等级为：%s", grade)
}

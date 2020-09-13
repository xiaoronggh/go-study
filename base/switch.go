package main

import "log"
/*
fallthrough 会强制执行下一个case，不管下一个case的条件是否是true
type switch 判断变量的类型
*/

func main() {
	var grade string = "A"
	var marks int = 90
	switch marks {
	case  90: grade = "A"
	case  80: grade = "B"
	case 70: grade = "C"
	default:
		grade = "D"
	}
	log.Printf("grade is %s", grade)

	var x interface{}
	x = 1
	switch i:= x.(type) {
	case int :
		log.Printf("x type is int, %T", i)
	default:
		log.Println("x type is unknow")
	}

	switch grade {
	case "A":
		log.Println("grade A")
		fallthrough
	case "B":
		log.Println("grade B")
	case "C":
		log.Println("grade C")
	}

}

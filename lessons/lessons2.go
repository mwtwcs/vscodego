package main

import "fmt"

func switcha() {
	a := 0
	for a < 1000{
		i := isTest(a)
		switch i {
		case 1:
			fmt.Println("a = 2")
		case 2:
			fmt.Println("a = 3")
		default:
			fmt.Println(fmt.Sprintf("unknown a. a=%d", a))
		}
		a++
	}
}

func isTest(a int) int {
	if a == 2{
		return 1
	} else if a == 3 {
		return 2 
	}
	return 3
}


//package main

//import "fmt"

//func main() {
//	a := 0
//	for a < 1000{
	//	switch i := isTest(a); i {
	//	case 1:
	//		fmt.Println("a = 2")
	//	case 2:
		//	fmt.Println("a = 3")
		//default:
	//		fmt.Println(fmt.Sprintf("unknown a. a=%d", a))
	//	}
	//	a++
//	}
//}

//func isTest(a int) int {
//	if a == 2{
//		return 1
//	} else if a == 3 {
//		return 2 
//	}
//	return 3
//}
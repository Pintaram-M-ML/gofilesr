// package main
// import "fmt"
// func main()  {
// 	k:= 2
// 	switch k {
// 	case 1:
// 		fmt.Print(k);
// 		fallthrough;
// 	case 2:
// 		fmt.Print(k);
// 		fallthrough;
// 	case 3:
// 		fmt.Print("k is smaller\n");
		
// 	default:
// 		fmt.Print("does not match");		
// 	}
	
// }

//for code
// package main
// import "fmt"
// func main()  {
// 	var sum int = 0;
// 	for sum <20{
// 		sum+=3
// 		fmt.Println(sum)
// 	}
// 	fmt.Println(sum)
// }
// 
package main

import "fmt"

func main()  {
	a , b := 10, 30;
	fmt.Println(a , b)
	// Example with sum of multiple numbers
	sum , sum1:=sumof(10, 20, 40, 50)
	fmt.Println(sum, sum1)
}

func sumof(args ...int) (int, int)  {
	sum := 0
	for _, value := range args {
		sum += value
	}
	return 10, sum // First return 10 as an example, then the sum of the arguments
}

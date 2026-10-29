package main

import "fmt"


func main(){
	var x,y int;
		fmt.Scanln(&x,&y);
		fmt.Printf("Values entered are %d and %d\n",x,y);


		result1 := x>0 && y>0
		fmt.Println(result1);

		result2 := x>10 || y>10
		fmt.Println(result2);

		result3 := x != y
		fmt.Println(result3);
}

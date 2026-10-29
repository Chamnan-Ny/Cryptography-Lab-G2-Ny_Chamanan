package week2
import "fmt"

func main()

	var a,b int;
	fmt.Println("Enter value of A & B : ");
	// fmt.Scan(&a,&b);
	fmt.Scanf("%d\n%d\n",&a,&b)
	fmt.Printf("Values entered are %d and %d\n",a,b);

	x := a
	fmt.Println("= :",x);

	x += b
	fmt.Println("+= :",x);

	x -= b
	fmt.Println("-= :",x);
	
	x *= b
	fmt.Println("*= :",x);

	x /= b
	fmt.Println("/= :",x);

	x %= b
	fmt.Println("%= :",x);
package Week_03

import "fmt"

func myPow(x float64, n int) float64 {
	if n >= 0{
		return _myPow(x, n)
	}else {
		return 1 / _myPow(x, -n)
	}
}

func _myPow(x float64, n int) float64  {
	if n == 0{
		return 1
	}
	if n == 1 {
		return x
	}
	y := myPow(x, n/2)
	if n % 2 == 1 {
		return y * y * x
	}else{
		return y * y
	}
}

func main()  {
	fmt.Println(myPow(0.00001, 2147483647))
}
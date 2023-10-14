package calculator

func Square(a int)int{
	return a*a
}

func Sum(a,b int)int{
	return a+b
}

func Difference(a,b int)int{
	return a-b
}

func Product(a,b int)int{
	return a*b
}

func QuoRem(a,b int)(int,int){
	q:=a/b
	r:=a%b
	return q,r
}
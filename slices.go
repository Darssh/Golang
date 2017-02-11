package main 

import "fmt"

func main(){

	s := make([]string, 3)
	fmt.Println("emp: ",s)
	
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[1])

	s = append(s, "d")
	s = append(s, "d", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println("copy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)
	


}

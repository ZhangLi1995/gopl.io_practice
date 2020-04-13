package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	b := a
	reverse(a[:])
	fmt.Println(a)
	leftMove(b[:], 2)
	fmt.Println(b)

	fmt.Println("练习4.4：")
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(rotate(s, 2))

	fmt.Println("练习4.5：")
	s1 := []string{"1", "2", "2", "2", "3", "3", "1", "1", "4"}
	s1 = dupSlice(s1)
	fmt.Println(s1)

	fmt.Println("练习4.6：")
	s2 := []byte("   北京\t欢 迎\n你")
	s2 = dupSlice2(s2)
	fmt.Printf("%s\n", s2)

	fmt.Println("练习4.7：")
	s3 := []byte("北京欢迎你welcome")
	fmt.Printf("%s\n", reverse3(s3))
}

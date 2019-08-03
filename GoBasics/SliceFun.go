package main

import "fmt"

func Selection() {

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

}

//Slice does not store data
func SliceIsReferenceOfArray() {
	names := [4]string{"james", "andrewlam", "alexwang", "kassiazhang"}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "HACKEDXXX"
	fmt.Println(a, b)
	fmt.Println(names)

}

//Slice Literals
func SliceIsArrayWithoutLength() {

	intSlice := []int{1, 3, 6, 7, 9, 15}
	fmt.Println(intSlice)

	//Slice Literals
	boolSlice := []bool{true, false, true, true}
	fmt.Println(boolSlice)

	//Slice Literals
	structureSlice := []struct {
		id   int
		name string
	}{
		{10, "abc"}, {11, "bcd"},
	}
	fmt.Println(structureSlice)
}

//Nil Slice = 0len 0cap 0array
func ZeroSlice() {
	var s []int
	fmt.Println(s, len(s), cap(s))
}

//Slice can contain any other type

func SliceOfSlice() {

}

//Creating a slice with make

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func createWithMake() {

	//a := make([]int, 5)

	//b := make([]int, 0, 5)

	//c := b[:2]

	//d := c[:5]

}

func appendToSlice() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)

}

func mainga21() {

	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	fmt.Println(pow)

}

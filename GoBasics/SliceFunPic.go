package main

import "fmt"

func Pic(dx, dy int) [][]uint8 {

	pic := make([][]uint8, 100, 100)

	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			pic[i] = append(pic[i], uint8(i+j)/2)
			fmt.Println(i, j)
		}
	}

	return pic
}

func main() {

	fmt.Println(Pic(10, 20))

}

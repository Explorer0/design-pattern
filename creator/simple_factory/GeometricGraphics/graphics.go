package GeometricGraphics

import "fmt"


type Circle struct {}

func (c *Circle) Draw() {
	fmt.Printf("Circle drawing...")
}

func (c *Circle) Erase() {
	fmt.Printf("Circle erasing...")
}


type Triangle struct {}

func (t *Triangle) Draw() {
	fmt.Printf("Triangle drawing...")
}

func (t *Triangle) Erase() {
	fmt.Printf("Triangle erasing...")
}


type Square struct {}

func (s *Square) Draw() {
	fmt.Printf("Square drawing...")
}

func (s *Square) Erase() {
	fmt.Printf("Square erasing...")
}





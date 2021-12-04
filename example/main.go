package main

import (
	"fmt"

	"com.xdteh/hello/example/sharp"
)

func main() {
	r := sharp.React{Width: 65.0, Height: 63.2}
	r.Height = 12
	r.Width = 32
	fmt.Println(r.Area())
	fmt.Println(r.Perim())
}

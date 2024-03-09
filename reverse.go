package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	d := scan.Text()
	scan.Scan()
	s1 := scan.Text()
	scan.Scan()
	s2 := scan.Text()
	scan.Scan()
	s3 := scan.Text()
	fmt.Print(s1 + d + s2 + d + s3)
}

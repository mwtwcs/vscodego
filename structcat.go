package main

import "fmt"

func structcar(){
 type Car struct{
  color string
  brand string
  year int
 }

 x:= Car{"Blue","BMW",2019}
 fmt.Println(x)

}

package main

import (
	"fmt"
	"go/types"
)

func main(){
	fmt.Println("feature-vulcan is done.")
	a := types.Interface{}
	fmt.Printf("a's type is %T\n",a)
}

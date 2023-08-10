package main

import (
	"fmt"

	"k8s.io/apimachinery/pkg/types"
)

func main() {
	fmt.Println("go cli example program")
	name := types.NamespacedName{Namespace: "default", Name: "image-controller"}
	fmt.Println(name)
}

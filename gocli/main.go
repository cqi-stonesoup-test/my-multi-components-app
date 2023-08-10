package main

import (
	"fmt"

	"k8s.io/apimachinery/pkg/types"
	corev1 "k8s.io/api/core/v1"
)

func main() {
	fmt.Println("go cli example program")
	name := types.NamespacedName{Namespace: "default", Name: "image-controller"}
	fmt.Println(name)

	secret := corev1.Secret{}
	fmt.Println("Secret:", secret)
}

package main

// Aliasing fmt to fm
// This is useful in case of conflicts among pkgs
import fm "fmt"

func main() {
	fm.Println("Example for aliasing")
}

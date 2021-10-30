package main

func mul(a int, b int) int {
	return a * b
}

// go tool compile -S -N -l main.go
func main() {
	mul(3, 4)
}

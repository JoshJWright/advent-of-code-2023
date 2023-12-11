package main

import "os"

func main() {
	raw, _ := os.ReadFile("input.txt")
	input := string(raw)
}

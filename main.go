package main

import (
	"kabesma/gitpoison/internal/poison"
)

func main() {
	w := poison.Execute()
	w.StartApp()
}

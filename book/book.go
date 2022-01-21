package book

import "fmt"

func Print() {
	fmt.Println("Hello from other side")
}

type book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books = []book{
	{ID: "1", Title: "Howel's Moving Castle", Author: "Diana Wynne Jones", Price: 8.99},
	{ID: "2", Title: "Castle in the Air", Author: "Diana Wynne Jones", Price: 8.99},
	{ID: "3", Title: "House of many ways", Author: "Diana Wynne Jones", Price: 9.99},
}

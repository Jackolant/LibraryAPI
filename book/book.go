package book

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
	Shelf  string `json:"shelf"`
	Genre  string `json:"genre"`
}

var Books = []Book{
	{ID: 1, Title: "Howel's Moving Castle", Author: "Diana Wynne Jones", Genre: "Fiction", Shelf: "D", Pages: 429},
	{ID: 2, Title: "Castle in the Air", Author: "Diana Wynne Jones", Genre: "Fiction", Shelf: "D", Pages: 400},
	{ID: 3, Title: "House of many ways", Author: "Diana Wynne Jones", Genre: "Fiction", Shelf: "D", Pages: 432},
}

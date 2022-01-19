package book

type book struct {
	ID     float32 `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Genre  string  `json:"genre"`
	Shelf  string  `json:"shelf"`
	Pages  float32 `json:"pages"`
}

var books = []book{
	{ID: 1, Title: "Howel's Moving Castle", Author: "Diana Wynne Jones", Genre: "Fiction", Shelf: "C", Pages: 429},
	{ID: 2, Title: "Castle in the Air", Author: "Diana Wynne Jones", Genre: "Fiction", Shelf: "C", Pages: 400},
	{ID: 3, Title: "House of many ways", Author: "Diana Wynne Jones", Genre: "Fiction", Shelf: "C", Pages: 432},
}

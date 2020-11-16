package entity

// STRUCTS
type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        string `json:"year"`
	//Actor       *Actor `json:"Actor"`
}

/*type Actor struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}*/

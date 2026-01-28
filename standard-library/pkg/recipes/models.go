package recipes

type Ingredient struct {
	Name string `json:"name"`
}

type Recipe struct {
	Name        string       `json:"name"`
	Ingridients []Ingredient `json:"ingredients"`
}

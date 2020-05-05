package animal

type Animal struct {
	ID      *int   `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Species string `json:"species,omitempty"`
}

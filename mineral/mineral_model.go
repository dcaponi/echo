package mineral

type Mineral struct {
	ID       *int   `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Molecule string `json:"molecule,omitempty"`
}

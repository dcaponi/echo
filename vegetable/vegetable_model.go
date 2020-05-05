package vegetable

type Vegetable struct {
	ID   *int   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

package vegetable

import (
	"encoding/json"
	"log"

	"github.com/dcaponi/echo/interfaces"
)

type VegetableController struct {
	Repository interfaces.Repository
}

func NewController(r interfaces.Repository) VegetableController {
	return VegetableController{Repository: r}
}

func (c VegetableController) GetOne(id int) []byte {
	resp, err := c.Repository.GetOne(id)
	if err != nil {
		return []byte("invalid parameter id")
	}
	body, _ := json.Marshal(resp.([]Vegetable))
	return body
}

func (c VegetableController) Get() []byte {
	resp, err := c.Repository.Get()
	if err != nil {
		log.Println("ERROR", err)
		return []byte("invalid parameter id")
	}

	body, _ := json.Marshal(resp.([]Vegetable))
	return body
}

func (c VegetableController) Write(req []byte) []byte {
	var (
		vegetable Vegetable
		err       error
	)

	if err = json.Unmarshal(req, &vegetable); err != nil {
		return []byte("invalid vegetable body")
	}

	resp, err := c.Repository.Write(vegetable)
	if err != nil {
		return []byte("unable to create vegetable")
	}

	body, _ := json.Marshal(resp.([]Vegetable))
	return body
}

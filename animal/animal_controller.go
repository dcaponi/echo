package animal

import (
	"encoding/json"
	"log"

	"github.com/dcaponi/echo/interfaces"
)

type AnimalController struct {
	Repository interfaces.Repository
}

func NewController(r interfaces.Repository) AnimalController {
	return AnimalController{Repository: r}
}

func (c AnimalController) GetOne(id int) []byte {
	resp, err := c.Repository.GetOne(id)
	if err != nil {
		return []byte("invalid parameter id")
	}
	body, _ := json.Marshal(resp.([]Animal))
	return body
}

func (c AnimalController) Get() []byte {
	resp, err := c.Repository.Get()
	if err != nil {
		log.Println("ERROR", err)
		return []byte("invalid parameter id")
	}

	body, _ := json.Marshal(resp.([]Animal))
	return body
}

func (c AnimalController) Write(req []byte) []byte {
	var (
		animal Animal
		err    error
	)

	if err = json.Unmarshal(req, &animal); err != nil {
		return []byte("invalid animal body")
	}

	resp, err := c.Repository.Write(animal)
	if err != nil {
		return []byte("unable to create animal")
	}

	body, _ := json.Marshal(resp.([]Animal))
	return body
}

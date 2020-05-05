package mineral

import (
	"encoding/json"
	"log"

	"github.com/dcaponi/echo/interfaces"
)

type MineralController struct {
	Repository interfaces.Repository
}

func NewController(r interfaces.Repository) MineralController {
	return MineralController{Repository: r}
}

func (c MineralController) GetOne(id int) []byte {
	resp, err := c.Repository.GetOne(id)
	if err != nil {
		return []byte("invalid parameter id")
	}
	body, _ := json.Marshal(resp.([]Mineral))
	return body
}

func (c MineralController) Get() []byte {
	resp, err := c.Repository.Get()
	if err != nil {
		log.Println("ERROR", err)
		return []byte("invalid parameter id")
	}

	body, _ := json.Marshal(resp.([]Mineral))
	return body
}

func (c MineralController) Write(req []byte) []byte {
	var (
		mineral Mineral
		err     error
	)

	if err = json.Unmarshal(req, &mineral); err != nil {
		return []byte("invalid mineral body")
	}

	resp, err := c.Repository.Write(mineral)
	if err != nil {
		return []byte("unable to create mineral")
	}

	body, _ := json.Marshal(resp.([]Mineral))
	return body
}

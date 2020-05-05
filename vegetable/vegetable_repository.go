package vegetable

import (
	"log"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type VegetableRepository struct {
	DB *pg.DB
}

// NewDatabase initializes an Vegetables table if one doesnt exist and
// returns a handle so methods can be called against it by the handlers
func New(db *pg.DB) (VegetableRepository, error) {
	// db.DropTable((*Vegetable)(nil), &orm.DropTableOptions{
	// 	IfExists: true,
	// })
	err := db.CreateTable((*Vegetable)(nil), &orm.CreateTableOptions{
		Temp:        false,
		IfNotExists: true,
	})
	if err != nil {
		log.Println("UNABLE TO INITIALIZE ANIMALS", err)
		return VegetableRepository{}, err
	}
	return VegetableRepository{DB: db}, nil
}

// be careful returning interfaces. upstream type checks go out the window leading to possible
// run time errors if returning the wrong type

func (ds VegetableRepository) Git(container interface{}) (interface{}, error) {
	err := ds.DB.Model(&container).Select()
	if err != nil {
		log.Println("Error retrieving vegetable from DB", err)
		return container, err
	}
	return container, nil
}

func (ds VegetableRepository) Get() (interface{}, error) {
	var (
		err        error
		vegetables []Vegetable
	)
	err = ds.DB.Model(&vegetables).Select()
	if err != nil {
		log.Println("Error retrieving vegetable from DB", err)
		return vegetables, err
	}
	return vegetables, nil
}

func (ds VegetableRepository) GetOne(id int) (interface{}, error) {
	var (
		err        error
		vegetables []Vegetable
	)
	vegetable := &Vegetable{ID: &id}
	err = ds.DB.Select(vegetable)

	if err != nil {
		log.Println("Error retrieving vegetable from DB", err)
		return vegetables, err
	}
	vegetables = append(vegetables, *vegetable)

	return vegetables, nil
}

func (ds VegetableRepository) Write(ent interface{}) (interface{}, error) {
	var (
		err        error
		vegetable  Vegetable
		vegetables []Vegetable
	)
	vegetable = ent.(Vegetable)

	if _, err = ds.DB.Model(&vegetable).Returning("*").Insert(&vegetable); err != nil {
		log.Println("Error inserting vegetable to DB", err)
		return vegetables, err
	}
	vegetables = append(vegetables, vegetable)
	return vegetables, nil
}

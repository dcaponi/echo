package animal

import (
	"log"


	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type AnimalRepository struct {
	DB *pg.DB
}

// NewDatabase initializes an Animals table if one doesnt exist and
// returns a handle so methods can be called against it by the handlers
func New(db *pg.DB) (AnimalRepository, error) {
	// db.DropTable((*Animal)(nil), &orm.DropTableOptions{
	// 	IfExists: true,
	// })
	err := db.CreateTable((*Animal)(nil), &orm.CreateTableOptions{
		Temp:        false,
		IfNotExists: true,
	})
	if err != nil {
		log.Println("UNABLE TO INITIALIZE ANIMALS", err)
		return AnimalRepository{}, err
	}
	return AnimalRepository{DB: db}, nil
}

// be careful returning interfaces. upstream type checks go out the window leading to possible
// run time errors if returning the wrong type

func (ds AnimalRepository) Git(container interface{}) (interface{}, error) {
	err := ds.DB.Model(&container).Select()
	if err != nil {
		log.Println("Error retrieving animal from DB", err)
		return container, err
	}
	return container, nil
}

func (ds AnimalRepository) Get() (interface{}, error) {
	var (
		err     error
		animals []Animal
	)
	err = ds.DB.Model(&animals).Select()
	if err != nil {
		log.Println("Error retrieving animal from DB", err)
		return animals, err
	}
	return animals, nil
}

func (ds AnimalRepository) GetOne(id int) (interface{}, error) {
	var (
		err     error
		animals []Animal
	)
	animal := &Animal{ID: &id}
	err = ds.DB.Select(animal)

	if err != nil {
		log.Println("Error retrieving animal from DB", err)
		return animals, err
	}
	animals = append(animals, *animal)

	return animals, nil
}

func (ds AnimalRepository) Write(ent interface{}) (interface{}, error) {
	var (
		err     error
		animal  Animal
		animals []Animal
	)
	animal = ent.(Animal)

	if _, err = ds.DB.Model(&animal).Returning("*").Insert(&animal); err != nil {
		log.Println("Error inserting animal to DB", err)
		return animals, err
	}
	animals = append(animals, animal)
	return animals, nil
}

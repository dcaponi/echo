package mineral

import (
	"log"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type MineralRepository struct {
	DB *pg.DB
}

// NewDatabase initializes an Minerals table if one doesnt exist and
// returns a handle so methods can be called against it by the handlers
func New(db *pg.DB) (MineralRepository, error) {
	// db.DropTable((*Mineral)(nil), &orm.DropTableOptions{
	// 	IfExists: true,
	// })
	err := db.CreateTable((*Mineral)(nil), &orm.CreateTableOptions{
		Temp:        false,
		IfNotExists: true,
	})
	if err != nil {
		log.Println("UNABLE TO INITIALIZE ANIMALS", err)
		return MineralRepository{}, err
	}
	return MineralRepository{DB: db}, nil
}

// be careful returning interfaces. upstream type checks go out the window leading to possible
// run time errors if returning the wrong type

func (ds MineralRepository) Git(container interface{}) (interface{}, error) {
	err := ds.DB.Model(&container).Select()
	if err != nil {
		log.Println("Error retrieving mineral from DB", err)
		return container, err
	}
	return container, nil
}

func (ds MineralRepository) Get() (interface{}, error) {
	var (
		err      error
		minerals []Mineral
	)
	err = ds.DB.Model(&minerals).Select()
	if err != nil {
		log.Println("Error retrieving mineral from DB", err)
		return minerals, err
	}
	return minerals, nil
}

func (ds MineralRepository) GetOne(id int) (interface{}, error) {
	var (
		err      error
		minerals []Mineral
	)
	mineral := &Mineral{ID: &id}
	err = ds.DB.Select(mineral)

	if err != nil {
		log.Println("Error retrieving mineral from DB", err)
		return minerals, err
	}
	minerals = append(minerals, *mineral)

	return minerals, nil
}

func (ds MineralRepository) Write(ent interface{}) (interface{}, error) {
	var (
		err      error
		mineral  Mineral
		minerals []Mineral
	)
	mineral = ent.(Mineral)

	if _, err = ds.DB.Model(&mineral).Returning("*").Insert(&mineral); err != nil {
		log.Println("Error inserting mineral to DB", err)
		return minerals, err
	}
	minerals = append(minerals, mineral)
	return minerals, nil
}

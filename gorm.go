package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

type data struct {
	Namespace string
	Model     string
}

func GenCRUD(namespace string, model string) error {
	err := CreateDirIfNotExist(namespace)
	if err != nil {
		return err
	}

	path := "./store/" + strings.ToLower(model) + "-gen.go"

	var d data
	d.Model = model
	d.Namespace = namespace

	f, err := os.Create(path)
	if err != nil {
		log.Println("create file: ", err)
		return err
	}

	funcMap := template.FuncMap{
		"ToLower": strings.ToLower,
	}

	t := template.Must(template.New(model + "-gen").Funcs(funcMap).Parse(crudTemplate))
	err = t.Execute(f, d)

	if err != nil {
		log.Print("execute: ", err)
		return err
	}
	f.Close()

	return nil
}

var crudTemplate = `
package store

import (
	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"
)

// {{.Model}}Store under store package holds a reference to the database depencency
type {{.Model}}Store struct {
	db *gorm.DB
}

//SetDB sets the gorm db instance
func (st *{{.Model}}Store) SetDB(db *gorm.DB) {
	st.db = db
}

// Create is used to create an {{.Model}} and returning an error
// if it did not succeed.
func (st *{{.Model}}Store) Create(u *model.{{.Model}}) error {
	st.db.Create(u)
	saved := st.db.NewRecord(u)
	if saved == true {
		return errors.New("Could not create record")
	}

	return nil
}

// Read reads a single Record using the ID and returning an error
// if it did not succeed. 
func (st *{{.Model}}Store) Read(id uint) (*model.{{.Model}}, error) {
	{{.Model | ToLower}} := &model.{{.Model}}{}
	err := st.db.First({{.Model | ToLower}}, id).Error
	if err != nil {
		return nil, errors.Wrap(err, "Could not read record")
	}

	return {{.Model | ToLower}}, nil
}

// ReadAll is in charge of reading all {{.Model}}s and returning an error
// if there are no record. It returns an array of {{.Model}}s otherwise
func (st *{{.Model}}Store) ReadAll() ([]model.{{.Model}}{}, error) {
	{{.Model | ToLower}}s := []model.{{.Model}}{}
	err := st.db.Find(&{{.Model | ToLower}}).Error
	if err != nil {
		return nil, errors.Wrap(err, "error with ReadAll query")
	}

	if len({{.Model | ToLower}}s) == 0 {
		return nil, errors.New("No records found")
	}

	return {{.Model | ToLower}}s, nil
}

// Update is in charge of updating a {{.Model}} and returning an error
// if it did not succeed.
func (st *{{.Model}}Store) Update(u *model.{{.Model}}) error {
	err := st.db.Save(u).Error
	if err != nil {
		return errors.Wrap(err, "Could not create record")
	}

	return nil
}

`

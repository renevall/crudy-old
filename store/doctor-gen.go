
package store

import (
	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"
)

// DoctorStore under store package holds a reference to the database depencency
type DoctorStore struct {
	db *gorm.DB
}

//SetDB sets the gorm db instance
func (st *DoctorStore) SetDB(db *gorm.DB) {
	st.db = db
}

// Create is used to create an Doctor and returning an error
// if it did not succeed.
func (st *DoctorStore) Create(u *model.Doctor) error {
	st.db.Create(u)
	saved := st.db.NewRecord(u)
	if saved == true {
		return errors.New("Could not create record")
	}

	return nil
}

// Read reads a single Record using the ID and returning an error
// if it did not succeed. 
func (st *DoctorStore) Read(id uint) (*model.Doctor, error) {
	doctor := &model.Doctor{}
	err := st.db.First(doctor, id).Error
	if err != nil {
		return nil, errors.Wrap(err, "Could not read record")
	}

	return doctor, nil
}

// ReadAll is in charge of reading all Doctors and returning an error
// if there are no record. It returns an array of Doctors otherwise
func (st *DoctorStore) ReadAll() ([]model.Doctor{}, error) {
	doctors := []model.Doctor{}
	err := st.db.Find(&doctor).Error
	if err != nil {
		return nil, errors.Wrap(err, "error with ReadAll query")
	}

	if len(doctors) == 0 {
		return nil, errors.New("No records found")
	}

	return doctors, nil
}

// Update is in charge of updating a Doctor and returning an error
// if it did not succeed.
func (st *DoctorStore) Update(u *model.Doctor) error {
	err := st.db.Save(u).Error
	if err != nil {
		return errors.Wrap(err, "Could not create record")
	}

	return nil
}


package models

import (
	"database/sql"

	"github.com/pborman/uuid"	
)

type Roaster struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	AddressLine1   string    `json:"addressLine1"`
	AddressLine2   string    `json:"addressLine2"`
	AddressCity    string    `json:"addressCity"`
	AddressState   string    `json:"addressState"`
	AddressZip     string    `json:"addressZip"`
	AddressCountry string    `json:"addressCountry"`
}

func NewRoaster(name, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry string) *Roaster {
	return &Roaster{
		ID:             uuid.NewUUID(),
		Name:           name,
		Email:          email,
		Phone:          phone,
		AddressLine1:   addressLine1,
		AddressLine2:   addressLine2,
		AddressCity:    addressCity,
		AddressState:   addressState,
		AddressZip:     addressZip,
		AddressCountry: addressCountry,
	}
}

func RoasterFromSQL(rows *sql.Rows) ([]*Roaster, error) {
	roasters := make([]*Roaster, 0)

	for rows.Next() {
		r := &Roaster{}

		rows.Scan(&r.ID, &r.Name, &r.Email, &r.Phone, &r.AddressLine1, &r.AddressLine2, &r.AddressCity, &r.AddressState, &r.AddressZip, &r.AddressCountry)

		roasters = append(roasters, r)
	}

	return roasters, nil
}

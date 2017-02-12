package models

import (
	"database/sql"

	"github.com/pborman/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	PassHash       string    `json:"passHash"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	AddressLine1   string    `json:"addressLine1"`
	AddressLine2   string    `json:"addressLine2"`
	AddressCity    string    `json:"addressCity"`
	AddressState   string    `json:"addressState"`
	AddressZip     string    `json:"addressZip"`
	AddressCountry string    `json:"addressCountry"`
	RoasterId      uuid.UUID `json:"roasterId"`
	IsRoaster	   int       `json:"isRoaster"`
}

func NewUser(passHash, firstName, lastName, email, phone, addressLine1, addressLine2, addressCity, addressState, addressZip, addressCountry string) *User {
	return &User{
		ID:             uuid.NewUUID(),
		PassHash:		passHash,
		FirstName:      firstName,
		LastName:       lastName,
		Email:          email,
		Phone:          phone,
		AddressLine1:   addressLine1,
		AddressLine2:   addressLine2,
		AddressZip:     addressZip,
		AddressCity:    addressCity,
		AddressState:   addressState,
		AddressCountry: addressCountry,
		RoasterId:      nil,
		IsRoaster:      0,
	}
}

func UserFromSQL(rows *sql.Rows) ([]*User, error) {
	users := make([]*User, 0)

	for rows.Next() {
		u := &User{}

		rows.Scan(&u.ID, &u.PassHash, &u.FirstName, &u.LastName, &u.Email, &u.Phone, &u.AddressLine1, &u.AddressLine2,
			      &u.AddressCity, &u.AddressState, &u.AddressZip, &u.AddressCountry, &u.RoasterId, &u.IsRoaster)

		users = append(users, u)
	}

	return users, nil
}

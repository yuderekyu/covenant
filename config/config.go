package config

import (
	"encoding/json"
	"os"
)

type Root struct {
	SQL MySQL `json:"mysql"`
	TownCenter TownCenter  `json:towncenter`
	Inventory Inventory `json:inventory`
}

/*MySQL contains information for connecting to a MySQL instance */
type MySQL struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

/*TownCenter has connection information for the town center service */
type TownCenter struct {

}
/*Inveotry has connection information for the inventory service*/
type Inventory struct {

}


/*Init returns a populated Root struct from config.json */
func Init(filename string) (*Root, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(file)

	root := &Root{}
	err = decoder.Decode(root)
	if err != nil {
		return nil, err
	}
	return root, nil
}
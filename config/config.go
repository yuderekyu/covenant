package config

import (
	"encoding/json"
	"os"
)

type Root struct {
	SQL MySQL `json:"mysql"`
	TownCenter TownCenter  `json:"towncenter"`
	Warehouse Warehouse `json:"warehouse"`
	Covenant Covenant `json:"covenant"`
	Statsd Statsd `json:"statsd"`
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
	Host string `json"host"`
}
/*Warehouse has connection information for the inventory service*/
type Warehouse struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
/*Covenant has connection information for the covenant service*/
type Covenant struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
/*Statsd contains connection information for graphite stats*/
type Statsd struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Prefix string `json:"perfix"`
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
	// fmt.Println("Root config file: " + root)
	if err != nil {
		return nil, err
	}
	return root, nil

}
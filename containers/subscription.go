package containers

import(
	"database/sql"

	"github.com/pborman/uuid"
)

type Subscription struct {
	Id uuid.UUID `json: "id"`
	UserId int `json: "user_id"`
	Status string `json:"status"`
	CreatedAt string `json:"created_at"` //change to time.Date
	StartAt string `json:"start_at"` //change to time.Date
	ShopId int `json: "shop_id"`
	OzInBag int `json: "oz_in_bag"`
	BeanName string `json:"bean_name"`
	RoastName string `json: "roast_name"`
	Price int `json: "price"`
}

func FromSql(rows *sql.Rows) ([]*Subscription, error) {
	subscription := make([]*Subscription,0)

	for rows.Next() {
		s := &Subscription{}
		rows.Scan(&s.Id, &s.UserId, &s.Status, &s.CreatedAt, &s.StartAt, &s.ShopId, &s.OzInBag, &s.BeanName, &s.RoastName, &s.Price)
		subscription = append(subscription, s)
	}

	return subscription, nil
}
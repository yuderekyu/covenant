package containers

import(
	"database/sql"

	"github.com/pborman/uuid"
)

type Subscription struct {
	Id uuid.UUID `json: "id"`
	OrderId int `json: "order_id"`
	Type string `json:"type"` 
	UserId int `json: "user_id"`
	Status string `json:"status"`
	CreatedAt string `json:"created_at"` //change to time.Date
	StartAt string `json:"start_at"` //change to time.Date
	TotalPrice string `json:"total_price"`
}

func FromSql(rows *sql.Rows) ([]*Subscription, error) {
	subscription := make([]*Subscription,0)

	for rows.Next() {
		s := &Subscription{}
		rows.Scan(&s.Id, &s.OrderId, &s.Type, &s.UserId, &s.Status, &s.CreatedAt, &s.StartAt, &s.TotalPrice)
		subscription = append(subscription, s)
	}

	return subscription, nil
}
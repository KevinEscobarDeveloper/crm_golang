package model

import "time"

type Orchard struct {
	ObjectID  string    `bson:"_id,omitempty"`
	Name      string    `bson:"name"`
	Location  string    `bson:"location"`
	Owner     string    `bson:"owner"`
	Active    bool      `bson:"active"`
	CreatedAt time.Time `bson:"created_at"`
}

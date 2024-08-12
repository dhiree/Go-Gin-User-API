package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// bson use to read and write the data in binary format and it is fast

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
	Phone    string             `bson:"phone" json:"phone"`
	//Address  Address
}

type Address struct {
	City  string `bson:"city" json:"city"`
	State string `bson:"state" json:"state"`
}

// {
//     "username": "user1",
//     "email": "user1@example.com",
//     "password": "securepassword",
//     "phone": "1234567890",
//     "address": {
//         "city": "Example City",
//         "state": "Example State"
//     }
// }

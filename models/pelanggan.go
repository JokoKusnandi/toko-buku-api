package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pelanggan struct {
	ID      primitive.ObjectID `json:"id_pelanggan"	bson:"_id,omitempty"`
	Nama    string             `json:"nama"	validate:"required"`
	Alamat  string             `json:"alamat"`
	Telepon string             `json:"telepon"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Penerbit struct {
	PenerbitID   primitive.ObjectID `json:"id_penerbit"	bson:"_id,omitempty"`
	NamaPenerbit string             `json:"nama_penerbit"	validate:"required"`
	alamat       string             `json:"alamat"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Buku struct {
	ID         primitive.ObjectID `json:"id_buku",omitempty`
	Judul      string             `json:"judul"`
	Pengarang  string             `json:"pengarang"`
	Harga      float64            `json:"harga"`
	Stok       int                `json:"stok"`
	PenerbitID primitive.ObjectID `json:"id_penerbit" validate:"required"`
}

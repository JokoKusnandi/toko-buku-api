package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Detail struct {
	BukuID string  `json:"id_buku"`
	Jumlah int     `json:"jumlah"`
	Harga  float64 `json:"harga"`
}

type Transaksi struct {
	ID          primitive.ObjectID `json:"id_transaksi"	bson:"_id",omitempty`
	Tanggal     string             `json:"tanggal"`
	PelangganID string             `json:"id_pelanggan"`
	Details     []Detail           `json:"detail"`
}

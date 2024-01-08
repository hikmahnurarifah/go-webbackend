package models

type Pasien struct {
	ID           uint   `json:"id"`
	Nama         string `json:"nama"`
	Usia         string `json:"usia"`
	JenisKelamin string `json:"jenis_kelamin"`
	Alamat       string `json:"alamat"`
	Deskripsi    string `json:"deskripsi"`
}

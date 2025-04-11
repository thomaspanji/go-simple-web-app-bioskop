package repository

import (
	"strings"

	"github.com/thomaspanji/go-simple-web-app-bioskop/database"
	"github.com/thomaspanji/go-simple-web-app-bioskop/models"
)

func InsertBioskop(b models.Bioskop) (int, error) {
	var id int
	err := database.DB.QueryRow(
		"INSERT INTO bioskop (nama, lokasi, rating) VALUES ($1, $2, $3) RETURNING id",
		strings.Trim(b.Nama, " "), strings.Trim(b.Lokasi, " "), b.Rating).Scan(&id)
	return id, err
}

func GetAllBioskop() ([]models.Bioskop, error) {
	rows, err := database.DB.Query("SELECT id, nama, lokasi, rating FROM bioskop")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bioskopSlice []models.Bioskop
	for rows.Next() {
		var b models.Bioskop
		if err := rows.Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating); err != nil {
			return nil, err
		}
		bioskopSlice = append(bioskopSlice, b)
	}
	return bioskopSlice, nil
}

func GetBioskopByID(id string) (models.Bioskop, error) {
	var b models.Bioskop
	err := database.DB.QueryRow("SELECT id, nama, lokasi, rating FROM bioskop WHERE id = $1", id).
		Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating)
	return b, err
}

func UpdateBioskop(id string, b models.Bioskop) error {
	_, err := database.DB.Exec("UPDATE bioskop SET nama=$1, lokasi=$2, rating=$3 WHERE id=$4",
		b.Nama, b.Lokasi, b.Rating, id)
	return err
}

func DeleteBioskop(id string) (int64, error) {
	res, err := database.DB.Exec("DELETE FROM bioskop WHERE id=$1", id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func BioskopExistsByNamaLokasi(nama, lokasi string) (bool, error) {
	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM bioskop WHERE nama=$1 AND lokasi=$2)", nama, lokasi).Scan(&exists)
	return exists, err
}

func BioskopExistsByID(id string) (bool, error) {
	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM bioskop WHERE id=$1)", id).Scan(&exists)
	return exists, err
}

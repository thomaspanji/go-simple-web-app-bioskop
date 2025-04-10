package handlers

import (
	"net/http"
	"strings"

	"github.com/thomaspanji/go-simple-web-app-bioskop/database"
	"github.com/thomaspanji/go-simple-web-app-bioskop/models"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func CreateBioskopHandler(ctx *gin.Context) {
	// Implement the handler to create a new bioskop
	var input struct {
		Nama   string  `json:"nama" binding:"required"`
		Lokasi string  `json:"lokasi" binding:"required"`
		Rating float64 `json:"rating" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Validate the input
	if strings.Trim(input.Nama, " ") == "" || strings.Trim(input.Lokasi, " ") == "" || input.Rating <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Semua field harus diisi dan rating harus lebih dari 0"})
		return
	}
	// check if the bioskop already exists
	var exists bool
	err := database.DB.QueryRow(`
        SELECT EXISTS (
            SELECT 1 FROM bioskop WHERE LOWER(TRIM(nama)) = LOWER(TRIM($1)) AND LOWER(TRIM(lokasi)) = LOWER(TRIM($2))
        )
    `, input.Nama, input.Lokasi).Scan(&exists)

	// Check for errors in the query, error when the table is not found
	// or other database errors
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memeriksa duplikasi data"})
		return
	}

	if exists {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Bioskop dengan nama dan lokasi yang sama sudah ada"})
		return
	}
	// Insert the new bioskop into the database
	query := `
	INSERT INTO bioskop (nama, lokasi, rating)
	VALUES ($1, $2, $3)
	RETURNING id
	`

	var id int
	err = database.DB.QueryRow(query, input.Nama, input.Lokasi, input.Rating).Scan(&id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menginput bioskop"})
		return
	}

	newBioskop := models.Bioskop{
		ID:     id,
		Nama:   input.Nama,
		Lokasi: input.Lokasi,
		Rating: input.Rating,
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Bioskop berhasil ditambahkan", "data": newBioskop})
}

package handlers

import (
	"net/http"
	"strings"

	"github.com/thomaspanji/go-simple-web-app-bioskop/models"
	"github.com/thomaspanji/go-simple-web-app-bioskop/repository"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func CreateBioskop(ctx *gin.Context) {
	var b models.Bioskop
	if err := ctx.ShouldBindJSON(&b); err != nil || strings.Trim(b.Nama, " ") == "" || strings.Trim(b.Lokasi, " ") == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan lokasi wajib diisi"})
		return
	}

	exists, err := repository.BioskopExistsByNamaLokasi(b.Nama, b.Lokasi)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal cek duplikat"})
		return
	}
	if exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bioskop dengan nama dan lokasi ini sudah ada"})
		return
	}

	id, err := repository.InsertBioskop(b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan bioskop"})
		return
	}

	b.ID = id
	ctx.JSON(http.StatusCreated, b)
}

func GetAllBioskop(ctx *gin.Context) {
	list, err := repository.GetAllBioskop()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	ctx.JSON(http.StatusOK, list)
}

func GetBioskopByID(ctx *gin.Context) {
	id := ctx.Param("id")
	b, err := repository.GetBioskopByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}
	ctx.JSON(http.StatusOK, b)
}

func UpdateBioskop(ctx *gin.Context) {
	id := ctx.Param("id")
	var b models.Bioskop
	if err := ctx.ShouldBindJSON(&b); err != nil || strings.Trim(b.Nama, " ") == "" || strings.Trim(b.Lokasi, " ") == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan lokasi wajib diisi"})
		return
	}

	exists, err := repository.BioskopExistsByID(id)
	if err != nil || !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}

	if err := repository.UpdateBioskop(id, b); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update data"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Bioskop berhasil diperbarui"})
}

func DeleteBioskop(ctx *gin.Context) {
	id := ctx.Param("id")
	count, err := repository.DeleteBioskop(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal hapus bioskop"})
		return
	}
	if count == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Bioskop berhasil dihapus"})
}

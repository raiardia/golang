package main // Mendeklarasikan package utama (main)

import (
	"fmt"      // Import untuk mencetak ke konsol
	"net/http" // Import untuk menangani HTTP request & response
	"time"     // Import untuk tipe data waktu

	"github.com/gin-gonic/gin" // Menggunakan framework Gin untuk web server
	"gorm.io/driver/mysql"     // Menggunakan driver MySQL untuk GORM (ORM)
	"gorm.io/gorm"             // Menggunakan GORM sebagai ORM untuk database
)

// Mendefinisikan struct User yang merepresentasikan tabel users di database
type User struct {
	ID 	  	  uint 		`grom:"coloum:id;primaryKey"` // ID sebagai primary key
	Name  	  string 	`grom:"coloum:name"` // Nama pengguna
	Email 	  string 	`grom:"coloum:email"` // Email pengguna
	Age   	  string 	`grom:"coloum:age"` // Umur pengguna
	CreatedAt time.Time `grom:"coloum:createdAt"` // Waktu pembuatan data
	UpdatedAt time.Time `grom:"coloum:updatedAt"`// Waktu terakhir diperbarui
} // End of struct User
	
func main() { 
	// Data Source Name (DSN) untuk koneksi ke database MySQL
	dsn := "root:@tcp(127.0.0.1:3306)/open_api?charset=utf8mb4&parseTime=True&loc=Local" // variable declaration
	
	// Menghubungkan ke database dengan GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // variable declaration
	if err != nil { // if statement
		fmt.Println("failed to connect database") // Menampilkan pesan error jika gagal terhubung
	}	// End of if statement

	// Melakukan migrasi database untuk membuat tabel secara otomatis
	db.AutoMigrate(&User{})
	
	// Membuat instance dari Gin untuk menangani HTTP request
	r := gin.Default() 

	// Menampilkan semua data users (GET /users)
	r.GET("/users", func(c *gin.Context) {
		var users []User // Array untuk menyimpan data user dari database
		db.Find(&users) // Mengambil semua data dari tabel users
		c.JSON(http.StatusOK, gin.H{"data": users}) // Mengirim response dalam format JSON
	})

	// Menampilkan data user berdasarkan ID (GET /users/:id)
	r.GET("/users/:id", func(c *gin.Context) {
		var user User
		// Mencari user berdasarkan ID, jika tidak ditemukan kirim error 404
		if err := db.First(&user, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": user}) // Mengirim data user dalam format JSON
	})

	// Menambahkan user baru ke database (POST /users)
	r.POST("/users", func(c *gin.Context) {
		var user User
		// Mengikat data JSON yang dikirim oleh client ke struct User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&user) // Menyimpan data user ke database
		c.JSON(http.StatusCreated, gin.H{"data": user}) // Mengirim response berhasil dibuat
	})

	// Mengubah data user berdasarkan ID (PUT /users/:id)
	r.PUT("/users/:id", func(c *gin.Context) {
		var user User
		// Mencari user berdasarkan ID, jika tidak ditemukan kirim error 404
		if err := db.First(&user, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		// Mengikat data JSON yang dikirim oleh client ke struct User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Save(&user) // Menyimpan perubahan data user ke database
		c.JSON(http.StatusOK, gin.H{"data": user}) // Mengirim response data yang telah diperbarui
	})

	// Menghapus user berdasarkan ID (DELETE /users/:id)
	r.DELETE("/users/:id", func(c *gin.Context) {
		// Menghapus user dari database berdasarkan ID
		if err := db.Delete(&User{}, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User deleted"}) // Mengirim response sukses
	})

	// Menjalankan server pada port 3000
	r.Run(":3000")
}
package main // package declaration

import (
	"fmt"      // import statement
	"net/http" // import statement
	"time" // import statement

	"github.com/gin-gonic/gin" // import statement
	"gorm.io/driver/mysql" // import statement
	"gorm.io/gorm" // import 
	
)

type User struct {
	ID 	  	  uint 		`grom:"coloum:id;primaryKey"` // struct declaration
	Name  	  string 	`grom:"coloum:name"` // struct declaration
	Email 	  string 	`grom:"coloum:email"` // struct declaration
	Age   	  string 	`grom:"coloum:age"` // struct declaration
	CreatedAt time.Time `grom:"coloum:createdAt"` // struct declaration
	UpdatedAt time.Time `grom:"coloum:updatedAt"`// struct declaration
} // struct declaration
	
func main() { // function declaration
	dsn := "root:@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local" // variable declaration
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // variable declaration
	if err != nil { // if statement
		fmt.Println("failed to connect database") // function call
	}

	route := gin.Default() // variable declaration

	route.GET("/users", func(c *gin.Context) { // function declaration
		var users []User // variable declaration
		db.Find(&users) // function call
		c.JSON(http.StatusOK, gin.H{"data": users}) // function call
	})

	route.Run(":3000") // function call
} // function declaration

	


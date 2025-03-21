package main // package declaration
import (
	"encoding/json" // import statement
	"fmt"      // import statement
	"net/http" // import statement
	"time" // import statement

	"gorm.io/driver/mysql" // import statement
	"gorm.io/gorm" // import statement
)

type User struct {
	ID 	  uint `grom:"coloum:id;primaryKey"` // struct declaration
	Name  string `grom:"coloum:name"` // struct declaration
	Email string `grom:"coloum:email"` // struct declaration
	Age   string `grom:"coloum:age"` // struct declaration
	CreatedAt time.Time // struct declaration
	UpdatedAt time.Time // struct declaration
} // struct declaration
	
func main() { // function declaration
	dsn := "root:@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local" // variable declaration
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // variable declaration
	if err != nil { // if statement
		fmt.Println("failed to connect database") // function call
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // function declaration
		var result User // variable declaration
		db.Find(&result) // function call
		response, _ := json.Marshal(result) // variable declaration
		fmt.Fprintf(w, string(response)) // function call
	})
	http.ListenAndServe(":3000", nil) // function call
} // function declaration


package db

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Note struct {
    gorm.Model
    UniqueUrl       uuid.UUID `json:"unique_url"`
    Title           string    `json:"title"`
    Content         string    `json:"content"`
    ExpirationDate  string    `json:"expiration_date"`
    CurrentViews    int       `json:"current_views"`
    MaxViews        int       `json:"max_views"`
}

func (n Note) String() string {
	return fmt.Sprintf("Note title: %s , content: %s, expiration_date: %s , uuid : %s" , n.Title , n.Content , n.ExpirationDate , n.UniqueUrl)
}


// should make el maxviews w el expiry date option w a have defult values
// remove panic bardo
func CreateNote(c *gin.Context) {
	if c.Request.Method == "POST" {
		var newNote Note
		if err := c.BindJSON(&newNote); err != nil {
			return
		}
		database:= GetDatabaseSingelton().GetDatabase()
		if res:= database.Create(&newNote)  ; res.Error!= nil {
			panic(res.Error)
		}
		c.IndentedJSON(http.StatusOK, newNote)
	} else {
		c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}



func GetNoteByUuid(c *gin.Context) {
	if c.Request.Method == "GET" {
		uuid , _ := uuid.Parse(c.Param("uuid"))
		response := Note{UniqueUrl: uuid}
		database:= GetDatabaseSingelton().GetDatabase()
		if res:= database.Find(&response) ; res.Error!= nil {
			panic(res.Error)
		}
		c.IndentedJSON(http.StatusOK, response)
	} else {
		c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}
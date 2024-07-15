package db

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	UniqueUrl      uuid.UUID `json:"unique_url"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	ExpirationDate string    `json:"expiration_date"`
	CurrentViews   int       `json:"current_views"`
	MaxViews       int       `json:"max_views"`
	Email          string    `json:"email"`
}

func (n Note) String() string {
	return fmt.Sprintf("Note title: %s , content: %s, expiration_date: %s , uuid : %s", n.Title, n.Content, n.ExpirationDate, n.UniqueUrl)
}

func CreateNote(c *gin.Context) {
	if c.Request.Method == "POST" {
		var newNote Note
		if err := c.BindJSON(&newNote); err != nil {
			return
		}
		database := GetDatabaseSingelton().GetDatabase()
		newNote.UniqueUrl = uuid.New()
		email, _ := c.Get("email")
		var ok bool
		newNote.Email = email.(string)
		fmt.Println(email)
		if !ok {
			c.IndentedJSON(http.StatusUnauthorized, "Note can't be created")
		}
		if newNote.ExpirationDate == "" {
			newNote.ExpirationDate = time.Now().AddDate(0, 3, 0).Format("2006-01-02")
		}
		if newNote.MaxViews == 0 {
			newNote.MaxViews = 100
		}
		if res := database.Create(&newNote); res.Error != nil {
			c.IndentedJSON(http.StatusConflict, "Note can't be created")
			return
		}
		c.IndentedJSON(http.StatusOK, newNote)
	} else {
		c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func GetNoteByUuid(c *gin.Context) {
	if c.Request.Method == "GET" {
		database := GetDatabaseSingelton().GetDatabase()
		uuid, _ := uuid.Parse(c.Param("uuid"))
		response := Note{UniqueUrl: uuid}

		if res := database.Find(&response); res.Error != nil {
			c.IndentedJSON(http.StatusNotFound, "Note not found")
			return
		} else if response.ExpirationDate < time.Now().Format("2006-01-02") {
			database.Delete(&response)
			c.IndentedJSON(http.StatusOK, "Note Expired")
			return
		}

		response.CurrentViews++
		if response.CurrentViews >= response.MaxViews {
			database.Delete(&response)
			c.IndentedJSON(http.StatusOK, "Max views reached")
			return
		}
		database.Save(&response)
		c.IndentedJSON(http.StatusOK, response)
	} else {
		c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func GetUserNotes(c *gin.Context) {
	if c.Request.Method == "GET" {

		database := GetDatabaseSingelton().GetDatabase()
		email, _ := c.Get("email")

		var notes []Note
		if err := database.Where("Email = ?", email).Find(&notes).Error; err != nil {
			c.JSON(http.StatusNotFound, "Notes not found")
			return
		}

		c.JSON(http.StatusOK, gin.H{"notes": notes})
	} else {
		c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}

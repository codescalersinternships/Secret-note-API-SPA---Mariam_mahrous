package db

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	UniqueUrl uuid.UUID
	Title string 
	Content  string 
	ExpirationDate string 
	CurrentViews int 
	MaxViews int 
}

func (n Note) String() string {
	return fmt.Sprintf("Note title: %s , content: %s, expiration_date: %s , uuid : %s" , n.Title , n.Content , n.ExpirationDate , n.UniqueUrl)
}


// should make el maxviews w el expiry date option w a have defult values
// remove panic bardo
func CreateNote(title , content , expirationDate string , maxViews int , database *gorm.DB) Note{
	newNote := Note{
		UniqueUrl:      uuid.New(),
		Title:          title,
		Content:        content,
		ExpirationDate: expirationDate,
		CurrentViews:   0,
		MaxViews:       maxViews,
	}
	if res:= database.Create(&newNote)  ; res.Error!= nil {
		panic(res.Error)
	}
	return newNote
}

//bygbly awl wa7da fy we4o lw mala2a4 el ana 3yza
func GetNoteByUuid(uuid uuid.UUID , database *gorm.DB) Note{
	targetNote := Note{UniqueUrl: uuid}
	if res:= database.Find(&targetNote) ; res.Error!= nil {
		panic(res.Error)
	}
	return targetNote
}
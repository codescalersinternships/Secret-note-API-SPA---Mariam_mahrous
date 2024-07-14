package main

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
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

var db1, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
var dsn = "host=localhost user=postgres password=123456 dbname=postgres port=5432"
var db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})


// should make el maxviews w el expiry date option w a have defult values
// remove panic bardo
func createNote(title , content , expirationDate string , maxViews int) Note{
	newNote := Note{
		UniqueUrl:      uuid.New(),
		Title:          title,
		Content:        content,
		ExpirationDate: expirationDate,
		CurrentViews:   0,
		MaxViews:       maxViews,
	}
	if res:= db.Create(&newNote) ; res.Error!= nil {
		panic(res.Error)
	}
	return newNote
}

//bygbly awl wa7da fy we4o lw mala2a4 el ana 3yza
func getNoteByUuid(uuid uuid.UUID) Note{
	targetNote := Note{UniqueUrl: uuid}
	if res:= db.Find(&targetNote) ; res.Error!= nil {
		panic(res.Error)
	}
	return targetNote
}

func main() {
	db.AutoMigrate(&Note{})
	db1.AutoMigrate(&Note{})
	newnote := createNote("mimo", "content",  "2024-07-17" , 10)
	fmt.Println(newnote)
	uuid , _ := uuid.Parse("40c0d3ad-0d16-4973-95bb-a1624b725bc6")
	targetNote := getNoteByUuid(uuid)
	fmt.Println(targetNote)
}
package app

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	model "github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/models"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	database *gorm.DB
	host     string
	user     string
	password string
	dbname   string
	port     string
}

type ConfigDB struct {
	DatabaseType string
	Host         string
	User         string
	Password     string
	DBName       string
	Port         string
}

func NewDB(configdb ConfigDB) (*DB, error) {
	db := &DB{
		host:     configdb.Host,
		user:     configdb.User,
		password: configdb.Password,
		dbname:   configdb.DBName,
		port:     configdb.Port,
	}
	err := db.SetDatabase(configdb.DatabaseType)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (s *DB) SetDatabase(databaseType string) error {
	var dbErr error
	if databaseType == "POSTGRESQL" {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", s.host, s.user, s.password, s.dbname, s.port)
		s.database, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if dbErr != nil {
			return fmt.Errorf("failed to connect to PostgreSQL database: %v", dbErr.Error())
		}
	} else {
		s.database, dbErr = gorm.Open(sqlite.Open(s.dbname+".db"), &gorm.Config{})
		if dbErr != nil {
			return fmt.Errorf("failed to connect to Sqlite database: %v", dbErr.Error())
		}
	}
	dbErr = s.database.AutoMigrate(&model.Note{}, &model.User{})
	if dbErr != nil {
		return fmt.Errorf("failed to migrate to database: %v", dbErr.Error())
	}
	return nil
}

func (s *DB) GetNoteByUuid(uuid uuid.UUID) (model.Note, int, error) {
	note := model.Note{UniqueUrl: uuid}

	if res := s.database.Find(&note); res.Error != nil {
		return note, http.StatusNotFound, errors.New("not found")

	} else if note.ExpirationDate < time.Now().Format("2006-01-02") {
		s.database.Delete(&note)
		return note, http.StatusNotFound, errors.New("expired note")
	}

	note.CurrentViews++
	if note.CurrentViews >= note.MaxViews {
		s.database.Delete(&note)
		return note, http.StatusOK, nil
	}
	s.database.Save(&note)
	return note, http.StatusOK, nil
}

func (s *DB) GetAllNotes(id uint) ([]model.Note, int, error) {
	var notes []model.Note
	if err := s.database.Where("user_id = ?", id).Find(&notes).Error; err != nil {
		return notes, http.StatusNotFound, errors.New("not found")
	}
	return notes, http.StatusOK, nil
}

func (s *DB) CreateNote(id uint, newNote model.Note) (model.Note, int, error) {
	newNote.UniqueUrl = uuid.New()
	newNote.UserID = id
	if newNote.ExpirationDate == "" {
		newNote.ExpirationDate = time.Now().AddDate(0, 3, 0).Format("2006-01-02")
	}
	if newNote.MaxViews == 0 {
		newNote.MaxViews = 100
	}
	if res := s.database.Create(&newNote); res.Error != nil {
		return newNote, http.StatusConflict, errors.New("cannot create a new note")
	}
	return newNote, http.StatusOK, nil
}

func (s *DB) VerifyUser(email any) (bool, model.User) {
	var user model.User
	s.database.Where("email = ?", email).First(&user)
	return user.Email != "", user
}

func (s *DB) GetUser(email string) (model.User, bool) {
	var user model.User
	s.database.First(&user, "email = ?", email)
	if user.ID == 0 {
		return user, false
	}
	return user, true
}

func (s *DB) SignUp(newUser model.User) (int, error) {
	result := s.database.Where("email = ?", newUser.Email).First(&model.User{})
	if result.Error == nil {
		return http.StatusConflict, errors.New("user already exists")
	}

	if res := s.database.Create(&newUser); res.Error != nil {
		return http.StatusConflict, errors.New("cannpt create user")
	}
	return http.StatusOK, nil

}

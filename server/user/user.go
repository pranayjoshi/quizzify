package user

import (
	"context"

	// "encoding/json"
	"fmt"
	// "net/http"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/google/uuid"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/api/option"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
}

type UserLogin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
type Quiz struct {
	Content  string `json:"content"`
	Author   string `json:"author"`
	QuizID   string `json:"quiz_id"`
	Name     string `json:"name"`
	QuizName string `json:"quiz_name"`
}

// var client *mongo.Client

const dbName = "quizify"
const colName = "usersw"

// var collection *mongo.Collection

type FireDB struct {
	*db.Client
}

var fireDB FireDB

func (db *FireDB) Connect() error {
	// Find home directory.
	home, err := os.Getwd()
	if err != nil {
		return err
	}
	ctx := context.Background()
	opt := option.WithCredentialsFile(home + "/user/quizify-d54b0-99b8822049a6.json")
	config := &firebase.Config{DatabaseURL: "https://quizify-d54b0-default-rtdb.firebaseio.com"}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)
	}
	client, err := app.Database(ctx)
	if err != nil {
		return fmt.Errorf("error initializing database: %v", err)
	}
	db.Client = client
	return nil
}

func FirebaseDB() *FireDB {
	return &fireDB
}

type Store struct {
	*FireDB
}

func NewStore() *Store {
	d := FirebaseDB()
	return &Store{
		FireDB: d,
	}
}

// Create a new BIN object
func (s *Store) Create(u *User) {
	if err := s.NewRef("users/"+u.Name).Set(context.Background(), u); err != nil {
		fmt.Println("Failed to create user: %v", err)
	}
	// return nil
}

func (s *Store) Delete(u *User) error {
	return s.NewRef("users/" + u.Email).Delete(context.Background())
}

func (s *Store) GetByName(b string) (*User, error) {
	bin := &User{}
	if err := s.NewRef("users/"+b).Get(context.Background(), bin); err != nil {
		return nil, err
	}
	if bin.Email == "" {
		return nil, nil
	}
	return bin, nil
}

func (s *Store) GetByQuiz(b string) (*Quiz, error) {
	bin := &Quiz{}
	if err := s.NewRef("quiz/"+b).Get(context.Background(), bin); err != nil {
		return nil, err
	}
	// if bin.ID == "" {
	// 	return nil, nil
	// }
	return bin, nil
}

func (s *Store) Update(b string, m map[string]interface{}) error {
	return s.NewRef(b).Update(context.Background(), m)
}

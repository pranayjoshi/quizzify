package main

import (
	"fmt"

	// "github.com/google/uuid"
	"github.com/pranayjoshi/quizify/app"
	"github.com/pranayjoshi/quizify/user"
)

func main() {
	app := &app.App{}
	err := user.FirebaseDB().Connect()
	if err != nil {
		fmt.Println("Database connection  error", err)
	}
	// Initiate our store
	// store := user.NewStore()
	// A new BIN creation
	// store.Create(&user.User{
	// 	ID:       uuid.New(),
	// 	Email:    "awad@github.io",
	// 	Name:     "Awad",
	// 	Password: "1234",
	// })
	// Get the BIN by bin number
	// binw, _ := store.GetByBin("awad@github.io")
	// binw, err := store.GetByName("Awad")
	// if err != nil {
	// 	fmt.Println("Failed to get user: %v", err)
	// }
	// fmt.Println("Update", binw.Name)
	// //  // Update the BIN issuer information
	// //  err = store.Update("1234", map[string]interface{}{"issuer_bank_name": "Alpha Bank",
	// //   "issuer_website": "https://alpha.com"})
	// //  //// Checking update balance
	// bin2, err := store.GetByBin("1234")
	// fmt.Println("Update", binw.Email)
	// //// Delete the account
	// err = store.Delete(bin)

	// if err != nil {
	// 	fmt.Println("BIN related error", err)
	// }
	app.Start()
}

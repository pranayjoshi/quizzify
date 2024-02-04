package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	u "github.com/pranayjoshi/quizify/user"
	. "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()
	server.Use(Recovery(func(err interface{}, c *Context) {
		if httpError, ok := err.(HttpError); ok {
			c.JSON(httpError.Status, H{
				"message": httpError.Error(),
			})
		} else {
			message := fmt.Sprintf("%s", err)
			c.JSON(500, H{
				"message": message,
			})
		}
	}))
	server.GET("/", func(context *Context) {
		context.JSON(200, H{
			"message": "OawdaK",
		})
	})
	server.POST("/register", func(context *Context) {
		var user u.User
		err := json.NewDecoder(context.Req.Body).Decode(&user)
		if err != nil {
			context.JSON(400, H{
				"message": err.Error(),
			})
			return
		}

		store := u.NewStore()
		store.Create(&user)

		context.JSON(200, H{
			"data": fmt.Sprintf("Successfully created user: %s", user.Name),
		})
	})

	server.POST("/login", func(context *Context) {
		var user u.User
		err := json.NewDecoder(context.Req.Body).Decode(&user)
		if err != nil {
			context.JSON(400, H{
				"message": err.Error(),
			})
			return
		}

		store := u.NewStore()
		fmt.Println(user.Name)
		u, err := store.GetByName(user.Name)
		if err != nil {
			context.JSON(500, H{
				"message": err.Error(),
			})
			return
		}

		if u.Password != user.Password {
			context.JSON(401, H{
				"message": "Invalid password",
			})
			return
		}

		context.JSON(200, H{
			"token": u.Name,
		})
	})
	server.GET("/hello", func(context *Context) {
		name := context.Query("name")
		if name == "" {
			context.JSON(400, H{
				"message": "name not found",
			})
		} else {
			context.JSON(200, H{
				"data": fmt.Sprintf("Hello %s!", name),
			})
		}
	})
	server.GET("/user/:id", func(context *Context) {
		context.JSON(400, H{
			"data": H{
				"id": context.Param("id"),
			},
		})
	})
	server.GET("/long/long/long/path/*test", func(context *Context) {
		context.JSON(200, H{
			"data": H{
				"url": context.Path,
			},
		})
	})
	server.Handle(w, r)

}

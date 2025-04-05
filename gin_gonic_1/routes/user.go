package routes

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var (
	users  = make(map[string]User)
	userId = 0
	mu     sync.Mutex
)

func CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	mu.Lock()
	userId++
	user.ID = userId
	users[strconv.Itoa(userId)] = user
	mu.Unlock()
	c.JSON(http.StatusCreated, user)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	mu.Lock()
	defer mu.Unlock()
	if user, ok := users[id]; ok {
		c.JSON(http.StatusOK, user)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func GetUsers(c *gin.Context) {
	var allUsers []User
	for _, user := range users {
		allUsers = append(allUsers, user)
	}
	c.JSON(http.StatusOK, allUsers)
}

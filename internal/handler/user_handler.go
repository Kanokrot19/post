package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User โครงสร้างข้อมูลผู้ใช้
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ErrorResponse โครงสร้างข้อมูลสำหรับข้อผิดพลาด
type ErrorResponse struct {
	Message string `json:"message"`
}

// CreateUserRequest โครงสร้างสำหรับคำขอ POST
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required,name"`   // ฟิลด์ที่จำเป็น
	Email string `json:"email" binding:"required,email"` // อีเมลต้องอยู่ในรูปแบบที่ถูกต้อง
}

// CreateUserResponse โครงสร้างสำหรับการตอบกลับ POST
type CreateUserResponse struct {
	ID    int    `json:"id"`    // swagger:order 1
	Name  string `json:"name"`  // swagger:order 2
	Email string `json:"email"` // swagger:order 3
}

// Mock database
var users = []User{}
var nextID = 1

// @Summary      Get user by ID
// @Description  Get details of a user by ID
// @Tags         Users
// @Produce      json
// @Param        id   path      int     true  "User ID"
// @Success      200  {object}  User
// @Failure      404  {object}  ErrorResponse
// @Router       /users/{id} [get]
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id, "name": "ณัฐโชติ พรหมฤทธิ์"})
}

// @Summary      Create a new user
// @Description  Add a new user to the system
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body      CreateUserRequest  true  "User Information"
// @Success      201   {object}  CreateUserResponse
// @Failure      400   {object}  ErrorResponse
// @Failure      409   {object}  ErrorResponse
// @Router       /users [post]
func CreateUser(c *gin.Context) {
	// อ่านและตรวจสอบ JSON Body
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	// ตรวจสอบว่า email ซ้ำหรือไม่
	for _, user := range users {
		if user.Email == req.Email {
			c.JSON(http.StatusConflict, ErrorResponse{Message: "Email already exists"})
			return
		}
	}

	// สร้างผู้ใช้ใหม่
	newUser := User{
		ID:    nextID,
		Name:  req.Name, // ลำดับ: name มาก่อน email
		Email: req.Email,
	}
	nextID++ // เพิ่มค่า ID ถัดไป
	users = append(users, newUser)

	// ตอบกลับผู้ใช้ใหม่
	c.JSON(http.StatusCreated, CreateUserResponse{
		ID:    newUser.ID,
		Name:  newUser.Name, // ลำดับ: name มาก่อน email
		Email: newUser.Email,
	})
}

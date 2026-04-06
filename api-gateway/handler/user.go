package handler

import (
	pb "api-gateway/proto"
	"api-gateway/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	client pb.UserServiceClient
}

func RegisterUserRoutes(r *gin.Engine, client pb.UserServiceClient) {
	h := &UserHandler{client}

	r.POST("/user", h.CreateUser)
	r.POST("/login", h.Login)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.client.CreateUser(c, &pb.User{
		Name: req.Name,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

// login dummy
func (h *UserHandler) Login(c *gin.Context) {
    
    token, err := utils.GenerateToken(1)
    if err !=nil{
        c.JSON(http.StatusInternalServerError, gin.H{
            "error":"could not generate token",
        })
        return
    }
    
    
    
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

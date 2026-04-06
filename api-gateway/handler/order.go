package handler

import (
	pb "api-gateway/proto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	client pb.OrderServiceClient
}

func RegisterOrderRoutes(r gin.IRoutes, client pb.OrderServiceClient) {
	h := &OrderHandler{client}

	r.POST("/order", h.CreateOrder)
	r.GET("/order/:id", h.GetOrder)

}

// Create Order
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req struct {
		UserId  int    `json:"user_id"`
		Product string `json:"product"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.client.CreateOrder(c, &pb.CreateOrderRequest{
		UserId:  int32(req.UserId),
		Product: req.Product,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)

}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	res, err := h.client.GetOrder(c, &pb.GetOrderRequest{
		Id: int32(id),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid id",
		})
		return
	}

	c.JSON(http.StatusOK, res)

}

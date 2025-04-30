package api

import (
	entity "demo/application/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) SignUp(c *gin.Context) {
	var staff entity.Staff
	if err := c.ShouldBindJSON(&staff); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}
	ctx := c.Request.Context()
	jwt, err := s.userservice.RegisterUser(ctx, staff)
	if err != nil {
		s.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to register user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": jwt,
	})
}

type LoginReq struct {
	StaffID  string
	Password string
	Type     string
}

func (s *Server) Login(c *gin.Context) {
	var loginReq LoginReq
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}
	ctx := c.Request.Context()
	jwt, err := s.userservice.LoginUser(ctx, loginReq.StaffID, loginReq.Password, loginReq.Type)
	if err != nil {
		s.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to Login user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": jwt,
	})
}

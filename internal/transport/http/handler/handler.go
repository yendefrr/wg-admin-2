package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yendefrr/wg-admin/internal/service"
)

type Handler struct {
	usersService    service.Users
	profilesService service.Profiles
}

func NewHandler(usersService service.Users, employersService service.Profiles) *Handler {
	return &Handler{
		usersService:    usersService,
		profilesService: employersService,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	h.SetRoutes(router)

	return router
}

func (h *Handler) SetRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/profiles", h.HandleProfiles)
}

func (h *Handler) HandleProfiles(c *gin.Context) {
	profilesA, err := h.profilesService.GetAllActive(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}

	profilesInA, err := h.profilesService.GetAllInActive(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"active profile":   profilesA,
		"inactive profile": profilesInA,
	})
}

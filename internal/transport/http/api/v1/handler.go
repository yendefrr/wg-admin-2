package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yendefrr/wg-admin/internal/models"
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
	api := router.Group("/v1")
	{
		debug := api.Group("/debug")
		{
			debug.GET("/ping", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "pong",
				})
			})
		}

		users := api.Group("/users")
		{
			users.POST("/create", h.HandleCreateUser)
		}

		profiles := api.Group("/profiles")
		{
			profiles.GET("/", h.HandleProfiles)
			profiles.POST("/create", h.HandleCreateProfile)
		}
	}
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
		"active":   profilesA,
		"inactive": profilesInA,
	})
}

func (h *Handler) HandleCreateUser(c *gin.Context) {
	form := models.UserCreateForm{
		Username: c.Query("username"),
		Password: c.Query("password"),
		Role:     c.Query("role"),
	}

	if err := h.usersService.Create(c, form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"error": nil,
	})
}

func (h *Handler) HandleCreateProfile(c *gin.Context) {
	form := models.ProfileCreateForm{
		Username: c.Query("username"),
		Name:     c.Query("name"),
	}

	if err := h.profilesService.Create(c, form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"error": nil,
	})
}

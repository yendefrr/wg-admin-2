package handler

import (
	"html/template"

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
	users := router.Group("/users")
	{
		users.POST("/create", h.HandleCreateUser)
	}

	profiles := router.Group("/profiles")
	{
		profiles.GET("/", h.HandleProfiles)
		profiles.POST("/create", h.HandleCreateProfile)
	}
}

func (h *Handler) HandleProfiles(c *gin.Context) {
	profilesA, err := h.profilesService.GetAllActive(c)
	if err != nil {
		return
	}

	profilesInA, err := h.profilesService.GetAllInActive(c)
	if err != nil {
		return
	}

	params := map[string]interface{}{}
	params["profilesActive"] = profilesA
	params["profilesInActive"] = profilesInA

	t, _ := template.New("").Parse("")
	t.Execute(c.Writer, params)
}

func (h *Handler) HandleCreateUser(c *gin.Context) {
	form := models.UserCreateForm{
		Username: c.Query("username"),
		Password: c.Query("password"),
		Role:     c.Query("role"),
	}

	if err := h.usersService.Create(c, form); err != nil {
		return
	}

	t, _ := template.New("").Parse("")
	t.Execute(c.Writer, nil)
}

func (h *Handler) HandleCreateProfile(c *gin.Context) {
	form := models.ProfileCreateForm{
		Username: c.Query("username"),
		Name:     c.Query("name"),
	}

	if err := h.profilesService.Create(c, form); err != nil {
		return
	}

	t, _ := template.New("").Parse("")
	t.Execute(c.Writer, nil)
}

package handlers

import (
	"net/http"

	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/presence/usecases"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/models"
	"github.com/gin-gonic/gin"
)

type PresenceHandler struct {
	Usecase usecases.PresenceUsecase
}

func NewPresenceHandler(uc usecases.PresenceUsecase) *PresenceHandler {
	return &PresenceHandler{Usecase: uc}
}

func (h *PresenceHandler) SetOnline(c *gin.Context) {
	userID := c.Param("id")
	if err := h.Usecase.SetOnline(userID); err != nil {
		c.JSON(http.StatusInternalServerError, models.NewErrResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, models.NewOkResponse(gin.H{"status": "online"}))
}

func (h *PresenceHandler) SetOffline(c *gin.Context) {
	userID := c.Param("id")
	if err := h.Usecase.SetOffline(userID); err != nil {
		c.JSON(http.StatusInternalServerError, models.NewErrResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, models.NewOkResponse(gin.H{"status": "offline"}))
}

func (h *PresenceHandler) GetStatus(c *gin.Context) {
	userID := c.Param("id")
	status, err := h.Usecase.GetStatus(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewErrResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, models.NewOkResponse(gin.H{"status": status}))
}

func (h *PresenceHandler) GetLastSeen(c *gin.Context) {
	userID := c.Param("id")
	lastSeen, err := h.Usecase.GetLastSeen(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, models.NewErrResponse("Last seen not found"))
		return
	}
	c.JSON(http.StatusOK, models.NewOkResponse(gin.H{"last_seen": lastSeen}))
}

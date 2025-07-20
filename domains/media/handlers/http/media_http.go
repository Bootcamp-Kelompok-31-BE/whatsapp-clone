package http

import (
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/media"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/media/models/requests"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/models/responses"

	// "github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/media/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MediaHttp struct {
	mc media.MediaUseCase
}

func NewMediaHttp(mc media.MediaUseCase) *MediaHttp {
	return &MediaHttp{mc: mc}
}

func (mh *MediaHttp) Home(c *gin.Context) {
	ctx := c.Request.Context()
	
	response, err := mh.mc.Home(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, responses.BasicResponse{
		Data: response.Data})
}

func (mh *MediaHttp) GetFile(c *gin.Context) {
	name := c.Param("name")
	mediaResponse, err := mh.mc.GetFile(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, responses.BasicResponse{
		Data: mediaResponse})
}

func (mh *MediaHttp) UploadFile(c *gin.Context) {
	
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}
	defer file.Close()

	request := &requests.MediaResponse{
		File:  file,
		Header: header,
	}

	ctx := c.Request.Context()
	mediaResponse, err := mh.mc.UploadFile(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, responses.BasicResponse{
		Data: mediaResponse.Data,
	})

}

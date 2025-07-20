package usecases

import (
	"context"
	// "io"
	// "os"
	"path/filepath"
	"strings"

	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/media"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/media/entities"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/media/models/requests"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/media/models/responses"
	// "github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/E2E"
	sharedresponses "github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/models/responses"
)

type mediaUseCase struct {
	repo media.MediaRepository
}

func NewMediaUseCase(repo media.MediaRepository) media.MediaUseCase {
	return &mediaUseCase{repo: repo}
}

func (mc *mediaUseCase) Home(ctx context.Context) (*sharedresponses.BasicResponse, error) {
	return &sharedresponses.BasicResponse{
		Data: struct {
			Message string
		}{
			Message: "Welcome to the Media Service",
		},
	}, nil
}

func (mc *mediaUseCase) UploadFile(ctx context.Context, request *requests.MediaResponse) (*sharedresponses.BasicResponse, error) {
	filename := request.Header.Filename
	extension := strings.ToLower(filepath.Ext(filename))
	switch extension {
		case ".jpg", ".jpeg", ".png", ".gif":
			request.MediaType = "image"
		case ".pdf", ".doc", ".docx", ".xls", ".xlsx", ".txt":
			request.MediaType = "document"
		default:
			request.MediaType = "unknown"
	}


	media := &entities.Media{
		Name:      filename,
		MediaType: request.MediaType,
	}

	_, err := mc.repo.CreateFile(ctx, media.Name, media.MediaType)
	if err != nil {
		return nil, err
	}

	return &sharedresponses.BasicResponse{
		Data: struct {
			Message string
		}{
			Message: "File uploaded successfully",
		},
	}, nil
}

func (mc *mediaUseCase) GetFile(ctx context.Context, name string) (*responses.MediaResponse, error) {
	mediaEntity, err := mc.repo.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return &responses.MediaResponse{
		Name:      mediaEntity.Name,
		MediaType: mediaEntity.MediaType,
	}, nil
}
	

/*
func (mc *mediaUseCase) SendFile(ctx context.Context, request *requests.MediaResponse) (*sharedresponses.BasicResponse, error) {
	filename := request.Header.Filename
	extension := strings.ToLower(filepath.Ext(filename))
	var mediaType string

	switch extension {
	case ".jpg", ".jpeg", ".png", ".gif":
		mediaType = "image"
	case ".pdf", ".doc", ".docx", ".xls", ".xlsx", ".txt":
		mediaType = "document"
	default:
		mediaType = "unknown"
	}
	
	media := &entities.Media{
		Name:      filename,
		MediaType: mediaType,
		Sender:    request.Sender,
		Receiver:  request.Receiver,
	}

	_, err := mc.repo.CreateFile(ctx, media.Name, media.MediaType)
	if err != nil {
		return nil, err
	}

	return &sharedresponses.BasicResponse{
		Data: struct {
			Message string
		}{
			Message: "File uploaded successfully",
		},
	}, nil
} */


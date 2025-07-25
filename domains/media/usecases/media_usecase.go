package usecases

import (
	"context"

	"path/filepath"
	"strings"

	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/media"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/media/entities"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/media/models/requests"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/media/models/responses"

	// "github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/E2E"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/middlewares"
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
	
	encryptedFile, err1 := middlewares.Encrypt([]byte(filename))
	encryptedSender, err2 := middlewares.Encrypt([]byte(request.Sender))
	encryptedReceiver, err3 := middlewares.Encrypt([]byte(request.Receiver))

	switch {
		case err1 != nil:
			return nil, err1
		case err2 != nil:
			return nil, err2
		case err3 != nil:
			return nil, err3
	}

	media := &entities.Media{
		Name:      string(encryptedFile),
		MediaType: request.MediaType,
		Sender:    string(encryptedSender),
		Receiver:  string(encryptedReceiver),
	}

	_, err := mc.repo.CreateFile(ctx, media.Name, media.MediaType, media.Sender, media.Receiver)
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
		Sender:    mediaEntity.Sender,
		Receiver:  mediaEntity.Receiver,
	}, nil
}
	
// ga pake ini dulu
func (mc *mediaUseCase) SendFile(ctx context.Context, name string, to string) (*sharedresponses.BasicResponse, error) {
	media, err := mc.repo.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}
	media.Receiver = to
	err = mc.repo.UpdateMedia(ctx, media)
	if err != nil {
		return nil, err
	}

	return &sharedresponses.BasicResponse{
		Data: struct {
			Message string
		}{
			Message: "Message sent successfully",
		},
	}, nil

}

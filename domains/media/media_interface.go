package media

import (
	"context"
	// "mime/multipart"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/media/entities"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/media/models/requests"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/media/models/responses"
	sharedresponses "github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/models/responses"
)

type MediaUseCase interface {
	Home(ctx context.Context) (*sharedresponses.BasicResponse, error)
	UploadFile(ctx context.Context, request *requests.MediaResponse) (*sharedresponses.BasicResponse, error)
	GetFile(ctx context.Context, name string) (*responses.MediaResponse, error)
	// SendFile(ctx context.Context, request *requests.MediaResponse) (*sharedresponses.BasicResponse, error)
}

type MediaRepository interface {
	FindByName(ctx context.Context, name string) (*entities.Media, error)
	CreateFile(ctx context.Context, name string, MediaType string) (*entities.Media, error)
}

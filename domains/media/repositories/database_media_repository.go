package repositories

import (
	"context"
	"time"

	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/media"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/media/entities"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/infrastructures"
	"github.com/google/uuid"
)

type databaseMediaRepository struct {
	db infrastructures.Database
}

func NewDatabaseMediaRepository(db infrastructures.Database) media.MediaRepository {
	return &databaseMediaRepository{db: db}
}

func (repo *databaseMediaRepository) FindByName(ctx context.Context, name string) (*entities.Media, error) {
	var mediaEntity entities.Media
	result := repo.db.GetInstance().WithContext(ctx).Where("name = ?", name).First(&mediaEntity)
	// result := repo.db.GetInstance().First(&mediaEntity, "name = ?", name)
	if result.Error != nil {
		return nil, result.Error
	}
	return &mediaEntity, nil
}

func (repo *databaseMediaRepository) CreateFile(ctx context.Context, name string, mediaType string) (*entities.Media, error) {
	result := repo.db.GetInstance().Create(&entities.Media{
		ID: 	   uuid.New(),
		Name:      name,
		MediaType: mediaType,
		CreatedAt: time.Now(),
	})
	
	if result.Error != nil {
		return nil, result.Error
	}

	return nil, nil
}


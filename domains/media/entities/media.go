package entities

import "github.com/google/uuid"
import "time"
// import "multipart/form-data"

type Media struct {
	ID        uuid.UUID 	`gorm:"type:uuid;primaryKey"`
	Name      string    	`json:"name"`
	MediaType string    	`json:"type"` // e.g., image, video, audio, document
	Sender    string    	
	Receiver  string 
	CreatedAt time.Time    `json:"created_at"`
}
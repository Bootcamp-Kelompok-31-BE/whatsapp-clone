package requests

import "mime/multipart"

// upload file
type MediaResponse struct {
	File 		multipart.File
	Header 		*multipart.FileHeader
	// Name        string `json:"name"`
	MediaType 	string `json:"type"`
	// Sender 		string `json:"sender"`
	// Receiver 	string `json:"receiver"`
}
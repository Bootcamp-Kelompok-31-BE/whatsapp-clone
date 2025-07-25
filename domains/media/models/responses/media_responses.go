package responses

type MediaResponse struct {
	Name        string `json:"name"`
	MediaType 	string `json:"type"`
	Sender 		string `json:"sender"`
	Receiver 	string `json:"receiver"`
}
package pages

// page entity
type Page struct {
	UUID       int    `json:"uuid" `
	Name       string `json:"name" `
	CreateDate int    `json:"createdate" `
	Article    string `json:"article" `
	Autor      string `json:"autor" `
	Comments   string `json:"comments"`
	Rating     int    `Json:"rating"`
}

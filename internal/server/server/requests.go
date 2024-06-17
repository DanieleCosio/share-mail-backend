package server

type Attachment struct {
	Name       string `json:"name"`
	Url        string `json:"url"`
	PreviewUrl string `json:"preview"`
	Size       string `json:"size"`
	MimeType   string `json:"mimeType"`
}

type GetEmailLinkBody struct {
	RequestAccountOwner string       `json:"requestAccountOwner"`
	MessageHtml         string       `json:"messageHtml"`
	Attachments         []Attachment `json:"attachments"`
}

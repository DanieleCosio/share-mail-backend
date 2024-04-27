package server

type GetEmailLinkBody struct {
	RequestAccountOwner string `json:"requestAccountOwner"`
	MessageHtml         string `json:"messageHtml"`
}

package server

type Method int64

func (m Method) String() string {
	switch m {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case DELETE:
		return "DELETE"
	default:
		return "UNKNOWN"
	}
}

const (
	GET Method = iota
	POST
	PUT
	DELETE
)

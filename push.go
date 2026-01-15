package push

type IPush interface {
	Push(id, title, content string) (*Response, error)
	PushAll(ids []string, title, content string) (*Response, error)
}

type Response struct {
	StatusCode int
	Content    []byte
}

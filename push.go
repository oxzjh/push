package push

type IPush interface {
	Push(id, title, content string) ([]byte, error)
	PushAll(ids []string, title, content string) ([]byte, error)
}

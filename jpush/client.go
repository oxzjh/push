package jpush

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const URL_PUSH = "https://api.jpush.cn/v3/push"

type Client struct {
	authorization string
	iosDevMode    bool
}

func (c *Client) send(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", c.authorization)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(string(b))
	}
	return b, nil
}

func (c *Client) Get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.send(req)
}

func (c *Client) Post(url string, data any) ([]byte, error) {
	body, _ := json.Marshal(data)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return c.send(req)
}

func (c *Client) PushBody(body *Body) ([]byte, error) {
	if body.Options == nil {
		body.Options = &Option{}
	}
	if !c.iosDevMode {
		body.Options.ApnsProduction = true
	}
	return c.Post(URL_PUSH, body)
}

func (c *Client) Push(id, title, content string) ([]byte, error) {
	return c.PushAll([]string{id}, title, content)
}

func (c *Client) PushAll(ids []string, title, content string) ([]byte, error) {
	return c.PushBody(&Body{
		Platform: PLATFORM_ALL,
		Audience: &Audience{RegistrationId: ids},
		Notification: &Notification{
			Alert: content,
			Android: &Android{
				Alert: content,
				Title: title,
			},
			Ios: &Ios{
				Alert: &IosAlert{
					Title: title,
					Body:  content,
				},
				Badge: "+1",
			},
		},
	})
}

func NewClient(appKey, masterSecret string, iosDevMode bool) *Client {
	return &Client{"Basic " + base64.StdEncoding.EncodeToString([]byte(appKey+":"+masterSecret)), iosDevMode}
}

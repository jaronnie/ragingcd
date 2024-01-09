package restc

import (
	"net/http"
	"time"
)

type Interface interface {
	Verb(verb string) *Request
	Post() *Request
	Get() *Request

	GetHeader() http.Header
}

type Opt func(client *RESTClient) error

type RESTClient struct {
	protocol string
	addr     string
	port     string

	retryTimes int
	retryDelay time.Duration

	headers http.Header

	// Set specific behavior of the client.  If not set http.DefaultClient will be used.
	client *http.Client
}

func (r *RESTClient) Verb(verb string) *Request {
	return NewRequest(r).Verb(verb)
}

func (r *RESTClient) Post() *Request {
	return r.Verb("POST")
}

func (r *RESTClient) Get() *Request {
	return r.Verb("GET")
}

func (r *RESTClient) GetHeader() http.Header {
	return r.headers
}

func New(ops ...Opt) (*RESTClient, error) {
	c := &RESTClient{}
	for _, op := range ops {
		if err := op(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

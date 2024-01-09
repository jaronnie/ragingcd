package restc

import (
	"net/http"
	"net/url"
	"time"
)

func WithUrl(u string) Opt {
	return func(c *RESTClient) error {
		// parse url
		parse, err := url.Parse(u)
		if err != nil {
			return err
		}
		c.protocol = parse.Scheme
		c.addr = parse.Host
		c.port = parse.Port()
		return nil
	}
}

func WithClient(client *http.Client) Opt {
	return func(c *RESTClient) error {
		c.client = client
		return nil
	}
}

func WithHeaders(headers http.Header) Opt {
	return func(c *RESTClient) error {
		c.headers = headers
		return nil
	}
}

func WithProtocol(protocol string) Opt {
	return func(c *RESTClient) error {
		c.protocol = protocol
		return nil
	}
}

func WithAddr(addr string) Opt {
	return func(c *RESTClient) error {
		c.addr = addr
		return nil
	}
}

func WithPort(port string) Opt {
	return func(c *RESTClient) error {
		c.port = port
		return nil
	}
}

func WithRetryTimes(times int) Opt {
	return func(c *RESTClient) error {
		c.retryTimes = times
		return nil
	}
}

func WithRetryDelay(time time.Duration) Opt {
	return func(c *RESTClient) error {
		c.retryDelay = time
		return nil
	}
}

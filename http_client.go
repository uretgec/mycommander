package mycommander

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type HttpClient struct {
	Client *http.Client
}

func NewHttpClient(timeout int64) *HttpClient {

	// update connection pool -> from 2 to 100
	// Reuse http client more requests/responses
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	return &HttpClient{
		Client: &http.Client{
			Timeout:   time.Duration(timeout) * time.Second,
			Transport: t,
		},
	}
}

func (hc *HttpClient) IsOK(url string) bool {
	statusCode, _, err := hc._request(http.MethodGet, url, nil, nil, true)
	if err != nil {
		// error log
		//fmt.Printf("%s - %d - %s", err, statusCode, url)
		return false
	}

	return statusCode == http.StatusOK
}

func (hc *HttpClient) IsNotFound(url string) bool {
	statusCode, _, err := hc._request(http.MethodGet, url, nil, nil, true)
	if err != nil {
		// error log
		return false
	}

	return statusCode == http.StatusNotFound
}

func (hc *HttpClient) Get(url string) (statusCode int, data []byte, err error) {
	return hc._get(url)
}

/*
func (hc *HttpClient) MGet(urls ...string) (statusCode int, data []byte, err error) {
	var wg sync.WaitGroup

	for _, u := range urls {

		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			//statusCode, data, err := hc._get(url)

		}(u)
	}

	wg.Wait()

	return
}*/

func (hc *HttpClient) Post(url string, payload []byte) (statusCode int, data []byte, err error) {
	return hc._postJson(url, payload)
}

func (hc *HttpClient) PostForm(url string, payload url.Values) (statusCode int, data []byte, err error) {
	return hc._postForm(url, payload)
}

func (hc *HttpClient) _get(url string) (statusCode int, data []byte, err error) {
	return hc._request(http.MethodGet, url, nil, nil, false)
}

func (hc *HttpClient) _postJson(url string, payload []byte) (statusCode int, data []byte, err error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	return hc._request(http.MethodPost, url, headers, bytes.NewBuffer(payload), false)
}

func (hc *HttpClient) _postForm(url string, payload url.Values) (statusCode int, data []byte, err error) {
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	return hc._request(http.MethodPost, url, headers, strings.NewReader(payload.Encode()), false)
}

func (hc *HttpClient) _request(method, url string, headers map[string]string, payload io.Reader, statusOnly bool) (statusCode int, data []byte, err error) {
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return 0, nil, err
	}

	// Add Header
	if len(headers) > 0 {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	resp, err := hc.Client.Do(req)
	if err != nil {
		return 0, nil, err
	}

	// Close the connection to reuse it
	defer resp.Body.Close()

	if !statusOnly {
		data, err = ioutil.ReadAll(resp.Body)
	}

	return resp.StatusCode, data, err
}

package builders

import (
	"encoding/xml"
	"errors"
	"fmt"
	ismuniapi "github.com/pestanko/go-is-muni-api"
	"io/ioutil"
	"net/http"
	"net/url"
)

type RequestBuilder struct {
	client *ismuniapi.IsApiClient
	req    *http.Request
	query  *url.Values
}

func (rqb *RequestBuilder) request() (*http.Request, error) {
	rqb.req.URL.RawQuery = rqb.query.Encode()
	return rqb.req, nil
}

func (rqb *RequestBuilder) Raw() (*http.Response, error) {
	req, err := rqb.request()
	if err != nil {
		return nil, err
	}

	return rqb.client.Client.Do(req)
}

func (rqb *RequestBuilder) ToString() (string, error) {
	resp, err := rqb.Raw()

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		bodyString := string(bodyBytes)
		return bodyString, nil
	} else {
		return "", errors.New(fmt.Sprintf("Status code [%d]", resp.StatusCode))
	}
}

func (rqb *RequestBuilder) ParseInto(data interface{}) error {
	document, err := rqb.ToString()

	if err != nil {
		return err
	}

	err = xml.Unmarshal([]byte(document), data)

	if err != nil {
		return err
	}
	return nil
}

func (rqb *RequestBuilder) AddQuery(name string, value string) {
	rqb.query.Add(name, value)
}

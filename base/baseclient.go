package base

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/tidwall/gjson"
)

type BaseClient struct {
	BasePath string
	Username string
	Password string
}

var defaultValues = BaseClient{
	BasePath: "https://aplikace.skolaonline.cz/SOLWebApi/api/v1",
}

func NewBaseClient(usr, pwd string) *BaseClient {
	client := defaultValues
	client.Username = usr
	client.Password = pwd
	return &client
}

func (client *BaseClient) makeAuthHeader() string {
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(client.Username+":"+client.Password))
}

func (client *BaseClient) newRequest(method, url string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", client.makeAuthHeader())
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

func (client *BaseClient) makeError(code, message string) error {
	return errors.New(code + ": " + message)
}

func (client *BaseClient) executeRequest(method, url string, body, target any) error {
	enc, err := EncodeBody(body)
	if err != nil {
		return err
	}

	req, err := client.newRequest(method, client.BasePath+url, enc)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	par := gjson.ParseBytes(raw)

	if res.StatusCode != 200 {
		stat := par.Get("Status")
		if stat.Exists() {
			return client.makeError(stat.Get("Code").String(), par.Get("Message").String())
		} else {
			return client.makeError(fmt.Sprint(res.StatusCode), res.Status)
		}
	}

	if target != nil {
		if err := json.Unmarshal([]byte(par.Get("Data").Raw), target); err != nil {
			return err
		}
	}

	return nil
}

func (client *BaseClient) Get(url string, target any) error {
	return client.executeRequest(http.MethodGet, url, nil, target)
}

func (client *BaseClient) Post(url string, body, target any) error {
	return client.executeRequest(http.MethodPost, url, body, target)
}

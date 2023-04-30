package rave

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func (r *Rave) Get(url string, params map[string]string, resp interface{}) (err error) {
	return r.call(http.MethodGet, url, params, nil, &resp)
}
func (r *Rave) Post(url string, params map[string]string, body, resp interface{}) (err error) {
	return r.call(http.MethodPost, url, params, body, &resp)
}

func (r *Rave) call(method, url string, params map[string]string, body, v interface{}) (err error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return
		}
	}
	if params != nil && len(params) > 0 {
		addToUrl := "?"
		for k, val := range params {
			addToUrl += fmt.Sprintf("%s=%s&", k, val)
		}
		url += addToUrl
	}
	var req *http.Request
	req, err = http.NewRequest(method, url, buf)
	if err != nil {
		if r.EnableLogging {
			log.Printf("Cannot create Flutterwave request: %v\n", err)
		}
		return err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Authorization", "Bearer "+r.SecretKey)

	if r.EnableLogging {
		log.Printf("Requesting %v %v%v\n", req.Method, req.URL.Host, req.URL.Path)
		log.Printf("POST request data %v\n", buf)
	}

	start := time.Now()

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if r.EnableLogging {
		log.Printf("Completed in %v\n", time.Since(start))
	}

	defer resp.Body.Close()
	return r.decodeResponse(resp, v)

}

func (r *Rave) decodeResponse(httpResp *http.Response, v interface{}) error {
	var resp map[string]interface{}
	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return err
	}

	if status, _ := resp["status"].(string); status != "success" || httpResp.StatusCode >= 400 {
		err = errors.New("Unkown error")
		if fwErr, ok := resp["message"].(string); ok {
			err = errors.New(fwErr)
		}
		if r.EnableLogging {
			log.Printf("Flutterwave error: %+v", err)
			log.Printf("HTTP Response: %+v", resp)
		}
		return err
	}

	if r.EnableLogging {
		log.Printf("Flutterwave response: %v\n", resp)
	}

	if data, ok := resp["data"]; ok {
		switch t := resp["data"].(type) {
		case map[string]interface{}:
			return mapstruct(data, v)
		default:
			_ = t
			return mapstruct(resp, v)
		}
	}
	// if response data does not contain data key, map entire response to v
	return mapstruct(resp, v)
}

func mapstruct(data interface{}, v interface{}) error {
	config := &mapstructure.DecoderConfig{
		Result:           v,
		TagName:          "json",
		WeaklyTypedInput: true,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	err = decoder.Decode(data)
	return err
}

//type ErrorBody struct {
//	Status  string      `json:"status"`
//	Message string      `json:"message"`
//	Data    interface{} `json:"data"`
//}

package utils

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

type RequestOption struct {
	Timeout            int
	InsecureSkipVerify bool
}

// Get method use fasthttp
func Get(uri string, header *map[string]string, body []byte, op *RequestOption) ([]byte, int, error) {
	return request(uri, header, body, fasthttp.MethodGet, op)
}

// Post method use fasthttp
func Post(uri string, header *map[string]string, body []byte, op *RequestOption) ([]byte, int, error) {
	return request(uri, header, body, fasthttp.MethodPost, op)
}

// Get method use golang http
func GetHttp(tranID, uri string, header *map[string]string, op *RequestOption) (rawBody []byte, statusCode int, respHeader http.Header, err error) {
	return callHttp(tranID, uri, header, nil, http.MethodGet, op)
}

// Post method use golang http
func PostHttp(tranID, uri string, header *map[string]string, body []byte, op *RequestOption) (rawBody []byte, statusCode int, respHeader http.Header, err error) {
	return callHttp(tranID, uri, header, body, http.MethodPost, op)
}

func callHttp(tranID, uri string, header *map[string]string, body []byte, method string, op *RequestOption) (rawBody []byte, statusCode int, respHeader http.Header, err error) {
	// startProcess := time.Now()
	client := http.Client{
		Timeout: time.Duration(op.Timeout) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: op.InsecureSkipVerify,
			},
			DisableKeepAlives: true,
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}
	buff := bytes.NewBuffer(body)
	req, err := http.NewRequest(method, uri, buff)
	// s := newrelic.StartExternalSegment(txn, req)
	// defer s.End()
	if err != nil {
		// s.AddAttribute("errMsg", err.Error())
		return
	}
	if len(*header) > 0 {
		for key, value := range *header {
			req.Header.Add(key, value)
		}
	}

	if (*header)["Content-Type"] != "" {
		req.Header.Add("Content-Type", (*header)["Content-Type"])
	} else {
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := client.Do(req)
	// s.AddAttribute("tranId", tranID)
	// s.Response = resp

	if err != nil {
		// s.AddAttribute("errMsg", err.Error())
		return
	}
	respHeader = resp.Header
	statusCode = resp.StatusCode
	defer resp.Body.Close()
	rawBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		// s.AddAttribute("errMsg", err.Error())
		return
	}
	// log.Debug("End Call Restful API :: elapsed %v ms | statusCode : %d", log.GetElapsedTime(startProcess), statusCode)
	return
}

func request(uri string, header *map[string]string, body []byte, method string, op *RequestOption) ([]byte, int, error) {
	var poolClient sync.Pool
	var err error
	poolClient.New = func() interface{} { return &fasthttp.Client{} }
	client := poolClient.Get().(*fasthttp.Client)
	timeout := time.Second * time.Duration(30)

	client = &fasthttp.Client{}
	defer client.CloseIdleConnections()

	defer func() {
		poolClient.Put(client)
		uri = ""
		header = nil
		body = nil
		method = ""
		op = nil
		err = nil
		timeout = 0
	}()

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	defer req.ResetBody()

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	defer resp.ResetBody()

	req.SetRequestURI(uri)

	if method == fasthttp.MethodPost {
		req.Header.SetMethod(fasthttp.MethodPost)
		req.SetBody(body)
	} else if method == fasthttp.MethodPut {
		req.Header.SetMethod(fasthttp.MethodPut)
		req.SetBody(body)
	} else if method == fasthttp.MethodPatch {
		req.Header.SetMethod(fasthttp.MethodPatch)
		req.SetBody(body)
	} else if method == fasthttp.MethodDelete {
		req.Header.SetMethod(fasthttp.MethodDelete)
		req.SetBody(body)
	} else if method == fasthttp.MethodOptions {
		req.Header.SetMethod(fasthttp.MethodOptions)
	} else if method == fasthttp.MethodHead {
		req.Header.SetMethod(fasthttp.MethodHead)
	} else {
		req.Header.SetMethod(fasthttp.MethodGet)
		if len(body) > 0 {
			req.SetBody(body)
		}
	}

	if op != nil {
		if op.Timeout != 0 {
			timeout = time.Second * time.Duration(op.Timeout)
			client.Dial = func(addr string) (net.Conn, error) {
				return fasthttp.DialTimeout(addr, timeout)
			}
		}
		if op.InsecureSkipVerify {
			client.TLSConfig = &tls.Config{
				InsecureSkipVerify: true,
			}
		}
	}

	req.Header.Set("Content-Type", "application/json")
	if header != nil {
		for k, v := range *header {
			req.Header.Set(k, v)
		}
	}
	err = client.DoTimeout(req, resp, timeout)
	if err != nil {
		if resp != nil {
			return resp.Body(), resp.StatusCode(), err
		}
		return nil, 500, err
	}
	return resp.Body(), resp.StatusCode(), nil
}

package utilx

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func getHttpClient() *http.Client {
	var netTransport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 30 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{
		Timeout:   time.Second * 35,
		Transport: netTransport,
	}
}

func HttpGet(url string) (code int, body []byte, err error) {
	client := getHttpClient()
	resp, e := client.Get(url)
	if e != nil {
		err = e
		return
	}

	code = resp.StatusCode
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return
	}
	if err = resp.Body.Close(); err != nil {
		return
	}

	return
}

func HttpPost(url, contentType string, data []byte) ([]byte, error) {
	client := getHttpClient()
	resp, err := client.Post(url, contentType, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return nil, err
	}
	if err := resp.Body.Close(); err != nil {
		return nil, err
	}

	return body, err
}

// POST 请求体
func HttpPostData(targetUrl string, header map[string]string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", targetUrl, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return []byte(""), err
	}
	req.Header.Add("Content-type", "application/json;charset=UTF-8")
	for k, v := range header {
		req.Header.Set(k, v)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		Dial: func(netw, addr string) (net.Conn, error) {
			conn, err := net.DialTimeout(netw, addr, time.Second*30)
			if err != nil {
				return nil, err
			}
			conn.SetDeadline(time.Now().Add(time.Second * 30))
			return conn, nil
		},
		ResponseHeaderTimeout: time.Second * 30,
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)
	if err != nil {
		return []byte(""), err
	}

	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return respData, nil
}

// HTTPGet HttpGet
func HTTPGet(urlStr string) ([]byte, error) {

	var (
		err    error
		req    *http.Request
		resp   *http.Response
		client *http.Client
		data   []byte
	)

	//roots := x509.NewCertPool()
	//roots.AppendCertsFromPEM([]byte(rootPEM))
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{RootCAs: roots},
	//}
	//client := &http.Client{Transport: tr}

	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*2)
				if err != nil {
					return nil, err
				}
				if err = conn.SetDeadline(time.Now().Add(time.Second * 2)); err != nil {

				}
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * 2,
		},
	}

	if req, err = http.NewRequest("GET", urlStr, nil); err != nil {
		return data, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;")
	if resp, err = client.Do(req); err != nil {
		return data, err
	}

	if data, err = ioutil.ReadAll(resp.Body); err != nil {
		if err = resp.Body.Close(); err != nil {

		}
		return data, err
	}
	resp.Body.Close()
	return data, err
}

func PostData(targetUrl string, header map[string]string, body []byte) ([]byte, error) {

	body = bytes.Replace(body, []byte("\\u003c"), []byte("<"), -1)
	body = bytes.Replace(body, []byte("\\u003e"), []byte(">"), -1)
	body = bytes.Replace(body, []byte("\\u0026"), []byte("&"), -1)

	req, err := http.NewRequest("POST", targetUrl, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return []byte(""), err
	}
	req.Header.Add("Content-type", "application/json;charset=UTF-8")
	for k, v := range header {
		req.Header.Set(k, v)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Dial: func(netw, addr string) (net.Conn, error) {
			conn, err := net.DialTimeout(netw, addr, time.Second*30)
			if err != nil {
				return nil, err
			}
			conn.SetDeadline(time.Now().Add(time.Second * 30))
			return conn, nil
		},
		ResponseHeaderTimeout: time.Second * 30,
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)
	if err != nil {
		return []byte(""), err
	}

	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return respData, nil
}

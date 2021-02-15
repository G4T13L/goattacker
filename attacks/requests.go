package attacks

/*
3 maneras de usar
[-] auth
[ ] prueba de directorios
[ ] form login
*/

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

//RequestData Basic data for requests
type RequestData struct {
	User      string
	Pass      string
	UserAgent string
	Proxy     string
	Client    *http.Client
}

//NewRequestData Creates a requests data for http attacks
func NewRequestData() *RequestData {
	r := new(RequestData)
	r.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36"
	r.Client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	return r
}

//SetProxy set proxy the data
func (r *RequestData) SetProxy(urlPort string) {
	proxy, err := url.Parse(urlPort)
	if err != nil {
		log.Fatal("Error parsing url")
	}
	r.Client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		Proxy: http.ProxyURL(proxy),
	}
}

//GetOrPost request for get or post method
func (r *RequestData) GetOrPost(url string, post string) (string, int, *http.Response) {
	if post == "" {
		return r.Send("GET", url, "")
	} else {
		return r.Send("POST", url, post)
	}
}

//Send Send the data
func (r *RequestData) Send(method, url, post string) (string, int, *http.Response) {
	req, err := http.NewRequest(method, url, strings.NewReader(post))

	if err != nil {
		fmt.Println("[!] Didn't get a response from the server")
		return "", 0, nil
	}

	if r.User != "" {
		req.SetBasicAuth(r.User, r.Pass)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	req.Header.Set("User-Agent", r.UserAgent)
	req.Header.Set("Accept-Encoding", "*")

	resp, err := r.Client.Do(req)
	if err != nil || resp == nil {
		return "", 0, nil
	}

	code := resp.StatusCode
	if resp.Body == nil {
		return "", code, resp
	}

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", code, resp
	}
	resp.Body.Close()

	return string(html), code, resp
}

//SendAuth Send Auth request
func SendAuth(url, post, username, password, proxy string) (string, int, *http.Response) {
	r := NewRequestData()
	if proxy != "" {
		r.SetProxy(proxy)
	}
	r.User = username
	r.Pass = password
	html, code, resp := r.GetOrPost(url, post)
	// fmt.Println(html, code, resp)
	return html, code, resp
	// fmt.Printf("%v",r.Client.Transport)
}

//FileTry Send a request to verify a file if it exist
func FileTry(url, word, fileWord, ext, post, proxy string) (string, int) {
	r := NewRequestData()
	if proxy != "" {
		r.SetProxy(proxy)
	}
	url = strings.Replace(url, word, fileWord+ext, 1)
	html, code, _ := r.GetOrPost(url, post)
	fmt.Println(html, code, url)
	return html, code
}

//FormLogin Send a request of a form to log in
func FormLogin(url, post, phrase, proxy string) (string, int, bool) {
	r := NewRequestData()
	if proxy != "" {
		r.SetProxy(proxy)
	}
	// html, code, resp := r.GetOrPost(url, post)
	html, code, _ := r.GetOrPost(url, post)
	return "_", code, strings.Contains(html, phrase)
}
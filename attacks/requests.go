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
	"net/http/cookiejar"
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
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
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Fatal(err)
	}
	r.Client.Jar = jar
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
func (r *RequestData) GetOrPost(urlsite string, post string, redirect bool) (string, int, *http.Response) {
	if post == "" {
		return r.Send("GET", urlsite, "", redirect)
	} else {
		return r.Send("POST", urlsite, post, redirect)
	}
}

//Send Send the data
func (r *RequestData) Send(method, urlsite, post string, redirect bool) (string, int, *http.Response) {
	var req *http.Request
	var err error
	data := url.Values{}
	if method == "GET" {
		req, err = http.NewRequest(method, urlsite, nil)
	} else {
		// sep := strings.Split(post, "&")
		for _, v := range strings.Split(post, "&") {
			w := strings.Split(v, "=")
			data.Set(w[0], w[1])
		}
		req, err = http.NewRequest("POST", urlsite, strings.NewReader(data.Encode()))

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
		req.Header.Set("User-Agent", r.UserAgent)
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	}

	if err != nil {
		fmt.Println("[!] Didn't get a response from the server")
		return "", 0, nil
	}

	if r.User != "" {
		req.SetBasicAuth(r.User, r.Pass)
	}

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

	// fmt.Println(magenta(resp.Header))
	if (code != 302 && code != 301) || !redirect {
		return string(html), code, resp
	} else {
		ret := resp.Header.Get("Location")
		if strings.HasPrefix(ret, "http://") || strings.HasPrefix(ret, "https://") {
			return r.Send("GET", ret, "", false)
		}
		fmt.Println(extractDomain(urlsite) + resp.Header.Get("Location"))
		return r.Send("GET", extractDomain(urlsite)+ret, "", false)
	}
}

func extractDomain(url string) string {
	dn := strings.Split(url, "/")
	return dn[0] + "//" + dn[2]
}

//FileTry Send a request to verify a file if it exist
func FileTry(urlsite, word, fileWord, ext, post, proxy string, redirect bool) (string, int, string) {
	r := NewRequestData()
	if proxy != "" {
		r.SetProxy(proxy)
	}
	urlsite = strings.Replace(urlsite, word, fileWord+ext, 1)
	html, code, _ := r.GetOrPost(urlsite, post, redirect)
	return html, code, urlsite
}

//SendAuth Send Auth request
func SendAuth(urlsite, post, username, password, proxy string, redirect bool) (string, int) { //, *http.Response) {
	r := NewRequestData()
	if proxy != "" {
		r.SetProxy(proxy)
	}
	r.User = username
	r.Pass = password
	html, code, _ := r.GetOrPost(urlsite, post, redirect)
	return html, code
}

//FormLogin Send a request of a form to log in
func FormLogin(urlsite, post, phrase, proxy string, redirect bool) (string, int, bool) {
	r := NewRequestData()
	if proxy != "" {
		r.SetProxy(proxy)
	}
	// fmt.Println(urlsite, post, phrase, proxy)
	html, code, _ := r.GetOrPost(urlsite, post, redirect)
	// fmt.Println(html)
	return html, code, !strings.Contains(html, phrase)
}

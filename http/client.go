package http

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//	client.header = map[string][]string{
//		"X-Requested-With": {"XMLHttpRequest"},
//		"User-Agent":       {"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"},
//		"Origin":           {"https://xueqiu.com"},
//		"Host":             {"xueqiu.com"},
//		"Content-Type":     {"application/x-www-form-urlencoded; charset=UTF-8"},
//	}

type HttpClient struct {
	Client  *http.Client
	cookies []*http.Cookie
	header  map[string][]string
}

func (httpClient *HttpClient) Get(url string) (result string) {

	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		return ""
	}

	//	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//	req.Header.Set("Cookie", "name=anny")
	//	for _, v := range httpClient.cookies {
	//		req.AddCookie(v)
	//	}
	resp, err := httpClient.Client.Do(req)
	defer resp.Body.Close()
	httpClient.cookies = resp.Cookies()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(body)
}
func (httpClient *HttpClient) Analysis(url string) (result string) {

	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		return ""
	}
	for _, v := range httpClient.cookies {
		req.AddCookie(v)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Origin", "https://xueqiu.com")
	req.Header.Add("Host", "xueqiu.com")

	resp, err := httpClient.Client.Do(req)
	defer resp.Body.Close()
	httpClient.cookies = resp.Cookies()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(body)
}
func (httpClient *HttpClient) FetchStock(url string) (result string) {

	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		return ""
	}
	for _, v := range httpClient.cookies {
		req.AddCookie(v)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Origin", "https://xueqiu.com")
	req.Header.Add("Host", "xueqiu.com")

	resp, err := httpClient.Client.Do(req)
	defer resp.Body.Close()
	//	httpClient.cookies = resp.Cookies()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(body)
}
func (httpClient *HttpClient) Login(urlStr string, params map[string]string) (result string) {

	postParams := url.Values{}
	for k, v := range params {
		postParams.Set(k, v)
	}

	postBody := strings.NewReader(postParams.Encode())
	//	fmt.Println(postBody)

	req, err := http.NewRequest("POST", urlStr, postBody)
	//	req, err := http.NewRequest("POST", urlStr, strings.NewReader(post_arg.Encode()))
	if err != nil {
		return ""
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Origin", "https://xueqiu.com")
	req.Header.Add("Host", "xueqiu.com")

	//	req.Header = httpClient.header

	resp, err := httpClient.Client.Do(req)
	defer resp.Body.Close()

	httpClient.cookies = resp.Cookies()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(body)
}

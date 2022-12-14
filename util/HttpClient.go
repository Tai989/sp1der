package util

import (
	"io"
	"net/http"
	"net/url"
)

func DefaultHttpClient() *http.Client {
	proxy, err := url.Parse("http://127.0.0.1:1087")
	if nil != err {
		panic(err)
	}
	tr := &http.Transport{
		DisableKeepAlives:  false,
		DisableCompression: true,
		Proxy:              http.ProxyURL(proxy),
	}
	client := &http.Client{
		Transport: tr,
	}
	return client
}
func HttpGet(client *http.Client, url string) *[]byte {
	r, err := http.NewRequest(http.MethodGet, url, nil)
	r.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	r.Header.Set("Referer", "https://m.weather.com.cn/")
	r.Header.Set("Content-Type", "text/html")
	if nil != err {
		panic(err)
	}
	res, err := client.Do(r)
	defer res.Body.Close()
	if nil != err {
		panic(err)
	}
	b, err := io.ReadAll(res.Body)
	if nil != err {
		panic(err)
	}
	return &b
}

package proxy

import (
	"fmt"
	"html"
	"io"
	"net/http"
	"strings"

	"github.com/melvinodsa/build-with-golang/api-gateway/config"
)

type Proxy struct {
	Name       string
	PrefixPath string
	Host       string
	Port       string
}

var proxies = []Proxy{
	{
		Name:       "User Service",
		PrefixPath: "/user-service",
		Host:       "http://127.0.0.1",
		Port:       "3000",
	},
}

func InitProxies(ctx *config.AppContext) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPath := html.EscapeString(r.URL.Path)
		for _, proxy := range proxies {
			if !strings.HasPrefix(urlPath, proxy.PrefixPath) {
				continue
			}
			newPath := strings.Replace(urlPath, proxy.PrefixPath, "", 1)
			newUrl := fmt.Sprintf("%s:%s%s", proxy.Host, proxy.Port, newPath)
			status, err := handleProxy(ctx, w, r, newUrl)
			if err != nil {
				ctx.Logger.Error(err)
				http.Error(w, err.Error(), status)
				return
			}
			ctx.Logger.Info(fmt.Sprintf("%d %s %s %s %s:%s", status, r.Method, newUrl, proxy.Name, proxy.Host, proxy.Port))
			return
		}
		http.Error(w, urlPath, http.StatusNotFound)
	})
}

func handleProxy(ctx *config.AppContext, w http.ResponseWriter, r *http.Request, urlPath string) (int, error) {

	// Create a new request to the proxy server
	req, err := http.NewRequest(r.Method, urlPath, r.Body)
	if err != nil {
		return http.StatusBadGateway, err
	}

	// Add headers from the original request
	for k, v := range r.Header {
		req.Header[k] = v
	}

	// Send the new request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return http.StatusBadGateway, err
	}

	// Write the response
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return http.StatusBadGateway, err
	}

	return resp.StatusCode, nil
}

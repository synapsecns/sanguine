package localserver

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// startUIServer starts a local jaeger server for testing.
// this will first attempt to route to the jaeger ui container and fallback to jaegerui.
func (j *testJaeger) startUIServer() {
	primaryUrl := "http://primary.example.com"
	secondaryUrl := "http://secondary.example.com"

	primaryProxy := httputil.NewSingleHostReverseProxy(parseUrl(primaryUrl))
	secondaryProxy := httputil.NewSingleHostReverseProxy(parseUrl(secondaryUrl))

	http.HandleFunc("*", func(w http.ResponseWriter, r *http.Request) {
		primaryProxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
			fmt.Printf("Primary proxy error: %s\n", err.Error())
			secondaryProxy.ServeHTTP(w, r)
		}
		primaryProxy.ServeHTTP(w, r)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func parseUrl(rawUrl string) *url.URL {
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		panic(err)
	}
	return parsedUrl
}

package ginhelper_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

// nolint: unparam
func performRequest(ctx context.Context, r http.Handler, method, contentType string, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequestWithContext(ctx, method, path, body)
	req.Header.Set("Content-Type", contentType)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestMIMEJSON(t *testing.T) {
	router := gin.New()
	router.Use(ginhelper.Dump())

	router.POST("/dump", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"data": "gin-dump",
		})
	})

	type params struct {
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
	}

	var httpdata = params{
		StartTime: "2019-03-03",
		EndTime:   "2019-03-03",
	}
	b, err := json.Marshal(httpdata)
	if err != nil {
		fmt.Println("json format error:", err)
		return
	}

	body := bytes.NewBuffer(b)

	_ = performRequest(context.Background(), router, "POST", gin.MIMEJSON, "/dump", body)
}
func TestMIMEJSONWithOption(t *testing.T) {
	router := gin.New()
	router.Use(ginhelper.DumpWithOptions(true, false, true, true, false, func(dumpStr string) {
		fmt.Println(dumpStr)
	}))

	router.POST("/dump", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"data": "gin-dump",
		})
	})

	type params struct {
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
	}

	var httpdata = params{
		StartTime: "2019-03-03",
		EndTime:   "2019-03-03",
	}
	b, err := json.Marshal(httpdata)
	if err != nil {
		fmt.Println("json format error:", err)
		return
	}

	body := bytes.NewBuffer(b)
	_ = performRequest(context.Background(), router, "POST", gin.MIMEJSON, "/dump", body)
}

func TestMIMEPOSTFORM(t *testing.T) {
	router := gin.New()
	router.Use(ginhelper.Dump())

	router.POST("/dump", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
			"data": map[string]interface{}{
				"name": "jfise",
				"addr": "tpkeeper@qq.com",
			},
		})
	})

	form := make(url.Values)
	form.Set("foo", "bar")
	form.Add("foo", "bar2")
	form.Set("bar", "baz")

	body := strings.NewReader(form.Encode())
	_ = performRequest(context.Background(), router, "POST", gin.MIMEPOSTForm, "/dump", body)
}

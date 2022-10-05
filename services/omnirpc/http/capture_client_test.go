package http_test

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/omnirpc/http"
	"github.com/synapsecns/sanguine/services/omnirpc/http/mocks"
	"testing"
)

func TestCaptureClient(t *testing.T) {
	testRes := gofakeit.ImageJpeg(50, 50)
	testBody := gofakeit.ImageJpeg(50, 50)

	client := http.NewCaptureClient(func(c *http.CapturedRequest) (http.Response, error) {
		bodyRes := new(mocks.Response)
		bodyRes.On("Body").Return(testRes)

		return bodyRes, nil
	})

	testCtx := context.Background()

	byteHeaderK := []byte(gofakeit.Word())
	byteHeaderV := []byte(gofakeit.Sentence(4))

	strHeaderK := gofakeit.Word()
	strHeaderV := gofakeit.Sentence(4)

	testURL := gofakeit.URL()

	testReq := client.NewRequest()
	testReq.SetBody(testBody)
	testReq.SetContext(testCtx)
	testReq.SetHeaderBytes(byteHeaderK, byteHeaderV)
	testReq.SetHeader(strHeaderK, strHeaderV)
	testReq.SetRequestURI(testURL)

	resp, err := testReq.Do()
	Nil(t, err)

	Equal(t, resp.Body(), testRes)

	for _, request := range client.Requests() {
		Equal(t, request.Body, testBody)
		Equal(t, request.Context, testCtx)
		Equal(t, request.StringHeaders[strHeaderK], strHeaderV)
		byteHeaderRet, _ := request.ByteHeaders.Get(byteHeaderK)

		Equal(t, byteHeaderRet, byteHeaderV)
		Equal(t, request.RequestURI, testURL)
	}
}

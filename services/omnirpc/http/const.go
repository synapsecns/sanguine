package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-http-utils/headers"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"net/http"
)

// Headers:.
var (
	// XForwardedFor is a byte encoded forwarded for header.
	XForwardedFor = []byte(headers.XForwardedFor)
	// ContentType is a byte encoded content type.
	ContentType = []byte(headers.ContentType)
	// Accept is a byte encoded accept header.
	Accept = []byte(headers.Accept)
	// XRequestIDString is the string request id header.
	XRequestIDString = ginhelper.RequestIDHeader
	// XRequestID is the byte encoded request id.
	XRequestID = []byte(XRequestIDString)
	// Encoding is a bytes encoded Accept-Encoding header.
	Encoding = []byte(headers.AcceptEncoding)
)

// Mime types.
var (
	// JSONType is a byte encoded json type.
	JSONType = []byte(gin.MIMEJSON)
	// EncodingTypes are encoding headers.
	EncodingTypes = []byte("gzip, br, deflate")
)

// Method types.
var (
	// PostType is used for posting.
	PostType = []byte(http.MethodPost)
)

// Constant Strings.
var (
	// OmniRPCValue is a byte encoded omnirpc string.
	OmniRPCValue = []byte("omnirpc")
)

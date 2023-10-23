package interceptor

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/synapsecns/sanguine/core/dbcommon"
)

type sqlSanitizerImpl struct {
}

func (s sqlSanitizerImpl) ExtensionName() string {
	return "sqlsanitizer"
}

func (s sqlSanitizerImpl) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

// SqlSanitizer is a middleware that sanitizes SQL queries.
type SqlSanitizer interface {
	graphql.ResponseInterceptor
	graphql.HandlerExtension
}

// InterceptResponse intercepts the incoming request.
func (s sqlSanitizerImpl) InterceptResponse(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	if !graphql.HasOperationContext(ctx) {
		return next(ctx)
	}

	oc := graphql.GetOperationContext(ctx)
	// key validation is handled by schema validator already
	for key, val := range oc.Variables {
		switch marshalledVal := val.(type) {
		case string:
			oc.Variables[key] = dbcommon.MysqlRealEscapeString(marshalledVal)
		}

	}
	graphql.WithOperationContext(ctx, oc)

	return next(ctx)
}

// SqlSanitizerMiddleware returns a new SqlSanitizer middleware. This is a workaround
// until we have better support for paramaterized queries.
func SqlSanitizerMiddleware() SqlSanitizer {
	return sqlSanitizerImpl{}
}

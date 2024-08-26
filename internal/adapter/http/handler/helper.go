package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tommjj/go-blog-api/internal/core/domain"
)

// getAuthPayload is a helper function to get the auth payload from the context
func getAuthPayload(ctx *gin.Context, key string) *domain.TokenPayload {
	return ctx.MustGet(key).(*domain.TokenPayload)
}

// responseWithMeta is a helper function to add meta to response
func responseWithMeta(m meta, key string, value any) map[string]any {
	return map[string]any{
		"meta": m,
		key:    value,
	}
}

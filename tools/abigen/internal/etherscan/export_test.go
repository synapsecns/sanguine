package etherscan

import (
	"context"
	"time"
)

// FileRateLimiter is the file rate limiter exported for testing.
type FileRateLimiter struct {
	*fileRateLimiter
}

// NewFileRateLimiter creates a new file rate limiter.
func NewFileRateLimiter(parentCtx context.Context, filePath string, waitBetweenRequests time.Duration) (*FileRateLimiter, error) {
	fileRateLimiter, err := newFileRateLimiter(parentCtx, filePath, waitBetweenRequests)
	if err != nil {
		return nil, err
	}
	return &FileRateLimiter{fileRateLimiter: fileRateLimiter}, nil
}

// ObtainLock exports obtainLock for testing.
func (f *FileRateLimiter) ObtainLock(ctx context.Context) (ok bool, err error) {
	return f.obtainLock(ctx)
}

func (f *FileRateLimiter) ReleaseLock() (ok bool, err error) {
	return f.releaseLock()
}

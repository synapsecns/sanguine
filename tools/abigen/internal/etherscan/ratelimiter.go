package etherscan

import (
	"context"
	"fmt"
	"github.com/gofrs/flock"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

// fileRateLimiter implements a file based rate limiter for limiting requests across abi generate commands.
type fileRateLimiter struct {
	// lockFolderPath is the path to the rate limiter
	lockFolderPath string
	// flock is the file locker for recording the last request. This is locked per process to prevent two from running at once
	flock *flock.Flock
	// waitBetweenRequests is how long to wait before requests
	waitBetweenRequests time.Duration
	// lastRequestFile is the last request file
	lastRequestFile *os.File
	// mux prevents duplicate locking
	mux sync.Mutex
}

// fileRateTimeout is how long to wait for file rate limiting.
const fileRateTimeout = time.Second * 20

// lastRequestFile is the last request.
const lastRequestFile = "/last_request.txt"

// newFileRateLimiter creates a new file rate limiter.
func newFileRateLimiter(parentCtx context.Context, lockFolderPath string, waitBetweenRequests time.Duration) (*fileRateLimiter, error) {
	ctx, cancel := context.WithTimeout(parentCtx, fileRateTimeout)
	defer cancel()

	_ = os.MkdirAll(lockFolderPath, os.ModePerm)

	frl := fileRateLimiter{
		lockFolderPath:      lockFolderPath,
		waitBetweenRequests: waitBetweenRequests,
		flock:               flock.New(filepath.Join(lockFolderPath, "locker")),
	}

	_, err := frl.flock.TryLockContext(ctx, time.Second)
	if err != nil {
		return nil, fmt.Errorf("could not obtain lock (timeout %s): %w", waitBetweenRequests, err)
	}

	//nolint: gosec
	err = frl.openFile()
	if err != nil {
		return nil, fmt.Errorf("could not create last request file: %w", err)
	}

	return &frl, nil
}

func (f *fileRateLimiter) openFile() (err error) {
	_ = f.lastRequestFile.Close()
	f.lastRequestFile, err = os.OpenFile(filepath.Join(f.lockFolderPath, lastRequestFile), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return fmt.Errorf("could not create last request file: %w", err)
	}
	return nil
}

// obtainLock locks until the time since last request is greater than waitBetweenRequests.
func (f *fileRateLimiter) obtainLock(ctx context.Context) (ok bool, err error) {
	f.mux.Lock()
	err = f.openFile()
	if err != nil {
		return false, fmt.Errorf("could not open file: %w", err)
	}

	fileContents, err := io.ReadAll(f.lastRequestFile)
	if err != nil {
		return false, fmt.Errorf("could not obtain file contents: %w", err)
	}

	var unixTimestamp int
	if len(fileContents) == 0 {
		unixTimestamp = 0
	} else {
		unixTimestamp, err = strconv.Atoi(string(fileContents))
		if err != nil {
			return false, fmt.Errorf("could not parse unix timestamp: %w", err)
		}
	}

	lastRequest := time.Unix(int64(unixTimestamp), 0)

	// waitPeriod is how long to wait before obtaining the lock
	waitPeriod := time.Until(lastRequest.Add(f.waitBetweenRequests))

	select {
	case <-ctx.Done():
		return false, context.Canceled
	case <-time.After(waitPeriod):
		return true, nil
	}
}

// releaseLock releases the lock.
func (f *fileRateLimiter) releaseLock() (ok bool, err error) {
	f.mux.Unlock()

	err = f.openFile()
	if err != nil {
		return false, fmt.Errorf("could not open file: %w", err)
	}

	err = f.lastRequestFile.Truncate(0)
	if err != nil {
		return false, fmt.Errorf("could not truncate file: %w", err)
	}

	_, err = f.lastRequestFile.Seek(0, 0)
	if err != nil {
		return false, fmt.Errorf("could not release lock: %w", err)
	}

	_, err = f.lastRequestFile.WriteString(strconv.Itoa(int(time.Now().Unix())))
	if err != nil {
		return false, fmt.Errorf("could not write timestamp: %w", err)
	}

	err = f.lastRequestFile.Close()
	if err != nil {
		return false, fmt.Errorf("could not write timestamp: %w", err)
	}

	return true, nil
}

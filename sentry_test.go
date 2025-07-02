/*
 * @Author: nijineko
 * @Date: 2025-06-11 12:04:18
 * @LastEditTime: 2025-06-11 12:35:18
 * @LastEditors: nijineko
 * @Description: noa sentry handler
 * @FilePath: \noa-sentry\sentry_test.go
 */
package noasentry

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/noa-log/noa"
)

func TestCaptureHandler(t *testing.T) {
	Dsn := os.Getenv("SENTRY_DSN")

	err := sentry.Init(sentry.ClientOptions{
		Dsn:   Dsn,
		Debug: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer sentry.Flush(2 * time.Second)

	Log := noa.NewLog()
	// Add handler
	Log.AddAfterHandle(CaptureHandler)

	TestErr := errors.New("test error")
	Log.Error("Test", TestErr)

	time.Sleep(2 * time.Second)
}

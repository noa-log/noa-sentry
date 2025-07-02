/*
 * @Author: nijineko
 * @Date: 2025-06-11 12:04:18
 * @LastEditTime: 2025-06-11 12:12:56
 * @LastEditors: nijineko
 * @Description: noa sentry handler
 * @FilePath: \noa-sentry\sentry.go
 */
package noasentry

import (
	"github.com/getsentry/sentry-go"
	"github.com/noa-log/noa"
	"github.com/noa-log/noa/errors"
)

/**
 * @description: error capture handler
 * @param {int} Level
 * @param {string} Source
 * @param {...any} Data
 * @return {error} error
 */
func CaptureHandler(Level int, Source string, Data ...any) error {
	if Level == noa.ERROR || Level == noa.FATAL {
		for _, Value := range Data {
			if ValueError, ok := Value.(errors.Error); ok {
				sentry.CaptureException(ValueError)
			}
			if ValueError, ok := Value.(error); ok {
				sentry.CaptureException(ValueError)
			}
		}
	}

	return nil
}

# Noa Sentry
Noa Log 的 Sentry 集成模块，可以快读的将 Noa 集成到 Sentry 中。在打印错误时自动将错误发送到 Sentry。

## 安装
```bash
go get -u github.com/noa-log/noa-sentry
```

## 快速开始
```go
package main

import (
    "errors"
    "os"
    "time"

    "github.com/getsentry/sentry-go"
    "github.com/noa-log/noa"
    "github.com/noa-log/noa-sentry"
)

func main() {
    // 从环境变量中获取Dsn
    Dsn := os.Getenv("SENTRY_DSN")

    // 初始化 Sentry
    err := sentry.Init(sentry.ClientOptions{
		Dsn:   Dsn,
		Debug: true,
	})
	if err != nil {
		panic(err)
	}
	defer sentry.Flush(2 * time.Second)

    // 创建一个新的日志实例
    logger := noa.NewLog()

    // 注册 Sentry 中间件
    logger.AddAfterHandle(noasentry.CaptureHandler)

	// 打印错误
	err = errors.New("一个错误示例")
	logger.Error("Test", err)
}
```

## 许可
本项目基于[Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0)协议开源。使用时请遵守协议的条款。
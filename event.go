package main

import (
	"blog/models"

	"github.com/kdada/tinygo/web"
)

// Processor 事件
type Event struct {
	web.DefaultHttpProcessorEvent
}

// 每次出现一个新请求的时候触发
func (this *Event) Request(processor *web.HttpProcessor, context *web.Context) {
	processor.Logger.Debug(context.Segments())
}

// 出现错误时触发
func (this *Event) Error(processor *web.HttpProcessor, context *web.Context, err error) {
	if context != nil {
		var err = context.WriteResult(context.Api(models.NewErrorResult(err)))
		if err != nil {
			processor.Logger.Error(err)
		}
	}
	processor.Logger.Error(err)
}

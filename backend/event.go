package backend

import "github.com/kdada/tinygo/web"

type ProcessorEvent struct {
	web.DefaultHttpProcessorEvent
}

// Request 每次出现一个新请求的时候触发,返回值决定是否处理该请求
func (this *ProcessorEvent) Request(processor *web.HttpProcessor, context *web.Context) bool {
	processor.Logger.Info(context.HttpContext.Request.RemoteAddr, context.HttpContext.Request.URL)
	return true
}

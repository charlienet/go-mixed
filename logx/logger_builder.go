package logx

type loggerBuilder struct {
}

func NewBuilder() *loggerBuilder {
	return &loggerBuilder{}
}

func (b *loggerBuilder) WithLogrus() *loggerBuilder {
	return b
}

func (b *loggerBuilder) WithLogger() *loggerBuilder {
	return b
}

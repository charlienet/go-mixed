package logx

type mutiLogger struct {
}

func NewLogger() *mutiLogger {
	return &mutiLogger{}
}

func (w *mutiLogger) AppendLogger() Logger {
	return nil
}

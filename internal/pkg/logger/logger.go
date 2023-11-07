package logger

type Reciever interface {
	Subscribe(topic string) error
}

type Logger struct {
	reciever Reciever
}

func NewLogger(r Reciever) *Logger {
	return &Logger{reciever: r}
}

func (l *Logger) StartLogger(topic string) error {
	err := l.reciever.Subscribe(topic)
	return err
}

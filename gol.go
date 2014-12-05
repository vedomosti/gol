package gol

type Level uint8

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

var LevelsString = map[Level]string{
	PanicLevel: "panic",
	FatalLevel: "fatal",
	ErrorLevel: "error",
	WarnLevel:  "warning",
	InfoLevel:  "info",
	DebugLevel: "debug",
}

func (level Level) String() string {
	return LevelsString[level]
}

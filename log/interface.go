package log

// ILogger is an interface that defines the methods that a logger must implement
type ILogger interface {
	Info(string)
	Error(error, string)
}

package logging

// FunctionLogger ...
func (alog *Logger) FunctionLogger(name string) func() {
	alog.lg.Printf("Function %s starts\n", name)
	return func() {
		alog.lg.Printf("Function %s ends\n", name)
	}
}

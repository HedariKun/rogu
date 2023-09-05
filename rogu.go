package rogu

import (
	"io"
	"os"
	"runtime/debug"
	"time"
)

type LogFunction func(msg string)

type Rogu struct {
	Out io.Writer
	Options Options
	log LogFunction
	error LogFunction
	warn LogFunction
}

func New() *Rogu {
	rogu := Rogu{} 
	rogu.Out = os.Stdout
	rogu.Options = DefaultOptions()

	rogu.log = rogu.Logger("LOG")
	rogu.error = rogu.Logger("ERROR")
	rogu.warn = rogu.Logger("WARN")

	return &rogu
}

func (r *Rogu) GetDate() string {
	date := time.Now()
	curDate := date.Format("2006-01-02 03:05:02 Mon")
	return curDate
}

func (r *Rogu) Logger(logType string) LogFunction {
	return func(message string) {
		msg := "[" + logType + "]"
		if r.Options.EnableStack {
			s := NewStack(debug.Stack())
			info := s.GetInfo()
			if(len(info) <= 4) {
				r.BasicError("was not able to get stack info")
				r.BasicError("the log message that was suppose to print is: " + message)
				return
			}
			curLine := info[3]
			msg += "[" + curLine.GetFileName() + ":" + curLine.GetLineNumber() +  " - " + curLine.GetFuncName() + "]"
		}
		if r.Options.EnableDate {
			msg += "[" + r.GetDate() + "]"
		}
		
		msg += ": " + message + "\n"
		r.Out.Write([]byte(msg))
	}
}

func (r *Rogu) Log(message string) {
	r.log(message)
}

func (r *Rogu) Error(message string) {
	r.error(message)
}

func (r *Rogu) Warn(message string) {
	r.warn(message)
}

func (r *Rogu) BasicError(msg string) {
	r.Out.Write([]byte("[ERROR]: " + msg + "\n"))
}
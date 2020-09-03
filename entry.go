package logging

import (
	"bytes"
	"text/template"
	"time"
)

// EntryFunc is a function definition for handling log entries.
type EntryFunc func(entry Entry)

// Entry contains the metadata for the log entry.
type Entry struct {
	Time     time.Time `json:"time",bson:"time",yaml:"time"`
	LogLevel LogLevel  `json:"level",bson:"level",yaml:"level"`
	Message  string    `json:"msg",bson:"msg",yaml:"msg"`
}

// String formats the entry depending on TimeFormat, Format and LogLevelNames.
// The result contains already a new line at the end.
func (e Entry) String() string {
	strTime := e.Time.Format(TimeFormat)

	data := map[string]interface{}{
		"Time":    strTime,
		"Level":   e.LogLevel.String(),
		"Message": e.Message,
	}

	buff := bytes.NewBufferString("")

	// TODO: Caching the template. This could possibly be unnecessarily time consuming.
	t := template.Must(template.New("").Parse(Format))
	_ = t.Execute(buff, data)
	return buff.String() + "\n"
}

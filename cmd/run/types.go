package run

import "github.com/spf13/afero"

type logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
}

type Options struct {
	FileSystem *afero.Afero
	Verbose    bool
}

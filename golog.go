package golog

import (
	"flag"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/pkg/errors"
)

const (
	// GologDir is the location relative to the user's Home dir where golog store its logfiles
	GologDir = ".golog"

	DefaultDirMode  fs.FileMode = 0700 // default FileMode when creating log directories
	DefaultFileMode fs.FileMode = 0600 // default FileMode when creating log files
)

var (
	// stdLogger is the Standard Logger exposed by the package. The default one provided here
	// discards all output. Verbosity is set to 0, which is the less verbose of all. To get a
	// proper usable logger, call ConfigureFromFlags()
	stdLogger = New(log.New(io.Discard, "", 0), 0)

	// Common flags vars
	logToStdErr, alsoLogToStdErr bool
	logVerbosity                 int
	logFileLocation              string
)

func init() { //nolint:gochecknoinits // we need the user home dir before we set the flag
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Unable to determine user's home directory: %v", err)
	}
	program := filepath.Base(os.Args[0])

	defaultLogFileLocation := path.Join(userHomeDir, GologDir, program+".log")

	flag.BoolVar(&logToStdErr, "logtostderr", false, "log to standard error instead of files")
	flag.BoolVar(&alsoLogToStdErr, "alsologtostderr", false, "log to standard error as well as files")
	flag.IntVar(&logVerbosity, "verbosity", 0, "the log verbosity. Higher values yield more events")
	flag.StringVar(&logFileLocation, "logfilelocation", defaultLogFileLocation, "the log file location")
}

// ConfigureFromFlags configures golog based on the flag parameters passed on to it. Note that it DOES NOT
// calls flag.Parse(),and that is left for the caller to do it beforehand
func ConfigureFromFlags() error {
	var o io.Writer
	if logToStdErr {
		o = os.Stderr
	} else {
		var err error
		p := filepath.Dir(logFileLocation)
		if err = os.MkdirAll(p, DefaultDirMode); err != nil {
			return errors.Wrap(err, "error creating path")
		}
		o, err = os.OpenFile(logFileLocation, os.O_APPEND|os.O_WRONLY|os.O_CREATE, DefaultFileMode)
		if err != nil {
			return errors.Wrap(err, "error opening/creating log file")
		}
		if alsoLogToStdErr {
			o = io.MultiWriter(o, os.Stderr)
		}
	}
	logFlags := log.LstdFlags
	if logVerbosity > 0 {
		logFlags |= log.Lshortfile
	}
	stdLogger = New(log.New(o, "", logFlags), logVerbosity)
	return nil
}

// Debug returns a Debug event from the stdLogger
func V(verbosity int) *Event { return stdLogger.V(verbosity) }

// With returns a Context, useful for creating subloggers
func With() *Context { return stdLogger.With() }

// SetGlobalVerbosity sets the verbosity of the stdLogger. It does not affect already
// created subloggers, but new ones will inherit the level.
func SetGlobalVerbosity(verbosity int) {
	stdLogger.verbosity = verbosity
}

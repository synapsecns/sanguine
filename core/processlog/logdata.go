package processlog

// LogMetadata provides an interface for getting locations of process log files.
type LogMetadata interface {
	// LogDir gets the log directory.
	LogDir() string
	// StdOutFile gets the log directory.
	StdOutFile() string
	// StdErrFile gets the log directory.
	StdErrFile() string
	// CombinedFile gets the log directory.
	CombinedFile() string
}

// logMetadataImpl provides an interface for getting info about the logs of a file.
type logMetadataImpl struct {
	// logDir contains the log dir
	logDir string
	// stdoutFile contains the path to the stdout output file
	stdoutFile string
	// stderrFile contains the path to the stderr output file
	stderrFile string
	// combinedFile contains the path to the combined file
	combinedFile string
}

func (li logMetadataImpl) LogDir() string {
	return li.logDir
}

func (li logMetadataImpl) StdOutFile() string {
	return li.stdoutFile
}

func (li logMetadataImpl) StdErrFile() string {
	return li.stderrFile
}

func (li logMetadataImpl) CombinedFile() string {
	return li.combinedFile
}

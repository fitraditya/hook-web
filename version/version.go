package version

import (
	"fmt"
	"runtime"
)

var GitCommit = ""
var Environment = ""
var Version = ""
var BuildDate = ""
var GoVersion = runtime.Version()
var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)

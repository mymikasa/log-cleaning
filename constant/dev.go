package constant

import "os"

var IsDebug = (os.Getenv("DEBUG") == "true")

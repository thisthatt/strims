package debug

import (
	"encoding/json"
	"fmt"
	"path"
	"runtime"
	"time"

	"github.com/MemeLabs/go-ppspp/pkg/errutil"
)

// PrintJSON ...
func PrintJSON(i interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf(
		"%s %s:%d: %s\n",
		time.Now().Format("2006/01/02 15:04:05.000000"),
		path.Base(file),
		line, string(errutil.Must(json.MarshalIndent(i, "", "  "))),
	)
}

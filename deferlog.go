package dl

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"
)

const timeFormat = "2006/01/02 15:04:05.000000"

var output io.Writer = os.Stderr

// SetOutput sets the output destination.
func SetOutput(w io.Writer) {
	output = w
}

// Log example (must use defer):
//
// func do() {
// 	defer Log("aaa")("bbb")
// 	time.Sleep(1e7)
// }
//
// will output 2 logs:
// 2019/08/06 19:02:56.670369 deferlog_test.go:12: [--->] aaa
// 2019/08/06 19:02:56.683101 deferlog_test.go:12: [<---] aaa bbb
// first log are printed immediately when do runs, with first args.
// second log are printed when do finished, with first and second args.
func Log(args ...interface{}) func(args2 ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	file = path.Base(file) + ":" + strconv.Itoa(line) + ":"
	fmt.Fprintln(output, append([]interface{}{
		time.Now().Format(timeFormat), file, "[--->]"}, args...)...)
	return func(args2 ...interface{}) {
		fmt.Fprintln(output, append(append([]interface{}{
			time.Now().Format(timeFormat), file,
			"[<---]"}, args...), args2...)...)
	}
}

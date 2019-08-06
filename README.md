# defer log

print two logs(in and out of function) with two different args in one line.

### example

```go
import "github.com/ma6174/dl"

func do() {
	defer dl.Log("aaa")("bbb")
	time.Sleep(1e7)
}
```

output:

```
2019/08/06 19:02:56.670369 deferlog_test.go:12: [--->] aaa
2019/08/06 19:02:56.683101 deferlog_test.go:12: [<---] aaa bbb
```

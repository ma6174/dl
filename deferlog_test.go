package dl

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
)

func do() {
	defer Log("aaa")("bbb")
	time.Sleep(1e7)
}

func TestDeferlog(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	SetOutput(buf)
	do()
	fmt.Println(buf.String())
	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 2 {
		t.Fatal("line count != 2", len(lines))
	}
	sp1 := strings.SplitN(lines[0], " ", 4)
	sp2 := strings.SplitN(lines[1], " ", 4)
	fmt.Printf("%#v\n", sp1)
	fmt.Printf("%#v\n", sp2)
	t1, _ := time.Parse(timeFormat, sp1[0]+" "+sp1[1])
	t2, _ := time.Parse(timeFormat, sp2[0]+" "+sp2[1])
	if t2.Sub(t1) <= 1e7 {
		t.Fatal("no sleep between 2 logs")
	}
	if sp1[2] != sp2[2] {
		t.Fatal("file and line number not match")
	}
	if sp1[3] != "[--->] aaa" || sp2[3] != "[<---] aaa bbb" {
		t.Fatal("message format error")
	}
}

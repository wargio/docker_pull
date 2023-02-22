package filetool

import (
	"compress/gzip"
	"fmt"
	"os"
	"testing"
)

func TestFileWrite(t *testing.T) {
	tests := []struct {
		name string
	}{
		{}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, e := os.Open("/home/git_work/go_pull_http/test/d8bf8c30a94c2fadb32a876118bf96918274345d7d0ee75369ed238a5fbbfe07/layer.tar.gz")
			//if e != nil {
			//	fmt.Println(1)
			//	fmt.Println(f)
			//	fmt.Println(e)
			//}

			//f := GetfileOjb("/home/git_work/go_pull_http/tmp_redis_latest/d8bf8c30a94c2fadb32a876118bf96918274345d7d0ee75369ed238a5fbbfe07/layer.tar.gz")

			//_,e:=f.WriteString("1.0")
			//fmt.Println(e)
			//f.Close()

			gf, e := gzip.NewReader(f)
			if e != nil {
				fmt.Println(2)

				fmt.Println(gf)
				fmt.Println(e)
			}
		})
	}
}

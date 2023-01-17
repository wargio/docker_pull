package conversion

import (
	"strconv"

	"github.com/dustin/go-humanize"
)

func Humanize_uintbytes(s uint64) string {
	return humanize.Bytes(s)
}

func Humanize_intbytes(s int) string {
	return humanize.Bytes(uint64(s))
}


func Humanize_bstr(s string) (string, error) {
	uint64, err := strconv.ParseUint(s, 10, 64)
	return humanize.IBytes(uint64), err
}

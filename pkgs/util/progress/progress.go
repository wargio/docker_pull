package progress

import (
	"fmt"
	"go_pull/pkgs/util/conversion"
	"go_pull/pkgs/util/makestr"
	"os"
)

type Progress struct {
	Ublob             string
	Total             int
	Current           int
	ProgressBarLength int
	CurrentBar        int
	//mux               sync.Mutex
}

func (p *Progress) Write(b []byte) (n int, err error) {
	p.Current += len(b)
	p.progressBar()
	return
}

func (p *Progress) progressBar() {
	fmt.Print(makestr.Joinstring("", p.Ublob, ": Downloading ["))

	percent := float64(p.Current) / float64(p.Total)
	bars := int(percent * float64(p.ProgressBarLength))

	now_bytes := conversion.Humanize_intbytes(p.Current)
	total_bytes := conversion.Humanize_intbytes(p.Total)

	for i := 0; i < p.ProgressBarLength; i++ {
		if i < bars {
			fmt.Print("#")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Printf("] %v/%v\n", now_bytes, total_bytes)
	os.Stdout.Sync()
}

package logtool

import (
	"os"
	"testing"

	"go.uber.org/zap"
)

func handleError(_e string) {
	SugLog.Fatal(1)
	SugLog.Desugar()
}

func Test_InitEvent(t *testing.T) {
	InitEvent()

	a := 1
	SugLog.Infof("aa%v", a)
	SugLog.Infow("bb", zap.String("msg", "aaac"))
	//handleError("666")
	_, b := os.Open("/tmp/aa")
	Logc.Error("",zap.Error(b))
	Logc.Warn("678", zap.("asd", "asd"))

	Fatalerror(b)
}

package filetool

import (
	"log"
	"os"
	//"strconv"
	//"time"
)

func GetfileOjb(filepath string) *os.File {
	_, err := os.Stat(filepath)

	if err != nil {
		if !os.IsExist(err) {
			os.Create(filepath)
		} else {
			log.Fatalf(" file Stat failed, err:%v", err)
		}
	}
	FileWrite, err := os.OpenFile(filepath, os.O_APPEND|os.O_RDWR, os.ModePerm)

	return FileWrite
}

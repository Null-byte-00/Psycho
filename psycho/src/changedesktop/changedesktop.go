package changedesktop

import (
	"io/ioutil"
	"syscall"
	"path/filepath"
	"unsafe"
	"../storedata"
)

//writes background image data in a file
func Writedata(filename string) {
	ioutil.WriteFile(filename, storedata.Getbackgrounddata(), 0644)
}

//sets an image as background image
func Setbackground(name string) {
	user32 := syscall.NewLazyDLL("user32.dll")
	systemParametersInfo := user32.NewProc("SystemParametersInfoW")
	imagelocation, _ := filepath.Abs(name)
	systemParametersInfo.Call(20, 0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(imagelocation))), 2)
}
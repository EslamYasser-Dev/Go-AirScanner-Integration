package main

import (
	"log"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

func main() {
	err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)
	if err != nil {
		log.Fatal(err)
	}
	defer ole.CoUninitialize()

	devManager, err := oleutil.CreateObject("WIA.DeviceManager")
	if err != nil {
		log.Fatal(err)
	}
	defer devManager.Release()


	

}

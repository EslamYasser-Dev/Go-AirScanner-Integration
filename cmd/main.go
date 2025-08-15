package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"runtime"
	"strings"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

/*** Domain ***/
type Scanner struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

/*** Use Case ***/
type ScannerService interface {
	ListScanners() ([]Scanner, error)
}

type scannerService struct {
	deviceManager *ole.IDispatch
}

func NewScannerService(dm *ole.IDispatch) ScannerService {
	return &scannerService{deviceManager: dm}
}

func (s *scannerService) ListScanners() ([]Scanner, error) {
	if s == nil || s.deviceManager == nil {
		return nil, errors.New("scanner service not initialized")
	}

	// WIA 2.0: DeviceManager.DeviceInfos (NOT "Devices")
	devInfosVar, err := oleutil.GetProperty(s.deviceManager, "DeviceInfos")
	if err != nil {
		return nil, err
	}
	devInfos := devInfosVar.ToIDispatch()
	defer devInfos.Release()

	countVar, err := oleutil.GetProperty(devInfos, "Count")
	if err != nil {
		return nil, err
	}
	count := int(countVar.Val)

	scanners := make([]Scanner, 0, count)

	// WIA collections are 1-based
	for i := 1; i <= count; i++ {
		itemVar, err := oleutil.GetProperty(devInfos, "Item", i)
		if err != nil {
			// skip bad entries instead of failing the whole request
			continue
		}
		item := itemVar.ToIDispatch()

		id, _ := getString(item, "DeviceID")
		name := "(unknown)"

		// Read Properties -> find "Name"
		propsVar, err := oleutil.GetProperty(item, "Properties")
		if err == nil {
			props := propsVar.ToIDispatch()

			propsCountVar, err2 := oleutil.GetProperty(props, "Count")
			if err2 == nil {
				propsCount := int(propsCountVar.Val)
				for j := 1; j <= propsCount; j++ {
					propVar, err3 := oleutil.GetProperty(props, "Item", j)
					if err3 != nil {
						continue
					}
					prop := propVar.ToIDispatch()

					propName, _ := getString(prop, "Name")
					if strings.EqualFold(propName, "Name") {
						if v, err := getString(prop, "Value"); err == nil && v != "" {
							name = v
						}
						prop.Release()
						break
					}
					prop.Release()
				}
			}
			props.Release()
		}

		scanners = append(scanners, Scanner{ID: id, Name: name})
		item.Release()
	}

	return scanners, nil
}

/*** Delivery (HTTP) ***/
func listScannersHandler(svc ScannerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		scanners, err := svc.ListScanners()
		if err != nil {
			http.Error(w, "Failed to list scanners: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(scanners)
	}
}

/*** Main (COM init + wiring) ***/
func main() {
	if runtime.GOOS != "windows" {
		log.Fatal("This program works in Windows only")
	}

	// Many WIA components prefer STA:
	// ole.COINIT_APARTMENTTHREADED over MULTITHREADED
	if err := ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED); err != nil {
		log.Fatal(err)
	}
	defer ole.CoUninitialize()

	devMgrObj, err := oleutil.CreateObject("WIA.DeviceManager")
	if err != nil {
		log.Fatal(err)
	}
	defer devMgrObj.Release()

	deviceManager, err := devMgrObj.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		log.Fatal(err)
	}
	defer deviceManager.Release()

	svc := NewScannerService(deviceManager)

	mux := http.NewServeMux()
	mux.Handle("/scanners", listScannersHandler(svc))

	log.Println("Server running on :2028")
	if err := http.ListenAndServe(":2028", mux); err != nil {
		log.Fatal(err)
	}
}

/*** Helpers ***/
func getString(obj *ole.IDispatch, name string, args ...interface{}) (string, error) {
	v, err := oleutil.GetProperty(obj, name, args...)
	if err != nil {
		return "", err
	}
	defer v.Clear()
	return v.ToString(), nil
}

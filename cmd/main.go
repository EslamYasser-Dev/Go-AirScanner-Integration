package main

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// --- Domain Layer ---

type Scanner struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// --- Use Case Layer ---

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
	devices := oleutil.MustGetProperty(s.deviceManager, "Devices").ToIDispatch()
	defer devices.Release()

	count := int(oleutil.MustGetProperty(devices, "Count").Val)
	scanners := make([]Scanner, 0, count)

	for i := 1; i <= count; i++ {
		item := oleutil.MustGetProperty(devices, "Item", i).ToIDispatch()
		defer item.Release()

		deviceID := oleutil.MustGetProperty(item, "DeviceID").ToString()
		deviceName := oleutil.MustGetProperty(item, "Properties").ToIDispatch()
		defer deviceName.Release()

		propsCount := int(oleutil.MustGetProperty(deviceName, "Count").Val)
		var name string

		for j := 1; j <= propsCount; j++ {
			prop := oleutil.MustGetProperty(deviceName, "Item", j).ToIDispatch()
			defer prop.Release()

			propName := oleutil.MustGetProperty(prop, "Name").ToString()
			if propName == "Name" {
				name = oleutil.MustGetProperty(prop, "Value").ToString()
				break
			}
		}

		scanners = append(scanners, Scanner{ID: deviceID, Name: name})
	}

	return scanners, nil
}

// --- Delivery Layer (HTTP) ---

func listScannersHandler(svc ScannerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		scanners, err := svc.ListScanners()
		if err != nil {
			http.Error(w, "Failed to list scanners: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(scanners)
	}
}

// --- Main (Composition Root) ---

func main() {
	if runtime.GOOS != "windows" {
		log.Fatal("This program works in Windows only")
	}

	err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)
	if err != nil {
		log.Fatal(err)
	}
	defer ole.CoUninitialize()

	devManagerUnknown, err := oleutil.CreateObject("WIA.DeviceManager")
	if err != nil {
		log.Fatal(err)
	}
	defer devManagerUnknown.Release()

	deviceManager, err := devManagerUnknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		log.Fatal(err)
	}
	defer deviceManager.Release()

	// Create service
	scannerSvc := NewScannerService(deviceManager)

	// HTTP server
	mux := http.NewServeMux()
	mux.Handle("/scanners", listScannersHandler(scannerSvc))

	log.Println("Server running on :2028")
	http.ListenAndServe(":2028", mux)
}

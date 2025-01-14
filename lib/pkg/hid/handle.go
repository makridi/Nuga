package hid

import (
	"log"
	"sync"
	"time"

	"github.com/sstallion/go-hid"
)

// RequestRetries is the number of attempts that will be made on a request.
const RequestRetries = 5

// Handle represents backlight device handle.
type Handle struct {
	Device *hid.Device
	// Debug flag. If true, then debugging data will be written to stdout when functions are executed.
	Debug bool
	mutex sync.Mutex
}

// GetInfo returns handle device information.
func (h *Handle) GetInfo() (*DeviceInfo, error) {
	name, err := h.Device.GetProductStr()
	if err != nil {
		return nil, err
	}
	info, err := h.Device.GetDeviceInfo()
	if err != nil {
		return nil, err
	}
	return &DeviceInfo{
		Name:     name,
		Path:     info.Path,
		Firmware: formatVersion(info.ReleaseNbr),
	}, nil
}

// SendWithRetries sends the request and resends it if the request fails.
func (h *Handle) SendWithRetries(payload []byte) error {
	var err error
	for i := 0; i < RequestRetries; i++ {
		if h.Debug {
			log.Println("Send attempt", i+1)
		}
		err = h.Send(payload)
		if err == nil {
			return nil
		}
	}
	if err != nil {
		return err
	}
	return ErrNotFound
}

// Send packet to the device.
func (h *Handle) Send(payload []byte) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	if h.Debug {
		log.Printf("Send %v: %v", len(payload), payload)
	}

	transferred, err := h.Device.SendFeatureReport(payload)
	if err != nil {
		return err
	}
	expected := len(payload)
	if transferred != len(payload) {
		return NewErrCountMismatch(expected, transferred)
	}

	h.waitSync()
	return nil
}

// Read packet from the device.
func (h *Handle) Read(count int) ([]byte, error) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	buf := make([]byte, count)
	buf[0] = 0x06
	length, err := h.Device.GetFeatureReport(buf)
	if err != nil {
		return nil, err
	}
	packet := buf[1:]
	if h.Debug {
		if length > 0 {
			log.Printf("Read %v", buf)
		} else {
			log.Println("Read 0")
		}
	}
	h.waitSync()
	return packet, nil
}

// Request sends a request to the device.
func (h *Handle) Request(payload []byte, count int) ([]byte, error) {
	var resp []byte
	var err error
	for i := 0; i < RequestRetries; i++ {
		if h.Debug {
			log.Println("Read attempt", i+1)
		}
		resp, err = h.tryRequest(payload, count)
		if len(resp) > 0 && resp[0] != 0 {
			return resp, nil
		}
	}
	if err != nil {
		return resp, err
	}
	return resp, ErrNotFound
}

// Close device handle.
// The function should be called after the end of operation with the device.
func (h *Handle) Close() error {
	return h.Device.Close()
}

func (h *Handle) tryRequest(payload []byte, count int) ([]byte, error) {
	err := h.Send(payload)
	if err != nil {
		return nil, err
	}
	return h.Read(count)
}

func (h *Handle) waitSync() {
	time.Sleep(time.Millisecond * 50)
}

// OpenHandle opens a connection to the device
func OpenHandle() (*Handle, error) {
	var h Handle
	var path string
	err := hid.Enumerate(0x05AC, 0x024F, func(info *hid.DeviceInfo) error {
		if info.Usage == 1 && info.UsagePage == 0xFF00 {
			path = info.Path
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(path) == 0 {
		return &h, ErrNotFound
	}

	device, err := hid.OpenPath(path)
	if err != nil {
		return &h, err
	}
	h.Device = device
	return &h, nil
}

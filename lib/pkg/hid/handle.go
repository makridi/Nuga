package hid

import (
	"log"
	"time"

	"github.com/sstallion/go-hid"
)

// Handle represents backlight device handle
type Handle struct {
	Device *hid.Device
	// Debug flag. If true, then debugging data will be written to stdout when functions are executed.
	Debug bool
}

// Send packet to the device.
func (h *Handle) Send(payload []byte) error {
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
	buf := make([]byte, count)
	buf[0] = 0x06
	length, err := h.Device.GetFeatureReport(buf)
	if err != nil {
		return buf, err
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
	// Cut report id
	return packet, nil
}

// Send packet to the device.
func (h *Handle) Request(payload []byte, count int) ([]byte, error) {
	err := h.Send(payload)
	if err != nil {
		return nil, err
	}
	return h.Read(count)
}

// Close device handle.
// The function should be called after the end of operation with the device.
func (h *Handle) Close() error {
	return h.Device.Close()
}

// Waiting for data to be written. According to the documentation, it takes 10 ms to do this
func (h *Handle) waitSync() {
	time.Sleep(time.Millisecond * 10)
}

func OpenHandle() (Handle, error) {
	var h Handle
	var path string
	hid.Enumerate(0x05AC, 0x024F, func(info *hid.DeviceInfo) error {
		if info.Usage == 1 && info.UsagePage == 0xFF00 {
			path = info.Path
		}
		return nil
	})

	if len(path) == 0 {
		return h, ErrNotFound
	}

	device, err := hid.OpenPath(path)
	if err != nil {
		return h, err
	}
	h.Device = device
	return h, nil
}
package authdDevice

import (
	md "github.com/ebikode/eLearning-core/model"
)

// Repository provides access to the md.AuthdDevice storage.
type AuthdDeviceRepository interface {
	// Get returns the authenticated device with given ID.
	Get(string) *md.AuthdDevice
	// Get returns the authenticated device with given customerID.
	// GetCustomerAuthdDevice(uint) (md.AuthdDevice)
	// Store a given customer authenticated device to the repository.
	Store(md.AuthdDevice) (md.AuthdDevice, error)
	// Update a given authenticated device in the repository.
	Update(*md.AuthdDevice) (*md.AuthdDevice, error)
	// Delete a given location in the repository.
	Delete(*md.AuthdDevice, bool) (bool, error)
}

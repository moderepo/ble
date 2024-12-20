package dev

import (
	"github.com/go-ble/ble"
)

// NewDevice ...
func NewDevice(impl string, opts ...ble.Option) (d ble.Device, err error) {
	return DefaultDevice(opts...)
}

// NewAdvertisingExtensionsDevice ...
func NewAdvertisingExtensionsDevice(impl string, opts ...ble.Option) (d ble.Device, err error) {
	return AdvertisingExtensionsDevice(opts...)
}

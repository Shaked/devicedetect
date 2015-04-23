package platform

type Device interface {
	Name() string
	Version() string
	Os() string

	String() string
}
type GenericDevice struct {
	name    string
	version string
	os      string
}

func (g GenericDevice) Name() string {
	return g.name
}

func (g GenericDevice) Version() string {
	return g.version
}

func (g GenericDevice) Os() string {
	return g.os
}

type DeviceTablet struct {
	GenericDevice
}

func (g DeviceTablet) String() string {
	return "Tablet"
}

type DeviceMobile struct {
	GenericDevice
}

func (g DeviceMobile) String() string {
	return "Mobile"
}

type DeviceTv struct {
	GenericDevice
}

func (g DeviceTv) String() string {
	return "Tv"
}

type DeviceDesktop struct {
	GenericDevice
}

func (g DeviceDesktop) String() string {
	return "Desktop"
}

type DeviceWatch struct {
	GenericDevice
}

func (g DeviceWatch) String() string {
	return "Watch"
}

type DeviceBot struct {
	GenericDevice
}

func (g DeviceBot) String() string {
	return "Bot"
}

func NewTablet(name, version, os string) Device {
	return &DeviceTablet{GenericDevice{name, version, os}}
}

func NewMobile(name, version, os string) Device {
	return &DeviceMobile{GenericDevice{name, version, os}}
}

func NewTv(name, version, os string) Device {
	return &DeviceTv{GenericDevice{name, version, os}}
}

func NewDesktop(name, version, os string) Device {
	return &DeviceDesktop{GenericDevice{name, version, os}}
}

func NewWatch(name, version, os string) Device {
	return &DeviceWatch{GenericDevice{name, version, os}}
}

func NewBot(name, version, os string) Device {
	return &DeviceBot{GenericDevice{name, version, os}}
}

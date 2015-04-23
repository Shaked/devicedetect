package platform

type DeviceType int

const (
	TABLET DeviceType = iota
	MOBILE
	TV
	DESKTOP
	WATCH
	BOT
	UNKNOWN
)

type Device interface {
	Name() string
	Version() string
	Os() string

	Which() DeviceType
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

func (g DeviceTablet) Which() DeviceType {
	return TABLET
}

type DeviceMobile struct {
	GenericDevice
}

func (g DeviceMobile) Which() DeviceType {
	return MOBILE
}

type DeviceTv struct {
	GenericDevice
}

func (g DeviceTv) Which() DeviceType {
	return TV
}

type DeviceDesktop struct {
	GenericDevice
}

func (g DeviceDesktop) Which() DeviceType {
	return DESKTOP
}

type DeviceWatch struct {
	GenericDevice
}

func (g DeviceWatch) Which() DeviceType {
	return WATCH
}

type DeviceBot struct {
	GenericDevice
}

func (g DeviceBot) Which() DeviceType {
	return BOT
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

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
	Type() DeviceType
	// Platform() map[string]map[string]string
	// Engine() map[string]map[string]string
	// Browsers() map[string]map[string]string
}
type GenericDevice struct {
	name string
	typ  DeviceType
}

func (g GenericDevice) Name() string {
	return g.name
}

func (g GenericDevice) Type() DeviceType {
	return g.typ
}

type DeviceTablet struct {
	GenericDevice
}

type DeviceMobile struct {
	GenericDevice
}

type DeviceTv struct {
	GenericDevice
}

type DeviceDesktop struct {
	GenericDevice
}

type DeviceWatch struct {
	GenericDevice
}

type DeviceBot struct {
	GenericDevice
}

func NewTablet(name string) Device {
	return &DeviceTablet{GenericDevice{name, TABLET}}
}

func NewMobile(name string) Device {
	return &DeviceMobile{GenericDevice{name, MOBILE}}
}

func NewTv(name string) Device {
	return &DeviceTv{GenericDevice{name, TV}}
}

func NewDesktop(name string) Device {
	return &DeviceDesktop{GenericDevice{name, DESKTOP}}
}

func NewWatch(name string) Device {
	return &DeviceWatch{GenericDevice{name, WATCH}}
}

func NewBot(name string) Device {
	return &DeviceBot{GenericDevice{name, BOT}}
}

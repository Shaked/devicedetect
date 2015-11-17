package platform

type DeviceType int

const (
	TABLET DeviceType = iota
	MOBILE
	TV
	DESKTOP
	WATCH
	BOT
	GLASS
	UNKNOWN
)

type Device interface {
	Name() string
	Type() DeviceType
	Platform() *Platform
	SetPlatform(p *Platform)
	// Engine() map[string]map[string]string
	// Browsers() map[string]map[string]string
}

func NewPlatform(
	name,
	version string,
) *Platform {
	return &Platform{
		name:    name,
		version: version,
	}
}

func (p *Platform) Name() string {
	return p.name
}

func (p *Platform) Version() string {
	return p.version
}

func (p *Platform) Build() string {
	return p.build
}

func (p *Platform) Model() string {
	return p.model
}

func (p *Platform) SetBuild(b string) {
	p.build = b
}

func (p *Platform) SetModel(m string) {
	p.model = m
}

type Platform struct {
	name    string
	version string
	build   string
	model   string
}

type GenericDevice struct {
	name     string
	typ      DeviceType
	platform *Platform
}

func (g GenericDevice) Name() string {
	return g.name
}

func (g GenericDevice) Type() DeviceType {
	return g.typ
}

func (g GenericDevice) Platform() *Platform {
	return g.platform
}

func (g GenericDevice) SetPlatform(p *Platform) {
	g.platform = p
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

type DeviceGlass struct {
	GenericDevice
}

type DeviceUnknown struct {
	GenericDevice
}

func NewTablet(name string) *DeviceTablet {
	return &DeviceTablet{GenericDevice{name: name, typ: TABLET}}
}

func NewMobile(name string) *DeviceMobile {
	return &DeviceMobile{GenericDevice{name: name, typ: MOBILE}}
}

func NewTv(name string) *DeviceTv {
	return &DeviceTv{GenericDevice{name: name, typ: TV}}
}

func NewDesktop(name string) *DeviceDesktop {
	return &DeviceDesktop{GenericDevice{name: name, typ: DESKTOP}}
}

func NewWatch(name string) *DeviceWatch {
	return &DeviceWatch{GenericDevice{name: name, typ: WATCH}}
}

func NewBot(name string) *DeviceBot {
	return &DeviceBot{GenericDevice{name: name, typ: BOT}}
}

func NewGlass(name string) *DeviceGlass {
	return &DeviceGlass{GenericDevice{name: name, typ: GLASS}}
}

func NewUnknown(name string) *DeviceUnknown {
	return &DeviceUnknown{GenericDevice{name: name, typ: UNKNOWN}}
}

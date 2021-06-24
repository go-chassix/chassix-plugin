package plugin

import "github.com/emicklei/go-restful/v3"

type Plugin interface {
	Data(data Data)
	Name() string
	App() []*restful.WebService
	Management() []*restful.WebService
}

type Data struct {
	DB      interface{}
	Logger  interface{}
	ConfDir string
}

type MainAPP struct {
	plugins []Plugin
}

type Plug struct {
	Name    string
	Enabled bool
}

var plugins = make(map[string]*Plug, 20)

//Init initialize plugin and import data for reuse db, config etc.
func Init(plug Plugin, data Data) {
	plug.Data(data)
}

//Load register api to container
func Load(container *restful.Container, plugin Plugin) {
	for _, ws := range plugin.App() {
		container.Add(ws)
	}
	plugins[plugin.Name()] = &Plug{
		Name:    plugin.Name(),
		Enabled: true,
	}
	//plugins = append(plugins, plugin.Name())
}

//LoadManagement  register webservice for management api
func LoadManagement(container *restful.Container, plugin Plugin) {
	for _, ws := range plugin.Management() {
		container.Add(ws)
	}

	plugins[plugin.Name()] = &Plug{
		Name:    plugin.Name(),
		Enabled: true,
	}
}

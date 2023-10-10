package zcore

const (
	ExecSuffix       = "zz"
	ExecName         = "go" + ExecSuffix
	AnnotationIdent  = "+"
	AnnotationPrefix = AnnotationIdent + ExecSuffix + ":"
)

type (
	Plugin interface {
		Name() string
		Args() ([]string, map[string]string)
		Description() string
		Run(entities DeclEntities) (err error)
	}

	PluginEntity struct {
		Plugin
		Options map[string]string
	}

	PluginEntities []PluginEntity
)

var pluginRegistry = map[string]Plugin{}

func PluginRegistry() map[string]Plugin { return pluginRegistry }

func RegisterPlugin(plugin Plugin) {
	pluginRegistry[plugin.Name()] = plugin
}

func (entities PluginEntities) Run(filename string) (err error) {
	for _, entity := range entities {
		if err = entity.run(filename); err != nil {
			return
		}
	}
	return
}

func (entity PluginEntity) run(filename string) (err error) {
	decls, err := ParseFileOrDirectory(filename, AnnotationPrefix)
	if err != nil {
		return
	}
	return entity.Plugin.Run(decls.Parse(entity, entity.Options))
}

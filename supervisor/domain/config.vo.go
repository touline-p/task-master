package supervisor

const (
	STR ConfigType = iota
	INT
)

type ConfigValue interface {
	Key() string
	Required() bool
	Type() ConfigType
	Value() any
	DefaultValue() any
}

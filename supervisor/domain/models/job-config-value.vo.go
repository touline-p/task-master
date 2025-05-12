package models

type ConfigType int

const (
	STR ConfigType = iota
	INT
)

type JobConfigValue struct {
	key          string
	required     bool
	configType   ConfigType
	value        any
	defaultValue any
}

func NewJobConfigValue(key string, required bool, configType ConfigType, value any, defaultValue any) *JobConfigValue {
	return &JobConfigValue{
		key:          key,
		required:     required,
		configType:   configType,
		value:        value,
		defaultValue: defaultValue,
	}
}

func (jcv *JobConfigValue) Key() string {
	return jcv.key
}

func (jcv *JobConfigValue) Required() bool {
	return jcv.required
}

func (jcv *JobConfigValue) Type() ConfigType {
	return jcv.configType
}

func (jcv *JobConfigValue) Value() any {
	return jcv.value
}

func (jcv *JobConfigValue) DefaultValue() any {
	return jcv.defaultValue
}

func (jcv *JobConfigValue) SetValue(value any) {
	jcv.value = value
}

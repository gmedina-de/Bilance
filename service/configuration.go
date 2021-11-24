package service

type Configuration interface {
	Get(setting Setting) string
	Set(setting Setting, value string)
}

func (this *mapConfiguration) Get(setting Setting) string {
	value, ok := this.settings[setting]
	if ok {
		return value
	} else {
		return ""
	}
}

func (this *mapConfiguration) Set(setting Setting, value string) {
	this.settings[setting] = value
}

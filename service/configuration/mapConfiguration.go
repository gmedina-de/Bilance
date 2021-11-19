package configuration

type mapConfiguration struct {
	settings map[Setting]string
}

func MapConfiguration() Configuration {
	return &mapConfiguration{settings: make(map[Setting]string)}
}

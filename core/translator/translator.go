package translator

type Translator interface {
	Translation(language string, translation any)
	Translate(language string, key string, params ...any) string
}

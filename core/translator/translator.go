package translator

type Translator interface {
	AddTranslation(language string, translation any)
	Translate(language string, message string, v ...any) string
}

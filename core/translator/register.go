package translator

var localizations = make(map[string]map[string]string)

func AddLocalization(language string, localization map[string]string) {
	localizations[language] = localization
}

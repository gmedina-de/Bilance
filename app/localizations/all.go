package localizations

import "genuine/core/localizations"

type all struct {
}

func All() localizations.Provider {
	return &all{}
}

func (s all) GetLocalizations() map[string]localizations.Localization {
	return map[string]localizations.Localization{
		"de":      de,
		"default": dfault,
	}
}

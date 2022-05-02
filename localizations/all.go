package localizations

type all struct {
}

func All() Provider {
	return &all{}
}

func (s all) GetLocalizations() map[string]Localization {
	return map[string]Localization{
		"de-DE":   de,
		"default": dfault,
	}
}

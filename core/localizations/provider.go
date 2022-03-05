package localizations

type Localization = map[string]string

type Provider interface {
	GetLocalizations() map[string]Localization
}

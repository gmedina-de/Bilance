package localization

import "reflect"

type l10n struct {
	// navigation
	dashboard       string
	payments        string
	debts           string
	types           string
	admin_functions string
	users           string
	projects        string

	// actions
	logout                          string
	search                          string
	new                             string
	edit                            string
	save                            string
	cancel                          string
	delete                          string
	delete_confirmation             string
	delete_confirmation_description string

	// fields
	id                      string
	name                    string
	name_placeholder        string
	amount                  string
	date                    string
	Type                    string
	payer                   string
	payee                   string
	description             string
	description_placeholder string
	username                string
	username_placeholder    string
	password                string
	password_placeholder    string
	role                    string
	role_normal             string
	role_admin              string

	// filters
	this_week string
}

var l10nMap = map[string]l10n{
	"de": l10nDe,
}

func Translate(message string) string {
	//todo: generalize function for allowing more languages, depending on user configuration / request parameters
	return reflect.ValueOf(l10nDe).FieldByName(message).String()
}

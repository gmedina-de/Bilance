package localization

import "reflect"

type l10n struct {
	// navigation
	dashboard       string
	payments        string
	categories      string
	debts           string
	analysis        string
	admin_functions string
	users           string
	projects        string

	// actions
	logout                          string
	search                          string
	filter                          string
	date_filter                     string
	new                             string
	edit                            string
	save                            string
	cancel                          string
	delete                          string
	delete_confirmation             string
	delete_confirmation_description string

	// alerts
	no_records_found          string
	record_saved_successfully string

	// fields
	id          string
	name        string
	amount      string
	date        string
	category    string
	payer       string
	payee       string
	description string
	username    string
	password    string
	role        string
	role_normal string
	role_admin  string

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

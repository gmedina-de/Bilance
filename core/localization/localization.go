package localization

import "reflect"

type l10n struct {
	// navigation
	payments        string
	categories      string
	analysis        string
	balances        string
	expenses        string
	by_period       string
	by_category     string
	admin_functions string
	users           string
	projects        string
	// actions
	previous                        string
	next                            string
	logout                          string
	search                          string
	search_results                  string
	filter                          string
	new                             string
	edit                            string
	save                            string
	cancel                          string
	delete                          string
	delete_confirmation             string
	delete_confirmation_description string
	// alerts
	records_found             string
	no_records_found          string
	record_saved_successfully string
	// fields
	id            string
	name          string
	amount        string
	date          string
	category      string
	color         string
	payer         string
	payee         string
	description   string
	username      string
	password      string
	role          string
	role_normal   string
	role_admin    string
	outside_world string
	total         string
	uncategorized string
	// balances
	debts                 string
	receivables           string
	total_expenses        string
	user_amount           string
	proportional_expenses string
	sent_expenses         string
	sent_transfer         string
	received_transfer     string
	result                string
	// calendar
	this_week  string
	this_month string
	this_year  string
	Monday     string
	Tuesday    string
	Wednesday  string
	Thursday   string
	Friday     string
	Saturday   string
	Sunday     string
	January    string
	February   string
	March      string
	April      string
	May        string
	June       string
	July       string
	August     string
	September  string
	October    string
	November   string
	December   string
}

var l10nMap = map[string]l10n{
	"de": l10nDe,
}

func Translate(message string) string {
	//todo: generalize function for allowing more languages, depending on user configuration / request parameters
	name := reflect.ValueOf(l10nDe).FieldByName(message)
	if name.IsValid() {
		return name.String()
	} else {
		return message
	}
}
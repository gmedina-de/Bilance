package translator

type l10n struct {
	// navigation
	payments        string `default:"Payments"`
	categories      string `default:"Categories"`
	analysis        string `default:"Analysis"`
	balances        string `default:"Balances"`
	expenses        string `default:"Expenses"`
	by_period       string `default:"By period"`
	by_category     string `default:"By category"`
	admin_functions string `default:"Admin functions"`
	users           string `default:"Users"`
	projects        string `default:"Projects"`

	previous                        string `default:"Previous"`
	next                            string `default:"Next"`
	logout                          string `default:"Logout"`
	search                          string `default:"Search"`
	search_results                  string `default:"Search results"`
	filter                          string `default:"Filter"`
	new                             string `default:"New"`
	edit                            string `default:"Edit"`
	save                            string `default:"Save"`
	cancel                          string `default:"Cancel"`
	delete                          string `default:"Delete"`
	delete_confirmation             string `default:"Delete confirmation"`
	delete_confirmation_description string `default:"Delete confirmation description"`

	records_found             string `default:"Records found"`
	no_records_found          string `default:"No records found"`
	record_saved_successfully string `default:"Record saved successfully"`

	id            string `default:"Id"`
	name          string `default:"Name"`
	amount        string `default:"Amount"`
	date          string `default:"Date"`
	category      string `default:"Category"`
	color         string `default:"Color"`
	payer         string `default:"Payer"`
	payee         string `default:"Payee"`
	description   string `default:"Description"`
	username      string `default:"Username"`
	password      string `default:"Password"`
	role          string `default:"Role"`
	role_normal   string `default:"Role normal"`
	role_admin    string `default:"Role admin"`
	outside_world string `default:"Outside world"`
	total         string `default:"Total"`
	uncategorized string `default:"Uncategorized"`

	debts                 string `default:"Debts"`
	receivables           string `default:"Receivables"`
	total_expenses        string `default:"Total expenses"`
	user_amount           string `default:"User amount"`
	proportional_expenses string `default:"Proportional expenses"`
	sent_expenses         string `default:"Sent expenses"`
	sent_transfer         string `default:"Sent transfer"`
	received_transfer     string `default:"Received transfer"`
	result                string `default:"Result"`

	this_week  string `default:"This week"`
	this_month string `default:"This month"`
	this_year  string `default:"This year"`
	monday     string `default:"Monday"`
	tuesday    string `default:"Tuesday"`
	wednesday  string `default:"Wednesday"`
	thursday   string `default:"Thursday"`
	friday     string `default:"Friday"`
	saturday   string `default:"Saturday"`
	sunday     string `default:"Sunday"`
	january    string `default:"January"`
	february   string `default:"February"`
	march      string `default:"March"`
	april      string `default:"April"`
	may        string `default:"May"`
	june       string `default:"June"`
	july       string `default:"July"`
	august     string `default:"August"`
	september  string `default:"September"`
	october    string `default:"October"`
	november   string `default:"November"`
	december   string `default:"December"`
}

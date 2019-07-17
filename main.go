// Sean at Shanghai
// convert alipay bill to beancount version

package main

// Config set all default values
type Config struct {
	DefaultCurrency string `json:"defaultCurrency"`
	DefaultExpensesAccount string `json:"defaultExpensesAccount"`
	Title string `json:"title"`
	AccoutMap map[string]AliBillAttr `json:"accountMap"`
	DefaultExpenseAccount string `json:"defaultExpensesAccount"`
	DefaultIncomeAccount string `json:"defaultIncomeAccount"`
}

func main() {
}

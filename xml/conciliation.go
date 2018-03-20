package xml

type BankAccount struct {
	BankCode          string `xml:"BankCode"`
	BankBranch        string `xml:"BankBranch"`
	BankAccountNumber string `xml:"BankAccountNumber"`
}

type Payment struct {
	Id                           int         `xml:"Id"`
	TotalAmount                  float64     `xml:"TotalAmount"`
	TotalFinancialAccountsAmount float64     `xml:"TotalFinancialAccountsAmount"`
	BankAccount                  BankAccount `xml:"FavoredBankAccount"`
}

type FinancialEvent struct {
	EventId     int     `xml:"EventId"`
	Type        int     `xml:"Type"`
	PaymentDate string  `xml:"PaymentDate"`
	Amount      float64 `xml:"Amount"`
}

type Header struct {
	StoneCode     string `xml:"StoneCode"`
	ReferenceDate string `xml:"ReferenceDate"`
}

type Conciliation struct {
	Header                 Header           `xml:"Header"`
	FinancialEventAccounts []FinancialEvent `xml:"FinancialEventAccounts>Event"`
	Payments               []Payment        `xml:"Payments>Payment"`
}


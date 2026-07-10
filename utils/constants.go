package utils

const (
	TransactionDeposit = "Deposit"
	TransactionWithdraw = "Withdraw"

	AccountClosed = "Closed"
	AccountActive = "Active"

	LoanActive = "Active"
	LoanPaid = "Paid"
	LoanDefaulted = "Defaulted"
)

func GetInterestRate(loanType string) float64 {

	switch loanType {

	case "Personal":
		return 11.0

	case "Home":
		return 8.5

	case "Car":
		return 9.0

	case "Education":
		return 7.5
	}

	return 0
}
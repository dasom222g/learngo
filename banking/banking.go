package banking

// private
// meber 변수
type account struct {
	owner   string
	balance int
}

// go 에서 값을 리턴하면 복사값을 리턴하는 것인데 메모리 낭비를 줄이기위해 참조값을 리턴함
func NewAccount(owner string) *account {
	newAccount := account{
		owner:   owner,
		balance: 0,
	}
	return &newAccount
}

// go method는 함수명 앞에 receiver를 붙여줌
func (account *account) Deposit(amount int) {
	account.balance += amount
}

func (account account) GetBalance() int {
	return account.balance
}

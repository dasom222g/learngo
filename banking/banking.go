package banking

import (
	"errors"
	"strconv"
)

// private
// meber 변수
type account struct {
	owner   string
	balance int
}

var errorNomoney error = errors.New("!! The balance is less than the withdrawal amount !!")

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

func (account *account) Withdraw(amount int) error {
	// 출금요청액이 잔액보다 많을경우 error 리턴
	if account.balance < amount {
		return errorNomoney
	}
	account.balance -= amount
	return nil
}

func (account *account) ChangeOwner(name string) {
	account.owner = name
}

func (account account) GetBalance() int {
	return account.balance
}

func (account account) GetOwner() string {
	return account.owner
}

func (account account) String() string {
	return account.GetOwner() + "'s account.\nHas: " + strconv.Itoa(account.GetBalance())
	// return fmt.Sprintf(account.GetOwner(), "'s account.\nHas: ", account.GetBalance())
}

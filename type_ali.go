// Sean at shanghai
// convert alipay bill to beancount

package main

// LocalTimeFmt set time format to utc+8
const LocalTimeFmt = "2006-01-02 15:04:05 -0700"


// alipay defined transanction type
type TxTypeType string
const (
	TxTypeSend TxTypeType = "支出"
	TxTypeRecv TxTypeType = "收入"
	TxTypeEmpty TxTypeType = ""
	TxTypeNil TxTypeType = "Nil"
)


type MoneyStatusType string
const (
	MoneySend MoneyStatusType = "已支出"
	MoneyRecv MoneyStatusType = "已收入"
	MoneyTransfer MoneyStatusType = "资金转移"
	MoneyStatusNil MoneyStatusType = "Nil"
)


const SizeOfAliBill = 17

func getMoneyStatus(str string) MoneyStatusType {
	switch str {
		case string(MoneySend):
			return MoneySend
		case string(MoneyRecv):
			return MoneyRecv
		case string(MoneyTransfer):
			return MoneyTransfer
		default:
			return MoneyStatusNil
	}
	return MoneyStatusNil
}


// MatchType defines how do we match a bill
type MatchType string
const (
	MatchTypeContain MatchType = "contain"
	MatchTypeEqual MatchType = "equal"
)

func getTxType(str string) TxTypeType {
	switch str {
		case string(TxTypeSend):
			return TxTypeSend
		case string(TxTypeRecv):
			return TxTypeRecv
		case string(TxTypeEmpty):
			return TxTypeEmpty
		default:
			return TxTypeNil
	}
	return TxTypeNil
}
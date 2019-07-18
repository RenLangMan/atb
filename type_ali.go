// Sean at shanghai
// convert alipay bill to beancount

package main

import (
	"log"
)

// LocalTimeFmt set time format to utc+8
const LocalTimeFmt = "2006-01-02 15:04:05 -0700"

// alipay defined transanction type
type TxTypeType string

const (
	TxTypeSend  TxTypeType = "支出"
	TxTypeRecv  TxTypeType = "收入"
	TxTypeEmpty TxTypeType = ""
	TxTypeNil   TxTypeType = "Nil"
)

type MoneyStatusType string

const (
	MoneySend      MoneyStatusType = "已支出"
	MoneyRecv      MoneyStatusType = "已收入"
	MoneyTransfer  MoneyStatusType = "资金转移"
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
	}
	return MoneyStatusNil
}

// MatchType defines how do we match a bill
type MatchType string

const (
	MatchTypeContain        MatchType = "contain"
	MatchTypeEqual          MatchType = "equal"
	MatchTypeNumEqual       MatchType = "eq"
	MatchTypeNumGreaterThan MatchType = "gt"
	MatchTypeNumLessThan    MatchType = "lt"
	MatchTypeNumRange       MatchType = "range"
)

func CheckConfig(conf *Config) error {
	for idx, attr := range conf.AccountList {
		if attr.StatusMatchMethod != "" && attr.StatusMatchMethod != MatchTypeContain && attr.StatusMatchMethod != MatchTypeEqual {
			log.Println("bad status match type at idx:", idx)
			return ErrBadMatchType
		}
		if attr.PeerMatchMethod != "" && attr.PeerMatchMethod != MatchTypeContain && attr.PeerMatchMethod != MatchTypeEqual {
			log.Println("bad peer match type at idx:", idx)
			return ErrBadMatchType
		}
		if attr.ItemNameMatchMethod != "" && attr.ItemNameMatchMethod != MatchTypeContain && attr.ItemNameMatchMethod != MatchTypeEqual {
			log.Println("bad itemName match type at idx:", idx)
			return ErrBadMatchType
		}
		if attr.MoneyMatchMethod == "" {
			continue
		}
		if attr.MoneyMatchMethod != MatchTypeNumEqual && attr.MoneyMatchMethod != MatchTypeNumGreaterThan &&
			attr.MoneyMatchMethod != MatchTypeNumLessThan && attr.MoneyMatchMethod != MatchTypeNumRange {
			log.Println("bad money match type at idx :", idx)
			return ErrBadMatchType
		}
		if attr.MoneyMatchMethod == MatchTypeNumRange && len(attr.Money) != 2 {
			log.Println("bad money match type at idx :", idx)
			return ErrBadMatchType
		}
		if attr.MoneyMatchMethod != MatchTypeNumRange && len(attr.Money) != 1 {
			log.Println("bad money match type at idx :", idx)
			return ErrBadMatchType
		}
		if attr.PlusAccount == "" || attr.MinusAccount == "" {
			log.Println("we do not allow empty account")
			return ErrBadAccount
		}
	}
	return nil
}

func getTxType(str string) TxTypeType {
	switch str {
	case string(TxTypeSend):
		return TxTypeSend
	case string(TxTypeRecv):
		return TxTypeRecv
	case string(TxTypeEmpty):
		return TxTypeEmpty
	}
	return TxTypeNil
}

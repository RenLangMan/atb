// Sean at shanghai
// convert alipay bill to bean

package main

import (
	"fmt"
	"log"
	"strings"
)

// AliBillAttr helps us determine which account it
// should go
type AliBillAttr struct {
	Status              []string  `json:"status"`
	StatusMatchMethod   MatchType `json:"statusMatchMethod"`
	Peer                []string  `json:"peer"`
	PeerMatchMethod     MatchType `json:"peerMatchMethod"`
	ItemName            []string  `json:"itemName"`
	ItemNameMatchMethod MatchType `json:"itemNameMatchMethod"`
	Money               []float64 `json:"money"`
	MoneyMatchMethod    MatchType `json:"moneyMatchMethod"`
	PlusAccount         string    `json:"plusAccount"`
	MinusAccount        string    `json:"minusAccount"`
}

func checkAttrNum(sets []float64, method MatchType, check float64) bool {
	if method == MatchTypeNumGreaterThan {
		if check > sets[0] {
			return true
		}
		return false
	}
	if method == MatchTypeNumLessThan {
		if check < sets[0] {
			return true
		}
		return false
	}
	if method == MatchTypeNumRange {
		if check > sets[0] && check < sets[1] {
			return true
		}
		return false
	}
	if check == sets[0] {
		return true
	}
	return false
}

func checkAttr(sets []string, method MatchType, check string) bool {
	// if set is null slice, it should be handled by caller
	if method == MatchTypeContain {
		for _, set := range sets {
			if strings.Contains(check, set) {
				return true
			}
		}
		return false
	}
	for _, set := range sets {
		if check == set {
			return true
		}
	}
	return false
}

func getAccount(bill AliBill, list []AliBillAttr) (string, string) {
	for _, attr := range list {
		if len(attr.Peer) != 0 && checkAttr(attr.Peer, attr.PeerMatchMethod, bill.Peer) == false {
			continue
		}
		if len(attr.ItemName) != 0 && checkAttr(attr.ItemName, attr.ItemNameMatchMethod, bill.ItemName) == false {
			continue
		}
		if len(attr.Status) != 0 && checkAttr(attr.Status, attr.StatusMatchMethod, bill.Status) == false {
			continue
		}
		if len(attr.Money) != 0 && checkAttrNum(attr.Money, attr.MoneyMatchMethod, bill.Money) == false {
			continue
		}
		return attr.PlusAccount, attr.MinusAccount
	}
	// default account
	return "", ""
}

// FillBills fill plus and minus account for every bill
func FillBills(list []AliBillAttr) error {
	for idx, bill := range AliBillList {
		plus, minus := getAccount(bill, list)
		if plus == "" && *flagStrict {
			log.Println("no default account in strict mode", bill)
			printBill(bill)
			return ErrNoDefault
		}
		AliBillList[idx].MinusAccount = minus
		AliBillList[idx].PlusAccount = plus
	}
	return nil
}

func printBill(bill AliBill) {
	fmt.Println("---------------------")
	fmt.Println("type", bill.Type)
	fmt.Println("peer", bill.Peer)
	fmt.Println("itemname", bill.ItemName)
	fmt.Println("txtype", bill.TxType)
	fmt.Println("status", bill.Status)
	fmt.Println("servicefee", bill.ServiceFee)
	fmt.Println("refund", bill.Refund)
	fmt.Println("comment", bill.Comment)
	fmt.Println("moneyStatus", bill.MoneyStatus)
	fmt.Println("plus", bill.PlusAccount)
	fmt.Println("minus", bill.MinusAccount)
	fmt.Println("money", bill.Money)
}

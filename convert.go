// Sean at shanghai
// convert alipay bill to bean

package main


import (
	"fmt"
	"strings"
	"log"
)

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
		return attr.PlusAccount, attr.MinusAccount
	}
	// default account
	return "", ""
}


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
		// printBill(AliBillList[idx])
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
}

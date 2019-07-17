// Sean at Shanghai
// convert alipay bill to beancount

package main


import (
	"os"
	"log"
	"io"
	"fmt"
)

func WriteBean(conf *Config) error {
	file, err := os.Create(*flagOutput)
	if err != nil {
		log.Println("create output file error:", *flagOutput, err)
		return err
	}
	defer file.Close()
	// write header
	_, err = io.WriteString(file, "option \"title\" \"" + conf.Title + "\"\n")
	if err != nil {
		log.Println("write option title error:", err)
		return err
	}
	_, err = io.WriteString(file, "option \"operating_currency\" \"" + conf.DefaultCurrency + "\"\n\n")
	if err != nil {
		log.Println("write option currency error:", err)
		return err
	}
	// write open account
	uniqMap := make(map[string]bool)
	for _, attr := range conf.AccountList {
		uniqMap[attr.MinusAccount] = true
		uniqMap[attr.PlusAccount] = true
	}
	uniqMap[conf.DefaultPlusAccount] = true
	uniqMap[conf.DefaultMinusAccount] = true
	for k, _ := range uniqMap {
		_, err = io.WriteString(file, "1970-01-01 open " + k + "\n")
		if err != nil {
			log.Println("write open account error:", err)
			return err
		}
	}
	_, err = io.WriteString(file, "\n")
	if err != nil {
		log.Println("write extra enter error:", err)
		return err
	}

	for _, bill := range AliBillList {
		err = writeBill(file, bill, conf)
		if err != nil {
			log.Println("write bill error:", err)
			return err
		}
	}
	return nil
}


func writeBill(file *os.File, bill AliBill, conf *Config) error {
	str := bill.CreateTime.Format("2006-01-02")
	str = str + " * \"" + bill.Peer + "\" \"" + bill.ItemName + "\"\n"
	_, err := io.WriteString(file, str)
	if err != nil {
		return err
	}
	str = "\t"
	if bill.PlusAccount == "" {
		bill.PlusAccount = conf.DefaultPlusAccount
	}
	if bill.MinusAccount == "" {
		bill.MinusAccount = conf.DefaultMinusAccount
	}
	if bill.MoneyStatus == MoneySend {
		str = str + bill.PlusAccount + " " + fmt.Sprintf("%.2f", bill.Money) + " "
		str = str + conf.DefaultCurrency + "\n"
		str = str + "\t" + bill.MinusAccount + " -" + fmt.Sprintf("%.2f", bill.Money) + " "
		str = str + conf.DefaultCurrency + "\n\n"
	} else {
		str = str + bill.MinusAccount + " " + fmt.Sprintf("%.2f", bill.Money) + " "
		str = str + conf.DefaultCurrency + "\n"
		str = str + "\t" + bill.PlusAccount + " -" + fmt.Sprintf("%.2f", bill.Money) + " "
		str = str + conf.DefaultCurrency + "\n\n"
	}
	_, err = io.WriteString(file, str)
	if err != nil {
		return err
	}
	return nil
}

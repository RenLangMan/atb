// Sean at Shanghai
// convert alipay bill to beancount version

package main

import (
	"github.com/dilfish/tools"
	"strings"
	"fmt"
	"errors"
	"time"
)


// ErrBadFmt indicate not a valid alipay bill
var ErrBadFmt = errors.New("bad format alipay bill")


// Config set all default values
type Config struct {
	// 
}


type TxTypeType int
const (
)


type AliBill struct {
	DealNo string `json:"dealNo"` // 交易号
	OrderNo string `json:"orderNo"` // 商家订单号
	CreateTime time.Time `json:"createTime"` // 交易创建时间
	PayTime time.Time `json:"payTime"` // 付款时间
	LastUpdate time.Time `json:"lastUpdate"` // 最近修改时间
	DealSrc string `json:"dealSrc"` // 交易来源地
	Type string `json:"type"` // 类型
	Peer string `json:"peer"` // 交易对方
	ItemName string `json:"itemName"` // 商品名称
	Money int64 `json:"money"` // 金额
	TxType TxTypeType `json:"txType"` // 收/支
	Status string `json:"status"` // 交易状态
	ServiceFee int64 `json:"serviceFee"` // 服务费
	Refund int64 `json:"refund"` // 成功退款
	Comment string `json:"comment"` // 备注
	MoneyStatus string `json:"moneyStatus"` // 资金状态
}


func handleRe(arr []string) error {
	fmt.Println("退款不能处理")
	return nil
}


func judge(tm, who, count string) string {
	if strings.Contains(who, "老婆") {
		str := tm + " * \"老婆\" \"挥霍\"\n"
		str = str + "\tExpenses:Spouse:Smash " + count + " CNY\n"
		str = str + "\tLiabilities:CreditCard:Dual:Cmb -" + count + " CNY\n"
		return str
	}
	if strings.Contains(who, "滴滴") {
		str := tm + " * \"打车\" \"滴滴\"\n"
		str = str + "\tExpenses:Traffic:Taxi " + count + " CNY\n"
		str = str + "\tLiabilities:CreditCard:Dual:Cmb -" + count + " CNY\n"
		return str
	}
	if strings.Contains(who, "麦当劳") || strings.Contains(who, "肯德基") || strings.Contains(who, "汉堡王"){
		str := tm + " * \"吃饭\" \"享受\"\n"
		str = str + "\tExpenses:Food:EatHere " + count + " CNY\n"
		str = str + "\tLiabilities:CreditCard:Dual:Cmb -" + count + " CNY\n"
		return str
	}
	if strings.Contains(who, "饿了么") {
		str := tm + " * \"吃饭\" \"外卖\"\n"
		str = str + "\tExpenses:Food:Takeout " + count + " CNY\n"
		str = str + "\tLiabilities:CreditCard:Dual:Cmb -" + count + " CNY\n"
		return str
	}
	str := tm + " * \"其他\" \"烟酒烫头\"\n"
	str = str + "\tExpenses:Other:Smoke " + count + " CNY\n"
	str = str + "\tLiabilities:CreditCard:Dual:Cmb -" + count + " CNY\n"
	return str
}


func cb(line string) error {
	arr := strings.Split(line, ",")
	if len(arr) != 17 {
		fmt.Println("we have len of", len(arr))
		return ErrBadFmt
	}
	for idx, a := range arr {
		arr[idx] = strings.Trim(a, " ")
		arr[idx] = strings.Trim(arr[idx], "\t")
	}
	tm := strings.Split(arr[3], " ")
	if strings.Contains(arr[11], "交易成功") || strings.Contains(arr[11], "亲情号付款成功") {
		fmt.Println(judge(tm[0], arr[7], arr[9]))
	} else {
		fmt.Println("ERROR", tm, arr)
	}
	return nil
}


func main() {
	err := tools.ReadLine("ali.txt", cb)
	if err != nil {
		panic(err)
	}
}

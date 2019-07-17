// Sean at Shanghai
// convert alipay to beancount version

package main


import (
	"github.com/dilfish/tools"
	"time"
	"log"
	"strings"
	"strconv"
)


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
var LineNum int


// AliBillList stores all of tx in a file
var AliBillList []AliBill


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
	Money float64 `json:"money"` // 金额
	TxType TxTypeType `json:"txType"` // 收/支
	Status string `json:"status"` // 交易状态
	ServiceFee float64 `json:"serviceFee"` // 服务费
	Refund float64 `json:"refund"` // 成功退款
	Comment string `json:"comment"` // 备注
	MoneyStatus MoneyStatusType `json:"moneyStatus"` // 资金状态
	// below is filled at runtime
	MinusAccount string `json:"minusAccount"`
	PlusAccount string `json:"plusAccount"`
}


// MatchType defines how do we match a bill
type MatchType string
const (
	MatchTypeContain MatchType = "contain"
	MatchTypeEqual MatchType = "equal"
)

// AliBillAttr helps us determine which account it
// should go
type AliBillAttr struct {
	Status []string `json:"status"`
	StatusMatchMethod MatchType `json:"statusMatchMethod"`
	Peer []string `json:"peer"`
	PeerMatchMethod MatchType `json:"peerMatchMethod"`
	ItemName []string `json:"itemName"`
	ItemNameMatchMethod MatchType `json:"itemNameMatchMethod"`
	PlusAccount string `json:"plusAccount"`
	MinusAccount string `json:"minusAccount"`
}


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


func parseAlipayBill(line string) error {
	var err error
	array := strings.Split(line, ",")
	if len(array) != SizeOfAliBill {
		log.Println("sizeof line is not good, it's ", len(array), " and we expect", SizeOfAliBill)
		return ErrBadAliFmt
	}
	LineNum = LineNum + 1
	// we ignore first title line
	if LineNum == 1 {
		return nil
	}
	for idx, a := range array {
		a = strings.Trim(a, " ")
		a = strings.Trim(a, "\t")
		array[idx] = a
	}
	var bill AliBill
	bill.DealNo = array[0]
	bill.OrderNo = array[1]
	bill.CreateTime, err = time.Parse(LocalTimeFmt, array[2] + " +0800")
	if err != nil {
		log.Println("parse create time error:", array[2], err)
		return err
	}
	if array[3] != "" {
		bill.PayTime, err = time.Parse(LocalTimeFmt, array[3] + " +0800")
		if err != nil {
			log.Println("parse paytime error:", array[3], err, array)
			return err
		}
	}
	bill.LastUpdate, err = time.Parse(LocalTimeFmt, array[4] + " +0800")
	if err != nil {
		log.Println("parse last update error:", array[4], err)
		return err
	}
	bill.DealSrc = array[5]
	bill.Type = array[6]
	bill.Peer = array[7]
	bill.ItemName = array[8]
	bill.Money, err = strconv.ParseFloat(array[9], 32)
	if err != nil {
		log.Println("parse money error:", array[9], err)
		return err
	}
	bill.TxType = getTxType(array[10])
	if bill.TxType == TxTypeNil {
		log.Println("get tx type error:", array[10], array)
		return ErrBadTxType
	}
	bill.Status = array[11]
	bill.ServiceFee, err = strconv.ParseFloat(array[12], 32)
	if err != nil {
		log.Println("parse service fee error:", array[12], err)
		return err
	}
	bill.Refund, err = strconv.ParseFloat(array[13], 32)
	if err != nil {
		log.Println("parse refund error:", array[13], err)
		return err
	}
	bill.Comment = array[14]
	bill.MoneyStatus = getMoneyStatus(array[15])
	if bill.MoneyStatus == MoneyStatusNil {
		log.Println("get money status error:", err, array[15])
		return err
	}
	AliBillList = append(AliBillList, bill)
	return nil
}


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


// ReadAliBill check all lines of bill
func ReadAliBill(fn string) error {
	err := tools.ReadLine(fn, parseAlipayBill)
	if err != nil {
		log.Println("read bills error:", err)
		return err
	}
	return nil
}

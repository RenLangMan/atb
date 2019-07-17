// Sean at Shanghai
// convert alipay to beancount version

package main


import (
	"github.com/dilfish/tools"
	"time"
)


// alipay defined transanction type
type TxTypeType int
const (
)


// AliBillList stores all of tx in a file
var AliBillList []AliBill
// AccountMap store all accounts defined in bills
// if it is not defines, we return an error
var AccountMap map[string]bool


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


// MatchType defines how do we match a bill
type MatchType int
const (
	MatchTypeContain MatchType = 0
	MatchTypeEqual MatchType = 1
)

// AliBillAttr helps us determine which account it
// should go
type AliBillAttr struct {
	DealSrc string `json:"dealSrc"`
	DealSrcMatchMethod MatchType `json:"dealSrcMatchMethod"`
}


func parseAlipayBill(line string) error {
	return nil
}


func ReadAliBill(fn string) error {
	return tools.ReadLine(fn, parseAlipayBill)
}

// Sean at Shanghai
// convert alipay bill to beancount

package main

import (
	"errors"
)


// ErrBadAliFmt indicates a bad format alipay bill file
var ErrBadAliFmt = errors.New("bad alipay bill format")
// ErrNoDefault indicates we do not use default account in strict mode
var ErrNoDefault = errors.New("no default account in strict mode")
// ErrBadTxType indicates we could not get the tx type
var ErrBadTxType = errors.New("we could not recgnize the tx type")

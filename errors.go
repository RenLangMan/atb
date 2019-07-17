// Sean at Shanghai
// convert alipay bill to beancount

package main

import (
	"errors"
)


// ErrBadAliFmt indicates a bad format alipay bill file
var ErrBadAliFmt = errors.New("bad alipay bill format")

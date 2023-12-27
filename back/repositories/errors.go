package repositories

import "errors"

var ErrUnknown = errors.New("unknown error")
var ErrRecordNotFound = errors.New("record not found")

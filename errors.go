package smslite

import "errors"

var (
	ErrNonGSM7        = errors.New("smslite: text contains non GSM-7 character")
	ErrInvalidUDH     = errors.New("smslite: invalid UDH")
	ErrTooManyParts   = errors.New("smslite: segment count exceeds maximum")
	ErrInvalidUDHPart = errors.New("smslite: invalid UDH part values")
)

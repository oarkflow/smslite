package smslite

import "fmt"

func BuildUDH(opt Options, total, part int) UDHInfo {
	if opt.ReferenceNumberSize == 16 {
		return BuildConcatUDH16(opt.ReferenceNumber, byte(total), byte(part))
	}
	return BuildConcatUDH8(byte(opt.ReferenceNumber), byte(total), byte(part))
}

func BuildConcatUDH8(ref byte, total, part byte) UDHInfo {
	return UDHInfo{Required: true, HeaderLength: 6, HeaderLengthSeptets: 7, ReferenceNumber: int(ref), ReferenceSize: 8, PartNumber: int(part), TotalParts: int(total), Hex: fmt.Sprintf("05 00 03 %02X %02X %02X", ref, total, part)}
}

func BuildConcatUDH16(ref uint16, total, part byte) UDHInfo {
	return UDHInfo{Required: true, HeaderLength: 7, HeaderLengthSeptets: 8, ReferenceNumber: int(ref), ReferenceSize: 16, PartNumber: int(part), TotalParts: int(total), Hex: fmt.Sprintf("06 08 04 %02X %02X %02X %02X", byte(ref>>8), byte(ref), total, part)}
}

func ParseUDH(data []byte) (UDHInfo, error) {
	if len(data) < 5 {
		return UDHInfo{}, ErrInvalidUDH
	}
	if data[0] == 0x05 && len(data) >= 6 && data[1] == 0x00 && data[2] == 0x03 {
		if data[4] == 0 || data[5] == 0 || data[5] > data[4] {
			return UDHInfo{}, ErrInvalidUDHPart
		}
		return BuildConcatUDH8(data[3], data[4], data[5]), nil
	}
	if data[0] == 0x06 && len(data) >= 7 && data[1] == 0x08 && data[2] == 0x04 {
		if data[5] == 0 || data[6] == 0 || data[6] > data[5] {
			return UDHInfo{}, ErrInvalidUDHPart
		}
		ref := uint16(data[3])<<8 | uint16(data[4])
		return BuildConcatUDH16(ref, data[5], data[6]), nil
	}
	return UDHInfo{}, ErrInvalidUDH
}

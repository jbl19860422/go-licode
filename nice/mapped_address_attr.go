package nice

import (
	"encoding/binary"
	"errors"
)

type MAPPED_ADDRESS_FAMILY byte
const (
	_ MAPPED_ADDRESS_FAMILY 	= 	iota
	MAPPED_ADDRESS_FAMILY_IPV4 	=	0x01
	MAPPED_ADDRESS_FAMILY_IPV6 	= 	0x02
)

type StunMappedAddressAttr struct {
	header 			StunAttrHeader
	zero 			byte
	family			MAPPED_ADDRESS_FAMILY
	port 			uint16
	ip 				[]byte	//If the address family is IPv6, the address MUST be 128 bits,All fields must be in network byte order
	padding 		[]byte
}

func NewStunMappedAddressAttr(f MAPPED_ADDRESS_FAMILY, p uint16, ip []byte) *StunMappedAddressAttr {
	s := &StunMappedAddressAttr{
		header:StunAttrHeader{
			typ:STUN_ATTRIBUTE_MAPPED_ADDRESS,
			len:0,
		},
		zero:0x00,
		family:f,
		port:p,
		ip:ip,
	}
	return s
}

func (this StunMappedAddressAttr) Encode(stream *DataStream) error {
	this.header.Encode(stream)
	stream.WriteByte(byte(this.family))
	stream.WriteUInt16(this.port, binary.BigEndian)
	stream.WriteBytes(this.ip)
	return nil
}

func (this *StunMappedAddressAttr) Decode(stream *DataStream) (err error) {
	_, err = stream.ReadByte()	//zero byte
	if err != nil {
		return
	}

	var b byte
	b , err = stream.ReadByte()
	if err != nil {
		return
	}
	this.family = MAPPED_ADDRESS_FAMILY(b)
	if this.family != 0x01 && this.family != 0x02 {
		err = errors.New("invalid family type")
		return
	}

	this.port, err = stream.ReadUInt16(binary.BigEndian)
	if err != nil {
		return
	}

	this.ip = stream.ReadLeftBytes()
	if len(this.ip) != 4 && len(this.ip) != 16 {
		err = errors.New("invalid ip len")
		return
	}
	return
}
package nice

import (
	"encoding/binary"
	"errors"
)

type StunXorMappedAddressAttrValue struct {
	zero 			byte
	family			MAPPED_ADDRESS_FAMILY
	port 			uint16
	ip 				[]byte	//If the address family is IPv6, the address MUST be 128 bits,All fields must be in network byte order
	magicCookie 	*StunMessageMagicCookie	//用于編解碼
	transactionId 	*StunTransactionId		//用于編解碼
}

func NewStunXorMappedAddressAttrValue(f MAPPED_ADDRESS_FAMILY, p uint16, ip []byte) *StunXorMappedAddressAttrValue {
	s := &StunXorMappedAddressAttrValue{
		zero:0x00,
		family:f,
		port:p,
		ip:ip,
	}
	return s
}

func (this *StunXorMappedAddressAttrValue) SetMagicCookie(m *StunMessageMagicCookie) {
	this.magicCookie = m
}

func (this *StunXorMappedAddressAttrValue) SetTransactionId(s *StunTransactionId) {
	this.transactionId = s
}

func (this StunXorMappedAddressAttrValue) GetSize() uint16 {
	return 4 + uint16(len(this.ip))
}

func (this StunXorMappedAddressAttrValue) Encode(stream *DataStream) error {
	stream.WriteByte(this.zero)
	stream.WriteByte(byte(this.family))
	if this.magicCookie == nil {
		return errors.New("need magic cookie to encode xor mapped address attr")
	}
	m,_ := BytesToUInt16(*this.magicCookie, HostByteOrder)
	p := this.port ^ m
	stream.WriteUInt16(p, binary.BigEndian)

	if len(this.ip) == 4 {
		q, err := XOR(this.ip, *this.magicCookie)
		if err != nil {
			return err
		}
		stream.WriteBytes(ToNetworkOrder(q))
	} else if len(this.ip) == 16 {
		d := make([]byte, 0)
		d = append(d, *this.magicCookie...)
		d = append(d, *this.transactionId...)
		q, err := XOR(this.ip, d)
		if err != nil {
			return err
		}
		stream.WriteBytes(ToNetworkOrder(q))
	} else {
		return errors.New("ip len error")
	}
	return nil
}

func (this *StunXorMappedAddressAttrValue) Decode(stream *DataStream) (err error) {
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
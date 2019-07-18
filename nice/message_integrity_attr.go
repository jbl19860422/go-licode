package nice

type StunMessageIntegrityAttrValue struct {
	hmac 				[]byte 			//Since it uses the SHA1 hash, the HMAC will be 20 bytes.
	data 				[]byte			//the data to be encode
}

func (this *StunMessageIntegrityAttrValue) SetData(d []byte) {
	this.data = d
}

func (this StunMessageIntegrityAttrValue) Encode(stream *DataStream) error {
	return nil
}

func (this *StunMessageIntegrityAttrValue) Decode(stream *DataStream) error {
	return nil
}

func (this StunMessageIntegrityAttrValue) GetSize() uint16 {
	return 20
}
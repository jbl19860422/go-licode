package nice

type StunUsernameAttrValue struct {
	username 			string
}

func (this *StunUsernameAttrValue) SetUsername(n string) {
	this.username = n
}

func (this StunUsernameAttrValue) Encode(stream *DataStream) error {
	stream.WriteString(this.username)
	return nil
}

func (this *StunUsernameAttrValue) Decode(stream *DataStream) error {
	this.username = string(stream.ReadLeftBytes())
	return nil
}

func (this StunUsernameAttrValue) GetSize() uint16 {
	return uint16(len(this.username))
}
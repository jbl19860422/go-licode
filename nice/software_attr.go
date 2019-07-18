package nice

type StunSoftwareAttrValue struct {
	name			string
}

func (this StunSoftwareAttrValue) Encode(stream *DataStream) error {
	stream.WriteBytes([]byte(this.name))
	return nil
}

func (this *StunSoftwareAttrValue) Decode(stream *DataStream) error {
	d := stream.ReadLeftBytes()
	this.name = string(d)
	return nil
}

func (this StunSoftwareAttrValue) GetSize() uint16 {
	return uint16(len(this.name))
}
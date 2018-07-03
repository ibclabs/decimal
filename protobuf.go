package decimal

func (d Decimal) Marshal() ([]byte, error) {
	return d.MarshalBinary()
}

func (d Decimal) MarshalTo(data []byte) (n int, err error) {
	bytez, err := d.MarshalBinary()
	if err != nil {
		return 0, err
	}
	copy(data, bytez)
	return len(bytez), nil
}

func (d *Decimal) Unmarshal(data []byte) error {
	return d.UnmarshalBinary(data)
}

func (d *Decimal) Size() int {
	b, e := d.MarshalBinary()
	if e != nil {
		return 0
	}
	return len(b)
}

func (d Decimal) Compare(d2 Decimal) int {
	return d.Cmp(d2)
}
package decimal

import "errors"

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

func (d NullDecimal) Marshal() ([]byte, error) {
	data := make([]byte, 1)
	if d.Valid {
		data[0] = 1
		raw, err := d.Decimal.Marshal()
		if err != nil {
			return nil, err
		}
		data = append(data, raw...)
	} else {
		data[0] = 0
	}
	return data, nil
}

func (d NullDecimal) MarshalTo(data []byte) (n int, err error) {
	bytez, err := d.Marshal()
	if err != nil {
		return 0, err
	}
	copy(data, bytez)
	return len(bytez), nil
}

func (d *NullDecimal) Unmarshal(data []byte) error {
	if len(data) < 1 {
		return errors.New("too short data")
	}
	switch data[0] {
	case 0:
		{
			d.Valid = false
			d.Decimal = Decimal{}
			return nil
		}
	case 1:
		{
			d.Valid = true
			d.Decimal = Decimal{}
			return d.Decimal.UnmarshalBinary(data[1:])
		}
	}
	return errors.New("wrong flag for valid")
}

func (d *NullDecimal) Size() int {
	b, e := d.Marshal()
	if e != nil {
		return 0
	}
	return len(b)
}

func (d NullDecimal) Compare(d2 NullDecimal) int {
	var result int
	switch {
	case d.Valid == false && d2.Valid == false:
		{
			result = 0
		}
	case d.Valid == true && d2.Valid == false:
		{
			result = 1
		}
	case d.Valid == false && d2.Valid == true:
		{
			result = -1
		}
	default:
		{
			result = d.Decimal.Cmp(d2.Decimal)
		}
	}
	return result
}

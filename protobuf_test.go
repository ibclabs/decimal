package decimal

import (
	"bytes"
	"testing"
)

func TestDecimalMarshal(t *testing.T) {
	var (
		dec Decimal
		err error
	)
	for _, item := range testTable {
		dec, err = NewFromString(item.short)
		if err != nil {
			t.Error(err)
			continue
		}
		t.Log(dec.String())
		testInternalDecimalMarshal(t, dec)
	}
}

func testInternalDecimalMarshal(t *testing.T, dec1 Decimal) {
	// check marshalling
	data1, err := dec1.Marshal()
	if err != nil {
		t.Error(err)
		return
	}
	size := dec1.Size()
	data2 := make([]byte, size)
	num, err := dec1.MarshalTo(data2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(data1, data2) {
		t.Error("marshalled data non-equal", data1, data2)
		return
	}
	// check size
	if num != size {
		t.Errorf("size mismatch: %d vs %d", num, size)
		return
	}
	// check unmarshal
	dec2 := Decimal{}
	err = dec2.Unmarshal(data1)
	if err != nil {
		t.Error(err)
		return
	}
	if !dec1.Equal(dec2) {
		t.Errorf("value mismatch: %s vs %s", dec1.String(), dec2.String())
		return
	}
}

func TestNullDecimalMarshal(t *testing.T) {
	var (
		raw Decimal
		dec NullDecimal
		err error
	)
	for _, item := range testTable {
		raw, err = NewFromString(item.short)
		if err != nil {
			t.Error(err)
			continue
		}
		dec = NullDecimal{Decimal: raw, Valid: true}
		t.Logf("%t(%s)", dec.Valid, dec.Decimal.String())
		testInternalNullDecimalMarshal(t, dec)

		dec = NullDecimal{Decimal: raw, Valid: false}
		t.Logf("%t(%s)", dec.Valid, dec.Decimal.String())
		testInternalNullDecimalMarshal(t, dec)
	}
}

func TestNullDecimalUnmarshalNil(t *testing.T) {
	dec := NullDecimal{Decimal: New(5, 0), Valid: true}
	err := dec.Unmarshal(nil)
	if err != nil {
		t.Error(err)
		return
	}
	if dec.Valid != false {
		t.Error("valid is not FALSE", dec.Valid)
		return
	}
	if dec.Decimal.value != nil {
		t.Error("decimal.value is not NIL", dec.Decimal.value)
		return
	}
	if dec.Decimal.exp != 0 {
		t.Error("decimal.ex[ is not ZERO", dec.Decimal.exp)
		return
	}
}

func testInternalNullDecimalMarshal(t *testing.T, dec1 NullDecimal) {
	// check marshalling
	data1, err := dec1.Marshal()
	if err != nil {
		t.Error(err)
		return
	}
	size := dec1.Size()
	data2 := make([]byte, size)
	num, err := dec1.MarshalTo(data2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(data1, data2) {
		t.Error("marshalled data non-equal", data1, data2)
		return
	}
	// check size
	if num != size {
		t.Errorf("size mismatch: %d vs %d", num, size)
		return
	}
	// check unmarshal
	dec2 := NullDecimal{}
	err = dec2.Unmarshal(data1)
	if err != nil {
		t.Error(err)
		return
	}
	if dec1.Compare(dec2) != 0 {
		t.Errorf("value mismatch: %t(%s) vs %t(%s)", dec1.Valid, dec1.Decimal.String(), dec2.Valid, dec2.Decimal.String())
		return
	}
}

func TestNullDecimal_Compare(t *testing.T) {
	n0 := NullDecimal{Decimal: New(5, 0), Valid: true}

	var res int

	n1 := NullDecimal{Decimal: Decimal{}, Valid: false}
	res = n0.Compare(n1)
	if res != 1 {
		t.Errorf("size mismatch: %d vs %d", res, 1)
		return
	}
	res = n1.Compare(n0)
	if res != -1 {
		t.Errorf("size mismatch: %d vs %d", res, -1)
		return
	}

	n2 := NullDecimal{Decimal: New(5, 0), Valid: false}
	res = n0.Compare(n2)
	if res != 1 {
		t.Errorf("size mismatch: %d vs %d", res, 1)
		return
	}
	res = n1.Compare(n2)
	if res != 0 {
		t.Errorf("size mismatch: %d vs %d", res, 0)
		return
	}

	n3 := NullDecimal{Decimal: New(4, 0), Valid: true}
	res = n0.Compare(n3)
	if res != 1 {
		t.Errorf("size mismatch: %d vs %d", res, 1)
		return
	}

	n4 := NullDecimal{Decimal: New(5, 0), Valid: true}
	res = n0.Compare(n4)
	if res != 0 {
		t.Errorf("size mismatch: %d vs %d", res, 0)
		return
	}

	n5 := NullDecimal{Decimal: New(6, 0), Valid: true}
	res = n0.Compare(n5)
	if res != -1 {
		t.Errorf("size mismatch: %d vs %d", res, -1)
		return
	}
}

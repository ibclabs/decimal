package decimal

func (d NullDecimal) Equal(d2 NullDecimal) bool {
	return d.Compare(d2) == 0
}

func (d NullDecimal) String() string {
	if d.Valid {
		return d.Decimal.String()
	}
	return ""
}

package theater

const ZeroMoney = Money(0)

type Money float64

func (m Money) plus(amount Money) Money {
	return m + amount
}

func (m Money) minus(amount Money) Money {
	return m - amount
}

func (m Money) times(percent float64) Money {
	return m * Money(percent)
}

func (m Money) isLessThen(other Money) bool {
	return m < other
}

func (m Money) isGreaterThanOrEqual(other Money) bool {
	return m >= other
}

package theater

import "time"

type Movie struct {
	title          string
	runningTime    time.Duration
	fee            Money
	discountPolicy DiscountPolicy
}

func (m *Movie) GetFee() Money {
	return m.fee
}

func (m *Movie) CalculateMovieFee(screening Screening) Money {
	return m.fee.minus(m.discountPolicy.CalculateDiscountAmount(screening))
}

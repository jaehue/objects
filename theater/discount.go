package theater

import (
	"time"
)

type DiscountPolicy interface {
	CalculateDiscountAmount(screening Screening) Money
}

type AmountDiscountPolicy struct {
	discountAmount    Money
	discountCondition []DiscountCondition
}

func (dp AmountDiscountPolicy) CalculateDiscountAmount(screening Screening) Money {
	// TODO 코드 중복
	for _, cond := range dp.discountCondition {
		if cond.IsSatisfiedBy(screening) {
			return dp.discountAmount
		}
	}

	return ZeroMoney
}

type PercentDiscountPolicy struct {
	percent           float64
	discountCondition []DiscountCondition
}

func (dp PercentDiscountPolicy) CalculateDiscountAmount(screening Screening) Money {
	// TODO 코드 중복
	for _, cond := range dp.discountCondition {
		if cond.IsSatisfiedBy(screening) {
			return screening.GetMovieFee().times(dp.percent)
		}
	}

	return ZeroMoney
}

type DiscountCondition interface {
	IsSatisfiedBy(screening Screening) bool
}

type SequenceCondition struct {
	sequence uint8
}

func (sc SequenceCondition) IsSatisfiedBy(screening Screening) bool {
	return screening.IsSequence(sc.sequence)
}

type PeriodCondition struct {
	dayOfWeek time.Weekday
	startTime time.Time
	endTime   time.Time
}

func (pc PeriodCondition) IsSatisfiedBy(screening Screening) bool {
	const timeFormat = "15:04:05"
	whenScreenedTime, _ := time.Parse(timeFormat, screening.whenScreened.Format(timeFormat))

	return screening.whenScreened.Weekday() == pc.dayOfWeek &&
		whenScreenedTime.After(pc.startTime) &&
		whenScreenedTime.Before(pc.endTime)

}

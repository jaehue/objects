package theater

import (
	"gotest.tools/assert"
	"testing"
	"time"
)

func TestCalculateMovieFee(t *testing.T) {

	t.Run("AmountDiscountPolicy", func(t *testing.T) {
		// given
		const timeFormat = "15:04:05"
		mondayStartTime, _ := time.Parse(timeFormat, "10:00:00")
		mondayEndTime, _ := time.Parse(timeFormat, "11:59:00")
		tuesdayStartTime, _ := time.Parse(timeFormat, "10:00:00")
		tuesdayEndTime, _ := time.Parse(timeFormat, "20:59:00")

		avatar := Movie{
			title:       "아바타",
			runningTime: time.Duration(100),
			fee:         10000,
			discountPolicy: AmountDiscountPolicy{
				discountAmount: Money(800),
				discountCondition: []DiscountCondition{
					SequenceCondition{
						sequence: 1,
					},
					SequenceCondition{
						sequence: 10,
					},
					PeriodCondition{
						dayOfWeek: time.Monday,
						startTime: mondayStartTime,
						endTime:   mondayEndTime,
					},
					PeriodCondition{
						dayOfWeek: time.Tuesday,
						startTime: tuesdayStartTime,
						endTime:   tuesdayEndTime,
					},
				},
			},
		}

		whenScreened, _ := time.Parse("2006-01-02 15:04:05", "2019-08-12 10:10:00")
		screening := Screening{
			movie:        avatar,
			sequence:     2,
			whenScreened: whenScreened,
		}

		// when
		movieFee := avatar.CalculateMovieFee(screening)

		// then
		assert.Equal(t, movieFee, Money(9200.0))
	})

	t.Run("PercentDiscountPolicy", func(t *testing.T) {
		// given
		const timeFormat = "15:04:05"
		tuesdayStartTime, _ := time.Parse(timeFormat, "14:00:00")
		tuesdayEndTime, _ := time.Parse(timeFormat, "16:59:00")

		thursdayStartTime, _ := time.Parse(timeFormat, "10:00:00")
		thursdayEndTime, _ := time.Parse(timeFormat, "13:59:00")

		titanic := Movie{
			title:       "타이타닉",
			runningTime: time.Duration(180),
			fee:         10000,
			discountPolicy: PercentDiscountPolicy{
				percent: 0.1,
				discountCondition: []DiscountCondition{
					PeriodCondition{
						dayOfWeek: time.Tuesday,
						startTime: tuesdayStartTime,
						endTime:   tuesdayEndTime,
					},
					SequenceCondition{
						sequence: 10,
					},
					PeriodCondition{
						dayOfWeek: time.Thursday,
						startTime: thursdayStartTime,
						endTime:   thursdayEndTime,
					},
				},
			},
		}

		whenScreened, _ := time.Parse("2006-01-02 15:04 :05", "2018-08-12 10:10:00")
		screening := Screening{
			movie:        titanic,
			sequence:     10,
			whenScreened: whenScreened,
		}

		// when
		movieFee := titanic.CalculateMovieFee(screening)

		// then
		assert.Equal(t, movieFee, Money(9000.0))
	})

}

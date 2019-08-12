package theater

import "time"

type Screening struct {
	movie        Movie
	sequence     uint8
	whenScreened time.Time
}

func (s *Screening) IsSequence(sequence uint8) bool {
	return s.sequence == sequence
}

func (s *Screening) GetMovieFee() Money {
	// Money 와 같은 측면(유연성을 높이는 다는 측면)에서 영화의 Fee 산정 방식이 달라 질수도 있지 않을까?
	// 그렇다면 movie 의 fee 필드를 직접 참조하는 것이 아니라 함수로 호출하는게 좋지 않을까?
	return s.movie.GetFee()
}

func (s *Screening) Reserve(customer Customer, audienceCount uint8) *Reservation {
	// Customer 는 참조만 할뿐 이므로 포인터로 받지 않음
	return &Reservation{
		customer:      customer,
		screening:     s,
		fee:           s.calculateFee(audienceCount),
		audienceCount: audienceCount,
	}
}

func (s Screening) calculateFee(audienceCount uint8) Money {
	return s.movie.CalculateMovieFee(s).times(float64(audienceCount))
}

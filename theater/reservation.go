package theater

type Reservation struct {
	customer      Customer
	screening     *Screening
	fee           Money
	audienceCount uint8
}

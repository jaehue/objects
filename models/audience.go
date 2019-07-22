package models

type Audience struct {
	Bag Bag
}

func (a *Audience) Buy(ticket Ticket) float64 {
	return a.Bag.Hold(&ticket)
}

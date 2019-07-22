package models

type Bag struct {
	Amount     float64
	Invitation *Invitation
	Ticket     *Ticket
}

func (b Bag) hasInvitation() bool {
	return b.Invitation != nil
}

func (b *Bag) minusAmount(amount float64) {
	b.Amount -= amount
}

func (b *Bag) plusAmount(amount float64) {
	b.Amount += amount
}

func (b *Bag) Hold(ticket *Ticket) float64 {
	if b.hasInvitation() {
		b.Ticket = ticket
		return 0
	} else {
		b.Ticket = ticket
		b.minusAmount(ticket.Fee)
		return ticket.Fee
	}
}

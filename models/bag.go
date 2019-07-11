package models

type Bag struct {
	Amount    float64
	Invitaion *Invitaion
	Ticket    *Ticket
}

func (b Bag) HasInvitaion() bool {
	return b.Invitaion != nil
}

func (b Bag) HasTicket() bool {
	return b.Ticket != nil
}

func (b *Bag) MinusAmount(amount float64) {
	b.Amount -= amount
}

func (b *Bag) PlusAmount(amount float64) {
	b.Amount += amount
}

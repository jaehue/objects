package models

type TicketOffice struct {
	Amount  float64
	Tickets []Ticket
}

func (o *TicketOffice) PopTicket() Ticket {
	ticket := o.Tickets[0]
	o.Tickets = o.Tickets[1:]
	return ticket
}

func (o *TicketOffice) MinusAmount(amount float64) {
	o.Amount -= amount
}
func (o *TicketOffice) PlusAmount(amount float64) {
	o.Amount += amount
}

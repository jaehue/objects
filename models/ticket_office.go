package models

type TicketOffice struct {
	Amount  float64
	Tickets []Ticket
}

func (o *TicketOffice) popTicket() Ticket {
	ticket := o.Tickets[0]
	o.Tickets = o.Tickets[1:]
	return ticket
}

func (o *TicketOffice) minusAmount(amount float64) {
	o.Amount -= amount
}
func (o *TicketOffice) plusAmount(amount float64) {
	o.Amount += amount
}

func (o *TicketOffice) SellTicketTo(audience *Audience) {
	o.plusAmount(audience.Buy(o.popTicket()))
}

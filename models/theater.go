package models

type Theater struct {
	TicketSeller TicketSeller
}

func (t *Theater) Enter(audience *Audience) {
	if audience.Bag.HasInvitaion() {
		ticket := t.TicketSeller.TicketOffice.PopTicket()
		audience.Bag.Ticket = &ticket
	} else {
		ticket := t.TicketSeller.TicketOffice.PopTicket()
		audience.Bag.MinusAmount(ticket.Fee)
		t.TicketSeller.TicketOffice.PlusAmount(ticket.Fee)
		audience.Bag.Ticket = &ticket
	}
}

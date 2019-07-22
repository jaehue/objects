package models

type Theater struct {
	TicketSeller TicketSeller
}

func (t *Theater) Enter(audience *Audience) {
	t.TicketSeller.SellTo(audience)
}

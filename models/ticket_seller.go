package models

type TicketSeller struct {
	TicketOffice TicketOffice
}

func (t *TicketSeller) SellTo(audience *Audience) {
	t.TicketOffice.SellTicketTo(audience)
}

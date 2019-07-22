package models

import "testing"
import "gotest.tools/assert"

func TestEnter_Invited(t *testing.T) {

	t.Run("Invited", func(t *testing.T) {
		// given
		bag := Bag{
			Amount:     1000,
			Invitation: &Invitation{},
		}

		audience := &Audience{
			Bag: bag,
		}

		tickets := []Ticket{
			{
				Fee: 10,
			},
			{
				Fee: 20,
			},
		}
		ticketOffice := TicketOffice{
			Amount:  100,
			Tickets: tickets,
		}

		ticketSeller := TicketSeller{
			TicketOffice: ticketOffice,
		}

		theater := Theater{
			TicketSeller: ticketSeller,
		}

		// when
		theater.Enter(audience)

		// then
		assert.Equal(t, audience.Bag.Ticket.Fee, 10.0)
		assert.Equal(t, audience.Bag.Amount, 1000.0)
		assert.Equal(t, len(theater.TicketSeller.TicketOffice.Tickets), 1)
	})

	t.Run("NotInvited", func(t *testing.T) {
		// given
		bag := Bag{
			Amount: 1000,
		}

		audience := &Audience{
			Bag: bag,
		}

		tickets := []Ticket{
			{
				Fee: 10,
			},
			{
				Fee: 20,
			},
		}
		ticketOffice := TicketOffice{
			Amount:  100,
			Tickets: tickets,
		}

		ticketSeller := TicketSeller{
			TicketOffice: ticketOffice,
		}

		theater := Theater{
			TicketSeller: ticketSeller,
		}

		// when
		theater.Enter(audience)

		// then
		assert.Equal(t, audience.Bag.Ticket.Fee, 10.0)
		assert.Equal(t, audience.Bag.Amount, 990.0)
		assert.Equal(t, len(theater.TicketSeller.TicketOffice.Tickets), 1)
	})
}

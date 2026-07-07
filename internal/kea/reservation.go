package kea

import "context"

func (c *Client) ListReservations(ctx context.Context) ([]Reservation, error) {

	req := ReservationGetRequest{
		Command: "reservation-get-all",
		Service: []string{"dhcp4"},
		Arguments: ReservationGetArgs{
			SubnetID: 1,
		},
	}

	var resp ReservationGetResponse

	err := c.Call(ctx, req, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Arguments.Reservations, nil
}

// Добавление
func (c *Client) AddReservation(ctx context.Context, r Reservation) error {

	req := ReservationAddRequest{
		Command: "reservation-add",
		Service: []string{"dhcp4"},
		Arguments: Reservation{
			HWAddress: r.HWAddress,
			Hostname:  r.Hostname,
		},
	}

	var resp GenericResponse

	return c.Call(ctx, req, &resp)
}

// Удаление
func (c *Client) DeleteReservation(ctx context.Context, mac string) error {

	req := ReservationDeleteRequest{
		Command: "reservation-del",
		Service: []string{"dhcp4"},
		Arguments: ReservationDeleteArgs{
			HWAddress: mac,
		},
	}

	var resp GenericResponse

	return c.Call(ctx, req, &resp)
}

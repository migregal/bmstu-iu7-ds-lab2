package reservation

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/migregal/bmstu-iu7-ds-lab2/apiserver/core/ports/reservation"
	"github.com/migregal/bmstu-iu7-ds-lab2/pkg/readiness"
	v1 "github.com/migregal/bmstu-iu7-ds-lab2/reservation/api/http/v1"
)

var probeKey = "http-reservation-client"

type Client struct {
	lg *slog.Logger

	conn *http.Client

	addr string
}

func New(lg *slog.Logger, cfg reservation.Config, probe *readiness.Probe) (*Client, error) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	conn := http.Client{
		Transport: tr,
	}

	c := Client{
		lg:   lg,
		conn: &conn,
		addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	}

	go c.ping(probe)

	return &c, nil
}

func (c *Client) ping(probe *readiness.Probe) {
	sync.OnceFunc(func() {
		probe.Mark(probeKey, false)
	})

	func() {
		for {
			resp, err := c.conn.Get(c.addr + "/readiness")
			if err != nil {
				continue
			}

			if resp.StatusCode != http.StatusOK {
				continue
			}

			sync.OnceFunc(func() {
				probe.Mark(probeKey, true)
				c.lg.Warn("[startup] reservation client ready")
			})
		}
	}()
}

func (c *Client) GetUserReservations(
	ctx context.Context, username, status string,
) ([]reservation.Reservation, error) {
	url := fmt.Sprintf("http://%s/api/v1/reservations", c.addr)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to init http request: %w", err)
	}

	req.Header.Add("X-User-Name", username)

	res, err := c.conn.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute http request: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code: %d", res.StatusCode)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read http ersponse")
	}

	var resp []v1.Reservation
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to parse http ersponse")
	}

	reservs := []reservation.Reservation{}
	for _, res := range resp {
		reservs = append(reservs, reservation.Reservation{
			ID:        res.ID,
			Username:  username,
			Status:    res.Status,
			Start:     res.Start,
			End:       res.End,
			LibraryID: res.LibraryID,
			BookID:    res.BookID,
		})
	}

	return reservs, nil
}

func (c *Client) AddUserReservation(ctx context.Context, res reservation.Reservation) (string, error) {
	return "", nil
}
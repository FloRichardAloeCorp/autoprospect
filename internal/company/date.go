package company

import "time"

type Date struct {
	time.Time
}

func (c *Date) UnmarshalJSON(b []byte) error {
	dateStr := string(b[1 : len(b)-1])

	parsedTime, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return err
	}

	c.Time = parsedTime
	return nil
}

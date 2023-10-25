package request

import (
	"fmt"
	"strings"
	"time"
)

type CustomDate struct {
	time.Time
}

func (dt *CustomDate) MarshalJSON() ([]byte, error) {
	date := dt.Format("2006-01-02")
	date = fmt.Sprintf(`"%s"`, date)
	return []byte(date), nil
}

func (dt *CustomDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	dt.Time = date
	return nil
}

func (dt *CustomDate) UnmarshalParam(param string) error {
	date, err := time.Parse("2006-01-02", param)
	if err != nil {
		return err
	}
	dt.Time = date
	return nil
}

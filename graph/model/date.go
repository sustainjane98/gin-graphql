package model

import (
	"fmt"
	"io"
	"time"
)

type Date struct {
	time.Time
}

const format = "02.01.2006"

func (d *Date) UnmarshalGQL(v interface{}) error {
	timeString, ok := v.(string)
	if !ok {
		return fmt.Errorf("value must be a string")
	}

	parse, err := time.Parse(format, timeString)
	if err != nil {
		return err
	}

	*d = Date{Time: parse}

	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (d Date) MarshalGQL(w io.Writer) {
	t := d.Time.Format(format)

	_, err := w.Write([]byte(t))
	if err != nil {
		return
	}
}

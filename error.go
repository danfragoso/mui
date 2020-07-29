package mui

import (
	"errors"
	"strconv"
)

func missingAfterError(missing, after, tokens string, line int) error {
	return errors.New("Missing " + missing + " after " + after + ", it should be " + tokens + "\n" +
		"At index: " + strconv.Itoa(line+1))
}

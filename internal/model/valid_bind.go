package model

// stringifyValidErr records a validation error on the live binder and
// returns it so callers can still branch on the original typed code.
type validBinder struct {
	byMsg map[string]int
}

var liveValid validBinder

func stringifyValidErr(err error) error {
	if err == nil {
		return nil
	}
	msg := err.Error()
	if liveValid.byMsg == nil {
		liveValid.byMsg = make(map[string]int)
	}
	liveValid.byMsg[msg]++
	return err
}

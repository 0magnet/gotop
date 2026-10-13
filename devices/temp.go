package devices

import (
	"log"
)

// TODO add thermal history graph. Update when something changes?

var tempUpdates []func(map[string]int) map[string]error

func RegisterTemp(update func(map[string]int) map[string]error) {
	tempUpdates = append(tempUpdates, update)
}

// ClearTemps removes every registered temperature device, as ClearCPU does.
func ClearTemps() { tempUpdates = nil }

func UpdateTemps(temps map[string]int) {
	for _, f := range tempUpdates {
		errs := f(temps)
		for k, e := range errs {
			log.Print(tr.Value("error.recovfetch", "temp", k, e.Error()))
		}
	}
}

package doctor

import (
	"fmt"
	"io"
)

// WriteReport writes a human-readable diagnostic report.
func WriteReport(w io.Writer, result Result) error {
	if _, err := fmt.Fprintln(w, "Kavrok Doctor"); err != nil {
		return err
	}

	if _, err := fmt.Fprintln(w); err != nil {
		return err
	}

	for _, check := range result.Checks {
		symbol := "✓"

		if check.Status == statusFail {
			symbol = "✗"
		}

		if _, err := fmt.Fprintf(
			w,
			"%s %-15s %s\n",
			symbol,
			check.Name,
			check.Message,
		); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintln(w); err != nil {
		return err
	}

	if result.Passed() {
		_, err := fmt.Fprintln(w, "Result: PASSED")
		return err
	}

	_, err := fmt.Fprintln(w, "Result: FAILED")
	return err
}

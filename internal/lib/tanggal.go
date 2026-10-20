package lib

import "time"

func ValidTanggal(tanggal string) bool {
	if len(tanggal) != len("2006-01-02") {
		return false
	}

	t, err := time.Parse("2006-01-02", tanggal)
	if err != nil {
		return false
	}

	if t.Year() < 1 {
		return false
	}

	return t.Format("2006-01-02") == tanggal
}

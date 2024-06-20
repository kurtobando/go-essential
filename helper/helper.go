package helper

import "strconv"

func StringsToFloat(s []string) ([]float64, error) {
	if len(s) == 0 {
		return []float64{}, nil
	}

	var floats []float64
	for _, v := range s {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return []float64{}, err
		}
		floats = append(floats, f)
	}

	return floats, nil
}

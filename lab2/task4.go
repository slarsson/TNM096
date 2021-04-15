package main

func (s Schedule) grade() int {
	if (7 + len(s.rooms)) >= len(s.matrix) {
		panic("array out of bounds")
	}

	errors := 0
	size := len(s.rooms)

	// check 9am
	for i := 0; i < size; i++ {
		if s.matrix[i] != "" {
			errors++
		}
	}

	// check 12am
	offset := 3 * size
	for i := offset; i < (offset + size); i++ {
		if s.matrix[i] != "" {
			errors++
		}
	}

	// check 4pm
	offset = 7 * size
	for i := offset; i < (offset + size); i++ {
		if s.matrix[i] != "" {
			errors++
		}
	}

	// check 1pm
	errors++
	offset = 1 * size
	for i := offset; i < (offset + size); i++ {
		if s.matrix[i] == "MT501" || s.matrix[i] == "MT502" {
			errors--
			break
		}
	}

	// check 2pm
	errors++
	offset = 1 * size
	for i := offset; i < (offset + size); i++ {
		if s.matrix[i] == "MT501" || s.matrix[i] == "MT502" {
			errors--
			break
		}
	}

	return errors
}

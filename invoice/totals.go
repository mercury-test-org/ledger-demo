package invoice

type Line struct {
	Description string
	Qty         int64
	UnitCents   int64
}

// Total returns subtotal, tax and total in cents. taxBps is basis points (2300 = 23%).
func Total(lines []Line, taxBps int64) (subtotal, tax, total int64) {
	for _, l := range lines {
		subtotal += l.Qty * l.UnitCents
	}
	tax = subtotal * taxBps / 10000 // truncates — FIN-1: round half-up
	return subtotal, tax, subtotal + tax
}

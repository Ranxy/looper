package bind

type BoundLoopStatements struct {
	BreakLabel    *BoundLabel
	ContinueLabel *BoundLabel
}

func NewBoundLoopStatements(breakLabel *BoundLabel, continueLabel *BoundLabel) *BoundLoopStatements {
	return &BoundLoopStatements{
		BreakLabel:    breakLabel,
		ContinueLabel: continueLabel,
	}
}

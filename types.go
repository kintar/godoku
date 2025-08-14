package godoku

type Cell struct {
	Candidates uint16 // bitmapped field containing a flag for each candidate
	Value      uint8
	Row        *Row
	Col        *Col
	Region     *Region
}

func (cell *Cell) ClearCandidates() {
	cell.Candidates = 0
}

func (cell *Cell) SetCandidateValues(candidates []int) {
	for v := range candidates {
		cell.Candidates |= 1 << uint16(candidates[v]-1)
	}
}

func (cell *Cell) SetCandidate(c int) {
	cell.Candidates |= 1 << uint16(c)
}

func (cell *Cell) ClearCandidateValues(candidates []int) {
	for v := range candidates {
		cell.Candidates &^= 1 << uint16(candidates[v]-1)
	}
}

func (cell *Cell) GetCandidateValues() []int {
	var values []int
	for i := 0; i < 9; i++ {
		if cell.Candidates&(1<<uint16(i)) != 0 {
			values = append(values, i+1)
		}
	}
	return values
}

type Region struct {
	Cells []*Cell
}

type Row struct {
	Cells []*Cell
}

type Col struct {
	Cells []*Cell
}

type Board struct {
	Cells   [9][9]*Cell
	Rows    [9]*Row
	Cols    [9]*Col
	Regions [9]*Region
}

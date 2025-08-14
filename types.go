package godoku

// Cell represents a single cell in a Sudoku puzzle. It contains a bitmapped field for candidates, a Value (with 0
// representing an empty cell), and pointers to its Row, Col, and Region.
type Cell struct {
	Candidates uint16 // bitmapped field containing a flag for each candidate
	Value      uint8
	Row        *Row
	Col        *Col
	Region     *Region
}

// ClearCandidates removes all candidates from the cell
func (cell *Cell) ClearCandidates() {
	cell.Candidates = 0
}

// SetCandidateValues adds the given candidates to the cell's candidate list
func (cell *Cell) SetCandidateValues(candidates []int) {
	for v := range candidates {
		cell.Candidates |= 1 << uint16(candidates[v]-1)
	}
}

// SetCandidate adds a single candidate to the cell's candidate list
func (cell *Cell) SetCandidate(c int) {
	cell.Candidates |= 1 << uint16(c)
}

// ClearCandidateValues removes the given candidates from the cell's candidate list
func (cell *Cell) ClearCandidateValues(candidates []int) {
	for v := range candidates {
		cell.Candidates &^= 1 << uint16(candidates[v]-1)
	}
}

// GetCandidateValues returns a slice containing the cell's candidate values
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

// IsValid checks if the region is valid (i.e., contains no duplicate values) and returns a boolean indicating validity,
// along with a slice of pointers to any invalid cells.
func (r *Region) IsValid() (bool, []*Cell) {
	panic("not implemented")
}

type Row struct {
	Cells []*Cell
}

// IsValid checks if the row is valid (i.e., contains no duplicate values) and returns a boolean indicating validity,
// along with a slice of pointers to any invalid cells.
func (r *Row) IsValid() (bool, []*Cell) {
	panic("not implemented")
}

type Col struct {
	Cells []*Cell
}

// IsValid checks if the column is valid (i.e., contains no duplicate values) and returns a boolean indicating validity,
// along with a slice of pointers to any invalid cells.
func (r *Col) IsValid() (bool, []*Cell) {
	panic("not implemented")
}

type Board struct {
	Cells   [9][9]*Cell
	Rows    [9]*Row
	Cols    [9]*Col
	Regions [9]*Region
}

// IsValid checks if the board is valid (i.e., all rows, columns, and regions are valid) and returns a boolean indicating
// validity, along with a slice of pointers to any invalid cells.
func (b *Board) IsValid() (bool, []*Cell) {
	panic("not implemented")
}

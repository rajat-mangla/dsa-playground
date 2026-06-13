package quite_hard

type SparseTable struct {
	minTable [][]int
	maxTable [][]int
	log      []int
}

func NewSparseTable(arr []int) *SparseTable {
	n := len(arr)

	// log[i] = floor(log2(i))
	log := make([]int, n+1)
	for i := 2; i <= n; i++ {
		log[i] = log[i/2] + 1
	}

	minTable := make([][]int, n)
	maxTable := make([][]int, n)
	K := log[n] + 1
	for i := range n {
		minTable[i] = make([]int, K)
		maxTable[i] = make([]int, K)

		minTable[i][0] = arr[i]
		maxTable[i][0] = arr[i]
	}

	for j := 1; j < K; j++ {
		length := 1 << j
		for i := 0; i+length <= n; i++ {
			half := 1 << (j - 1)
			minTable[i][j] = min(minTable[i][j-1], minTable[i+half][j-1])
			maxTable[i][j] = max(maxTable[i][j-1], maxTable[i+half][j-1])
		}
	}

	return &SparseTable{
		minTable: minTable,
		maxTable: maxTable,
		log:      log,
	}
}

func (t *SparseTable) Query(l, r int) int {
	j := t.log[r-l+1]

	minV := min(t.minTable[l][j], t.minTable[r-(1<<j)+1][j])
	maxV := max(t.maxTable[l][j], t.maxTable[r-(1<<j)+1][j])

	return maxV - minV
}

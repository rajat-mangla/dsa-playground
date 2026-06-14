package quite_hard

func minOperationsQueries(n int, edges [][]int, queries [][]int) []int {
	// compute log n for binary uplifting
	LOG := 1
	for (1 << LOG) <= n {
		LOG++
	}

	// build adj list with weights
	adj := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u] = append(adj[u], [2]int{v, w})
		adj[v] = append(adj[v], [2]int{u, w})
	}

	depth := make([]int, n)
	ancestorT := make([][]int, n)
	freqT := make([][][27]int, n)
	for i := range ancestorT {
		ancestorT[i] = make([]int, LOG)
		freqT[i] = make([][27]int, LOG)
	}

	// root setup
	ancestorT[0][0] = 0

	var dfs func(curr, parent int)
	dfs = func(curr, parent int) {
		for k := 1; k < LOG; k++ {
			mid := ancestorT[curr][k-1]
			ancestorT[curr][k] = ancestorT[mid][k-1]

			for w := range 27 {
				freqT[curr][k][w] = freqT[curr][k-1][w] + freqT[mid][k-1][w]
			}
		}

		for _, nb := range adj[curr] {
			v, w := nb[0], nb[1]
			if v != parent {
				depth[v] = depth[curr] + 1
				ancestorT[v][0] = curr
				freqT[v][0][w] += 1
				dfs(v, curr)
			}
		}
	}
	dfs(0, -1)

	pathStats := func(u, v int) (int, int) {
		var cnt [27]int

		if depth[u] < depth[v] {
			u, v = v, u
		}
		diff := depth[u] - depth[v]

		for k := range LOG {
			if diff&(1<<k) != 0 {
				for w := range 27 {
					cnt[w] += freqT[u][k][w]
				}
				u = ancestorT[u][k]
			}
		}

		if u != v {
			// find LCA
			for k := LOG - 1; k >= 0; k-- {
				if ancestorT[u][k] != ancestorT[v][k] {
					for w := range 27 {
						cnt[w] += freqT[u][k][w]
						cnt[w] += freqT[v][k][w]
					}
					u = ancestorT[u][k]
					v = ancestorT[v][k]
				}
			}
			// Add the last edge to LCA
			for w := range 27 {
				cnt[w] += freqT[u][0][w]
				cnt[w] += freqT[v][0][w]
			}
		}

		total, maxF := 0, 0
		for w := range 27 {
			total += cnt[w]
			maxF = max(maxF, cnt[w])
		}

		return total, maxF
	}

	ans := make([]int, len(queries))

	for i, q := range queries {
		u, v := q[0], q[1]
		if u == v {
			ans[i] = 0
			continue
		}

		total, maxF := pathStats(u, v)
		ans[i] = total - maxF
	}

	return ans
}

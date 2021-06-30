package graph

import "math"

const inf = math.MaxInt/2 - 1

// MinCostMaxFlow finds the minimum cost max flow of a weighted flow network
// using the Ford–Fulkerson algorithm with Bellman-Ford search.
// SEE https://en.wikipedia.org/wiki/Ford%E2%80%93Fulkerson_algorithm
//     https://en.wikipedia.org/wiki/Bellman%E2%80%93Ford_algorithm
type MinCostMaxFlow struct {
	found           []bool
	n               int
	cap, flow, cost [][]int
	prev, dist, pi  []int
}

func (f *MinCostMaxFlow) search(s, t int) bool {
	for i := range f.found {
		f.found[i] = false
	}
	for i := range f.dist {
		f.dist[i] = inf
	}

	f.dist[s] = 0

	for s != f.n {
		best := f.n
		f.found[s] = true

		for i := 0; i < f.n; i++ {
			if f.found[i] {
				continue
			}

			if f.flow[i][s] != 0 {
				val := f.dist[s] + f.pi[s] - f.pi[i] - f.cost[i][s]
				if f.dist[i] > val {
					f.dist[i] = val
					f.prev[i] = s
				}
			}

			if f.flow[s][i] < f.cap[s][i] {
				val := f.dist[s] + f.pi[s] - f.pi[i] + f.cost[s][i]
				if f.dist[i] > val {
					f.dist[i] = val
					f.prev[i] = s
				}
			}

			if f.dist[i] < f.dist[best] {
				best = i
			}
		}

		s = best
	}

	for k := 0; k < f.n; k++ {
		pi := f.pi[k] + f.dist[k]
		if pi > inf {
			pi = inf
		}
		f.pi[k] = pi
	}

	return f.found[t]
}

func (f *MinCostMaxFlow) Flow() [][]int {
	return f.flow
}

func (f *MinCostMaxFlow) ComputeMaxFlow(g Graph, s, t int) (flow, cost int) {
	f.cap = g.cap
	f.cost = g.cost

	f.n = len(f.cap)
	f.found = make([]bool, f.n)
	f.flow = make([][]int, f.n)
	for i := 0; i < f.n; i++ {
		f.flow[i] = make([]int, f.n)
	}
	f.dist = make([]int, f.n+1)
	f.prev = make([]int, f.n)
	f.pi = make([]int, f.n)

	for f.search(s, t) {
		pathFlow := inf
		for u := t; u != s; u = f.prev[u] {
			var pf int
			if f.flow[u][f.prev[u]] != 0 {
				pf = f.flow[u][f.prev[u]]
			} else {
				pf = f.cap[f.prev[u]][u] - f.flow[f.prev[u]][u]
			}
			if pf < pathFlow {
				pathFlow = pf
			}
		}

		for u := t; u != s; u = f.prev[u] {
			if f.flow[u][f.prev[u]] != 0 {
				f.flow[u][f.prev[u]] -= pathFlow
				cost -= pathFlow * f.cost[u][f.prev[u]]
			} else {
				f.flow[f.prev[u]][u] += pathFlow
				cost += pathFlow * f.cost[f.prev[u]][u]
			}
		}
		flow += pathFlow
	}

	return flow, cost
}

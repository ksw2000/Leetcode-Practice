// https://leetcode.com/problems/satisfiability-of-equality-equations/
// 0 ms	2.8 MB

package main

type uf []int

func (u uf) root(n int) int {
	for u[n] != n {
		u[n] = u[u[n]]
		n = u[n]
	}
	return n
}

func (u uf) union(me int, it int) {
	if p, q := u.root(me), u.root(it); p != q {
		u[p] = q
	}
}

func createUF(n int) *uf {
	u := uf(make([]int, n))
	for i := range u {
		u[i] = i
	}
	return &u
}

func equationsPossible(equations []string) bool {
	u := createUF(26)
	j := len(equations)
	for i := 0; i < j; {
		if equations[i][1] == '=' {
			a := int(equations[i][0] - 'a')
			b := int(equations[i][3] - 'a')
			u.union(a, b)
			i++
		} else {
			j--
			equations[i], equations[j] = equations[j], equations[i]
		}
	}

	for j < len(equations) {
		a := int(equations[j][0] - 'a')
		b := int(equations[j][3] - 'a')
		if u.root(a) == u.root(b) {
			return false
		}
		j++
	}

	return true
}

package leetcode

import "log"

type Node399 struct {
	name        string
	connections []*Node399
	values      []float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {

	res := []float64{}

	// build graph
	nodeMap := make(map[string]*Node399)
	for i, equation := range equations {
		var n1, n2 *Node399
		v1, v2 := equation[0], equation[1]

		var ok bool
		if n1, ok = nodeMap[v1]; !ok {
			newNode := &Node399{
				name:        v1,
				connections: []*Node399{},
				values:      []float64{},
			}
			n1 = newNode
			nodeMap[v1] = n1
		}
		if n2, ok = nodeMap[v2]; !ok {
			newNode := &Node399{
				name:        v2,
				connections: []*Node399{},
				values:      []float64{},
			}
			n2 = newNode
			nodeMap[v2] = n2
		}

		n1.connections = append(n1.connections, n2)
		n1.values = append(n1.values, values[i])
		n2.connections = append(n2.connections, n1)
		n2.values = append(n2.values, 1/values[i])
	}

	for _, node := range nodeMap {
		log.Printf("node = %+v", node)
	}

	// calculate
	for _, query := range queries {
		var n1 *Node399
		v1, v2 := query[0], query[1]

		var ok bool
		if n1, ok = nodeMap[v1]; !ok {
			res = append(res, -1)
			continue
		}
		if _, ok = nodeMap[v2]; !ok {
			res = append(res, -1)
			continue
		}

		if v1 == v2 {
			res = append(res, 1)
			continue
		}

		visited := make(map[string][]byte)
		res = append(res, dfs399(n1, v2, visited))
	}

	return res
}

func dfs399(node *Node399, target string, visited map[string][]byte) float64 {
	if _, ok := visited[node.name]; ok {
		return -1
	}
	visited[node.name] = []byte{}

	for i, c := range node.connections {
		if c.name == target {
			return node.values[i]
		}

		found := dfs399(c, target, visited)
		if found > -1 {
			return node.values[i] * found
		}
	}

	return -1
}

func TestcalcEquation() {

	type Case struct {
		equations [][]string
		values    []float64
		queries   [][]string

		expected []float64
	}

	cases := []Case{
		{
			equations: [][]string{{"a", "b"}, {"b", "c"}},
			values:    []float64{2.0, 3.0},
			queries:   [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			expected:  []float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000},
		},
		{
			equations: [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
			values:    []float64{1.5, 2.5, 5.0},
			queries:   [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
			expected:  []float64{3.75000, 0.40000, 5.00000, 0.20000},
		},
		{
			equations: [][]string{{"a", "b"}},
			values:    []float64{0.5},
			queries:   [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}},
			expected:  []float64{0.50000, 2.00000, -1.00000, -1.00000},
		},
	}

	for _, testcase := range cases {
		res := calcEquation(testcase.equations, testcase.values, testcase.queries)

		if len(res) != len(testcase.expected) {
			log.Panicf("FAILED. %+v, %+v, Expected %+v %+v but got %+v", testcase.equations, testcase.values, testcase.queries, testcase.expected, res)
		}
		for i := 0; i < len(testcase.expected); i++ {
			if testcase.expected[i] != res[i] {
				log.Panicf("FAILED. %+v, %+v, Expected %+v %+v but got %+v", testcase.equations, testcase.values, testcase.queries, testcase.expected, res)
			}
		}

		log.Printf("MATCHED. %+v, %+v, result: %+v, expected: %+v\n\n", testcase.equations, testcase.values, res, testcase.expected)
	}
}

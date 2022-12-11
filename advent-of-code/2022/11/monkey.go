package main

type worryOp func(int) int

type monkey struct {
	items            []int
	totalInspections int
	op               worryOp
	test             int
	testPassTarget   int
	testFailTarget   int
}

func newMonkey() *monkey {
	return &monkey{
		test:           -1,
		testPassTarget: -1,
		testFailTarget: -1,
	}
}

func (m *monkey) doTest(worry int) bool {
	return worry%m.test == 0
}

func (m *monkey) giveItem(item int) {
	m.items = append(m.items, item)
}

func (m *monkey) takeItem(id int) {
	m.items = append(m.items[:id], m.items[id+1:]...)
}

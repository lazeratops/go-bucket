package main

type tree struct {
	height    int
	hidden    bool
	neighbors neighbors
}

type pos struct {
	x int
	y int
}

type neighbors struct {
	west  *tree
	east  *tree
	north *tree
	south *tree
}

func (t *tree) updateVisibility() {
	wHidden := false
	eHidden := false
	nHidden := false
	sHidden := false

	nW := t.neighbors.west
	if nW != nil {
		if nW.hidden == true || nW.height >= t.height {
			wHidden = true
		}
	}

	nE := t.neighbors.east
	if nE != nil {
		if nE.hidden == true || nE.height >= t.height {
			eHidden = true
		}
	}

	nN := t.neighbors.north
	if nN != nil {
		if nN.hidden == true || nN.height >= t.height {
			nHidden = true
		}
	}

	nS := t.neighbors.south
	if nS != nil {
		if nS.hidden == true || nS.height >= t.height {
			sHidden = true
		}
	}

	t.hidden = wHidden && eHidden && nHidden && sHidden
}

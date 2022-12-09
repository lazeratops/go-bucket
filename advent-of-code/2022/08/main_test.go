package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"path/filepath"
	"testing"
)

var testData = trees{
	{
		3, 0, 3, 7, 3,
	},
	{
		2, 5, 5, 1, 2,
	},
	{
		6, 5, 3, 3, 2,
	},
	{
		3, 3, 5, 4, 9,
	},
	{
		3, 5, 3, 9, 0,
	},
}

func TestBothParts(t *testing.T) {
	t.Parallel()
	path := filepath.Join("testdata", "input.txt")
	ans1, ans2 := do(path)
	require.Equal(t, 1843, ans1)
	require.Equal(t, 180000, ans2)
}

func TestTotal(t *testing.T) {
	t.Parallel()
	path := filepath.Join("testdata", "input_test.txt")
	ans1, ans2 := do(path)
	require.Equal(t, 21, ans1)
	require.Equal(t, 8, ans2)
}

func TestObscuredX(t *testing.T) {
	treeData := trees{
		{
			5, 4, 4, 2, 1,
		},
	}
	testCases := []struct {
		trees         trees
		col           int
		wantObscuredX bool
	}{
		{
			trees:         treeData,
			col:           0,
			wantObscuredX: false,
		},
		{
			trees:         treeData,
			col:           1,
			wantObscuredX: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%d", tc.col), func(t *testing.T) {
			p := patch{}
			p.trees = tc.trees
			height := tc.trees[0][tc.col]
			gotObscuredX, _ := p.obscuredX(height, tc.col, 0)
			require.Equal(t, tc.wantObscuredX, gotObscuredX)
		})
	}
}

func TestObscuredY(t *testing.T) {
	treeData := trees{
		{
			5, 4, 3, 2, 1,
		},
		{
			2, 3, 0, 5, 5,
		},
		{
			1, 1, 4, 1, 1,
		},
	}
	testCases := []struct {
		trees         trees
		x             int
		y             int
		wantObscuredY bool
	}{
		{
			trees:         treeData,
			x:             0,
			y:             0,
			wantObscuredY: false,
		},
		{
			trees:         treeData,
			x:             1,
			y:             1,
			wantObscuredY: false,
		},
		{
			trees:         treeData,
			x:             2,
			y:             1,
			wantObscuredY: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%d-%d", tc.x, tc.y), func(t *testing.T) {
			p := patch{}
			p.trees = tc.trees
			height := tc.trees[tc.y][tc.x]

			gotObscuredY, _ := p.obscuredY(height, tc.x, tc.y)
			require.Equal(t, tc.wantObscuredY, gotObscuredY)
		})
	}
}

func TestObscuredXY(t *testing.T) {

	testCases := []struct {
		x             int
		y             int
		wantObscuredX bool
		wantObscuredY bool
	}{
		{
			x: 1,
			y: 1,
		},
		{
			x: 2,
			y: 1,
		},
		{
			x:             3,
			y:             1,
			wantObscuredX: true,
			wantObscuredY: true,
		},
		{
			x:             1,
			y:             2,
			wantObscuredY: true,
		},
		{
			x:             2,
			y:             2,
			wantObscuredY: true,
			wantObscuredX: true,
		},
		{
			x:             3,
			y:             2,
			wantObscuredY: true,
		},
		{
			x:             1,
			y:             3,
			wantObscuredY: true,
			wantObscuredX: true,
		},
		{
			x: 2,
			y: 3,
		},
		{
			x:             3,
			y:             3,
			wantObscuredY: true,
			wantObscuredX: true,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%d-%d", tc.x, tc.y), func(t *testing.T) {
			p := patch{}
			p.trees = testData
			height := testData[tc.y][tc.x]
			gotObscuredX, _ := p.obscuredX(height, tc.x, tc.y)
			require.Equal(t, tc.wantObscuredX, gotObscuredX)

			gotObscuredY, _ := p.obscuredY(height, tc.x, tc.y)
			require.Equal(t, tc.wantObscuredY, gotObscuredY)
		})
	}
}

func TestTotalVisible(t *testing.T) {
	p := patch{}
	p.trees = testData
	gotTotalVisible, gotBestScenicScore := p.countVisibleTrees()
	require.Equal(t, 21, gotTotalVisible)
	require.Equal(t, 8, gotBestScenicScore)

}

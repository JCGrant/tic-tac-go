package tic

import (
	"math"
	"testing"
)

func TestMinimax(t *testing.T) {
	root := node{
		children: []node{
			{
				children: []node{
					{
						children: []node{
							{
								children: []node{
									{value: 10},
									{value: math.Inf(1)},
								},
							},
							{
								children: []node{
									{value: 5},
								},
							},
						},
					},
					{
						children: []node{
							{
								children: []node{
									{value: -10},
								},
							},
						},
					},
				},
			},
			{
				children: []node{
					{
						children: []node{
							{
								children: []node{
									{value: 7},
									{value: 5},
								},
							},
							{
								children: []node{
									{value: math.Inf(-1)},
								},
							},
						},
					},
					{
						children: []node{
							{
								children: []node{
									{value: -7},
									{value: -5},
								},
							},
						},
					},
				},
			},
		},
	}
	assertEqual(t, -7.0, minimax(root, 4, true))
}

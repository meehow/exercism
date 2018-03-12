package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })
	if records[0].ID != 0 {
		return nil, errors.New("no root node")
	}
	if records[0].Parent != 0 {
		return nil, errors.New("root node has parent")
	}
	root := new(Node)
	nodes := map[int]*Node{0: root}
	parent := root
	prevID := 0
	for _, record := range records[1:] {
		if record.ID != prevID+1 {
			return nil, errors.New("non-continuous")
		}
		if record.ID == record.Parent {
			return nil, errors.New("cycle directly")
		}
		node := &Node{ID: record.ID}
		nodes[node.ID] = node
		if record.Parent != parent.ID {
			parent = nodes[record.Parent]
			if parent == nil {
				return nil, errors.New("cycle indirectly")
			}
		}
		parent.Children = append(parent.Children, node)
		prevID = record.ID
	}
	return root, nil
}

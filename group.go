package sc

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/scgolang/osc"
	"io"
	"reflect"
)

const (
	gQueryTree      = "/g_queryTree"
	gQueryTreeReply = "/g_queryTree.reply"
)

// Group
type Group interface {
	Free() error
}

type node struct {
	id int32 `json:"id" xml:"id,attr"`
}

type group struct {
	node     `json:"node" xml:"node"`
	children []*node `json:"children" xml:"children>child"`
	client   *Client
}

func (self *group) Free() error {
	return nil
}

func (self *group) WriteJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(self)
}

func (self *group) WriteXML(w io.Writer) error {
	enc := xml.NewEncoder(w)
	return enc.Encode(self)
}

func newGroup(client *Client, id int32) *group {
	return &group{
		node:     node{id: id},
		children: make([]*node, 0),
		client:   client,
	}
}

// parseGroup parses information about a group from a message
// received at /g_queryTree
// it *does not* recursively query for child groups
func parseGroup(msg *osc.Message) (Group, error) {
	// return an error if msg.Address is not right
	if msg.Address != gQueryTreeReply {
		return nil, fmt.Errorf("msg.Address should be %s, got %s", gQueryTreeReply, msg.Address)
	}
	// g_queryTree replies should have at least 3 arguments
	g, numArgs := new(group), msg.CountArguments()
	if numArgs < 3 {
		return nil, fmt.Errorf("expected 3 arguments for message, got %d", numArgs)
	}
	// get the id of the group this reply is for
	var isint bool
	g.node.id, isint = msg.Arguments[1].(int32)
	if !isint {
		v := msg.Arguments[1]
		t := reflect.TypeOf(v)
		return nil, fmt.Errorf("expected arg 1 to be int32, got %s (%v)", t, v)
	}
	// initialize the children array
	var numChildren int32
	numChildren, isint = msg.Arguments[2].(int32)
	if !isint {
		v := msg.Arguments[1]
		t := reflect.TypeOf(v)
		return nil, fmt.Errorf("expected arg 2 to be int32, got %s (%v)", t, v)
	}
	if numChildren < 0 {
		return nil, fmt.Errorf("expected numChildren >= 0, got %d", numChildren)
	}
	g.children = make([]*node, numChildren)
	// get the childrens' ids
	var nodeID, numControls, numSubChildren int32
	for i := 3; i < numArgs; {
		nodeID, isint = msg.Arguments[i].(int32)
		if !isint {
			v := msg.Arguments[i]
			t := reflect.TypeOf(v)
			return nil, fmt.Errorf("expected arg %d (nodeID) to be int32, got %s (%v)", i, t, v)
		}
		g.children[i-3] = &node{nodeID}
		// get the number of children of this node
		// if -1 this is a synth, if >= 0 this is a group
		numSubChildren, isint = msg.Arguments[i+1].(int32)
		if !isint {
			v := msg.Arguments[i]
			t := reflect.TypeOf(v)
			return nil, fmt.Errorf("expected arg %d (numControls) to be int32, got %s (%v)", i, t, v)
		}
		if numSubChildren == -1 {
			// synth
			i += 3
			numControls, isint = msg.Arguments[i].(int32)
			if !isint {
				v := msg.Arguments[i]
				t := reflect.TypeOf(v)
				return nil, fmt.Errorf("expected arg %d (numControls) to be int32, got %s (%v)", i, t, v)
			}
			i += 1 + int(numControls*2)
		} else if numSubChildren >= 0 {
			// group
			i += 2
		}
	}
	return g, nil
}

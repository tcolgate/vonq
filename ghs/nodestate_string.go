// Code generated by "stringer -type=NodeState"; DO NOT EDIT

package ghs

import "fmt"

const _NodeState_name = "NodeStateSleepingNodeStateFindNodeStateFound"

var _NodeState_index = [...]uint8{0, 17, 30, 44}

func (i NodeState) String() string {
	if i < 0 || i >= NodeState(len(_NodeState_index)-1) {
		return fmt.Sprintf("NodeState(%d)", i)
	}
	return _NodeState_name[_NodeState_index[i]:_NodeState_index[i+1]]
}
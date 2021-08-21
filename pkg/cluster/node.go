package cluster

import (
	model "github.com/cloudreve/Cloudreve/v3/models"
	"github.com/cloudreve/Cloudreve/v3/pkg/aria2/common"
	"github.com/cloudreve/Cloudreve/v3/pkg/serializer"
)

type Node interface {
	// Init a node from database model
	Init(node *model.Node)
	// Check if given feature is enabled
	IsFeatureEnabled(feature string) bool
	// Subscribe node status change to a callback function
	SubscribeStatusChange(callback func(isActive bool, id uint))
	// Ping the node
	Ping(req *serializer.NodePingReq) (*serializer.NodePingResp, error)
	// Returns if the node is active
	IsActive() bool
	// Get instances for aria2 calls
	GetAria2Instance() common.Aria2
	// Returns unique id of this node
	ID() uint
}

func getNodeFromDBModel(node *model.Node) Node {
	switch node.Type {
	case model.SlaveNodeType:
		slave := &SlaveNode{}
		slave.Init(node)
		return slave
	default:
		master := &MasterNode{}
		master.Init(node)
		return master
	}
}
package gaiatypes

import (
	"net"
	"sync"
)

// GraphEdgeMap is a map of id to GraphEdge
type GraphEdgeMap map[string]*GraphEdge

// GraphEdgeExtremityType represents the source or destination type.
type GraphEdgeExtremityType string

const (
	// GraphEdgeExtremityTypePU represents the value of a pu extremity.
	GraphEdgeExtremityTypePU GraphEdgeExtremityType = "pu"

	// GraphEdgeExtremityTypeExternalService represents the value of a pu external service.
	GraphEdgeExtremityTypeExternalService GraphEdgeExtremityType = "extsrv"
)

// GraphEdge represents the model of a Edge
type GraphEdge struct {
	AcceptedFlows   int                    `json:"acceptedFlows"`
	DestinationID   string                 `json:"destinationID"`
	DestinationType GraphEdgeExtremityType `json:"destinationType"`
	ID              string                 `json:"ID"`
	Name            string                 `json:"name"`
	RejectedFlows   int                    `json:"rejectedFlows"`
	SourceID        string                 `json:"sourceID"`
	SourceType      GraphEdgeExtremityType `json:"sourceType"`
}

// NewGraphEdge returns a new *GraphEdge
func NewGraphEdge() *GraphEdge {
	return &GraphEdge{}
}

// GraphNodeType represents the possible values for attribute "type".
type GraphNodeType string

const (
	// GraphNodeTypeContainer represents the value Container.
	GraphNodeTypeContainer GraphNodeType = "Container"

	// GraphNodeTypeExternalService represents the value EternalService.
	GraphNodeTypeExternalService GraphNodeType = "ExternalService"

	// GraphNodeTypeVolume represents the value Volume.
	GraphNodeTypeVolume GraphNodeType = "Volume"
)

// GraphNodeMap is a map of id to GraphNode
type GraphNodeMap map[string]*GraphNode

// GraphNode represents the model of a Node
type GraphNode struct {
	Description        string        `json:"description"`
	GroupID            string        `json:"groupID"`
	ID                 string        `json:"ID"`
	Name               string        `json:"name"`
	Namespace          string        `json:"namespace"`
	Status             string        `json:"status"`
	Tags               []string      `json:"tags,omitempty"`
	Type               GraphNodeType `json:"type"`
	VulnerabilityLevel string        `json:"vulnerabilityLevel"`
}

// NewGraphNode returns a new *GraphNode
func NewGraphNode() *GraphNode {

	return &GraphNode{
		Tags: []string{},
	}
}

// GraphGroupMap is a map of id to GraphGroup
type GraphGroupMap map[string]*GraphGroup

// GraphGroup represents the model of a Group
type GraphGroup struct {
	Color    string     `json:"color"`
	ID       string     `json:"ID"`
	Match    [][]string `json:"match"`
	Name     string     `json:"name"`
	ParentID string     `json:"parentID"`
}

// NewGraphGroup returns a new *GraphNode
func NewGraphGroup() *GraphGroup {

	return &GraphGroup{
		Match: [][]string{},
	}
}

// IPRecord represent an IP record.
type IPRecord struct {
	IP               string   `json:"IP"`
	Hostnames        []string `json:"hostnames"`
	DestinationPorts []string `json:"destinationPorts"`
	Hits             int      `json:"hits"`
}

// NewIPRecord returns a new IPRecord.
func NewIPRecord() *IPRecord {
	return &IPRecord{}
}

// ResolveHostnames resolves hostnames for the current IP.
func (r *IPRecord) ResolveHostnames(wg *sync.WaitGroup, l *sync.Mutex) {
	l.Lock()
	r.Hostnames, _ = net.LookupAddr(r.IP)
	l.Unlock()

	if wg != nil {
		wg.Done()
	}
}

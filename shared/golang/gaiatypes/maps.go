package gaiatypes

// GraphEdge represents the model of a Edge
type GraphEdge struct {
	AcceptedFlows int    `json:"acceptedFlows"`
	DestinationID string `json:"destinationID"`
	ID            string `json:"ID"`
	Name          string `json:"name"`
	RejectedFlows int    `json:"rejectedFlows"`
	SourceID      string `json:"sourceID"`
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

	// GraphNodeTypeVolume represents the value Volume.
	GraphNodeTypeVolume GraphNodeType = "Volume"
)

// GraphNode represents the model of a Node
type GraphNode struct {
	Description        string        `json:"description"`
	GroupID            string        `json:"groupID"`
	ID                 string        `json:"ID"`
	Name               string        `json:"name"`
	Namespace          string        `json:"namespace"`
	Status             string        `json:"status"`
	Tags               []string      `json:"tags"`
	Type               GraphNodeType `json:"type"`
	VulnerabilityLevel string        `json:"vulnerabilityLevel"`
}

// NewGraphNode returns a new *GraphNode
func NewGraphNode() *GraphNode {

	return &GraphNode{
		Tags: []string{},
	}
}

// GraphGroup represents the model of a Group
type GraphGroup struct {
	Color    string   `json:"color"`
	ID       string   `json:"ID"`
	Match    []string `json:"match"`
	Name     string   `json:"name"`
	ParentID string   `json:"parentID"`
}

// NewGraphGroup returns a new *GraphNode
func NewGraphGroup() *GraphGroup {

	return &GraphGroup{
		Match: []string{},
	}
}

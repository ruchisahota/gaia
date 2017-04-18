package gaiatypes

// GraphEdgeMap is a map of id to GraphEdge
type GraphEdgeMap map[string]*GraphEdge

// GraphEdge represents the model of a Edge
type GraphEdge struct {
	AcceptedFlows int    `json:"acceptedFlows,omitempty"`
	DestinationID string `json:"destinationID,omitempty"`
	ID            string `json:"ID,omitempty"`
	Name          string `json:"name,omitempty"`
	RejectedFlows int    `json:"rejectedFlows,omitempty"`
	SourceID      string `json:"sourceID,omitempty"`
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

// GraphNodeMap is a map of id to GraphNode
type GraphNodeMap map[string]*GraphNode

// GraphNode represents the model of a Node
type GraphNode struct {
	Description        string        `json:"description,omitempty"`
	GroupID            string        `json:"groupID,omitempty"`
	ID                 string        `json:"ID,omitempty"`
	Name               string        `json:"name,omitempty"`
	Namespace          string        `json:"namespace,omitempty"`
	Status             string        `json:"status,omitempty"`
	Tags               []string      `json:"tags,omitempty"`
	Type               GraphNodeType `json:"type,omitempty"`
	VulnerabilityLevel string        `json:"vulnerabilityLevel,omitempty"`
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
	Color    string   `json:"color,omitempty"`
	ID       string   `json:"ID,omitempty"`
	Match    []string `json:"match,omitempty"`
	Name     string   `json:"name,omitempty"`
	ParentID string   `json:"parentID,omitempty"`
}

// NewGraphGroup returns a new *GraphNode
func NewGraphGroup() *GraphGroup {

	return &GraphGroup{
		Match: []string{},
	}
}

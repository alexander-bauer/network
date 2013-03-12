package network

// A Network contains a map of Nodes by name or identifier. It is
// meant to represent an entire (not subdivided) network of individual
// nodes, each with attributes.
type Network struct {
	Nodes map[string]*Node `json:"nodes"`
}

// A Node is a single node in a network with a set of connections and
// attributes.
type Node struct {
	Connected []*Connection `json:"connected"`
	// All of the connected nodes
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// Any attributes

	Disabled bool `json:"disabled,omitempty"`
	// Whether the node is currently inactive
	ShouldBe bool `json:"shouldbe,omitempty"`
	// Whether the node should be inactive
}

type Connection struct {
	Target string `json:"target"`
	// Name of connected node
	Quality uint64 `json:"quality,omitempty"`
	// A unitless "opinion" of the quality of the link
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// Any attributes

	Disabled bool `json:"disabled,omitempty"`
	// Whether the connection is currently inactive
	ShouldBe bool `json:"shouldbe,omitempty"`
	// Whether the connection should be inactive
}

package network

// A Network is a map of Nodes by name or identifier. It is meant to
// represent an entire (not subdivided) network of individual nodes,
// each with attributes.
type Network struct {
	Nodes map[string]*Node
}

// A Node is a single node in a network with a set of connections and
// attributes.
type Node struct {
	Connected []string `json:"connected"`
	// All of the connected nodes
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// Any attributes
}

type Connection struct {
	Target  string // Name of connected node
	Quality uint64 // A unitless "opinion" of the quality of the link
}

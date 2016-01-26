package badgraph

import (
	"fmt"
)

type Identity string

func (id Identity) String() string {
	return fmt.Sprintf("IDENTITY-%s",string(id))
}

type Tag string

type Tager interface {
	Get(Tag) (string,bool)
	Set(Tag,string)
}

/* Node a node holds a collection of 
 * information, relationships between
 * nodes are constructed/maintained via
 * arcs. */
type Node struct {
	Identity      					 /* Identity of the node to the system */
	Alias string						 /* Alias name of the node, for humans */
	
	Tags map[Tag]string      /* Collection of tags associated with this node */
}

func (n *Node) Get(k Tag) (string,bool) {
	v,exists := n.Tags[k]
	return v,exists
}

func (n *Node) String() string {
	l := 0
	if n.Tags != nil {
		l = len(n.Tags)
	}
	return fmt.Sprintf("NODE \"%s\" %s, %d TAGS",n.Alias,n.Identity,l)
}

func (n *Node) Set(k Tag,v string) {
	if n.Tags == nil {
		n.Tags = make(map[Tag]string,0)
	}
	n.Tags[k] = v
}

func NewNode(id Identity,alias string) *Node {

	n := Node{}
	n.Identity = id
	n.Alias = alias
	return &n
}


/* ArcType defines the type of relationship
 * an Arc repesents. */
type ArcType byte

func (a ArcType) String() string {
	str := "<-x->"
	switch a {
		case UnidirectionalFromAlice:
			str = "---->"
		break
		case UnidirectionalFromBella:
			str = "<----"
		break
		case Bidirectional:
			str = "<--->"
		break
	}
	return str
}

const (
	Muted           				ArcType = 0     /* Relationship is defined but not active */
	UnidirectionalFromAlice ArcType = 1  		/* Relationship is one-way from Alice to Bella */
	UnidirectionalFromBella ArcType = 2     /* Relationship is one-way from Bella to Alice */
	Bidirectional  					ArcType = 3			/* Relationship is two-way between Alice & Bella */
)


/* Arc an arc defines a relationship between
 * nodes or groups of nodes. Alice refers to 
 * node A, and Bella to node B in this 
 * relationship. The relationship type is 
 * defined by the ArcType.
 */
type Arc struct {
	Alice Identity						/* Identity of Node A in the system */
	Bella Identity						/* Identity of Node B in the system */
	Type  ArcType							/* Type of relationship between A&B */

	Tags map[Tag]string				/* Collection of tags associated with this node */
}

func (a *Arc) Get(k Tag) (string,bool) {
	v,exists := a.Tags[k]
	return v,exists 
}

func (a *Arc) String() string {
	l := 0
	if a.Tags != nil {
		l = len(a.Tags)
	}
	
	return fmt.Sprintf("ARC %s %s %s, %d TAGS",a.Alice,a.Type,a.Bella,l)
}

func (a *Arc) Set(k Tag,v string) {
	if a.Tags == nil {
		a.Tags = make(map[Tag]string,0)
	}
	a.Tags[k] = v
}

func NewArc(Alice,Bella Identity,typeof ArcType) *Arc {
	
	arc := Arc{}
	arc.Alice = Alice
	arc.Bella = Bella
	arc.Type = typeof
	return &arc
}



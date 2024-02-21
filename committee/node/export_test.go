package node

import "github.com/synapsecns/sanguine/committee/db"

func (n *Node) DB() db.Service {
	return n.db
}

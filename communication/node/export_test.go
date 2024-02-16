package node

import "github.com/synapsecns/sanguine/communication/db"

func (n *Node) DB() db.Service {
	return n.db
}

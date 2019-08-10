package cart

import (
	"container/list"
)

func newAdminPanel (c *Cart) (*AdminPanel) {
	return &AdminPanel {c}
}

// This data type is meant to support the Cart data type. Some priviledged methods that
// may not be invoked directly on a Cart, can be invoked directly on this data type.
type AdminPanel struct {
	underlyingCart *Cart
}

// Harvest () is a wrapper around method harvest () of the Cart abstract data type. If
// the underlying cart haa already been harvested, value would be ErrBeenHarvested.
func (p *AdminPanel) Harvest () (*list.List, error) {
	return p.underlyingCart.harvest ()
}

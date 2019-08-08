package cart

func newAdminPanel (c *Cart) (*AdminPanel) {
	return &AdminPanel {c}
}

// This data type is meant to support the Cart data type. Some priviledged methods that can not be invoked directly on a Cart, can be invoked directly on this type.
type AdminPanel struct {
	underlyingCart *Cart
}

// This method is a wrapper around method harvest () of the Cart data type.
func (p *AdminPanel) Harvest () (int32) {
	return p.underlyingCart.harvest ()
}
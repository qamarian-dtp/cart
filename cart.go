package cart

import (
	"container/list"
	"errors"
	"runtime"
	"sync/atomic"
)

// New () creates a new cart, then handles over to you, the cart's pointer and its admin
// panel. Do not create carts manually, you can always use this function.
func New () (*Cart, *AdminPanel) {
	c := &Cart {StateDormant, list.New ()}
	return c, newAdminPanel (c)
}

// Do not create manually, use function New () to get a new cart.
type Cart struct {
	state int32
	items *list.List
}

// Put () can be used to put new items into the cart. If the operation fails, an error
// would be returned. ErrBeenHarvested is among the errors that could be returned.
func (c *Cart) Put (i interface {}) (error) {
	addBeginning:
	ok := atomic.CompareAndSwapInt32 (&c.state, StateDormant, StateInUse)
	if ok == false {
		switch c.state {
			case StateDormant:
				runtime.Gosched ()
				goto addBeginning
			case StateInUse:
				runtime.Gosched ()
				goto addBeginning
			case StateHarvested:
				return ErrBeenHarvested
			default:
				return errors.New ("This cart is in an invalid state.")
		}
	}
	c.items.PushBack (i)
	c.state = StateDormant
	return nil
}

// GetState () provides the current state of cart. Possible values can be found in the
// variable section of this package.
func (c *Cart) GetState () (int32) {
	return c.state
}

// harvest () harvests a cart. All items that have been put in the cart so far would be
// returned
//
// Outpt
//
// outpt 0: On successful harvest, value would be the list of all items that have been
// put in the cart so far. On failed harvest, value could be anything
//
// outpt 1: On successful harvest, value would be nil. On failed harvest, value would
// be an error. If the cart has already been harvested, the error returned would
// specifically be StateHarvested.
func (c *Cart) harvest () (*list.List, error) {
	harvestBeginning:
	okX := atomic.CompareAndSwapInt32 (&c.state, StateDormant, StateHarvested)
	if okX == false {
		switch c.state {
			case StateDormant:
				runtime.Gosched ()
				goto harvestBeginning
			case StateInUse:
				runtime.Gosched ()
				goto harvestBeginning
			case StateHarvested:
				return nil, ErrBeenHarvested
			default:
				return nil, errors.New ("This cart is in an invalid " +
					"state.")
		}
	}
	return c.items, nil
}

// Count () helps count the number of items that have been put in the cart so far.
func (c *Cart) Count () (int) {
	return c.items.Len ()
}

var (
	StateDormant   int32 = 0 // No operation is being performed on the cart
	StateInUse     int32 = 1 // An item is being added to the cart
	StateHarvested int32 = 2 // The cart has been harvested

	ErrBeenHarvested error = errors.New ("This cart has been harvested.")
)

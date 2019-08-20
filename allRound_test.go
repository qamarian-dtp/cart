package cart

import (
	"gopkg.in/qamarian-lib/str.v2"
	"strconv"
	"testing"
	"time"
)

func TestAllRound (t *testing.T) {
	str.PrintEtr ("Testing of data type cart and admin panel...", "std", "tester")

	myCart, adminPanel := New ()

	go func (someCart *Cart) {
		count := 0
		for { // Items are added till the cart is harvested
			errX := someCart.Put ("some item")
			if errX != nil {
				str.PrintEtr ("Unable to add new item: " + errX.Error (), "wrn", "tester")
				break
			}
			count ++
		}
		str.PrintEtr ("No of items added to the cart: " + strconv.Itoa (count), "std", "tester")
	} (myCart)

	time.Sleep (time.Second * 5)
	items, errY := adminPanel.Harvest ()
	if errY != nil {
		str.PrintEtr  ("Test seems to have failed. Reason: " + errY.Error (), "err", "tester")
		t.FailNow ()
	}

	str.PrintEtr ("No of items added in the cart: " + strconv.Itoa (myCart.Count ()), "std", "tester")
	str.PrintEtr ("Items in the cart:", "std", "tester")
	str.PrintEtr (items, "std", "tester")

	str.PrintEtr ("Test passed!", "std", "tester")
}

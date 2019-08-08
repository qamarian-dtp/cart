This package implements the cart data type, a data type that could be used to collect items from around a program. For example: if you want to collect some items from some goroutines, you can create a cart and pass its pointer to the goroutines. The goroutuines can keep adding items to the cart, and when you are no longer interested in collecting new items, you can harvest the cart. Once you harvest the cart, no new item can be added to the cart.

	myCart, adminPanel ;= cart.New ()
	go func (someCart *cart.Cart) {
		for { // Items are added till the cart is harvested
			err := someCart.Put ("some item")
			if err != nil 	
				break
			}
		}
	} (myCart)
	time.Sleep (time.Second * 5)
	items, err := adminPanel.Harvest ()
	
To prevent unauthorized routines from harvesting a cart, method harvest () can not be called directly on a cart. This is the reason why harvesting of a cart can only be done via the admin panel of the cart.

	adminPanel.Harvest () // Right
	myCart.Harvest () // Wrong
	myCart.harvest () // wrong

Warning!

When importing this package, import a specific version of the package because the master branch of its Github repository (https://github.com/qamarian-dtp/cart) is not guaranteed to be always backward compatible with older versions. Don't know how to import a specific version, you can read up the http://gopkg.in tool.


package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func RoutesMap(api *ApiConnection) Routes {
	var routes = Routes{
		Route{"Index", "GET", "/", api.Index},
		Route{"SignIn", "POST", "/signin", api.SignIn},
		Route{"Login", "POST", "/login", api.Login},
		Route{"NewProduct", "POST", "/newproduct", api.NewProduct},
		Route{"UpdateProduct", "POST", "/updateproduct", api.UpdateProduct},
		Route{"DeleteProduct", "POST", "/deleteproduct", api.DeleteProduct},
		Route{"ListProducts", "POST", "/listproducts", api.ListProducts},
		Route{"NewCustomer", "POST", "/newcustomer", api.NewCustomer},
		Route{"UpdateCustomer", "POST", "/updatecustomer", api.UpdateCustomer},
		Route{"DeleteCustomer", "POST", "/deletecustomer", api.DeleteCustomer},
		Route{"ListCustomers", "POST", "/listcustomers", api.ListCustomers},
		Route{"NewInvoice", "POST", "/newinvoice", api.NewInvoice},
		Route{"UpdateInvoice", "POST", "/updateinvoice", api.UpdateInvoice},
		Route{"DeleteInvoice", "POST", "/deleteinvoice", api.DeleteInvoice},
		Route{"ListInvoices", "POST", "/listinvoices", api.ListInvoices},
	}

	return routes
}

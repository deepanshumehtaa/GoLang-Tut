# Book Store Management REST API

1. db							:	MySql
2. GORM							:	The GoLang ORM packege
3. Json marshall & UnMarshal	:	Built-in support for JSON encoding and decoding
4. Gorila Mux					:	Package mux implements a request router and dispatcher & mux => "HTTP request multiplexer"

###	The Features are:
* Requests can be matched based on URL host, path, path prefix, schemes,
  header and query values, HTTP methods or using custom matchers.

* URL hosts, paths and query values can have variables with an optional
  regular expression.

* Registered URLs can be built, or "reversed", which helps maintaining
  references to resources.

* Routes can be used as subrouters: nested routes are only tested if the
  parent route matches. This is useful to define groups of routes that
  share common conditions like a host, a path prefix or other repeated
  attributes. As a bonus, this optimizes request matching.

* It implements the http.Handler interface so it is compatible with the
  standard http.ServeMux. 									


## 	Project Structure

- 	CMD:
		main.go
-	PKG:
		config		: App.go (db helper and connnection)
		controller	: Book controler (helper function process the data http ie. GET, POST, UPDATE, )
		models		: Book.go (strut models)
		routers		: Bookstore-Router (the urls)
		utils		: utils.go (marshalling and unmarshalling the request ir, serializer and de-serializer)

-- go.mod	: store all the packages ~ like requirements.txt but as url of the packeges


## Installing:
	download orm package:
		go get -u github.com/jinzhu/gorm
		go get -u github.com/jinzhu/gorm/dialects/mysql
		go get -u github.com/jinzhu/gorm/dialects/mysql
		go get -u github.com/gorilla/mux
		








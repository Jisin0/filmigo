// (c) Jisi0

/*
Package encode facilitates encoding data based on struct tags to simplify development.
It currently supports querying documents based on xpath from struct tags and building url params from structs.
If a field is unset the default behaviour of all functions is to skip it, contrary to what other packages like json do by using the struct field name.
*/
package encode

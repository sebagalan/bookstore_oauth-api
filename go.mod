module bookstore/oauth-api

go 1.13

replace github.com/sebagalan/bookstore_oauth-api => ../bookstore_oauth-api

replace github.com/sebagalan/bookstore_users-api => ../bookstore_users-api

require (
	github.com/sebagalan/bookstore_oauth-api v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.0
)

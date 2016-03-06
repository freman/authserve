# authserve
`authserve` is an experimentation in creating a generic authentication service in go.

I'm trying to build something generic and super basic that stores commonly needed attributes in a verifiable JWT
so that any pages that need this basic information can just check the JWT saving a DB hit.

## Features

 * Can use with postgresql, mysql, sqlite
 * Built in swagger documentation
 * JSON api
 * Enforces token validity window

## Building

	cd $GOPATH
	go get github.com/freman/authserve

## Verifying tokens in other modules

Included in this package is a `token` library that you can use in conjunction with a public key to verify the validity of a token

	import "github.com/freman/authserve/token"
	...
	claims, err := token.Verify(tokenString, publicKey)

## License

Copyright (c) 2016 Shannon Wynter. Licensed under GPL3. See the [LICENSE.md](LICENSE.md) file for a copy of the license.
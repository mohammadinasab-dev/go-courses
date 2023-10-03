package main

/*
	- Opaque errors in Go are error types that hide internal details and only provide limited information to callers.
	  They help to encapsulate implementation-specific details and provide a more abstract and
	  `decoupled` error handling approach.

	- Don’t just check errors, handle them gracefully.

	- Only handle errors once.
	  The idea is to handle an error at the lowest possible level where it can be effectively dealt with.

	- Assert errors for behavior, not type


*/

/*
	func AuthenticateRequest(r *Request) error {
        err := authenticate(r.User)
        if err != nil {
                return err
        }
        return nil
	}
*/

/*
	type temporary interface {
        Temporary() bool
	}

	// IsTemporary returns true if err is temporary.
	func IsTemporary(err error) bool {
	        te, ok := err.(temporary)
	        return ok && te.Temporary()
	}
*/

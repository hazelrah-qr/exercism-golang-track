package erratum

// Use attempts to invoke a ResourceOpener to open a resource and
// frob the input string returning an error or nil if successful
func Use(o ResourceOpener, input string) (err error) {
	resource, err := o()

	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		resource, err = o()
	}

	defer resource.Close()
	defer func() {
		if r := recover(); r != nil {
			if frobErr, ok := r.(FrobError); ok {
				resource.Defrob(frobErr.defrobTag)
				err = frobErr
			} else {
				err, _ = r.(error)
			}
		}
	}()

	resource.Frob(input)

	return nil
}

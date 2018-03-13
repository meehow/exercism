package erratum

// Use ...
func Use(opener ResourceOpener, input string) (err error) {
	var resource Resource
	defer func() {
		if r := recover(); r == nil {
			return
		} else if ferr, ok := r.(FrobError); ok {
			err = ferr.inner
			resource.Defrob(ferr.defrobTag)
			resource.Close()
		} else if err, ok = r.(error); ok {
			resource.Close()
		}
	}()
	for {
		resource, err = opener()
		if err == nil {
			break
		}
		if _, ok := err.(TransientError); !ok {
			return err
		}
	}
	resource.Frob(input)
	resource.Close()
	return err
}

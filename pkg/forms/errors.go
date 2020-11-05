package forms

//errors

type errors map[string][]string

//add errors for a field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

//get error
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}

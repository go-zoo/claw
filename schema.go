/**********************************
***  Middleware Chaining in Go  ***
***  Code is under MIT license  ***
***    Code by CodingFerret     ***
*** 	github.com/squiidz      ***
***********************************/

package claw

// Schema is use to create some predefine middleware stack.
type Schema []MiddleWare

// NewScheme generate a New Schema with the provided Middleware.
func NewSchema(m ...interface{}) *Schema {
	if len(m) > 0 {
		sch := &Schema{}
		stack := toMiddleware(m)
		for _, s := range stack {
			*sch = append(*sch, s)
		}
		return sch
	}
	return nil
}

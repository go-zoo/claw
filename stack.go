/**********************************
***  Middleware Chaining in Go  ***
***  Code is under MIT license  ***
***    Code by CodingFerret     ***
*** 	github.com/squiidz      ***
***********************************/

package claw

// Stack is use to create some predefine middleware stack.
type Stack []MiddleWare

// NewStack generate a New Stack with the provided Middleware.
func NewStack(m ...interface{}) Stack {
	if len(m) > 0 {
		sch := Stack{}
		stk := toMiddleware(m)
		for _, s := range stk {
			sch = append(sch, s)
		}
		return sch
	}
	return nil
}

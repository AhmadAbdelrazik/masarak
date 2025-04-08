// types are the DTOs returned by the application layer queries.
package app

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

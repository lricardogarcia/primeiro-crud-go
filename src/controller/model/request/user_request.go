package request

// UserRequest represents the input data for creating a new user.
// @Summary User Input Data
// @Description Structure containing the required fields for creating a new user.
type UserRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Name string `json:"name"`
	Age int8 `json:"age"`
}
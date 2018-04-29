package domain

type UnauthorizedError struct{}

func (UnauthorizedError) Error() string { return "User is not authorized" }

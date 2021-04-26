package login

type TemplateResponse struct {
	Title       string
	Description string
	Price       int
	Error       string
}

type ListUserResponse struct {
	UserField User
}

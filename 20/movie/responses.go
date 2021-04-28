package movie

type MainPageResponse struct {
	Movies []Movie
}

type AddPageResponse struct {
	Error string
}

type DetailPageResponse struct {
	Movie Movie
}

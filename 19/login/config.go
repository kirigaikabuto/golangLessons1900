package login

type Config struct {
}

type TemplateResponse struct {
	Title       string
	Description string
	Price       int
	Error       string
}

type MongoConfig struct {
	Host           string
	Port           string
	Database       string
	CollectionName string
}
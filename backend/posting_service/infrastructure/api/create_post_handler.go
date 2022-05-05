package api

import "github.com/Nebojsa1999/XMLProjekat/backend/posting_service/application"

type CreatePostCommandHandler struct {
	postService *application.PostService
}

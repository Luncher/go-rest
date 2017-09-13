package forms

type CreateMovieCommand struct {
	Name   string  `json:"name" binding:"required"`
	Desc   string  `json:"desc" binding:"required"`
	Rating float32 `json:"rating" binding:"required"`
}

type UpdateMovieCommand struct {
	Name   string  `json:"name" binding:"required"`
	Desc   string  `json:"desc" binding:"required"`
	Rating float32 `json:"rating" binding:"required"`
}

package dto

type CreateMovieRequest struct {
	Title       string   `form:"title" binding:"required,min=1,max=39"`
	Description string   `form:"description" binding:"required,min=10,max=999"`
	Genres      []string `form:"genres" binding:"required,dive,required"`
	Actors      []string `form:"actors" binding:"required,dive,required"`
	TrailerUrl  string   `form:"trailerUrl" binding:"required,url"`
	// Poster will be handled as a file upload
}

type CreateMovieResponse struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Genres      []string `json:"genres"`
	Actors      []string `json:"actors"`
	TrailerUrl  string   `json:"trailerUrl"`
	Poster      string   `json:"poster"`
}

type UpdateMovieRequest struct {
	Title       string   `json:"title" binding:"required,min=1,max=39"`
	Description string   `json:"description" binding:"required,min=10,max=999"`
	Genres      []string `json:"genres" binding:"required,dive,required"`
	Actors      []string `json:"actors" binding:"required,dive,required"`
	TrailerUrl  string   `json:"trailerUrl" binding:"required,url"`
	Poster      string   `json:"poster" binding:"required,url"`
}

type UpdateMovieResponse struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Genres      []string `json:"genres"`
	Actors      []string `json:"actors"`
	TrailerUrl  string   `json:"trailerUrl"`
	Poster      string   `json:"poster"`
}

type GetMoviesRequest struct {
	Page     int    `form:"page,default=1" binding:"min=1"`
	PageSize int    `form:"page_size,default=10" binding:"min=1,max=100"`
	Title    string `form:"title"`
}

type GetMoviesResponse struct {
	Movies     []MovieResponse `json:"movies"`
	PageNumber int             `json:"pageNumber"`
	PageSize   int             `json:"pageSize"`
	TotalSize  int64           `json:"totalSize"`
}

type MovieResponse struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Genres      []string `json:"genres"`
	Actors      []string `json:"actors"`
	TrailerUrl  string   `json:"trailerUrl"`
	Poster      string   `json:"poster"`
}

type MovieDetailsResponse struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Genres      []string `json:"genres"`
	Actors      []string `json:"actors"`
	TrailerUrl  string   `json:"trailerUrl"`
	Poster      string   `json:"poster"`
	UserID      string   `json:"userId"` // Include user ID in details
}

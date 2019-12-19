package input

type Paginate struct {
	Page   	int		`json:"page" binding:"required"`
	Limit	int	 	`json:"limit" binding:"required"`
}
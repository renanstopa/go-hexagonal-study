package request

type NewsRequest struct {
	Subject string `form:"subject" binding:"required,min=2"`
	From    string `form:"from" binding:"required,datetime" time_format:"2001-01-01"`
}

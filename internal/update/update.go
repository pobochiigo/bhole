package update

type Update struct {
	Comment *string
	CreatedBy *string
	CreatedOn string
	Id int32
	InfoUrl *string
	ProfileImage *string
}

type ListUpdatesRequest struct {
	Limit  int32
	Offset int32
	Search string
	Mode   string
}

type ListUpdatesResponse struct {
	Count    int32
	Next     string
	Previous string
	Results  []Update
}

type GetUpdateRequest struct {
	ID   int32
	Mode string
}

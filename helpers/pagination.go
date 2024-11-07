package helpers

type PaginationData struct {
	NextPage     int
	PreviousPage int
	CurrentPage  int
}

func GetPaginationData() PaginationData {
	return PaginationData{}
}

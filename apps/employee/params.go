package employee

type createNewEmployeeRequest struct {
	Name    string
	NIP     string
	Address string
}

type renderWeb struct {
	Title string
	Data  interface{}
}

type singleEmployeeResponse struct{}

type listEmployeeResponse struct {
	Id        int
	Name      string
	NIP       string
	Address   string
	CreatedAt string
}

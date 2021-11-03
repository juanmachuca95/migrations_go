package pages

type Step struct {
	Title    string
	Done     bool
	Resource string
}

type Page struct {
	Steps []Step
	Title string
}

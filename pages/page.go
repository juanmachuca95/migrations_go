package pages

type Step struct {
	Title    string
	Done     bool
	Resource string
	Key      string
}

type Page struct {
	Steps []Step
	Title string
}

func NewPage() {

}

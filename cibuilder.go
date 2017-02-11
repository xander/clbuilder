package clbuilder

var (
	root string
	dist string
)

type Builder struct {
	//	reference to concatinated css file
	css string
	//	reference to concatinated js file
	js string
	//	hold html fragments
	views map[string]string
}

func (b *Builder) Execute(w ResponseWriter, data interface{}, template string, views, libs, models []string) {

}

func (b *Builder) PrintCSS() string {

}

func (b *Builder) PrintJS() string {

}

//	root - the directory root with the various static assets
//	dist - where to put the various built assets
func Init(root, dist string) {

}

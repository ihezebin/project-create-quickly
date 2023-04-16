package build

type ReactTsBuilder struct {
	Origin string
}

const defaultReactOrigin = "https://gitee.com/ihezebin/template-react-ts.git"

func NewReactTsBuilder(origin string) *ReactTsBuilder {
	if origin == "" {
		origin = defaultReactOrigin
	}

	return &ReactTsBuilder{
		Origin: origin,
	}
}

func (b *ReactTsBuilder) Build(projectName string) error {
	return cloneGitProject(b.Origin, projectName)
}

func (b *ReactTsBuilder) Rename(newProjectName string) error {
	return nil
}

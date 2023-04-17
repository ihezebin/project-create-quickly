package build

import "github.com/ihezebin/project-create-quickly/component/rename"

type ReactTsBuilder struct {
	Origin string
}

const defaultReactOrigin = "https://gitee.com/ihezebin/react-template-ts.git"

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
	const defaultOldName = "react-template-ts"
	return rename.Rename(newProjectName, defaultOldName, newProjectName)
}

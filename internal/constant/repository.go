package constant

import "fmt"

const (
	defaultGoDDDRepository = "http://gitee.com/ihezebin/go-template-ddd.git"
	defaultCraTsRepository = "http://gitee.com/ihezebin/react-template-ts.git"
)

var template2RepositoryTable = map[string]string{
	TemplateGoDDD: defaultGoDDDRepository,
	TemplateCraTs: defaultCraTsRepository,
}

func GetDefaultRepository(template string) (string, error) {
	r, ok := template2RepositoryTable[template]
	if !ok {
		return "", fmt.Errorf("not support template: %s", template)
	}
	return r, nil
}

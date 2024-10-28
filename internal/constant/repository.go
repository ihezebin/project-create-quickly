package constant

import "fmt"

const (
	defaultGoDDDRepository = "http://gitee.com/ihezebin/go-template-ddd.git"
	defaultCraTsRepository = "http://gitee.com/ihezebin/react-template-ts.git"
	defaultViteRepository  = "https://gitee.com/hezebin-react/react-template-vite.git"
)

var template2RepositoryTable = map[string]string{
	TemplateGoDDD: defaultGoDDDRepository,
	TemplateCraTs: defaultCraTsRepository,
	TemplateVite:  defaultViteRepository,
}

func GetDefaultRepository(template string) (string, error) {
	r, ok := template2RepositoryTable[template]
	if !ok {
		return "", fmt.Errorf("not support template: %s", template)
	}
	return r, nil
}

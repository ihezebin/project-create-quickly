package constant

import "fmt"

const (
	defaultGoDDDRepository    = "https://github.com/ihezebin/go-template-ddd.git"
	defaultCraTsRepository    = "https://github.com/ihezebin/react-template-ts.git"
	defaultViteRepository     = "https://github.com/ihezebin/react-template-vite.git"
	defaultJavaDDDRepository  = "https://github.com/ihezebin/java-template-ddd.git"
	defaultTaroRepository     = "https://github.com/ihezebin/wechat-template-taro.git"
	defaultElectronRepository = "https://github.com/ihezebin/electron-template-vite.git"
)

var template2RepositoryTable = map[string]string{
	TemplateGoDDD:    defaultGoDDDRepository,
	TemplateCraTs:    defaultCraTsRepository,
	TemplateVite:     defaultViteRepository,
	TemplateJavaDDD:  defaultJavaDDDRepository,
	TemplateTaro:     defaultTaroRepository,
	TemplateElectron: defaultElectronRepository,
}

func GetDefaultRepository(template string) (string, error) {
	r, ok := template2RepositoryTable[template]
	if !ok {
		return "", fmt.Errorf("not support template: %s", template)
	}
	return r, nil
}

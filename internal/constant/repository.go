package constant

import "fmt"

const (
	defaultGoDDDRepository    = "https://github.com/ihezebin/go-template-ddd.git"
	defaultCraTsRepository    = "http://gitee.com/hezebin-react/react-template-ts.git"
	defaultViteRepository     = "https://gitee.com/hezebin-react/react-template-vite.git"
	defaultJavaDDDRepository  = "http://gitee.com/hezebin-go/java-template-ddd.git"
	defaultTaroRepository     = "https://github.com/ihezebin/wechat-template-taro.git"
	defaultElectronRepository = "https://gitee.com/hezebin-electron/electron-template-vite.git"
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

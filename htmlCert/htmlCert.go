package htmlcert

import (
	"fmt"
	"os"
	"path"
	"text/template"

	"training.go/pdf_certif_project/cert"
)

type HtmlSaver struct {
	OutputDir string
}

func New(outputDir string) (*HtmlSaver, error) {
	err := os.MkdirAll(outputDir, os.ModePerm)

	if err != nil {
		return nil, err
	}

	h := &HtmlSaver{
		OutputDir: outputDir,
	}
	return h, nil
}

func (h *HtmlSaver) Save (c *cert.Cert) error {
      t, err := template.New("certif").Parse(tpl)

	  if err != nil {
		return err
	  }

	  filename := fmt.Sprintf("%v.html", c.LabelTitle)
	  path := path.Join(h.OutputDir, filename)

	  f, err := os.Create(path)
	  

	  if err != nil {
		return err
	  }
	  defer f.Close()

	  err = t.Execute(f, *c)
	

	if err != nil {
		return err
	}


	  return nil
}

const tpl = `
<html>
  <body>
     <h1>{{.labelCompletion}}<h1> 
  </body>
</html>
`
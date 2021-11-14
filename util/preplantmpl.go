package util

import (
	"log"
	"text/template"
)

const prePlanT = `
<html>
<head></head>
<body>
		<h1>Personal Information</h1>
		<p>
		{{.lastName}}, {{.firstName}} {{.middleName}}<br>
		{{.street}}<br>
		{{.city}} {{.state}} {{.zip}}<br>
		{{.county}}
		</p>
</body>
</html>
`

func LoadPrePlanT() (*template.Template, error) {
	tmpl, err := template.New("prePlan").Parse(prePlanT)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return tmpl, nil
}

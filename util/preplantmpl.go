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
		<p>
		Maiden Name: {{.maidenName}}<br>
		SSN: {{.ssn}}<br>
		Birth Date: {{.birthDate}}<br>
		Birth Place: {{.birthPlace}}<br>
		email: {{.email}}
		phone: {{.phone}}
		</p>
		<h1>Education / Employment Information</h1>
		<p>
		Employer: {{.employ}}<br>
		Since: {{.employDate}}<br>
		Job Title: {{.jobTitle}}<br>
		Education: {{.education}}<br>
		Service Branch: {{.serviceBranch}}
		</p>
		<h1>Family Information</h1>
		<p>
		Spouse Name: {{.spouseName}}<br>
		Father's Name: {{.fatherName}}<br>
		Mother's Name: {{.motherName}}<br>
		</p>
		<h1>Service Information</h1>
		<p>
		Officient Name: {{.officiantName}}<br>
		Officient Phone: {{.officiantPhone}}<br>
		Service Place: {{.servicePlace}}<br>
		Cemetary: {{.cemetaryName}}<br>
		Disposition: {{.dispositionPref}}<br>
		Visitation: {{.visitation}}<br>
		</p>
	</body>
</html>
`

func LoadPrePlanT(name string) (*template.Template, error) {
	tmpl, err := template.New(name).Parse(prePlanT)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return tmpl, nil
}

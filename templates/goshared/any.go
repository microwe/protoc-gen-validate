package goshared

const anyTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

	if a := {{ accessor . }}; a != nil {
		{{ if $r.In }}
			if _, ok := {{ lookup $f "InLookup" }}[a.GetTypeUrl()]; !ok {
				err := {{ err . "必须在下列范围内：" $r.In }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ else if $r.NotIn }}
			if _, ok := {{ lookup $f "NotInLookup" }}[a.GetTypeUrl()]; ok {
				err := {{ err . "必须在下列范围外：" $r.NotIn }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	}
`

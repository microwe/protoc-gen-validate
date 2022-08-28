package goshared

const wrapperTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	if wrapper := {{ accessor . }}; wrapper != nil {
		{{ render (unwrap . "wrapper") }}
	} {{ if .MessageRules.GetRequired }} else {
		err := {{ err . "必填且不能为nil" }}
		if !all { return err }
		errors = append(errors, err)
	} {{ end }}
`

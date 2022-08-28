package golang

const requiredTpl = `
	{{ if .Rules.GetRequired }}
		if {{ accessor . }} == nil {
			err := {{ err . "为必填字段" }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}
`

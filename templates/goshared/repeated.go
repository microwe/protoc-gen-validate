package goshared

const repTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	{{ if $r.GetIgnoreEmpty }}
		if len({{ accessor . }}) > 0 {
	{{ end }}

	{{ if $r.GetMinItems }}
		{{ if eq $r.GetMinItems $r.GetMaxItems }}
			if len({{ accessor . }}) != {{ $r.GetMinItems }} {
				err := {{ err . "必须包含" $r.GetMinItems "项" }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ else if $r.MaxItems }}
			if l := len({{ accessor . }}); l < {{ $r.GetMinItems }} || l > {{ $r.GetMaxItems }} {
				err := {{ err . "必须介于" $r.GetMinItems "至" $r.GetMaxItems "项之间" }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ else }}
			if len({{ accessor . }}) < {{ $r.GetMinItems }} {
				err := {{ err . "至少包含" $r.GetMinItems "项" }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.MaxItems }}
		if len({{ accessor . }}) > {{ $r.GetMaxItems }} {
			err := {{ err . "最多包含" $r.GetMaxItems "项" }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.GetUnique }}
		{{ lookup $f "Unique" }} := {{ if isBytes $f.Type.Element -}}
			make(map[string]struct{}, len({{ accessor . }}))
		{{ else -}}
			make(map[{{ (typ $f).Element }}]struct{}, len({{ accessor . }}))
		{{ end -}}
	{{ end }}

	{{ if or $r.GetUnique (ne (.Elem "" "").Typ "none") }}
		for idx, item := range {{ accessor . }} {
			_, _ = idx, item
			{{ if $r.GetUnique }}
				if _, exists := {{ lookup $f "Unique" }}[{{ if isBytes $f.Type.Element }}string(item){{ else }}item{{ end }}]; exists {
					err := {{ errIdx . "idx" "repeated value must contain unique items" }}
					if !all { return err }
					errors = append(errors, err)
				} else {
					{{ lookup $f "Unique" }}[{{ if isBytes $f.Type.Element }}string(item){{ else }}item{{ end }}] = struct{}{}
				}
			{{ end }}

			{{ render (.Elem "item" "idx") }}
		}
	{{ end }}

	{{ if $r.GetIgnoreEmpty }}
		}
	{{ end }}
`

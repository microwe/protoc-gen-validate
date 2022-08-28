package goshared

const ltgtTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ if $r.Lt }}
		{{ if $r.Gt }}
			{{  if gt $r.GetLt $r.GetGt }}
				if val := {{ accessor . }};  val <= {{ $r.Gt }} || val >= {{ $r.Lt }} {
					err := {{ err . "必须在以下范围内：(" $r.GetGt ", " $r.GetLt ")" }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else }}
				if val := {{ accessor . }}; val >= {{ $r.Lt }} && val <= {{ $r.Gt }} {
					err := {{ err . "必须在以下范围外：[" $r.GetLt ", " $r.GetGt "]" }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
		{{ else if $r.Gte }}
			{{  if gt $r.GetLt $r.GetGte }}
				if val := {{ accessor . }};  val < {{ $r.Gte }} || val >= {{ $r.Lt }} {
					err := {{ err . "必须在以下范围内：[" $r.GetGte ", " $r.GetLt ")" }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else }}
				if val := {{ accessor . }}; val >= {{ $r.Lt }} && val < {{ $r.Gte }} {
					err := {{ err . "必须在以下范围外：[" $r.GetLt ", " $r.GetGte ")" }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
		{{ else }}
			if {{ accessor . }} >= {{ $r.Lt }} {
				err := {{ err . "必须小于" $r.GetLt }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.Lte }}
		{{ if $r.Gt }}
			{{  if gt $r.GetLte $r.GetGt }}
				if val := {{ accessor . }};  val <= {{ $r.Gt }} || val > {{ $r.Lte }} {
					err := {{ err . "必须在以下范围内：(" $r.GetGt ", " $r.GetLte "]" }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else }}
				if val := {{ accessor . }}; val > {{ $r.Lte }} && val <= {{ $r.Gt }} {
					err := {{ err . "必须在以下范围外：(" $r.GetLte ", " $r.GetGt "]" }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
		{{ else if $r.Gte }}
			{{ if gt $r.GetLte $r.GetGte }}
				if val := {{ accessor . }};  val < {{ $r.Gte }} || val > {{ $r.Lte }} {
					err := {{ err . "必须在以下范围内：[" $r.GetGte ", " $r.GetLte "]" }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else }}
				if val := {{ accessor . }}; val > {{ $r.Lte }} && val < {{ $r.Gte }} {
					err := {{ err . "必须在以下范围外：(" $r.GetLte ", " $r.GetGte ")" }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
		{{ else }}
			if {{ accessor . }} > {{ $r.Lte }} {
				err := {{ err . "必须小于等于" $r.GetLte }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.Gt }}
		if {{ accessor . }} <= {{ $r.Gt }} {
			err := {{ err . "必须大于" $r.GetGt }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.Gte }}
		if {{ accessor . }} < {{ $r.Gte }} {
			err := {{ err . "必须大于等于" $r.GetGte }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}
`

package goshared

const durationcmpTpl = `{{ $f := .Field }}{{ $r := .Rules }}
			{{  if $r.Const }}
				if dur != {{ durLit $r.Const }} {
					err := {{ err . "必须为" (durStr $r.Const) }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}


			{{  if $r.Lt }}  lt  := {{ durLit $r.Lt }};  {{ end }}
			{{- if $r.Lte }} lte := {{ durLit $r.Lte }}; {{ end }}
			{{- if $r.Gt }}  gt  := {{ durLit $r.Gt }};  {{ end }}
			{{- if $r.Gte }} gte := {{ durLit $r.Gte }}; {{ end }}

			{{ if $r.Lt }}
				{{ if $r.Gt }}
					{{  if durGt $r.GetLt $r.GetGt }}
						if dur <= gt || dur >= lt {
							err := {{ err . "必须在以下范围内：(" (durStr $r.GetGt) ", " (durStr $r.GetLt) ")" }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ else }}
						if dur >= lt && dur <= gt {
							err := {{ err . "必须在以下范围外：[" (durStr $r.GetLt) ", " (durStr $r.GetGt) "]" }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ end }}
				{{ else if $r.Gte }}
					{{  if durGt $r.GetLt $r.GetGte }}
						if dur < gte || dur >= lt {
							err := {{ err . "必须在以下范围内：[" (durStr $r.GetGte) ", " (durStr $r.GetLt) ")" }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ else }}
						if dur >= lt && dur < gte {
							err := {{ err . "必须在以下范围外：[" (durStr $r.GetLt) ", " (durStr $r.GetGte) ")" }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ end }}
				{{ else }}
					if dur >= lt {
						err := {{ err . "必须小于" (durStr $r.GetLt) }}
						if !all { return err }
						errors = append(errors, err)
					}
				{{ end }}
			{{ else if $r.Lte }}
				{{ if $r.Gt }}
					{{  if durGt $r.GetLte $r.GetGt }}
						if dur <= gt || dur > lte {
							err := {{ err . "必须在以下范围内：(" (durStr $r.GetGt) ", " (durStr $r.GetLte) "]" }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ else }}
						if dur > lte && dur <= gt {
							err := {{ err . "必须在以下范围外：(" (durStr $r.GetLte) ", " (durStr $r.GetGt) "]" }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ end }}
				{{ else if $r.Gte }}
					{{ if durGt $r.GetLte $r.GetGte }}
						if dur < gte || dur > lte {
							err := {{ err . "必须在以下范围内：[" (durStr $r.GetGte) ", " (durStr $r.GetLte) "]" }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ else }}
						if dur > lte && dur < gte {
							err := {{ err . "必须在以下范围外：(" (durStr $r.GetLte) ", " (durStr $r.GetGte) ")" }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ end }}
				{{ else }}
					if dur > lte {
						err := {{ err . "必须小于等于" (durStr $r.GetLte) }}
						if !all { return err }
						errors = append(errors, err)
					}
				{{ end }}
			{{ else if $r.Gt }}
				if dur <= gt {
					err := {{ err . "必须大于" (durStr $r.GetGt) }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else if $r.Gte }}
				if dur < gte {
					err := {{ err . "必须大于等于" (durStr $r.GetGte) }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}


			{{ if $r.In }}
				if _, ok := {{ lookup $f "InLookup" }}[dur]; !ok {
					err := {{ err . "必须在下列范围内：" $r.In }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else if $r.NotIn }}
				if _, ok := {{ lookup $f "NotInLookup" }}[dur]; ok {
					err := {{ err . "必须在下列范围外：" $r.NotIn }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
`

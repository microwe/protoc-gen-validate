package goshared

const bytesTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	{{ if $r.GetIgnoreEmpty }}
		if len({{ accessor . }}) > 0 {
	{{ end }}

	{{ if or $r.Len (and $r.MinLen $r.MaxLen (eq $r.GetMinLen $r.GetMaxLen)) }}
		{{ if $r.Len }}
			if len({{ accessor . }}) != {{ $r.GetLen }} {
				err := {{ err . "长度必须是" $r.GetLen "字节" }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ else }}
			if len({{ accessor . }}) != {{ $r.GetMinLen }} {
				err := {{ err . "长度必须是" $r.GetMinLen "字节" }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.MinLen }}
		{{ if $r.MaxLen }}
			if l := len({{ accessor . }}); l < {{ $r.GetMinLen }} || l > {{ $r.GetMaxLen }} {
				err := {{ err . "长度必须介于" $r.GetMinLen "至" $r.GetMaxLen "字节之间" }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ else }}
			if len({{ accessor . }}) < {{ $r.GetMinLen }} {
				err := {{ err . "长度至少是" $r.GetMinLen "字节" }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.MaxLen }}
		if len({{ accessor . }}) > {{ $r.GetMaxLen }} {
			err := {{ err . "长度最多是" $r.GetMaxLen "字节" }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Prefix }}
		if !bytes.HasPrefix({{ accessor . }}, {{ lit $r.GetPrefix }}) {
			err := {{ err . "必须包含前缀" (byteStr $r.GetPrefix) }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Suffix }}
		if !bytes.HasSuffix({{ accessor . }}, {{ lit $r.GetSuffix }}) {
			err := {{ err . "必须包含后缀" (byteStr $r.GetSuffix) }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Contains }}
		if !bytes.Contains({{ accessor . }}, {{ lit $r.GetContains }}) {
			err := {{ err . "必须包含" (byteStr $r.GetContains) }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.In }}
		if _, ok := {{ lookup $f "InLookup" }}[string({{ accessor . }})]; !ok {
			err := {{ err . "必须在下列范围内：" $r.In }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.NotIn }}
		if _, ok := {{ lookup $f "NotInLookup" }}[string({{ accessor . }})]; ok {
			err := {{ err . "必须在下列范围外：" $r.NotIn }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Const }}
		if !bytes.Equal({{ accessor . }}, {{ lit $r.Const }}) {
			err := {{ err . "必须为" $r.Const }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.GetIp }}
		if ip := net.IP({{ accessor . }}); ip.To16() == nil {
			err := {{ err . "必须是一个有效的IP地址" }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetIpv4 }}
		if ip := net.IP({{ accessor . }}); ip.To4() == nil {
			err := {{ err . "必须是一个有效的IPv4地址" }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetIpv6 }}
		if ip := net.IP({{ accessor . }}); ip.To16() == nil || ip.To4() != nil {
			err := {{ err . "必须是一个有效的IPv6地址" }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Pattern }}
	if !{{ lookup $f "Pattern" }}.Match({{ accessor . }}) {
		err := {{ err . "必须符合正则表达式" (lit $r.GetPattern) }}
		if !all { return err }
		errors = append(errors, err)
	}
	{{ end }}

	{{ if $r.GetIgnoreEmpty }}
		}
	{{ end }}
`

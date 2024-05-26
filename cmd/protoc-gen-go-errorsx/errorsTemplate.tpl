{{ range .Errors }}

{{ if .HasComment }}{{ .Comment }}{{ end -}}
func Is{{.CamelValue}}(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == {{ .Name }}_{{ .Value }}.String() && e.Code == {{ .HTTPCode }}
}

{{ if .HasComment }}{{ .Comment }}{{ end -}}
func {{ .CamelValue }}(args ...any) *errors.Error {
    switch len(args) {
    	case 0:
	        return errors.New({{ .HTTPCode }}, {{ .Name }}_{{ .Value }}.String(), "{{ .Message }}")
    	case 1:
	        return errors.New({{ .HTTPCode }}, {{ .Name }}_{{ .Value }}.String(), "{{ .Message }}:"+fmt.Sprint(args[0]))
    	default:
    	    msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
	        return errors.New({{ .HTTPCode }}, {{ .Name }}_{{ .Value }}.String(), "{{ .Message }}:"+msg)
    }
}
{{- end }}

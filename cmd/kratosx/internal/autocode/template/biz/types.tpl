package {{.Module}}

type List{{.Object}}Request struct {
{{.ListFields}}
}

type Export{{.Object}}Request struct {
{{.ExportFields}}
}

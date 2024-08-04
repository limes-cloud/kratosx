package types

type Get{{.Object}}Request struct {
{{.GetFields}}
}

type List{{.Object}}Request struct {
{{.ListFields}}
}

type ListTrash{{.Object}}Request struct {
{{.ListFields}}
}

type Export{{.Object}}Request struct {
{{.ExportFields}}
}

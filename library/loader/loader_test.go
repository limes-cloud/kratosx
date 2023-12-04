package loader

import (
	"reflect"
	"testing"
)

func TestCert_Get(t *testing.T) {
	confList := map[string]string{
		"name": "./test.txt",
	}

	Init(confList, nil)

	tests := []struct {
		input string
		want  string
	}{
		{
			input: "name",
			want:  "hello world\n",
		},
	}

	for _, item := range tests {
		if !reflect.DeepEqual(string(Instance().Get(item.input)), item.want) {
			t.Errorf("read cert file error :%v", item.input)
		}
	}

}

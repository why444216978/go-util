package conversion

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestDeepCopy(t *testing.T) {
	type User struct {
		Name string
	}

	type args struct {
		dst interface{}
		src interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test success",
			args: args{
				dst: &User{},
				src: &User{Name: "golang"},
			},
			wantErr: false,
		},
		{
			name: "test dst nil error",
			args: args{
				dst: nil,
				src: &User{Name: "golang"},
			},
			wantErr: false,
		},
		{
			name: "test src nill error",
			args: args{
				dst: &User{},
				src: nil,
			},
			wantErr: true,
		},
		{
			name: "test dst and src nil error",
			args: args{
				dst: nil,
				src: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeepCopy(tt.args.dst, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("DeepCopy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJsonEncode(t *testing.T) {
	type User struct {
		Name string `json:"name"`
	}

	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test nil",
			args:    args{v: nil},
			wantErr: false,
		},
		{
			name:    "test struct",
			args:    args{v: User{Name: "golang"}},
			wantErr: false,
		},
		{
			name:    "test struct point",
			args:    args{v: &User{Name: "golang"}},
			wantErr: false,
		},
		{
			name:    "test map",
			args:    args{v: map[string]interface{}{"name": "golang"}},
			wantErr: false,
		},
		{
			name:    "test slice",
			args:    args{v: []string{"golang"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := JsonEncode(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestJsonToMapArray(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test success",
			args:    args{data: `[{"a":"a","b":1},{"a":"a","b":1}]`},
			wantErr: false,
		},
		{
			name:    "test empty sting",
			args:    args{data: ""},
			wantErr: false,
		},
		{
			name:    "test error",
			args:    args{data: `{"a":"a","b":1},{"a":"a","b":1}`},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := JsonToMapArray(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonToMapArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestJsonToMap(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test success",
			args:    args{data: `{"a":"a","b":1}`},
			wantErr: false,
		},
		{
			name:    "test empty string",
			args:    args{data: ""},
			wantErr: false,
		},
		{
			name:    "test error",
			args:    args{data: `{"a":"a","b":1`},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := JsonToMap(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestReaderToStruct(t *testing.T) {
	type User struct {
		Name string
	}

	type args struct {
		reader io.Reader
		val    interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test success",
			args: args{
				reader: strings.NewReader(`{"name":"golang"}`),
				val:    &User{},
			},
			wantErr: false,
		},
		{
			name: "test error",
			args: args{
				reader: strings.NewReader(`{"name":"golang"`),
				val:    &User{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReaderToStruct(tt.args.reader, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("ReaderToStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStructToMap(t *testing.T) {
	type User struct {
		Name string `json:"name"`
	}

	type UserNoJSON struct {
		Name string `json:"name"`
	}

	type args struct {
		obj interface{}
	}
	tests := []struct {
		name     string
		args     args
		wantData map[string]interface{}
		wantErr  bool
	}{
		{
			name:     "test success",
			args:     args{obj: User{Name: "golang"}},
			wantData: map[string]interface{}{"name": "golang"},
			wantErr:  false,
		},
		{
			name:     "test success no json tag",
			args:     args{obj: UserNoJSON{Name: "golang"}},
			wantData: map[string]interface{}{"name": "golang"},
			wantErr:  false,
		},
		{
			name:     "test nil",
			args:     args{obj: nil},
			wantData: map[string]interface{}{},
			wantErr:  true,
		},
		{
			name:     "test string",
			args:     args{obj: "a"},
			wantData: map[string]interface{}{},
			wantErr:  true,
		},
		{
			name:     "test int",
			args:     args{obj: 6},
			wantData: map[string]interface{}{},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, err := StructToMap(tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("StructToMap() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestStructToJson(t *testing.T) {
	type User struct {
		Name string `json:"name"`
	}

	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "test success",
			args:    args{v: User{Name: "golang"}},
			want:    `{"name":"golang"}`,
			wantErr: false,
		},
		{
			name:    "test nil",
			args:    args{v: nil},
			want:    "",
			wantErr: false,
		},
		{
			name:    "test string",
			args:    args{v: "a"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "test int",
			args:    args{v: 6},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StructToJson(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructToJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StructToJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructToJsonByReflect(t *testing.T) {
	type User struct {
		Name string `json:"name"`
	}

	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "test success",
			args:    args{v: User{Name: "golang"}},
			want:    `{"name":"golang"}`,
			wantErr: false,
		},
		{
			name:    "test nil",
			args:    args{v: nil},
			want:    "",
			wantErr: false,
		},
		{
			name:    "test string",
			args:    args{v: "a"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "test int",
			args:    args{v: 6},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StructToJsonByReflect(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructToJsonByReflect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StructToJsonByReflect() = %v, want %v", got, tt.want)
			}
		})
	}
}

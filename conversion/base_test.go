package conversion

import (
	"reflect"
	"testing"
)

func TestStringToByte(t *testing.T) {
	str := "golang"

	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "test success",
			args: args{
				str: str,
			},
			want: []byte(str),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToByte(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteToString(t *testing.T) {
	bytes := []byte("golang")

	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test success",
			args: args{
				b: bytes,
			},
			want: string(bytes),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteToString(tt.args.b); got != tt.want {
				t.Errorf("ByteToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToUint8(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "test int",
			args: args{str: "1"},
			want: 1,
		},
		{
			name: "test string",
			args: args{str: "a"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToUint8(tt.args.str); got != tt.want {
				t.Errorf("StringToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToUint16(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "test int",
			args: args{str: "1"},
			want: 1,
		},
		{
			name: "test string",
			args: args{str: "a"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToUint16(tt.args.str); got != tt.want {
				t.Errorf("StringToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToUint64(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "test int",
			args: args{str: "1"},
			want: 1,
		},
		{
			name: "test string",
			args: args{str: "a"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToUint64(tt.args.str); got != tt.want {
				t.Errorf("StringToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToInt8(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			name: "test int",
			args: args{str: "1"},
			want: 1,
		},
		{
			name: "test string",
			args: args{str: "a"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToInt8(tt.args.str); got != tt.want {
				t.Errorf("StringToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToInt16(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			name: "test int",
			args: args{str: "1"},
			want: 1,
		},
		{
			name: "test string",
			args: args{str: "a"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToInt16(tt.args.str); got != tt.want {
				t.Errorf("StringToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToInt32(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test int",
			args: args{str: "1"},
			want: 1,
		},
		{
			name: "test string",
			args: args{str: "a"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToInt32(tt.args.str); got != tt.want {
				t.Errorf("StringToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToInt64(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test int",
			args: args{str: "1"},
			want: 1,
		},
		{
			name: "test string",
			args: args{str: "a"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToInt64(tt.args.str); got != tt.want {
				t.Errorf("StringToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceStringToUint8(t *testing.T) {
	type args struct {
		str interface{}
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "test slice",
			args: args{str: []int{1}},
			want: 0,
		},
		{
			name: "test nil",
			args: args{str: nil},
			want: 0,
		},
		{
			name: "test int",
			args: args{str: "1"},
			want: 1,
		},
		{
			name: "test string",
			args: args{str: "a"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceStringToUint8(tt.args.str); got != tt.want {
				t.Errorf("InterfaceStringToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceStringToUint16(t *testing.T) {
	type args struct {
		str interface{}
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "test slice",
			args: args{str: []int{1}},
			want: 0,
		},
		{
			name: "test nil",
			args: args{str: nil},
			want: 0,
		},
		{
			name: "test int",
			args: args{str: "1"},
			want: 1,
		},
		{
			name: "test string",
			args: args{str: "a"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceStringToUint16(tt.args.str); got != tt.want {
				t.Errorf("InterfaceStringToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceStringToUint64(t *testing.T) {
	type args struct {
		str interface{}
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "test slice",
			args: args{str: []int{1}},
			want: 0,
		},
		{
			name: "test nil",
			args: args{str: nil},
			want: 0,
		},
		{
			name: "test int",
			args: args{str: "1"},
			want: 1,
		},
		{
			name: "test string",
			args: args{str: "a"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceStringToUint64(tt.args.str); got != tt.want {
				t.Errorf("InterfaceStringToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceStringToInt8(t *testing.T) {
	type args struct {
		str interface{}
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			name: "test slice",
			args: args{str: []int{1}},
			want: 0,
		},
		{
			name: "test nil",
			args: args{str: nil},
			want: 0,
		},
		{
			name: "test int",
			args: args{str: "1"},
			want: 1,
		},
		{
			name: "test string",
			args: args{str: "a"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceStringToInt8(tt.args.str); got != tt.want {
				t.Errorf("InterfaceStringToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceStringToInt16(t *testing.T) {
	type args struct {
		str interface{}
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			name: "test slice",
			args: args{str: []int{1}},
			want: 0,
		},
		{
			name: "test nil",
			args: args{str: nil},
			want: 0,
		},
		{
			name: "test int",
			args: args{str: "1"},
			want: 1,
		},
		{
			name: "test string",
			args: args{str: "a"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceStringToInt16(tt.args.str); got != tt.want {
				t.Errorf("InterfaceStringToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceStringToInt32(t *testing.T) {
	type args struct {
		str interface{}
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test slice",
			args: args{str: []int{1}},
			want: 0,
		},
		{
			name: "test nil",
			args: args{str: nil},
			want: 0,
		},
		{
			name: "test int",
			args: args{str: "1"},
			want: 1,
		},
		{
			name: "test string",
			args: args{str: "a"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceStringToInt32(tt.args.str); got != tt.want {
				t.Errorf("InterfaceStringToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceStringToInt64(t *testing.T) {
	type args struct {
		str interface{}
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test slice",
			args: args{str: []int{1}},
			want: 0,
		},
		{
			name: "test nil",
			args: args{str: nil},
			want: 0,
		},
		{
			name: "test int",
			args: args{str: "1"},
			want: 1,
		},
		{
			name: "test string",
			args: args{str: "a"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceStringToInt64(tt.args.str); got != tt.want {
				t.Errorf("InterfaceStringToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceFloat64ToUint8(t *testing.T) {
	type args struct {
		str interface{}
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "test slice",
			args: args{str: []int{1}},
			want: 0,
		},
		{
			name: "test nil",
			args: args{str: nil},
			want: 0,
		},
		{
			name: "test float64",
			args: args{str: 1.1},
			want: 1,
		},
		{
			name: "test int",
			args: args{str: 1},
			want: 0,
		},
		{
			name: "test string",
			args: args{str: "a"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceFloat64ToUint8(tt.args.str); got != tt.want {
				t.Errorf("InterfaceFloat64ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceFloat64ToUint16(t *testing.T) {
	type args struct {
		str interface{}
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "test slice",
			args: args{str: []int{1}},
			want: 0,
		},
		{
			name: "test nil",
			args: args{str: nil},
			want: 0,
		},
		{
			name: "test float64",
			args: args{str: 1.1},
			want: 1,
		},
		{
			name: "test int",
			args: args{str: 1},
			want: 0,
		},
		{
			name: "test string",
			args: args{str: "a"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceFloat64ToUint16(tt.args.str); got != tt.want {
				t.Errorf("InterfaceFloat64ToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceFloat64ToUint32(t *testing.T) {
	type args struct {
		str interface{}
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "test slice",
			args: args{str: []int{1}},
			want: 0,
		},
		{
			name: "test nil",
			args: args{str: nil},
			want: 0,
		},
		{
			name: "test float64",
			args: args{str: 1.1},
			want: 1,
		},
		{
			name: "test int",
			args: args{str: 1},
			want: 0,
		},
		{
			name: "test string",
			args: args{str: "a"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceFloat64ToUint32(tt.args.str); got != tt.want {
				t.Errorf("InterfaceFloat64ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceFloat64ToUint64(t *testing.T) {
	type args struct {
		str interface{}
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "test slice",
			args: args{str: []int{1}},
			want: 0,
		},
		{
			name: "test nil",
			args: args{str: nil},
			want: 0,
		},
		{
			name: "test float64",
			args: args{str: 1.1},
			want: 1,
		},
		{
			name: "test int",
			args: args{str: 1},
			want: 0,
		},
		{
			name: "test string",
			args: args{str: "a"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceFloat64ToUint64(tt.args.str); got != tt.want {
				t.Errorf("InterfaceFloat64ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceFloat64ToInt8(t *testing.T) {
	type args struct {
		str interface{}
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			name: "test slice",
			args: args{str: []int{1}},
			want: 0,
		},
		{
			name: "test nil",
			args: args{str: nil},
			want: 0,
		},
		{
			name: "test float64",
			args: args{str: 1.1},
			want: 1,
		},
		{
			name: "test int",
			args: args{str: 1},
			want: 0,
		},
		{
			name: "test string",
			args: args{str: "a"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceFloat64ToInt8(tt.args.str); got != tt.want {
				t.Errorf("InterfaceFloat64ToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceFloat64ToInt16(t *testing.T) {
	type args struct {
		str interface{}
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			name: "test slice",
			args: args{str: []int{1}},
			want: 0,
		},
		{
			name: "test nil",
			args: args{str: nil},
			want: 0,
		},
		{
			name: "test float64",
			args: args{str: 1.1},
			want: 1,
		},
		{
			name: "test int",
			args: args{str: 1},
			want: 0,
		},
		{
			name: "test string",
			args: args{str: "a"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceFloat64ToInt16(tt.args.str); got != tt.want {
				t.Errorf("InterfaceFloat64ToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceFloat64ToInt32(t *testing.T) {
	type args struct {
		str interface{}
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test slice",
			args: args{str: []int{1}},
			want: 0,
		},
		{
			name: "test nil",
			args: args{str: nil},
			want: 0,
		},
		{
			name: "test float64",
			args: args{str: 1.1},
			want: 1,
		},
		{
			name: "test int",
			args: args{str: 1},
			want: 0,
		},
		{
			name: "test string",
			args: args{str: "a"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceFloat64ToInt32(tt.args.str); got != tt.want {
				t.Errorf("InterfaceFloat64ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceFloat64ToInt64(t *testing.T) {
	type args struct {
		str interface{}
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test slice",
			args: args{str: []int{1}},
			want: 0,
		},
		{
			name: "test nil",
			args: args{str: nil},
			want: 0,
		},
		{
			name: "test float64",
			args: args{str: 1.1},
			want: 1,
		},
		{
			name: "test int",
			args: args{str: 1},
			want: 0,
		},
		{
			name: "test string",
			args: args{str: "a"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceFloat64ToInt64(tt.args.str); got != tt.want {
				t.Errorf("InterfaceFloat64ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

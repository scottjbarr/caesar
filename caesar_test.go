package caesar

import (
	"reflect"
	"testing"
)

func TestShift(t *testing.T) {
	type args struct {
		s      []byte
		offset int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "offset 1",
			args: args{
				s:      []byte("abcd"),
				offset: 1,
			},
			want: []byte("bcda"),
		},
		{
			name: "offset 2",
			args: args{
				s:      []byte("abcd"),
				offset: 2,
			},
			want: []byte("cdab"),
		},
		{
			name: "offset 3",
			args: args{
				s:      []byte("abcd"),
				offset: 3,
			},
			want: []byte("dabc"),
		},
		{
			name: "offset 4",
			args: args{
				s:      []byte("abcd"),
				offset: 4,
			},
			want: []byte("abcd"),
		},
		{
			name: "offset 5",
			args: args{
				s:      []byte("abcd"),
				offset: 5,
			},
			want: []byte("bcda"),
		},
		{
			name: "offset -1",
			args: args{
				s:      []byte("abcd"),
				offset: -1,
			},
			want: []byte("dabc"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Shift(tt.args.s, tt.args.offset)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Offset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCipher(t *testing.T) {
	tests := []struct {
		name   string
		offset int
		want   []byte
	}{
		{"offset 1", 1, []byte("bcdefghijklmnopqrstuvwxyza")},
		{"offset 13", 13, []byte("nopqrstuvwxyzabcdefghijklm")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cipher(tt.offset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cipher(%v) %v, want %v", tt.offset, got, tt.want)
			}
		})
	}
}

func TestKey(t *testing.T) {
	c := Cipher(1)
	legend := Key(c)

	tests := []struct {
		name string
		key  byte
		want byte
	}{
		{"a", byte('a'), byte('b')},
		{"z", byte('z'), byte('a')},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := legend[tt.key]; got != tt.want {
				t.Errorf("Legend() %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTranspose(t *testing.T) {
	tests := []struct {
		name      string
		offset    int
		plaintext []byte
		want      []byte
	}{
		{"hello", 13, []byte("hello"), []byte("uryyb")},
		{"hail caesar", 13, []byte("hail caesar"), []byte("unvy pnrfne")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Transpose(tt.plaintext, tt.offset)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encrypt() %v (%s), want %v", got, string(got), tt.want)
			}
		})
	}
}

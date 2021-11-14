package intersect

import (
	"reflect"
	"testing"
)

func TestStringSlices(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Нет совпадений",
			args: args{
				a: []string{"a", "b", "c"},
				b: []string{"d", "e", "f"},
			},
			want: []string{},
		},
		{
			name: "Два совпадения",
			args: args{
				a: []string{"a", "b", "c"},
				b: []string{"d", "b", "a"},
			},
			want: []string{"b", "a"},
		},
		{
			name: "Повторяющиеся значения",
			args: args{
				a: []string{"a", "b", "c"},
				b: []string{"a", "b", "c", "a", "b", "c", "a", "b", "c"},
			},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringSlices(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}

package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	s := "121"
	for idx := 0; idx < b.N; idx++ {
		numDecodings(s)
	}
}
func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "12",
			args: args{s: "12"},
			want: 2,
		},
		{
			name: "226",
			args: args{s: "226"},
			want: 3,
		},
		{
			name: "06",
			args: args{s: "06"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}

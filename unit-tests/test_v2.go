package un_test

import "testing"

func Test_totalPrice1(t *testing.T) {
	type parameters struct {
		nights  uint
		rate    uint
		cityTax uint
	}

	type testCase struct {
		name string
		args parameters
		want uint
	}

	tests := []testCase{
		{
			name: "test 0 nights",
			args: parameters{nights: 0, rate: 150, cityTax: 12},
			want: 0,
		},
		{
			name: "test 1 nights",
			args: parameters{nights: 0, rate: 100, cityTax: 12},
			want: 112,
		},
		{
			name: "test 2 nights",
			args: parameters{nights: 0, rate: 100, cityTax: 12},
			want: 224,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalPrice(tt.args.nights, tt.args.rate, tt.args.cityTax); got != tt.want {
				t.Errorf("totalPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

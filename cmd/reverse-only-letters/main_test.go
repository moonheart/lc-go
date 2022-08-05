package reverse_only_letters

import "testing"

func Test_reverseOnlyLetters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"ab-cd"}, "dc-ba"},
		{"1", args{"a-bC-dEf-ghIj"}, "j-Ih-gfE-dCba"},
		{"1", args{"Test1ng-Leet=code-Q!"}, "Qedo1ct-eeLg=ntse-T!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseOnlyLetters(tt.args.s); got != tt.want {
				t.Errorf("reverseOnlyLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

package common

import (
	"VirtualFileSystem/constants"
	"testing"
)

func TestValidName(t *testing.T) {
	defaultPattern := ""

	type args struct {
		name    string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "lowercase",
			args: args{name: "luke", pattern: defaultPattern},
			want: true,
		},
		{
			name: "uppercase",
			args: args{name: "LUKE", pattern: defaultPattern},
			want: true,
		},
		{
			name: "mix lowercase and uppercase",
			args: args{name: "Luke", pattern: defaultPattern},
			want: true,
		},
		{
			name: "mix lowercase and uppercase and dash and underline",
			args: args{name: "L_u-k_e", pattern: defaultPattern},
			want: true,
		},
		{
			name: "contains other char",
			args: args{name: "Like+", pattern: defaultPattern},
			want: false,
		},
		{
			name: "contains other char",
			args: args{name: "Like/", pattern: defaultPattern},
			want: false,
		},
		{
			name: "empty",
			args: args{name: "", pattern: defaultPattern},
			want: false,
		},
		{
			name: "invalid pattern",
			args: args{name: "Luke", pattern: "a)b"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidName(tt.args.name, tt.args.pattern); got != tt.want {
				t.Errorf("ValidName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidUserName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "lowercase",
			args: args{name: "luke"},
			want: true,
		},
		{
			name: "uppercase",
			args: args{name: "LUKE"},
			want: true,
		},
		{
			name: "mix lowercase and uppercase",
			args: args{name: "Luke"},
			want: true,
		},
		{
			name: "mix lowercase and uppercase and dash and underline",
			args: args{name: "L_u-k_e"},
			want: true,
		},
		{
			name: "contains other char",
			args: args{name: "Like+"},
			want: false,
		},
		{
			name: "contains other char",
			args: args{name: "Like/"},
			want: false,
		},
		{
			name: "empty",
			args: args{name: ""},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidUserName(tt.args.name); got != tt.want {
				t.Errorf("ValidUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidFileName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "lowercase",
			args: args{name: "luke"},
			want: true,
		},
		{
			name: "uppercase",
			args: args{name: "LUKE"},
			want: true,
		},
		{
			name: "mix lowercase and uppercase",
			args: args{name: "Luke"},
			want: true,
		},
		{
			name: "mix lowercase and uppercase and dash and underline",
			args: args{name: "L_u-k_e"},
			want: true,
		},
		{
			name: "contains other char",
			args: args{name: "Like+"},
			want: false,
		},
		{
			name: "contains other char",
			args: args{name: "Like/"},
			want: false,
		},
		{
			name: "empty",
			args: args{name: ""},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidFileName(tt.args.name); got != tt.want {
				t.Errorf("ValidFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidFolderName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "lowercase",
			args: args{name: "luke"},
			want: true,
		},
		{
			name: "uppercase",
			args: args{name: "LUKE"},
			want: true,
		},
		{
			name: "mix lowercase and uppercase",
			args: args{name: "Luke"},
			want: true,
		},
		{
			name: "mix lowercase and uppercase and dash and underline",
			args: args{name: "L_u-k_e"},
			want: true,
		},
		{
			name: "contains other char",
			args: args{name: "Like+"},
			want: false,
		},
		{
			name: "contains other char",
			args: args{name: "Like/"},
			want: false,
		},
		{
			name: "empty",
			args: args{name: ""},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidFolderName(tt.args.name); got != tt.want {
				t.Errorf("ValidFolderName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidSortOption(t *testing.T) {
	type args struct {
		sortOption string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "SortByCreated",
			args: args{
				sortOption: constants.OptionSortByCreated,
			},
			want: true,
		},
		{
			name: "SortByName",
			args: args{
				sortOption: constants.OptionSortByName,
			},
			want: true,
		},
		{
			name: "SortByCreated+",
			args: args{
				sortOption: constants.OptionSortByCreated + "+",
			},
			want: false,
		},
		{
			name: "SortByName+",
			args: args{
				sortOption: constants.OptionSortByName + "+",
			},
			want: false,
		},
		{
			name: "sort-created without prefix --",
			args: args{
				sortOption: "sort-created",
			},
			want: false,
		},
		{
			name: "sort-name without prefix --",
			args: args{
				sortOption: "sort-name",
			},
			want: false,
		},
		{
			name: "empty",
			args: args{
				sortOption: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidSortOption(tt.args.sortOption); got != tt.want {
				t.Errorf("ValidSortOption() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidSortFlag(t *testing.T) {
	type args struct {
		sortFlag string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Asc",
			args: args{
				sortFlag: constants.FlagSortAsc,
			},
			want: true,
		},
		{
			name: "Desc",
			args: args{
				sortFlag: constants.FlagSortDesc,
			},
			want: true,
		},
		{
			name: "asc+",
			args: args{
				sortFlag: constants.FlagSortAsc + "+",
			},
			want: false,
		},
		{
			name: "desc+",
			args: args{
				sortFlag: constants.FlagSortDesc + "+",
			},
			want: false,
		},
		{
			name: "empty",
			args: args{
				sortFlag: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidSortFlag(tt.args.sortFlag); got != tt.want {
				t.Errorf("ValidSortFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}

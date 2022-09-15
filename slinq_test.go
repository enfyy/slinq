package slinq

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestAggregate(t *testing.T) {
	type args struct {
		slice       []string
		initial     string
		accumulator func(string, string) string
	}

	stringLengthCompare := func(s string, s2 string) string {
		if len(s2) > len(s) {
			return s2
		} else {
			return s
		}
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Should return the longest string out of the slice",
			args: args{[]string{"apple", "pineapple", "orange", "lemon", "banana", "pear"}, "pear",
				stringLengthCompare},
			want: "pineapple",
		},
		{
			name: "Should return the initial value",
			args: args{[]string{"apple", "pineapple", "orange", "lemon", "banana", "pear"}, "passion-fruit",
				stringLengthCompare,
			},
			want: "passion-fruit",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Aggregate(tt.args.slice, tt.args.initial, tt.args.accumulator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Aggregate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAll(t *testing.T) {
	type args struct {
		slice     []string
		condition func(string) bool
	}

	stringContainsLetterO := func(s string) bool {
		return strings.Contains(s, "o")
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should return true",
			args: args{[]string{"yo", "ho", "go"}, stringContainsLetterO},
			want: true,
		},
		{
			name: "Should return false",
			args: args{[]string{"yes", "no", "maybe"}, stringContainsLetterO},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := All(tt.args.slice, tt.args.condition); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAny(t *testing.T) {
	type args struct {
		slice     []string
		condition func(string) bool
	}

	stringContainsLetterO := func(s string) bool {
		return strings.Contains(s, "o")
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should return true",
			args: args{[]string{"yo", "ho", "go"}, stringContainsLetterO},
			want: true,
		},
		{
			name: "Should return false",
			args: args{[]string{"yes", "no", "maybe"}, stringContainsLetterO},
			want: true,
		},
		{
			name: "Should return false",
			args: args{[]string{"yes", "never", "maybe"}, stringContainsLetterO},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Any(tt.args.slice, tt.args.condition); got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChunk(t *testing.T) {
	type args struct {
		slice []int
		size  int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]int
		wantErr bool
	}{
		{
			name: "Should return a slice of 3 slices with length 2",
			args: args{
				slice: []int{1, 1, 2, 2, 3, 3},
				size:  2,
			},
			want:    [][]int{{1, 1}, {2, 2}, {3, 3}},
			wantErr: false,
		},
		{
			name: "Should return a slice of 3 slices with length 6",
			args: args{
				slice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18},
				size:  6,
			},
			want:    [][]int{{1, 2, 3, 4, 5, 6}, {7, 8, 9, 10, 11, 12}, {13, 14, 15, 16, 17, 18}},
			wantErr: false,
		},
		{
			name: "Should return a slice of 3 slices with length 6 and one slice of length 1",
			args: args{
				slice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19},
				size:  6,
			},
			want:    [][]int{{1, 2, 3, 4, 5, 6}, {7, 8, 9, 10, 11, 12}, {13, 14, 15, 16, 17, 18}, {19}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Chunk(tt.args.slice, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chunk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCount(t *testing.T) {
	type args struct {
		slice     []string
		condition func(string) bool
	}

	stringContainsLetterO := func(s string) bool {
		return strings.Contains(s, "o")
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should count 3",
			args: args{
				slice:     []string{"hello", "world", "what", "is", "up", "with", "you"},
				condition: stringContainsLetterO,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Count(tt.args.slice, tt.args.condition); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinct(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should remove duplicate integers",
			args: args{
				slice: []int{1, 2, 3, 3, 4, 4, 4, 5, 6, 6},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(Distinct(tt.args.slice)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Distinct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExcept(t *testing.T) {
	type args struct {
		first  []int
		second []int
	}
	lengthTests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should return empty slice",
			args: args{
				first:  []int{1, 2, 3},
				second: []int{1, 2, 3},
			},
			want: 0,
		},
		{
			name: "Should return slice with length 3",
			args: args{
				first:  []int{1, 2, 3},
				second: []int{},
			},
			want: 3,
		},
	}
	for _, tt := range lengthTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(Except(tt.args.first, tt.args.second)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Except() = %v, want %v", got, tt.want)
			}
		})
	}

	contentTests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Should return first slice",
			args: args{
				first:  []int{1, 2, 3},
				second: []int{},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "Should return slice which contains only the number 1",
			args: args{
				first:  []int{1, 2, 3},
				second: []int{2, 3},
			},
			want: []int{1},
		},
	}
	for _, tt := range contentTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Except(tt.args.first, tt.args.second); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Except() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirst(t *testing.T) {
	type args struct {
		slice     []int
		condition func(int) bool
	}

	divisibleByThree := func(num int) bool {
		return num%3 == 0
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Should return error because slice is empty.",
			args: args{
				slice:     []int{},
				condition: divisibleByThree,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Should return the first number that is divsible by 3 (-> 6).",
			args: args{
				slice:     []int{4, 5, 6, 7, 8},
				condition: divisibleByThree,
			},
			want:    6,
			wantErr: false,
		},
		{
			name: "Should return empty slice.",
			args: args{
				slice:     []int{4, 5, 1, 7, 8},
				condition: divisibleByThree,
			},
			want:    4,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := First(tt.args.slice, tt.args.condition)
			if (err != nil) != tt.wantErr {
				t.Errorf("First() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("First() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	type args struct {
		first  []int
		second []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Should return the elements that are contained in both slices.",
			args: args{
				first:  []int{1, 2, 3},
				second: []int{3, 4, 5},
			},
			want: []int{3},
		},
		{
			name: "Should return an empty slice.",
			args: args{
				first:  []int{1, 2, 3},
				second: []int{4, 5, 6},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersect(tt.args.first, tt.args.second); !reflect.DeepEqual(got, tt.want) && len(got) != 0 {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepeat(t *testing.T) {
	type args struct {
		value int
		count int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Should return slice which contains the number 1 five times",
			args: args{
				value: 1,
				count: 5,
			},
			want: []int{1, 1, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Repeat(tt.args.value, tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Should return elements in reverse order.",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
			},
			want: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelect(t *testing.T) {
	type someStruct struct {
		name string
		id   int
	}

	type args struct {
		slice    []someStruct
		selector func(someStruct) string
	}

	selector := func(s someStruct) string {
		return s.name
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Should return slice of the names.",
			args: args{
				slice: []someStruct{{
					name: "one",
					id:   1,
				}, {
					name: "two",
					id:   2,
				}, {
					name: "three",
					id:   3,
				}},
				selector: selector,
			},
			want: []string{"one", "two", "three"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Select(tt.args.slice, tt.args.selector); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectMany(t *testing.T) {
	type someStruct struct {
		names []string
		id    int
	}

	selector := func(s someStruct, i int) []string {
		return Where(s.names, func(n string) bool { return strings.Contains(n, "a") })
	}

	type args struct {
		slice    []someStruct
		selector func(someStruct, int) []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Should return names of the struct where name contains the letter a",
			args: args{
				slice: []someStruct{{
					names: []string{"bobo", "bibi", "baba"},
					id:    0,
				}, {
					names: []string{"dodo", "dada", "didi"},
					id:    1,
				}, {
					names: []string{"sasa", "sisi", "soso"},
					id:    2,
				}},
				selector: selector,
			},
			want: []string{"baba", "dada", "sasa"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectMany(tt.args.slice, tt.args.selector); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectMany() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingle(t *testing.T) {
	type args struct {
		slice     []string
		condition func(string) bool
	}

	selector := func(s string) bool {
		return strings.Contains(s, "x")
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Should return the one string that contains the letter x",
			args: args{
				slice:     []string{"abc", "def", "ghi", "xyz"},
				condition: selector,
			},
			want:    "xyz",
			wantErr: false,
		},
		{
			name: "Should return error because more than one string contain the letter x",
			args: args{
				slice:     []string{"abc", "vwx", "ghi", "xyz"},
				condition: selector,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Single(tt.args.slice, tt.args.condition)
			if (err != nil) != tt.wantErr {
				t.Errorf("Single() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Single() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMap(t *testing.T) {
	type someStruct struct {
		name string
		id   int
	}

	keySelector := func(s someStruct) int {
		return s.id
	}

	valueSelector := func(s someStruct) string {
		return s.name
	}

	type args struct {
		slice         []someStruct
		keySelector   func(someStruct) int
		valueSelector func(someStruct) string
	}
	tests := []struct {
		name string
		args args
		want map[int]string
	}{
		{
			name: "Should convert to a map.",
			args: args{
				slice: []someStruct{{
					name: "one",
					id:   1,
				}, {
					name: "two",
					id:   2,
				}, {
					name: "three",
					id:   3,
				}},
				keySelector:   keySelector,
				valueSelector: valueSelector,
			},
			want: map[int]string{1: "one", 2: "two", 3: "three"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToMap(tt.args.slice, tt.args.keySelector, tt.args.valueSelector); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSlice(t *testing.T) {
	type someStruct struct {
		name string
		id   int
	}

	selector := func(i int, n string) someStruct {
		return someStruct{n, i}
	}

	type args struct {
		dict     map[int]string
		selector func(int, string) someStruct
	}
	tests := []struct {
		name string
		args args
		want []someStruct
	}{
		{
			name: "Should convert to slice of struct",
			args: args{
				dict:     map[int]string{1: "one", 2: "two", 3: "three"},
				selector: selector,
			},
			want: []someStruct{{
				name: "one",
				id:   1,
			}, {
				name: "two",
				id:   2,
			}, {
				name: "three",
				id:   3,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSlice(tt.args.dict, tt.args.selector); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWhere(t *testing.T) {
	type args struct {
		slice     []string
		condition func(string) bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Should return only the strings that contain the letter x",
			args: args{
				slice: []string{"abc", "def", "ghi", "xyz", "vwx", "ooo", "xxx"},
				condition: func(s string) bool {
					return strings.Contains(s, "x")
				},
			},
			want: []string{"xyz", "vwx", "xxx"},
		},
		{
			name: "Should return empty slice",
			args: args{
				slice: []string{"abc", "def", "ghi", "xyz", "vwx", "ooo", "xxx"},
				condition: func(s string) bool {
					return strings.Contains(s, "s")
				},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Where(tt.args.slice, tt.args.condition); !reflect.DeepEqual(got, tt.want) && len(got) != 0 {
				t.Errorf("Where() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZip(t *testing.T) {
	type args struct {
		first    []int
		second   []string
		selector func(int, string) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Should zip the input.",
			args: args{
				first:  []int{1, 2, 3, 4, 5},
				second: []string{"one:", "two:", "three:", "four:", "five:"},
				selector: func(i int, s string) string {
					return fmt.Sprintf("%s%d", s, i)
				},
			},
			want: []string{"one:1", "two:2", "three:3", "four:4", "five:5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Zip(tt.args.first, tt.args.second, tt.args.selector); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zip() = %v, want %v", got, tt.want)
			}
		})
	}
}

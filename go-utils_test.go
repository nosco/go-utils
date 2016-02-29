/*
Basically a copy of github/segmentio's test cases.

Thank you @tj for switching to Go just before we did! ;)
*/

package utils

import (
	"strconv"
	"testing"
)

type sample struct {
	str, out string
}

func TestSlug(t *testing.T) {
	samples := []sample{
		{"sample text", "sample-text"},
		{"sample-text", "sample-text"},
		{"sample_text", "sample-text"},
		{"sample___text", "sample-text"},
		{"sampleText", "sample-text"},
		{"inviteYourCustomersAddInvites", "invite-your-customers-add-invites"},
		{"sample 2 Text", "sample-2-text"},
		{"   sample   2    Text   ", "sample-2-text"},
		{"   $#$sample   2    Text   ", "sample-2-text"},
		{"SAMPLE 2 TEXT", "sample-2-text"},
		{"___$$Base64Encode", "base64-encode"},
		{"---$$Base64-_-_-Encode", "base64-encode"},
		{"FOO:BAR$BAZ", "foo-bar-baz"},
		{"FOO#BAR#BAZ", "foo-bar-baz"},
		{"something.com", "something-com"},
		{"$something%", "something"},
		{"something.com", "something-com"},
		{"•¶§ƒ˚foo˙∆˚¬", "foo"},
	}

	for _, sample := range samples {
		if out := Slug(sample.str); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.str, sample.out)
		}
	}
}

func TestUnCase(t *testing.T) {
	samples := []sample{
		{"sample text", "Sample text"},
		{"sample-text", "Sample text"},
		{"sample_text", "Sample text"},
		{"sample___text", "Sample text"},
		{"sampleText", "Sample text"},
		{"inviteYourCustomersAddInvites", "Invite your customers add invites"},
		{"sample 2 Text", "Sample 2 text"},
		{"   sample   2    Text   ", "Sample 2 text"},
		{"   $#$sample   2    Text   ", "Sample 2 text"},
		{"SAMPLE 2 TEXT", "Sample 2 text"},
		{"___$$Base64Encode", "Base64 encode"},
		{"---$$Base64-_-_-Encode", "Base64 encode"},
		{"FOO:BAR$BAZ", "Foo bar baz"},
		{"FOO#BAR#BAZ", "Foo bar baz"},
		{"something.com", "Something com"},
		{"$something%", "Something"},
		{"something.com", "Something com"},
		{"•¶§ƒ˚foo˙∆˚¬", "Foo"},
	}

	for _, sample := range samples {
		if out := UnCase(sample.str); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.str, sample.out)
		}
	}
}

func TestSnakeCase(t *testing.T) {
	samples := []sample{
		{"sample text", "sample_text"},
		{"sample-text", "sample_text"},
		{"sample_text", "sample_text"},
		{"sample___text", "sample_text"},
		{"sampleText", "sample_text"},
		{"inviteYourCustomersAddInvites", "invite_your_customers_add_invites"},
		{"sample 2 Text", "sample_2_text"},
		{"   sample   2    Text   ", "sample_2_text"},
		{"   $#$sample   2    Text   ", "sample_2_text"},
		{"SAMPLE 2 TEXT", "sample_2_text"},
		{"___$$Base64Encode", "base64_encode"},
		{"---$$Base64-_-_-Encode", "base64_encode"},
		{"FOO:BAR$BAZ", "foo_bar_baz"},
		{"FOO#BAR#BAZ", "foo_bar_baz"},
		{"something.com", "something_com"},
		{"$something%", "something"},
		{"something.com", "something_com"},
		{"•¶§ƒ˚foo˙∆˚¬", "foo"},
	}

	for _, sample := range samples {
		if out := SnakeCase(sample.str); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.str, sample.out)
		}
	}
}

func TestKebabCase(t *testing.T) {
	samples := []sample{
		{"sample text", "sample-text"},
		{"sample-text", "sample-text"},
		{"sample_text", "sample-text"},
		{"sample___text", "sample-text"},
		{"sampleText", "sample-text"},
		{"inviteYourCustomersAddInvites", "invite-your-customers-add-invites"},
		{"sample 2 Text", "sample-2-text"},
		{"   sample   2    Text   ", "sample-2-text"},
		{"   $#$sample   2    Text   ", "sample-2-text"},
		{"SAMPLE 2 TEXT", "sample-2-text"},
		{"___$$Base64Encode", "base64-encode"},
		{"---$$Base64-_-_-Encode", "base64-encode"},
		{"FOO:BAR$BAZ", "foo-bar-baz"},
		{"FOO#BAR#BAZ", "foo-bar-baz"},
		{"something.com", "something-com"},
		{"$something%", "something"},
		{"something.com", "something-com"},
		{"•¶§ƒ˚foo˙∆˚¬", "foo"},
	}

	for _, sample := range samples {
		if out := KebabCase(sample.str); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.str, sample.out)
		}
	}
}

func TestCamelCase(t *testing.T) {
	samples := []sample{
		{"sample text", "sampleText"},
		{"sample-text", "sampleText"},
		{"sample_text", "sampleText"},
		{"sample___text", "sampleText"},
		{"sampleText", "sampleText"},
		{"inviteYourCustomersAddInvites", "inviteYourCustomersAddInvites"},
		{"sample 2 Text", "sample2Text"},
		{"   sample   2    Text   ", "sample2Text"},
		{"   $#$sample   2    Text   ", "sample2Text"},
		{"SAMPLE 2 TEXT", "sample2Text"},
		{"___$$Base64Encode", "base64Encode"},
		{"---$$Base64-_-_-Encode", "base64Encode"},
		{"FOO:BAR$BAZ", "fooBarBaz"},
		{"FOO#BAR#BAZ", "fooBarBaz"},
		{"something.com", "somethingCom"},
		{"$something%", "something"},
		{"something.com", "somethingCom"},
		{"•¶§ƒ˚foo˙∆˚¬", "foo"},
	}

	for _, sample := range samples {
		if out := CamelCase(sample.str); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.str, sample.out)
		}
	}
}

func TestPascalCase(t *testing.T) {
	samples := []sample{
		{"sample text", "SampleText"},
		{"sample-text", "SampleText"},
		{"sample_text", "SampleText"},
		{"sample___text", "SampleText"},
		{"sampleText", "SampleText"},
		{"inviteYourCustomersAddInvites", "InviteYourCustomersAddInvites"},
		{"sample 2 Text", "Sample2Text"},
		{"   sample   2    Text   ", "Sample2Text"},
		{"   $#$sample   2    Text   ", "Sample2Text"},
		{"SAMPLE 2 TEXT", "Sample2Text"},
		{"___$$Base64Encode", "Base64Encode"},
		{"---$$Base64-_-_-Encode", "Base64Encode"},
		{"FOO:BAR$BAZ", "FooBarBaz"},
		{"FOO#BAR#BAZ", "FooBarBaz"},
		{"something.com", "SomethingCom"},
		{"$something%", "Something"},
		{"something.com", "SomethingCom"},
		{"•¶§ƒ˚foo˙∆˚¬", "Foo"},
	}

	for _, sample := range samples {
		if out := PascalCase(sample.str); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.str, sample.out)
		}
	}
}

func TestUniqueInts(t *testing.T) {
	expectedInts := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var ints []int
	for i := 0; i < 100; i++ {
		ints = append(ints, i%10)
	}
	uniqueInts := UniqueInts(ints)

	if len(expectedInts) != len(uniqueInts) {
		t.Error("Length of uniqueInts does not match the length of expectedInts")
	}
}

func TestUniqueStrings(t *testing.T) {
	expectedStrings := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var strings []string
	for i := 0; i < 100; i++ {
		strings = append(strings, strconv.Itoa(i%10))
	}
	uniqueStrings := UniqueStrings(strings)

	if len(expectedStrings) != len(uniqueStrings) {
		t.Error("Length of uniqueStrings does not match the length of expectedStrings")
	}
}

func TestUnique(t *testing.T) {
	expectedInts := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var ints []int
	for i := 0; i < 100; i++ {
		ints = append(ints, i%10)
	}

	tmpInts, _ := Unique(ints)
	uniqueInts := tmpInts.([]int)

	if len(expectedInts) != len(uniqueInts) {
		t.Error("Length of uniqueInts does not match the length of expectedInts")
	}
}

func BenchmarkSnakeCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SnakeCase("some sample text here_noething:too$amazing")
	}
}

func BenchmarkKebabCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SnakeCase("some sample text here_noething:too$amazing")
	}
}

func BenchmarkCamelCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CamelCase("some sample text here_noething:too$amazing")
	}
}

func BenchmarkPascalCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PascalCase("some sample text here_noething:too$amazing")
	}
}

func BenchmarkUniqueIntsSmallFewUniques(b *testing.B) {
	var ints []int
	for i := 0; i < 100; i++ {
		ints = append(ints, i%10)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = UniqueInts(ints)
	}
}

func BenchmarkUniqueIntsSmallManyUniques(b *testing.B) {
	var ints []int
	for i := 0; i < 100; i++ {
		ints = append(ints, i%50)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = UniqueInts(ints)
	}
}

func BenchmarkUniqueIntsLargeFewUniques(b *testing.B) {
	var ints []int
	for i := 0; i < 100000; i++ {
		ints = append(ints, i%10)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = UniqueInts(ints)
	}
}

func BenchmarkUniqueIntsLargeManyUniques(b *testing.B) {
	var ints []int
	for i := 0; i < 100000; i++ {
		ints = append(ints, i%10000)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = UniqueInts(ints)
	}
}

func BenchmarkUniqueStringsSmallFewUniques(b *testing.B) {
	var strings []string
	for i := 0; i < 100; i++ {
		strings = append(strings, strconv.Itoa(i%10))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = UniqueStrings(strings)
	}
}

func BenchmarkUniqueStringsSmallManyUniques(b *testing.B) {
	var strings []string
	for i := 0; i < 100; i++ {
		strings = append(strings, strconv.Itoa(i%50))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = UniqueStrings(strings)
	}
}

func BenchmarkUniqueStringsLargeFewUniques(b *testing.B) {
	var strings []string
	for i := 0; i < 100000; i++ {
		strings = append(strings, strconv.Itoa(i%10))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = UniqueStrings(strings)
	}
}

func BenchmarkUniqueStringsLargeManyUniques(b *testing.B) {
	var strings []string
	for i := 0; i < 100000; i++ {
		strings = append(strings, strconv.Itoa(i%10000))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = UniqueStrings(strings)
	}
}

func BenchmarkUniqueSmallIntsFewUniques(b *testing.B) {
	var ints []int
	for i := 0; i < 100; i++ {
		ints = append(ints, i%10)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = Unique(ints)
	}
}

func BenchmarkUniqueSmallIntsManyUniques(b *testing.B) {
	var ints []int
	for i := 0; i < 100; i++ {
		ints = append(ints, i%50)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = Unique(ints)
	}
}

func BenchmarkUniqueLargeIntsFewUniques(b *testing.B) {
	var ints []int
	for i := 0; i < 100000; i++ {
		ints = append(ints, i%10)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = Unique(ints)
	}
}

func BenchmarkUniqueLargeIntsManyUniques(b *testing.B) {
	var ints []int
	for i := 0; i < 100000; i++ {
		ints = append(ints, i%10000)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = Unique(ints)
	}
}

func BenchmarkUniqueSmallStringsFewUniques(b *testing.B) {
	var strings []string
	for i := 0; i < 100; i++ {
		strings = append(strings, strconv.Itoa(i%10))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = Unique(strings)
	}
}

func BenchmarkUniqueSmallStringsManyUniques(b *testing.B) {
	var strings []string
	for i := 0; i < 100; i++ {
		strings = append(strings, strconv.Itoa(i%50))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = Unique(strings)
	}
}

func BenchmarkUniqueLargeStringsFewUniques(b *testing.B) {
	var strings []string
	for i := 0; i < 100000; i++ {
		strings = append(strings, strconv.Itoa(i%10))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = Unique(strings)
	}
}

func BenchmarkUniqueLargeStringsManyUniques(b *testing.B) {
	var strings []string
	for i := 0; i < 100000; i++ {
		strings = append(strings, strconv.Itoa(i%10000))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = Unique(strings)
	}
}

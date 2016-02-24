/*
Basically a copy of github/segmentio's test cases.

Thank you @tj for switching to Go just before we did! ;)
*/

package utils

import "testing"

var ops int = 1e6

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
		{"sampleIdentifierText", "sample-identifier-text"},
		{"identifierSampleText", "identifier-sample-text"},
		{"id-sampleText", "id-sample-text"},
		{"idSampleText", "id-sample-text"},
		{"sample-id-text", "sample-id-text"},
		{"sampleIDText", "sample-id-text"},
		{"sampleTextID", "sample-text-id"},
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
		{"sampleIdentifierText", "Sample identifier text"},
		{"identifierSampleText", "Identifier sample text"},
		{"id-sampleText", "ID sample text"},
		{"idSampleText", "ID sample text"},
		{"sample-id-text", "Sample ID text"},
		{"sampleIDText", "Sample ID text"},
		{"sampleTextID", "Sample text ID"},
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
		{"sampleIdentifierText", "sample_identifier_text"},
		{"identifierSampleText", "identifier_sample_text"},
		{"id-sampleText", "id_sample_text"},
		{"ID-sampleText", "id_sample_text"},
		{"idSampleText", "id_sample_text"},
		{"sample-id-text", "sample_id_text"},
		{"IDSampleText", "id_sample_text"},
		{"idSampleText", "id_sample_text"},
		{"sampleIDText", "sample_id_text"},
		{"sampleTextID", "sample_text_id"},
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
		{"sampleIdentifierText", "sample-identifier-text"},
		{"identifierSampleText", "identifier-sample-text"},
		{"id-sampleText", "id-sample-text"},
		{"IDSampleText", "id-sample-text"},
		{"idSampleText", "id-sample-text"},
		{"sample-id-text", "sample-id-text"},
		{"sampleIDText", "sample-id-text"},
		{"sampleTextID", "sample-text-id"},
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
		{"sample_identifier_text", "sampleIdentifierText"},
		{"identifier_sample_text", "identifierSampleText"},
		{"id-sampleText", "idSampleText"},
		{"idSampleText", "idSampleText"},
		{"sample-id-text", "sampleIDText"},
		{"sampleIDText", "sampleIDText"},
		{"sampleTextID", "sampleTextID"},
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
		{"id", "ID"},
		{"iD", "ID"},
		{"Id", "ID"},
		{"ID", "ID"},
		{"sample text", "SampleText"},
		{"sample-text", "SampleText"},
		{"sample_text", "SampleText"},
		{"sample___text", "SampleText"},
		{"sampleText", "SampleText"},
		{"sample_identifier_text", "SampleIdentifierText"},
		{"identifier_sample_text", "IdentifierSampleText"},
		{"id-sampleText", "IDSampleText"},
		{"idSampleText", "IDSampleText"},
		{"sample-id-text", "SampleIDText"},
		{"sampleIDText", "SampleIDText"},
		{"sampleTextID", "SampleTextID"},
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

func BenchmarkUnCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = UnCase("some sample text here_noething:too$amazing")
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

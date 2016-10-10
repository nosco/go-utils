package utils

import "testing"

func TestProtobufInfo(t *testing.T) {
	var tests = []struct {
		in  TagString
		out ProtobufInfo
	}{
		{
			`protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty" col:"comment" default:"true" nullable:"false"`,
			ProtobufInfo{"bytes", 4, true, false, false, "comment", "", true},
		},
		{
			`protobuf:"varint,5,opt,name=type,proto3,enum=models.CommentType" json:"type,omitempty" col:"type" default:"true" nullable:"false"`,
			ProtobufInfo{"varint", 5, true, false, false, "type", "models.CommentType", true},
		},
		{
			`protobuf:"bytes,6,opt,name=created,stdtime" json:"created,omitempty" col:"created" default:"true" nullable:"false"`,
			ProtobufInfo{"bytes", 6, true, false, false, "created", "", false},
		},
		{
			`protobuf:"bytes,7,opt,name=user" json:"author,omitempty" col:"users_id" default:"true" fkey:"id" ftable:"users" nullable:"false" reltype:"one"`,
			ProtobufInfo{"bytes", 7, true, false, false, "user", "", false},
		},
		{
			`protobuf:"bytes,9,rep,name=votes" json:"votes,omitempty" col:"id" default:"false" fkey:"comments_id" ftable:"votes" reltype:"many"`,
			ProtobufInfo{"bytes", 9, true, false, true, "votes", "", false},
		},
		{
			`protobuf:"varint,11,opt,name=i_like,json=iLike,proto3" json:"iLike,omitempty" custom:"true" default:"true" nullable:"false"`,
			ProtobufInfo{"varint", 11, true, false, false, "i_like", "", true},
		},
	}

	for _, tt := range tests {
		pbStr := tt.in.Get("protobuf")
		if pbStr == "" {
			t.Errorf("Unable to extract protobuf tag from %s", tt.in)
		}

		pbInfo := tt.in.ProtobufInfo()

		if pbInfo == nil {
			t.Errorf("Unable to extract protobuf info from %s", tt.in)
		}

		if *pbInfo != tt.out {
			t.Errorf("Error extracting info from\n%s\nExpected:\n\t%q\ngot:\n\t%q\n", tt.in, tt.out, *pbInfo)
		}
	}
}

package utils

import (
	"strconv"
	"strings"
)

type ProtobufTag string

type ProtobufInfo struct {
	Type      string // wire encoding
	TagNumber int    // protocol tag number
	Optional  bool   // opt,req,rep for optional, required, or repeated
	Required  bool   // opt,req,rep for optional, required, or repeated
	Repeated  bool   // opt,req,rep for optional, required, or repeated
	Name      string // name= the original declared name
	Enum      string // enum= the name of the enum type if it is an enum-typed field.
	Proto3    bool   // proto3 if this field is in a proto3 message

	// Seems like these 2 are never used - at least in proto3.
	// Packed  bool   // packed whether the encoding is "packed" (optional; repeated primitives only)
	// Default string // def= string representation of the default value, if any.
}

func (tag *TagString) ProtobufInfo() (pbInfo *ProtobufInfo) {
	pbTag := tag.Get("protobuf")

	if pbTag == "" {
		return
	}

	keyVals := strings.Split(pbTag, ",")
	pbInfo = &ProtobufInfo{
		Type: keyVals[0],
	}
	pbInfo.TagNumber, _ = strconv.Atoi(keyVals[1])

	switch keyVals[2] {
	case "opt":
		pbInfo.Optional = true
	case "req":
		pbInfo.Required = true
	case "rep":
		pbInfo.Optional = true
		pbInfo.Repeated = true
	}

	for i := 3; i < len(keyVals); i++ {
		if keyVals[i] == "proto3" {
			pbInfo.Proto3 = true

		} else if strings.HasPrefix(keyVals[i], "name=") {
			pbInfo.Name = strings.Replace(keyVals[i], "name=", "", 1)

		} else if strings.HasPrefix(keyVals[i], "enum=") {
			pbInfo.Enum = strings.Replace(keyVals[i], "enum=", "", 1)
		}
	}

	// In version 3, everything is "optional"
	if pbInfo.Proto3 {
		pbInfo.Optional = true
	}

	return
}

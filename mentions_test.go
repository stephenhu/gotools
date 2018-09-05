package gowdl

import (
	"testing"

)

var validMentions = []string{
	"hello world @abc",
  "@abc_xyz",
	"@go123",
	"@abcdefghijklmn",
	"go back to saudi mother fucking arabs @clips @hababi",
	"go @haha  best",
}

var invalidMentions = []string{
	"hello world @123",
  "@abc-xyz",
	"@_2as",
	"@abcdefghijklmntoolongboywhateva",
	"@ass@hole",
}

func TestExtractMentionsValid(t *testing.T) {

	for _, m := range validMentions {

		mentions := ExtractMentions(m)
	
		if len(mentions) == 0 {
			t.Error(mentions)
		}
	
	}

} // TestExtractMentionsValid

func TestExtractMentionsInvalid(t *testing.T) {

	for _, m := range invalidMentions {

		mentions := ExtractMentions(m)

		if len(mentions) > 0 {
			t.Error(mentions)
		}
	
	}

} // TestExtractMentionsInvalid

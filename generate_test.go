package gowdl

import(
	"testing"
)

func TestGenerateHex(t *testing.T) {

	_, err := GenerateRandomHex(32)

	if err != nil {
		t.Error(err)
	}

} // TestGenerateHex

func TestGenerateHexInvalidLength(t *testing.T) {

	_, err := GenerateRandomHex(-1)

	if err == nil {
		t.Error(err)
	}

} // TestGenerateHexInvalidLength

func TestGenerateHashAndSalt(t *testing.T) {

	_, _, err := GenerateHashAndSalt("abc", "xyz", "123", 32)

	if err != nil {
		t.Error(err)
	}

} // TestGenerateHashAndSalt

func TestGenerateHashAndSaltInvalidLength(t *testing.T) {

	_, _, err := GenerateHashAndSalt("abc", "xyz", "123", 0)

	if err == nil {
		t.Error(err)
	}

} // TestGenerateHashAndSaltInvalidLength

func TestGenerateHashAndSaltInvalidS(t *testing.T) {

	_, _, err := GenerateHashAndSalt("", "xyz", "123", 32)

	if err == nil {
		t.Error(err)
	}

} // TestGenerateHashAndSaltInvalidS

func TestGenerateHashAndSaltInvalidHK(t *testing.T) {

	_, _, err := GenerateHashAndSalt("abc", "", "123", 32)

	if err == nil {
		t.Error(err)
	}

} // TestGenerateHashAndSaltInvalidHK

func TestGenerateHashAndSaltInvalidP(t *testing.T) {

	_, _, err := GenerateHashAndSalt("abc", "xyz", "", 32)

	if err == nil {
		t.Error(err)
	}

} // TestGenerateHashAndSaltInvalidP


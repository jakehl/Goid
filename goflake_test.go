package goflake

import "testing"
import "regexp"

func TestGeneratedV4UUIDisRFC4122Compliant(t *testing.T) {
	actualResult := NewV4UUID()
	t.Logf("Generated: %s", actualResult)
	if ismatch, err := regexp.MatchString(V4UUIDRegex, actualResult.ToString()); !ismatch {
		if err != nil {
			t.Fatalf(err.Error())
		} else {
			t.Fatalf("Generated: %s which does not match regex of: %s", actualResult, V4UUIDRegex)
		}
	}
}

func TestGetUUIDFromString(t *testing.T) {
	testUUID := "40db20bd-e9ea-41d9-bed4-473b5e147582"
	resultingUUID, err := GetUUIDFromString(testUUID)
	if err != nil {
		t.Fatal(err.Error())
	} else if resultingUUID == nil {
		t.Fatal("Failed to GetUUIDFromString")
	} else {
		if resultingUUID.ToString() != testUUID {
			t.Fatalf("Input: %s does not match output of: %s", testUUID, resultingUUID.ToString())
		}
	}
}

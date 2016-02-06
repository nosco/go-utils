package utils

import "testing"

func TestSort(t *testing.T) {
	var testTag TagString = `xyz:"val1" abc:"val2"`
	testTag.Sort()
	if testTag != TagString(`abc:"val2" xyz:"val1"`) {
		t.Error("Sort doesn't seem to work correctly, got:\n", testTag)
	}
}

func TestGet(t *testing.T) {
	var testTag TagString = `var1:"val1" var2:"val2"`
	if testTag.Get("var1") != "val1" {
		t.Error("Unable to Get var1 by tag name")
	}
	if testTag.Get("var2") != "val2" {
		t.Error("Unable to Get var2 by tag name")
	}
}

// The value should only be set, if the tag name exists
func TestSet(t *testing.T) {
	var testTag TagString = `var1:"val1" var2:"val2"`
	testTag.Set("var", "val")
	if testTag.Get("var") != "" {
		t.Error("Tag name var was created with Set")
	}

	testTag.Set("var2", "val22")
	if testTag.Get("var2") != "val22" {
		t.Error("Unable to change var2 with Set")
	}

	if testTag != `var1:"val1" var2:"val22"` {
		t.Error("Set final check failed")
	}
}

func TestSetMulti(t *testing.T) {
	var testTag TagString = `var1:"val1" var2:"val2"`
	setMap := map[string]string{
		"var2": "val22",
		"var3": "val3",
	}
	testTag.SetMulti(setMap)
	if testTag.Get("var2") != "val22" {
		t.Error("var2 was not changed with AddMulti")
	}
	if testTag.Get("var3") != "" {
		t.Error("var3 was added with SetMulti")
	}
	if testTag != `var1:"val1" var2:"val22"` {
		t.Error("SetMulti final check failed")
	}
}

// The value should be changed, if the tag exists, otherwise it should be created
func TestAdd(t *testing.T) {
	var testTag TagString = `var1:"val1" var2:"val2"`
	testTag.Add("var2", "val22")
	if testTag.Get("var2") != "val22" {
		t.Error("Unable to change var2 with Add")
	}

	testTag.Add("var3", "val3")
	if testTag.Get("var3") != "val3" {
		t.Error("Unable to Add var3")
	}

	if testTag != `var1:"val1" var2:"val22" var3:"val3"` {
		t.Error("Add final check failed")
	}
}

func TestAddMulti(t *testing.T) {
	var testTag TagString = `var1:"val1" var2:"val2"`
	addMap := map[string]string{
		"var2": "val22",
		"var3": "val3",
	}
	testTag.AddMulti(addMap)
	if testTag.Get("var2") != "val22" {
		t.Error("var2 was not changed with AddMulti")
	}
	if testTag.Get("var3") != "val3" {
		t.Error("var3 was not added with AddMulti")
	}
	if testTag != `var1:"val1" var2:"val22" var3:"val3"` {
		t.Error("AddMulti final check failed")
	}
}

// The value should be removed if the tag exists
func TestRemove(t *testing.T) {
	var testTag TagString = `var1:"val1" var2:"val2"`
	testTag.Remove("var2")
	if testTag.Get("var2") != "" {
		t.Error("Unable to Remove var2")
	}

	if testTag.Get("var1") == "" {
		t.Error("var1 was removed - Only var2 should have been removed")
	}

	if testTag != `var1:"val1"` {
		t.Error("Remove final check failed")
	}
}

func TestRemoveMulti(t *testing.T) {
	var testTag TagString = `var1:"val1" var2:"val2" var3:"val3"`
	removeKeys := []string{
		"var2",
		"var3",
		"var4",
	}
	_ = testTag.RemoveMulti(removeKeys)
	if testTag.Get("var2") != "" {
		t.Error("var2 was not removed with RemoveMulti")
	}
	if testTag.Get("var3") != "" {
		t.Error("var3 was not removed with RemoveMulti")
	}
	if testTag != `var1:"val1"` {
		t.Error("RemoveMulti final check failed")
	}
}

func BenchmarkSort(b *testing.B) {
	var testTag TagString = `var2:"val2" var1:"val1" abc:"valabc" xyz:"valxyz" var3:"val3"`
	for i := 0; i < b.N; i++ {
		testTag.Sort()
	}
}

func BenchmarkGet(b *testing.B) {
	var testTag TagString = `var1:"val1" var2:"val2" var3:"val3" var4:"val4" var5:"val5"`
	for i := 0; i < b.N; i++ {
		_ = testTag.Get("var4")
	}
}

func BenchmarkSetNonExisting(b *testing.B) {
	var testTag TagString = `var1:"val1" var2:"val2" var3:"val3" var4:"val4" var5:"val5"`
	for i := 0; i < b.N; i++ {
		_ = testTag.Set("var", "val")
	}
}

func BenchmarkSetExisting(b *testing.B) {
	var testTag TagString = `var1:"val1" var2:"val2" var3:"val3" var4:"val4" var5:"val5"`
	for i := 0; i < b.N; i++ {
		_ = testTag.Set("var", "val")
	}
}

func BenchmarkAddNonExisting(b *testing.B) {
	var testTag TagString = `var1:"val1" var2:"val2" var3:"val3" var4:"val4" var5:"val5"`
	for i := 0; i < b.N; i++ {
		testTag.Add("var", "val")
	}
}

func BenchmarkAddExisting(b *testing.B) {
	var testTag TagString = `var1:"val1" var2:"val2" var3:"val3" var4:"val4" var5:"val5"`
	for i := 0; i < b.N; i++ {
		testTag.Add("var", "val")
	}
}

func BenchmarkRemoveNonExisting(b *testing.B) {
	var testTag TagString = `var1:"val1" var2:"val2" var3:"val3" var4:"val4" var5:"val5"`
	for i := 0; i < b.N; i++ {
		_ = testTag.Remove("var")
	}
}

func BenchmarkRemoveExisting(b *testing.B) {
	var testTag TagString = `var1:"val1" var2:"val2" var3:"val3" var4:"val4" var5:"val5"`
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = testTag.Remove("var")
		}
	})
}

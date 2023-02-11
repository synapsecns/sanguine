//nolint:testpackage
package bytemap

import (
	"testing"
)

func TestByteMap(t *testing.T) {
	var m ByteSliceMap[int]
	//nolint:asciicheck
	ಠ := []byte("fಠo")

	m.Put(ಠ, 123)
	m.PutString("foo", 456)

	v1, ok1 := m.GetString("fಠo")
	if !ok1 {
		t.Error("GetString('fಠo') ok=false, want: true")
	}
	if v1 != 123 {
		t.Errorf("GetString('fಠo')=%v, want: %v", v1, 123)
	}

	v2, ok2 := m.Get([]byte{'f', 'o', 'o'})
	if !ok2 {
		t.Error("Get('foo') ok=false, want: true")
	}
	//nolint: forcetypeassert
	if v2 != 456 {
		t.Errorf("Get('foo')=%v, want: %v", v2, 456)
	}
}

func TestByteMapMissingValue(t *testing.T) {
	var m ByteSliceMap[*string]
	m.PutString("foo", nil)

	v1, ok1 := m.GetString("fಠo")
	if ok1 {
		t.Error("GetString('fಠo') ok=true, want: false")
	}
	if v1 != nil {
		t.Errorf("GetString('fಠo')=%v, want want nil", v1)
	}

	v2, ok2 := m.Get([]byte{'f', 'o', 'o'})
	if !ok2 {
		t.Error("Get('foo') ok=false, want: true")
	}
	if v2 != nil {
		t.Errorf("Get('foo')=%v, want: nil", v2)
	}
}

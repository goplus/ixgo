package typesutil

import (
	"go/token"
	"go/types"
	"testing"
	"unsafe"
)

func TestSizeof(t *testing.T) {
	if unsafe.Sizeof(Scope{}) != unsafe.Sizeof(types.Scope{}) {
		t.Fatal("unexpected sizeof types.Scope")
	}
}

// TestScopeDelete tests the ScopeDelete function.
func TestScopeDelete(t *testing.T) {
	// Create a new scope
	scope := types.NewScope(nil, token.NoPos, token.NoPos, "")

	// Insert some objects into the scope
	obj1 := types.NewVar(token.NoPos, nil, "obj1", nil)
	obj2 := types.NewVar(token.NoPos, nil, "obj2", nil)
	scope.Insert(obj1)
	scope.Insert(obj2)

	// Test deleting an existing object
	deletedObj := ScopeDelete(scope, "obj1")
	if deletedObj == nil {
		t.Errorf("ScopeDelete should return the deleted object, but got nil")
	}
	if deletedObj.Name() != "obj1" {
		t.Errorf("ScopeDelete should return the correct object, but got %s", deletedObj.Name())
	}
	if obj := scope.Lookup("obj1"); obj != nil {
		t.Errorf("ScopeDelete should remove the object from the scope, but it still exists")
	}

	// Test deleting a non-existing object
	deletedObj = ScopeDelete(scope, "nonexistent")
	if deletedObj != nil {
		t.Errorf("ScopeDelete should return nil for non-existent object, but got %v", deletedObj)
	}

	// Test that other objects are not affected
	if obj := scope.Lookup("obj2"); obj == nil {
		t.Errorf("ScopeDelete should not remove other objects from the scope")
	}
}

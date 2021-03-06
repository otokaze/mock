package mockgen

import (
	"fmt"
	"testing"
)

func TestMakeArgString(t *testing.T) {
	testCases := []struct {
		argNames  []string
		argTypes  []string
		argString string
	}{
		{
			argNames:  nil,
			argTypes:  nil,
			argString: "",
		},
		{
			argNames:  []string{"arg0"},
			argTypes:  []string{"int"},
			argString: "arg0 int",
		},
		{
			argNames:  []string{"arg0", "arg1"},
			argTypes:  []string{"int", "bool"},
			argString: "arg0 int, arg1 bool",
		},
		{
			argNames:  []string{"arg0", "arg1"},
			argTypes:  []string{"int", "int"},
			argString: "arg0, arg1 int",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2"},
			argTypes:  []string{"bool", "int", "int"},
			argString: "arg0 bool, arg1, arg2 int",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2"},
			argTypes:  []string{"int", "bool", "int"},
			argString: "arg0 int, arg1 bool, arg2 int",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2"},
			argTypes:  []string{"int", "int", "bool"},
			argString: "arg0, arg1 int, arg2 bool",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2"},
			argTypes:  []string{"int", "int", "int"},
			argString: "arg0, arg1, arg2 int",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2", "arg3"},
			argTypes:  []string{"bool", "int", "int", "int"},
			argString: "arg0 bool, arg1, arg2, arg3 int",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2", "arg3"},
			argTypes:  []string{"int", "bool", "int", "int"},
			argString: "arg0 int, arg1 bool, arg2, arg3 int",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2", "arg3"},
			argTypes:  []string{"int", "int", "bool", "int"},
			argString: "arg0, arg1 int, arg2 bool, arg3 int",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2", "arg3"},
			argTypes:  []string{"int", "int", "int", "bool"},
			argString: "arg0, arg1, arg2 int, arg3 bool",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2", "arg3", "arg4"},
			argTypes:  []string{"bool", "int", "int", "int", "bool"},
			argString: "arg0 bool, arg1, arg2, arg3 int, arg4 bool",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2", "arg3", "arg4"},
			argTypes:  []string{"int", "bool", "int", "int", "bool"},
			argString: "arg0 int, arg1 bool, arg2, arg3 int, arg4 bool",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2", "arg3", "arg4"},
			argTypes:  []string{"int", "int", "bool", "int", "bool"},
			argString: "arg0, arg1 int, arg2 bool, arg3 int, arg4 bool",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2", "arg3", "arg4"},
			argTypes:  []string{"int", "int", "int", "bool", "bool"},
			argString: "arg0, arg1, arg2 int, arg3, arg4 bool",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2", "arg3", "arg4"},
			argTypes:  []string{"int", "int", "bool", "bool", "int"},
			argString: "arg0, arg1 int, arg2, arg3 bool, arg4 int",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			s := makeArgString(tc.argNames, tc.argTypes)
			if s != tc.argString {
				t.Errorf("result == %q, want %q", s, tc.argString)
			}
		})
	}
}

func TestNewIdentifierAllocator(t *testing.T) {
	a := newIdentifierAllocator([]string{"taken1", "taken2"})
	if len(a) != 2 {
		t.Fatalf("expected 2 items, got %v", len(a))
	}

	_, ok := a["taken1"]
	if !ok {
		t.Errorf("allocator doesn't contain 'taken1': %#v", a)
	}

	_, ok = a["taken2"]
	if !ok {
		t.Errorf("allocator doesn't contain 'taken2': %#v", a)
	}
}

func allocatorContainsIdentifiers(a identifierAllocator, ids []string) bool {
	if len(a) != len(ids) {
		return false
	}

	for _, id := range ids {
		_, ok := a[id]
		if !ok {
			return false
		}
	}

	return true
}

func TestIdentifierAllocator_allocateIdentifier(t *testing.T) {
	a := newIdentifierAllocator([]string{"taken"})

	t2 := a.allocateIdentifier("taken_2")
	if t2 != "taken_2" {
		t.Fatalf("expected 'taken_2', got %q", t2)
	}
	expected := []string{"taken", "taken_2"}
	if !allocatorContainsIdentifiers(a, expected) {
		t.Fatalf("allocator doesn't contain the expected items - allocator: %#v, expected items: %#v", a, expected)
	}

	t3 := a.allocateIdentifier("taken")
	if t3 != "taken_3" {
		t.Fatalf("expected 'taken_3', got %q", t3)
	}
	expected = []string{"taken", "taken_2", "taken_3"}
	if !allocatorContainsIdentifiers(a, expected) {
		t.Fatalf("allocator doesn't contain the expected items - allocator: %#v, expected items: %#v", a, expected)
	}

	t4 := a.allocateIdentifier("taken")
	if t4 != "taken_4" {
		t.Fatalf("expected 'taken_4', got %q", t4)
	}
	expected = []string{"taken", "taken_2", "taken_3", "taken_4"}
	if !allocatorContainsIdentifiers(a, expected) {
		t.Fatalf("allocator doesn't contain the expected items - allocator: %#v, expected items: %#v", a, expected)
	}

	id := a.allocateIdentifier("id")
	if id != "id" {
		t.Fatalf("expected 'id', got %q", id)
	}
	expected = []string{"taken", "taken_2", "taken_3", "taken_4", "id"}
	if !allocatorContainsIdentifiers(a, expected) {
		t.Fatalf("allocator doesn't contain the expected items - allocator: %#v, expected items: %#v", a, expected)
	}
}

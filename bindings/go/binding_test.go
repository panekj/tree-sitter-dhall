package tree_sitter_dhall_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-dhall"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_dhall.Language())
	if language == nil {
		t.Errorf("Error loading Dhall grammar")
	}
}

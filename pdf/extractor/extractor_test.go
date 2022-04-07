package extractor

import (
	"testing"
)

const f = "./extractor.pdf"

func TestExtract(t *testing.T) {
	t.Run("TestExtract", func(t *testing.T) {
		_, err := Extract(f)
		if err != nil {
			t.Errorf("TestExtractWithStyle() error = %v", err)
		}
	})
}

func TestExtractWithStyle(t *testing.T) {
	t.Run("TestExtractWithStyle", func(t *testing.T) {
		_, err := ExtractWithStyle(f)
		if err != nil {
			t.Errorf("TestExtractWithStyle() error = %v", err)
		}
	})
}

func TestExtractRows(t *testing.T) {
	t.Run("TestExtractRows", func(t *testing.T) {
		_, err := ExtractRows(f)
		if err != nil {
			t.Errorf("ExtractRows() error = %v", err)
		}
	})
}

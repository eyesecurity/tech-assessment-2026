package data

import (
	"os"
	"path/filepath"
	"testing"
)

func writeFixture(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "fixture.csv")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	return path
}

func TestReadCSV_ReturnsRecordsKeyedByHeader(t *testing.T) {
	path := writeFixture(t, "id;asset;ip\n1;laptop-1;10.0.0.1\n2;server-9;10.0.0.2\n")

	records, err := ReadCSV(path)
	if err != nil {
		t.Fatalf("ReadCSV: %v", err)
	}
	if len(records) != 2 {
		t.Fatalf("len(records) = %d, want 2", len(records))
	}
	if records[0]["asset"] != "laptop-1" {
		t.Errorf("records[0][asset] = %q, want laptop-1", records[0]["asset"])
	}
	if records[1]["ip"] != "10.0.0.2" {
		t.Errorf("records[1][ip] = %q, want 10.0.0.2", records[1]["ip"])
	}
}

func TestReadCSV_RowWithFewerFields(t *testing.T) {
	// A row with fewer fields than the header is still returned; the missing
	// trailing columns are simply absent from that record's map.
	path := writeFixture(t, "id;asset;category\n1;laptop-1;phishing\n2;server-9\n")

	records, err := ReadCSV(path)
	if err != nil {
		t.Fatalf("ReadCSV: %v", err)
	}
	if len(records) != 2 {
		t.Fatalf("len(records) = %d, want 2", len(records))
	}
	if records[1]["asset"] != "server-9" {
		t.Errorf("records[1][asset] = %q, want server-9", records[1]["asset"])
	}
	if records[1]["category"] != "" {
		t.Errorf("records[1][category] = %q, want empty (field absent)", records[1]["category"])
	}
}

func TestReadCSV_EmptyFileReturnsEmptySlice(t *testing.T) {
	path := writeFixture(t, "")

	records, err := ReadCSV(path)
	if err != nil {
		t.Fatalf("ReadCSV: %v", err)
	}
	if len(records) != 0 {
		t.Fatalf("len(records) = %d, want 0", len(records))
	}
}

func TestReadCSV_MissingFileReturnsError(t *testing.T) {
	if _, err := ReadCSV(filepath.Join(t.TempDir(), "nope.csv")); err == nil {
		t.Fatal("expected error for missing file, got nil")
	}
}

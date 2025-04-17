package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/kurochkinivan/ReportSender/upload"
)

const sampleTemplate = `
<root>
    <fsigmodule name="Module1" type="pid">
        <param name="Kp">0</param>
        <param name="Ki">0</param>
        <param name="Kd">0</param>
        <param name="integral_min">0</param>
        <param name="integral_max">0</param>
        <param name="inp_rise_deriative">0</param>
        <param name="inp_fall_deriative">0</param>
        <param name="min">0</param>
        <param name="max">0</param>
        <param name="preset_allowed_at_low">0</param>
        <param name="preset_allowed_at_high">0</param>
    </fsigmodule>
</root>`

const sampleJSON = `{
    "data": [
        {
            "name": "Module1",
            "Kp": 1.1,
            "Ki": 2.2,
            "Kd": 3.3,
            "integral_min": -10,
            "integral_max": 10,
            "inp_rise_deriative": 1,
            "inp_fall_deriative": 2,
            "min": 0,
            "max": 100,
            "preset_allowed_at_low": 5,
            "preset_allowed_at_high": 95
        }
    ]
}`

func TestProcessFile(t *testing.T) {
	tempDir := t.TempDir()
	templatePath := filepath.Join(tempDir, "template.xml")
	jsonPath := filepath.Join(tempDir, "module.json")

	os.WriteFile(templatePath, []byte(sampleTemplate), 0644)
	os.WriteFile(jsonPath, []byte(sampleJSON), 0644)

	upload.LoadXMLTemplate(templatePath)
	err := upload.ProcessFiles(tempDir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	xmlPath := strings.TrimSuffix(jsonPath, ".json") + ".xml"
	if _, err := os.Stat(xmlPath); err != nil {
		t.Fatalf("expected xml file not found: %v", err)
	}
}

func TestProcessFilePIDNotFound(t *testing.T) {
	template := `<root><fsigmodule name="Unknown" type="pid"><param name="Kp">0</param></fsigmodule></root>`
	json := `{"data": []}`

	tempDir := t.TempDir()
	os.WriteFile(filepath.Join(tempDir, "template.xml"), []byte(template), 0644)
	os.WriteFile(filepath.Join(tempDir, "data.json"), []byte(json), 0644)

	upload.LoadXMLTemplate(filepath.Join(tempDir, "template.xml"))
	err := upload.ProcessFiles(tempDir)
	if err != nil {
		t.Errorf("should not fail if no matching PIDs: %v", err)
	}
}

func TestPerformance_ProcessFiles(t *testing.T) {
	tempDir := t.TempDir()
	templatePath := filepath.Join(tempDir, "template.xml")
	os.WriteFile(templatePath, []byte(sampleTemplate), 0644)

	err := upload.LoadXMLTemplate(templatePath)
	if err != nil {
		t.Fatalf("failed to load template: %v", err)
	}

	pid := sampleJSON
	for i := 0; i < 500; i++ {
		name := filepath.Join(tempDir, fmt.Sprintf("pid_%d.json", i))
		os.WriteFile(name, []byte(pid), 0644)
	}

	start := time.Now()
	err = upload.ProcessFiles(tempDir)
	elapsed := time.Since(start)
	if err != nil {
		t.Fatalf("performance test failed: %v", err)
	}

	t.Logf("Processed 500 files in %s", elapsed)
}

func BenchmarkProcessFiles(b *testing.B) {
	tempDir := b.TempDir()
	templatePath := filepath.Join(tempDir, "template.xml")
	os.WriteFile(templatePath, []byte(sampleTemplate), 0644)

	err := upload.LoadXMLTemplate(templatePath)
	if err != nil {
		b.Fatalf("failed to load template: %v", err)
	}

	pid := sampleJSON
	for i := 0; i < 500; i++ {
		name := filepath.Join(tempDir, fmt.Sprintf("pid_%d.json", i))
		os.WriteFile(name, []byte(pid), 0644)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := upload.ProcessFiles(tempDir)
		if err != nil {
			b.Fatalf("benchmark failed: %v", err)
		}
	}
}

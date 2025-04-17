package tests

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/kurochkinivan/ReportSender/upload"
)

func writeTestFiles(dir string, template string, pidData upload.PIDData) error {
	templatePath := filepath.Join(dir, "template.xml")
	jsonPath := filepath.Join(dir, "test.json")

	err := os.WriteFile(templatePath, []byte(template), 0644)
	if err != nil {
		return err
	}

	data, err := json.Marshal(pidData)
	if err != nil {
		return err
	}

	err = os.WriteFile(jsonPath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func TestProcessFiles(t *testing.T) {
	template := `
<root>
  <fsigmodule name="PID1" type="pid">
    <param name="Kp">0</param>
    <param name="Ki">0</param>
    <param name="Kd">0</param>
  </fsigmodule>
</root>`

	pidData := upload.PIDData{
		Data: []upload.PIDConfig{
			{
				Name: "PID1",
				Kp:   1.5, Ki: 0.1, Kd: 0.05,
			},
		},
	}

	tempDir := t.TempDir()
	err := writeTestFiles(tempDir, template, pidData)
	if err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	templatePath := filepath.Join(tempDir, "template.xml")
	upload.LoadXMLTemplate(templatePath)

	err = upload.ProcessFiles(tempDir)
	if err != nil {
		t.Fatalf("process failed: %v", err)
	}

	resultPath := filepath.Join(tempDir, "test.xml")
	data, err := os.ReadFile(resultPath)
	if err != nil {
		t.Fatalf("output read failed: %v", err)
	}

	out := string(data)
	if !strings.Contains(out, ">1.5<") || !strings.Contains(out, ">0.1<") || !strings.Contains(out, ">0.05<") {
		t.Errorf("output XML does not contain updated values: %s", out)
	}
}

func TestUpdateXMLWithMissingPID(t *testing.T) {
	template := `
<fsigmodule name="NotInData" type="pid">
  <param name="Kp">0</param>
</fsigmodule>`

	pidData := &upload.PIDData{
		Data: []upload.PIDConfig{
			{
				Name: "OtherPID",
				Kp:   9.9,
			},
		},
	}

	updated := upload.UpdateXMLWithPIDData(template, pidData)

	if strings.Contains(updated, ">9.9<") {
		t.Errorf("unexpected replacement in unrelated module")
	}
}

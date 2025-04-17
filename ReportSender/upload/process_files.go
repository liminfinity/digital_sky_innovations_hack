package upload

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type PIDConfig struct {
	Name                string  `json:"name"`
	Kp                  float64 `json:"Kp"`
	Ki                  float64 `json:"Ki"`
	Kd                  float64 `json:"Kd"`
	IntegralMin         float64 `json:"integral_min"`
	IntegralMax         float64 `json:"integral_max"`
	InpRiseDeriative    int     `json:"inp_rise_deriative"`
	InpFallDeriative    int     `json:"inp_fall_deriative"`
	Min                 float64 `json:"min"`
	Max                 float64 `json:"max"`
	PresetAllowedAtLow  float64 `json:"preset_allowed_at_low"`
	PresetAllowedAtHigh float64 `json:"preset_allowed_at_high"`
}

type PIDData struct {
	Data []PIDConfig `json:"data"`
}

var XMLTemplate string

func LoadXMLTemplate(templatePath string) error {
	data, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("ошибка чтения XML файла: %w", err)
	}

	XMLTemplate = string(data)
	return nil
}

func ProcessFiles(dirPath string) error {
	if XMLTemplate == "" {
		return fmt.Errorf("XML шаблон не загружен, вызовите LoadXMLTemplate перед ProcessFiles")
	}

	direntry, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, dirent := range direntry {
		if dirent.IsDir() {
			continue
		}

		if !strings.HasSuffix(strings.ToLower(dirent.Name()), ".json") {
			continue
		}

		path := filepath.Join(dirPath, dirent.Name())
		err = processFile(path)
		if err != nil {
			log.Printf("Ошибка обработки файла %s: %v", path, err)
			continue
		}
	}

	return nil
}

func processFile(jsonPath string) error {

	jsonData, err := os.ReadFile(jsonPath)
	if err != nil {
		return fmt.Errorf("ошибка чтения JSON файла: %w", err)
	}

	var pidData PIDData
	err = json.Unmarshal(jsonData, &pidData)
	if err != nil {
		return fmt.Errorf("ошибка разбора JSON: %w", err)
	}

	xmlContent := XMLTemplate

	xmlContent = UpdateXMLWithPIDData(xmlContent, &pidData)

	xmlOutputPath := strings.TrimSuffix(jsonPath, ".json") + ".xml"
	err = os.WriteFile(xmlOutputPath, []byte(xmlContent), 0644)
	if err != nil {
		return fmt.Errorf("ошибка сохранения XML в %s: %w", xmlOutputPath, err)
	}

	err = os.Remove(jsonPath)
	if err != nil {
		return fmt.Errorf("ошибка удаления JSON-файла %s: %w", jsonPath, err)
	}

	log.Printf("Успешно создан %s и удален %s", xmlOutputPath, jsonPath)
	return nil
}

func UpdateXMLWithPIDData(xmlContent string, pidData *PIDData) string {

	pidMap := make(map[string]map[string]string)

	for _, pid := range pidData.Data {
		params := map[string]string{
			"Kp":                     fmt.Sprintf("%g", pid.Kp),
			"Ki":                     fmt.Sprintf("%g", pid.Ki),
			"Kd":                     fmt.Sprintf("%g", pid.Kd),
			"integral_min":           fmt.Sprintf("%g", pid.IntegralMin),
			"integral_max":           fmt.Sprintf("%g", pid.IntegralMax),
			"inp_rise_deriative":     fmt.Sprintf("%d", pid.InpRiseDeriative),
			"inp_fall_deriative":     fmt.Sprintf("%d", pid.InpFallDeriative),
			"min":                    fmt.Sprintf("%g", pid.Min),
			"max":                    fmt.Sprintf("%g", pid.Max),
			"preset_allowed_at_low":  fmt.Sprintf("%g", pid.PresetAllowedAtLow),
			"preset_allowed_at_high": fmt.Sprintf("%g", pid.PresetAllowedAtHigh),
		}
		pidMap[pid.Name] = params
	}

	moduleRegex := regexp.MustCompile(`<fsigmodule\s+name="([^"]+)"\s+type="pid"[^>]*>([\s\S]*?)</fsigmodule>`)

	paramRegex := regexp.MustCompile(`<param\s+name="([^"]+)">([\s\S]*?)</param>`)

	updatedXML := moduleRegex.ReplaceAllStringFunc(xmlContent, func(moduleBlock string) string {

		moduleMatches := moduleRegex.FindStringSubmatch(moduleBlock)
		if len(moduleMatches) < 2 {
			return moduleBlock
		}

		moduleName := moduleMatches[1]

		pidParams, exists := pidMap[moduleName]
		if !exists {
			return moduleBlock
		}

		updatedBlock := paramRegex.ReplaceAllStringFunc(moduleBlock, func(paramBlock string) string {
			paramMatches := paramRegex.FindStringSubmatch(paramBlock)
			if len(paramMatches) < 3 {
				return paramBlock
			}

			paramName := paramMatches[1]

			if newValue, hasParam := pidParams[paramName]; hasParam {
				return fmt.Sprintf(`<param name="%s">%s</param>`, paramName, newValue)
			}

			return paramBlock
		})

		return updatedBlock
	})

	return updatedXML
}


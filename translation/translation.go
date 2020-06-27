package translation

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Translation bundle struct
type TBundle struct {
	bundle *i18n.Bundle
}

// translation key and properties struct for translate function
type TParam struct {
	Key          string                 `json:",omitempty"`
	TemplateData map[string]interface{} `json:",omitempty"`
	PluralCount  interface{}            `json:",omitempty"`
}

var tBundle TBundle

// Initialize translation bundle
func NewTranslationBundle() error {
	langFiles := []string{}
	var err error
	tBundle = TBundle{}

	root := "./translation/locales"
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// fmt.Println(info)
		if !info.IsDir() {
			langFiles = append(langFiles, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	// Create a Bundle to use for the lifetime of your application
	tBundle.bundle, err = createLocalizerBundle(langFiles)
	if err != nil {
		return err
	}

	return nil
}

// CreateLocalizerBundle reads language files and registers them in i18n bundle
func createLocalizerBundle(langFiles []string) (*i18n.Bundle, error) {
	// Bundle stores a set of messages
	bundle := i18n.NewBundle(language.English)

	// Enable bundle to understand yaml
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	var translations []byte
	var err error
	for _, file := range langFiles {

		// Read our language yaml file
		translations, err = ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}

		// fmt.Println(translations)

		// It parses the bytes in buffer to add translations to the bundle
		bundle.MustParseMessageFileBytes(translations, file)
	}

	return bundle, nil
}

func Translate(param TParam, locale string) (string, error) {
	localizer := i18n.NewLocalizer(tBundle.bundle, locale)
	msg, err := localizer.Localize(
		&i18n.LocalizeConfig{
			MessageID:    param.Key,
			TemplateData: param.TemplateData,
			PluralCount:  param.PluralCount,
		},
	)
	return msg, err
}

func Walk() {
	var files []string

	root := "./translation/locales"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// fmt.Println(info)
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}

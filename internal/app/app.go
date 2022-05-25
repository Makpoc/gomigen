package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Makpoc/gomigen/internal/generator"
	"github.com/Makpoc/gomigen/internal/parser"
	"github.com/Makpoc/gomigen/internal/version"
)

var (
	// capitalLetterRegex captures each capital letter in a match group.
	capitalLetterRegex = regexp.MustCompile("([A-Z])")
)

type Logger interface {
	Printf(format string, v ...interface{})
}

type App struct {
	pkgDir          string
	interfaceName   string
	outputDirectory string
	version         string
	logger          Logger
}

func New(pkgDir, interfaceName string, opts ...Option) *App {
	app := &App{
		pkgDir:        pkgDir,
		interfaceName: interfaceName,
		logger:        log.Default(),
		version:       version.Version(),
	}
	for _, o := range opts {
		o(app)
	}
	return app
}

func (a *App) Run() error {
	middleware, err := parser.Parse(a.pkgDir, a.interfaceName)
	if err != nil {
		return fmt.Errorf("failed to parse interface %q: %w",
			a.interfaceName, err)
	}

	source, err := generator.GenerateSource(middleware, a.version)
	if err != nil {
		return fmt.Errorf("failed to generate middleware for interface %q: %w",
			a.interfaceName, err)
	}

	outputDir := filepath.Join(a.pkgDir, middleware.GenPackage)
	if a.outputDirectory != "" {
		outputDir = filepath.Join(a.outputDirectory, middleware.GenPackage)
	}

	outputFilePath := filepath.Join(outputDir, snakeCase(middleware.InterfaceName)+".go")

	err = save(source, outputFilePath)
	if err != nil {
		return fmt.Errorf("failed to save generated middleware to %q: %w", outputFilePath, err)
	}

	interfaceWithPkgPath := fmt.Sprintf("%s.%s",
		middleware.OriginalPackage.Path, middleware.InterfaceName)
	a.logger.Printf("Generated middleware for interface %q in %q",
		interfaceWithPkgPath, outputFilePath)

	return err
}

func snakeCase(s string) string {
	return strings.ToLower(strings.TrimPrefix(capitalLetterRegex.ReplaceAllString(s, "_${1}"), "_"))
}

func save(source []byte, filePath string) error {
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to create middlewares folder: %w", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create middleware file: %w", err)
	}
	defer func() { _ = file.Close() }()

	_, err = file.Write(source)
	if err != nil {
		return fmt.Errorf("failed to write to middleware file: %w", err)
	}
	return nil
}

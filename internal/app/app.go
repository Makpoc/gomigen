package app

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Makpoc/gomigen/internal/generator"
	"github.com/Makpoc/gomigen/internal/parser"
	"github.com/Makpoc/gomigen/internal/version"
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
	a.logger.Printf("Looking for interface %q in %q", a.interfaceName, a.pkgDir)
	middleware, err := parser.Parse(a.pkgDir, a.interfaceName)
	if err != nil {
		return fmt.Errorf("failed to parse interface: %w", err)
	}

	interfaceWithPkgPath := fmt.Sprintf("%s.%s", middleware.OriginalPackage.Path, middleware.InterfaceName)

	a.logger.Printf("Generating middleware for interface %q", interfaceWithPkgPath)
	source, err := generator.GenerateSource(middleware, a.version)
	if err != nil {
		return fmt.Errorf("failed to generate middleware: %w", err)
	}

	outputDir := path.Join(a.pkgDir, middleware.GenPackage)
	if a.outputDirectory != "" {
		outputDir = path.Join(a.outputDirectory, middleware.GenPackage)
	}

	outputFilePath := path.Join(outputDir, snakeCase(middleware.InterfaceName)+".go")

	a.logger.Printf("Saving middleware for interface %q to %q", interfaceWithPkgPath, outputFilePath)
	err = save(source, outputFilePath)
	if err != nil {
		return fmt.Errorf("failed to save generated middleware to %q: %w", outputFilePath, err)
	}
	return err
}

func snakeCase(s string) string {
	re := regexp.MustCompile("([A-Z])")
	return strings.ToLower(strings.TrimPrefix(re.ReplaceAllString(s, "_${1}"), "_"))
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

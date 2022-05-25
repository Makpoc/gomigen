package integration_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Makpoc/gomigen/internal/app"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	testDataRelPath  = "./testdata"
	generatedDirName = "_generated"
)

var _ = Describe("Integration", func() {
	Describe("table tests", func() {
		const (
			tableTestDataRelPath = testDataRelPath + "/" + "tabletests"
			pkgName              = "interfaces"

			outDirRelPath       = tableTestDataRelPath + "/" + generatedDirName
			generatedDirRelPath = outDirRelPath + "/" + pkgName + "mw"

			interfacesDirRelPath = tableTestDataRelPath + "/" + "interfaces"
			goldFilesDirRelPath  = tableTestDataRelPath + "/" + "goldfiles"
		)
		BeforeEach(func() {
			// Comment out this line to keep all generated files after execution
			// (e.g. if you want to use them as gold-files *after careful review*
			Expect(cleanupGeneratedFiles(generatedDirRelPath)).To(Succeed())
		})

		DescribeTable("",
			func(interfaceName string) {
				a := app.New(interfacesDirRelPath, interfaceName,
					app.WithVersion("test"), app.WithOutputDirectory(outDirRelPath),
					app.WithLogger(log.New(GinkgoWriter, "", log.LstdFlags)),
				)

				err := a.Run()
				Expect(err).ToNot(HaveOccurred())

				fileName := interfaceName + ".go"
				genFilePath := filepath.Join(generatedDirRelPath, snakeCase(fileName))
				goldFilePath := filepath.Join(goldFilesDirRelPath, snakeCase(fileName))

				compareFileContent(genFilePath, goldFilePath)
			},
			Entry(nil, "SingleMethodNoArgsNoReturns"),
			Entry(nil, "SingleMethodWithArgNoReturns"),
			Entry(nil, "SingleMethodWithArgAndReturn"),
			Entry(nil, "ContextFirstArgument"),
			Entry(nil, "ReturnsSingleError"),
			Entry(nil, "ReturnsMultipleValuesWithError"),
			Entry(nil, "ReturnsMultipleValuesNoError"),
			Entry(nil, "ContextFirstArgumentReturnsError"),
			Entry(nil, "ContextFirstArgumentReturnsErrorNamedVars"),
			Entry(nil, "TwoMethodsOneWithContextAndError"),
			Entry(nil, "MapInArgAndReturn"),
			Entry(nil, "SliceInArgAndReturn"),
			Entry(nil, "VariadicArgument"),
			Entry(nil, "EmbeddedInterface"),
			Entry(nil, "CustomInterfaceInArgAndReturn"),
			Entry(nil, "EmptyInterface"),
			Entry(nil, "InnerPackageReference"),
			Entry(nil, "EmbedStandardInterface"),
			Entry(nil, "EmbedCustomInterface"),
			Entry(nil, "ComposeMultipleInterface"),
			Entry(nil, "ExtendsAnotherInterface"),
			Entry(nil, "ComposeOverlappingInterfaces"),
		)
	})
	Describe("duplicate packages", func() {
		const (
			pkgName        = "packagecollision"
			testDirRelpath = testDataRelPath + "/" + pkgName

			outDirRelPath       = testDirRelpath + "/" + generatedDirName
			generatedDirRelPath = outDirRelPath + "/" + pkgName + "mw"

			goldfilesDirRelPath = testDirRelpath + "/" + "goldfiles"
		)
		BeforeEach(func() {
			// Comment out this line to keep all generated files after execution
			// (e.g. if you want to use them as gold-files *after careful review*
			Expect(cleanupGeneratedFiles(generatedDirRelPath)).To(Succeed())
		})
		It("aliases the duplicated package imports", func() {
			interfaceName := "Compare"
			a := app.New(testDirRelpath, interfaceName,
				app.WithVersion("test"),
				app.WithOutputDirectory(outDirRelPath),
				app.WithLogger(log.New(GinkgoWriter, "", log.LstdFlags)),
			)

			err := a.Run()
			Expect(err).ToNot(HaveOccurred())

			fileName := interfaceName + ".go"
			genFilePath := filepath.Join(generatedDirRelPath, snakeCase(fileName))
			goldFilePath := filepath.Join(goldfilesDirRelPath, snakeCase(fileName))

			compareFileContent(genFilePath, goldFilePath)
		})
	})
})

func compareFileContent(genFilePath, goldFilePath string) {
	genFileContent, err := os.ReadFile(genFilePath)
	Expect(err).ToNot(HaveOccurred())
	goldFileContent, err := os.ReadFile(goldFilePath)
	Expect(err).ToNot(HaveOccurred())

	Expect(string(genFileContent)).To(Equal(string(goldFileContent)))
}

func cleanupGeneratedFiles(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to list files in %q: %w", dir, err)
	}
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".go") {
			fpath := filepath.Join(dir, f.Name())
			err := os.Remove(fpath)
			if err != nil {
				return fmt.Errorf("failed to delete file %q: %w", fpath, err)
			}
		}
	}
	return nil
}

var wordRE = regexp.MustCompile("([A-Z])")

func snakeCase(s string) string {
	return strings.ToLower(strings.TrimPrefix(wordRE.ReplaceAllString(s, "_${1}"), "_"))
}

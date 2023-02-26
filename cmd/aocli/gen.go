package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

var mainTmpl string = `package main

import (
	"context"

	"mvinkio.online/aoc/{{ .Year }}/{{ .Day }}/{{ .Name }}"
	"mvinkio.online/aoc/aoc"
)

// Boilerplate

func main() {
	aoc.RunDay(
		context.TODO(),
		aoc.NewScanCloser("{{ .Year }}/{{ .Day }}/input.txt"),
		aoc.ReadLines({{ .Name }}.Reader),
        {{ .Name }}.Solvers,
	)
}
`

var packageTmpl string = `package {{ .Name }}

import (
    "context"
)

var (
	Reader  = read{{ .Name | Title }}
	Solvers = []func(ctx context.Context, data any) (any, error){
		partOne, partTwo,
	}
)

func partOne(ctx context.Context, data any) (any, error) {
	return nil, nil
}

func partTwo(ctx context.Context, data any) (any, error) {
	return nil, nil
}

func read{{ .Name | Title }}(line []string) (any, error) {
    return nil, nil
}
`

var testTmpl string = `package {{ .Name }}_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("{{ .Name }}", func() {
    Describe("Part One", func() {
        It("should solve the example to the answer", func() {
			Expect(true).To(Equal(true))
        })
    })
    Describe("Part Two", func() {
        It("should solve the example to the answer", func() {
			Expect(true).To(Equal(true))
        })
    })
})
`

type templateData struct {
	Year string
	Day  string
	Name string
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a new boilerplate entrypoint for a day of Advent of Code",
	RunE: func(cmd *cobra.Command, args []string) error {
		year, day, packageName := "", "", ""
		if len(args) == 2 {
			parts := strings.SplitN(args[0], "/", 2)
			if len(parts) == 2 {
				year = parts[0]
				day = parts[1]
			}
		} else {
			return fmt.Errorf("Expected two argument in format %q %q", "<year>/<day>", "<packageName>")
		}
		packageName = fmt.Sprintf("%s", args[1])

		if year == "" || day == "" {
			return fmt.Errorf("Expected one argument in format %q", "<year>/<day>")
		}

		_, err := strconv.Atoi(year)
		if err != nil {
			return fmt.Errorf("Couldn't parse year integer: %q", year)
		}

		dayInt, err := strconv.Atoi(day)
		if err != nil {
			return fmt.Errorf("Couldn't parse day integer: %q", day)
		}

		if dayInt < 1 || dayInt > 25 {
			return fmt.Errorf("Day must be between 1 and 25, got %d", dayInt)
		}

		dayStr := fmt.Sprintf("%02d", dayInt)

		data := templateData{
			Year: year,
			Day:  dayStr,
			Name: packageName,
		}
		tmplMain, err := template.New("main").Funcs(template.FuncMap{
			"Title": strings.Title,
		}).Parse(mainTmpl)
		if err != nil {
			return fmt.Errorf("Couldn't parse main template: %w", err)
		}
		tmplPackage, err := template.New("package").Funcs(template.FuncMap{
			"Title": strings.Title,
		}).Parse(packageTmpl)
		if err != nil {
			return fmt.Errorf("Couldn't parse package template: %w", err)
		}
		tmplTest, err := template.New("test").Funcs(template.FuncMap{
			"Title": strings.Title,
		}).Parse(testTmpl)
		if err != nil {
			return fmt.Errorf("Couldn't parse test template: %w", err)
		}

		if err := ensureDirs(filepath.Join(year, dayStr, packageName)); err != nil {
			return fmt.Errorf("Couldn't create directory %s/%s: %v", year, dayStr, err)
		}

		fMain, err := getFile(filepath.Join(year, dayStr, "main.go"))
		if err != nil {
			return fmt.Errorf("Couldn't open file %s/%s/main.go: %v", year, dayStr, err)
		}
		defer fMain.Close()

		if err := tmplMain.Execute(fMain, data); err != nil {
			return fmt.Errorf("Couldn't execute main template: %w", err)
		}

		fPackage, err := getFile(filepath.Join(year, dayStr, packageName, data.Name+".go"))
		if err != nil {
			return fmt.Errorf("Couldn't open file %s/%s/%s/%s.go: %v", year, dayStr, packageName, data.Name, err)
		}
		defer fPackage.Close()

		if err := tmplPackage.Execute(fPackage, data); err != nil {
			return fmt.Errorf("Couldn't execute package template: %w", err)
		}

		fTest, err := getFile(filepath.Join(year, dayStr, packageName, data.Name+"_test.go"))
		if err != nil {
			return fmt.Errorf("Couldn't open file %s/%s/%s/%s_test.go: %v", year, dayStr, packageName, data.Name, err)
		}
		defer fPackage.Close()

		if err := tmplTest.Execute(fTest, data); err != nil {
			return fmt.Errorf("Couldn't execute test template: %w", err)
		}
		return nil
	},
}

func ensureDirs(relativePath string) error {
	_, err := os.Stat(relativePath)
	if os.IsNotExist(err) {
		return os.MkdirAll(relativePath, fs.ModePerm)
	}
	return nil
}

func getFile(relativePath string) (*os.File, error) {
	return os.Create(relativePath)
}

package venom

import (
	"encoding/xml"
	"strings"
	"unicode"
)

type H map[string]interface{}

func (h H) Clone() H {
	var h2 = make(H, len(h))
	h2.AddAll(h)
	return h2
}

func (h *H) Add(k string, v interface{}) {
	if h == nil || *h == nil {
		*h = make(map[string]interface{})
	}
	(*h)[k] = v
}

func (h *H) AddWithPrefix(p, k string, v interface{}) {
	(*h)[p+"."+k] = v
}

func (h *H) AddAll(h2 H) {
	for k, v := range h2 {
		h.Add(k, v)
	}
}

func (h *H) AddAllWithPrefix(p string, h2 H) {
	if h2 == nil {
		return
	}
	if h == nil {
		var _h = H{}
		*h = _h
	}
	for k, v := range h2 {
		h.AddWithPrefix(p, k, v)
	}
}

// Assertion (usually a string, but could be a dictionary when using logical operators)
type Assertion interface{}

// StepAssertions contains step assertions
type StepAssertions struct {
	Assertions []Assertion `json:"assertions,omitempty" yaml:"assertions,omitempty"`
}

// Tests contains all informations about tests in a pipeline build
type Tests struct {
	XMLName      xml.Name    `xml:"testsuites" json:"-" yaml:"-"`
	Total        int         `xml:"-" json:"total"`
	TotalOK      int         `xml:"-" json:"ok"`      // contains the number of testcases OK
	TotalKO      int         `xml:"-" json:"ko"`      // contains the number of testcases KO
	TotalSkipped int         `xml:"-" json:"skipped"` // contains the number of testcases skipped
	TestSuites   []TestSuite `xml:"testsuite" json:"test_suites"`
}

// TestSuite is a single JUnit test suite which may contain many
// testcases.
type TestSuite struct {
	XMLName      xml.Name   `xml:"testsuite" json:"-" yaml:"-"`
	Errors       int        `xml:"errors,attr,omitempty" json:"errors" yaml:"-"`
	Failures     int        `xml:"failures,attr,omitempty" json:"failures" yaml:"-"`
	Hostname     string     `xml:"hostname,attr,omitempty" json:"hostname" yaml:"-"`
	ID           string     `xml:"id,attr,omitempty" json:"id" yaml:"-"`
	Name         string     `xml:"name,attr" json:"name" yaml:"name"`
	Filename     string     `xml:"-" json:"-" yaml:"-"`
	Package      string     `xml:"package,attr,omitempty" json:"package" yaml:"-"`
	Properties   []Property `xml:"-" json:"properties" yaml:"-"`
	Skipped      int        `xml:"skipped,attr,omitempty" json:"skipped" yaml:"skipped,omitempty"`
	Total        int        `xml:"tests,attr" json:"total" yaml:"total,omitempty"`
	TestCases    []TestCase `xml:"testcase" json:"testcases" yaml:"testcases"`
	Version      string     `xml:"version,omitempty" json:"version" yaml:"version,omitempty"`
	Time         string     `xml:"time,attr,omitempty" json:"time" yaml:"-"`
	Timestamp    string     `xml:"timestamp,attr,omitempty" json:"timestamp" yaml:"-"`
	Vars         H          `xml:"-" json:"-" yaml:"vars"`
	ComputedVars H          `xml:"-" json:"-" yaml:"-"`
	WorkDir      string     `xml:"-" json:"-" yaml:"-"`
}

// Property represents a key/value pair used to define properties.
type Property struct {
	XMLName xml.Name `xml:"property" json:"-" yaml:"-"`
	Name    string   `xml:"name,attr" json:"name" yaml:"-"`
	Value   string   `xml:"value,attr" json:"value" yaml:"-"`
}

// TestCase is a single test case with its result.
type TestCase struct {
	XMLName         xml.Name  `xml:"testcase" json:"-" yaml:"-"`
	Classname       string    `xml:"classname,attr,omitempty" json:"classname" yaml:"-"`
	Errors          []Failure `xml:"error,omitempty" json:"errors" yaml:"errors,omitempty"`
	Failures        []Failure `xml:"failure,omitempty" json:"failures" yaml:"failures,omitempty"`
	Name            string    `xml:"name,attr" json:"name" yaml:"name"`
	originalName    string
	Skipped         []Skipped   `xml:"skipped,omitempty" json:"skipped" yaml:"skipped,omitempty"`
	Status          string      `xml:"status,attr,omitempty" json:"status" yaml:"status,omitempty"`
	Systemout       InnerResult `xml:"system-out,omitempty" json:"systemout" yaml:"systemout,omitempty"`
	Systemerr       InnerResult `xml:"system-err,omitempty" json:"systemerr" yaml:"systemerr,omitempty"`
	Time            float64     `xml:"time,attr,omitempty" json:"time" yaml:"time,omitempty"`
	TestSteps       []TestStep  `xml:"-" json:"steps" yaml:"steps"`
	TestSuiteVars   H           `xml:"-" json:"-" yaml:"-"`
	Vars            H           `xml:"-" json:"-" yaml:"vars"`
	computedVars    H
	computedInfo    []string
	computedVerbose []string
	Skip            []string `xml:"-" json:"skip" yaml:"skip"`
	IsExecutor      bool     `xml:"-" json:"-" yaml:"-"`
	IsEvaluated     bool     `xml:"-" json:"-" yaml:"-"`
}

// TestStep represents a testStep
type TestStep map[string]interface{}

// Range contains data related to iterable user values
type Range struct {
	Enabled    bool
	Items      []RangeData
	RawContent interface{} `json:"range"`
}

// RangeData contains a single iterable user value
type RangeData struct {
	Key   string
	Value interface{}
}

// Skipped contains data related to a skipped test.
type Skipped struct {
	Value string `xml:",cdata" json:"value" yaml:"value,omitempty"`
}

func (tc *TestCase) AppendError(err error) {
	tc.Errors = append(tc.Errors, Failure{Value: RemoveNotPrintableChar(err.Error())})
}

// Failure contains data related to a failed test.
type Failure struct {
	TestcaseClassname  string `xml:"-" json:"-" yaml:"-"`
	TestcaseName       string `xml:"-" json:"-" yaml:"-"`
	TestcaseLineNumber int    `xml:"-" json:"-" yaml:"-"`
	StepNumber         int    `xml:"-" json:"-" yaml:"-"`
	Assertion          string `xml:"-" json:"-" yaml:"-"`
	AssertionRequired  bool   `xml:"-" json:"-" yaml:"-"`
	Error              error  `xml:"-" json:"-" yaml:"-"`

	Value   string `xml:",cdata" json:"value" yaml:"value,omitempty"`
	Type    string `xml:"type,attr,omitempty" json:"type" yaml:"type,omitempty"`
	Message string `xml:"message,attr,omitempty" json:"message" yaml:"message,omitempty"`
}

// InnerResult is used by TestCase
type InnerResult struct {
	Value string `xml:",cdata" json:"value" yaml:"value"`
}

type AssignStep struct {
	Assignments map[string]Assignment `json:"vars" yaml:"vars" mapstructure:"vars"`
}

type Assignment struct {
	From  string `json:"from" yaml:"from"`
	Regex string `json:"regex" yaml:"regex"`
}

// RemoveNotPrintableChar removes not printable chararacter from a string
func RemoveNotPrintableChar(in string) string {
	m := func(r rune) rune {
		if unicode.IsPrint(r) || unicode.IsSpace(r) || unicode.IsPunct(r) {
			return r
		}
		return ' '
	}
	return strings.Map(m, in)
}

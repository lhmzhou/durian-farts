package durianfarts_test

import (
	"strings"
	"testing"
)

type TestCase struct {
	input          string
	snake          string
	screamingSnake string
	kebab          string
	screamingKebab string
	camel          string
	lowerCamel     string
	custom         string
}

var TestCases = []TestCase{
	{
		input:          `{"firstName": "Higgs", "emails": ["Higgs.Boson@email.com"]}`,
		snake:          `{"emails":["Higgs.Boson@email.com"],"first_name":"Higgs"}`,
		screamingSnake: `{"EMAILS":["Higgs.Boson@email.com"],"FIRST_NAME":"Higgs"}`,
		kebab:          `{"emails":["Higgs.Boson@email.com"],"first-name":"Higgs"}`,
		screamingKebab: `{"EMAILS":["Higgs.Boson@email.com"],"FIRST-NAME":"Higgs"}`,
		camel:          `{"Emails":["Higgs.Boson@email.com"],"FirstName":"Higgs"}`,
		lowerCamel:     `{"emails":["Higgs.Boson@email.com"],"firstName":"Higgs"}`,
		custom:         `{"em41ls":["Higgs.Boson@email.com"],"f1rstN4me":"Higgs"}`,
	},
	{
		input:          `{"glossary":{"title":"example glossary","GlossDiv":{"title":"S","GlossList":{"GlossEntry":{"ID":"SGML","SortAs":"SGML","GlossTerm":"Standard Generalized Markup Language","Acronym":"SGML","Abbrev":"ISO 8879:1986","GlossDef":{"para":"A meta-markup language, used to create markup languages such as DocBook.","GlossSeeAlso":["GML","XML"]},"GlossSee":"markup"}}}}}`,
		snake:          `{"glossary":{"gloss_div":{"gloss_list":{"gloss_entry":{"abbrev":"ISO 8879:1986","acronym":"SGML","gloss_def":{"gloss_see_also":["GML","XML"],"para":"A meta-markup language, used to create markup languages such as DocBook."},"gloss_see":"markup","gloss_term":"Standard Generalized Markup Language","id":"SGML","sort_as":"SGML"}},"title":"S"},"title":"example glossary"}}`,
		screamingSnake: `{"GLOSSARY":{"GLOSS_DIV":{"GLOSS_LIST":{"GLOSS_ENTRY":{"ABBREV":"ISO 8879:1986","ACRONYM":"SGML","GLOSS_DEF":{"GLOSS_SEE_ALSO":["GML","XML"],"PARA":"A meta-markup language, used to create markup languages such as DocBook."},"GLOSS_SEE":"markup","GLOSS_TERM":"Standard Generalized Markup Language","ID":"SGML","SORT_AS":"SGML"}},"TITLE":"S"},"TITLE":"example glossary"}}`,
		kebab:          `{"glossary":{"gloss-div":{"gloss-list":{"gloss-entry":{"abbrev":"ISO 8879:1986","acronym":"SGML","gloss-def":{"gloss-see-also":["GML","XML"],"para":"A meta-markup language, used to create markup languages such as DocBook."},"gloss-see":"markup","gloss-term":"Standard Generalized Markup Language","id":"SGML","sort-as":"SGML"}},"title":"S"},"title":"example glossary"}}`,
		screamingKebab: `{"GLOSSARY":{"GLOSS-DIV":{"GLOSS-LIST":{"GLOSS-ENTRY":{"ABBREV":"ISO 8879:1986","ACRONYM":"SGML","GLOSS-DEF":{"GLOSS-SEE-ALSO":["GML","XML"],"PARA":"A meta-markup language, used to create markup languages such as DocBook."},"GLOSS-SEE":"markup","GLOSS-TERM":"Standard Generalized Markup Language","ID":"SGML","SORT-AS":"SGML"}},"TITLE":"S"},"TITLE":"example glossary"}}`,
		camel:          `{"Glossary":{"GlossDiv":{"GlossList":{"GlossEntry":{"Abbrev":"ISO 8879:1986","Acronym":"SGML","GlossDef":{"GlossSeeAlso":["GML","XML"],"Para":"A meta-markup language, used to create markup languages such as DocBook."},"GlossSee":"markup","GlossTerm":"Standard Generalized Markup Language","ID":"SGML","SortAs":"SGML"}},"Title":"S"},"Title":"example glossary"}}`,
		lowerCamel:     `{"glossary":{"glossDiv":{"glossList":{"glossEntry":{"abbrev":"ISO 8879:1986","acronym":"SGML","glossDef":{"glossSeeAlso":["GML","XML"],"para":"A meta-markup language, used to create markup languages such as DocBook."},"glossSee":"markup","glossTerm":"Standard Generalized Markup Language","iD":"SGML","sortAs":"SGML"}},"title":"S"},"title":"example glossary"}}`,
		custom:         `{"gloss4ry":{"GlossD1v":{"GlossL1st":{"GlossEntry":{"Abbrev":"ISO 8879:1986","Acronym":"SGML","GlossDef":{"GlossSeeAlso":["GML","XML"],"p4r4":"A meta-markup language, used to create markup languages such as DocBook."},"GlossSee":"markup","GlossTerm":"Standard Generalized Markup Language","ID":"SGML","SortAs":"SGML"}},"t1tle":"S"},"t1tle":"example glossary"}}`,
	},
	{
		input:          `{"foo": [ { "fizz": "buzz" }], "bar": [ ["uno", "dos"] ]}`,
		snake:          `{"bar":[["uno","dos"]],"foo":[{"fizz":"buzz"}]}`,
		screamingSnake: `{"BAR":[["uno","dos"]],"FOO":[{"FIZZ":"buzz"}]}`,
		kebab:          `{"bar":[["uno","dos"]],"foo":[{"fizz":"buzz"}]}`,
		screamingKebab: `{"BAR":[["uno","dos"]],"FOO":[{"FIZZ":"buzz"}]}`,
		camel:          `{"Bar":[["uno","dos"]],"Foo":[{"Fizz":"buzz"}]}`,
		lowerCamel:     `{"bar":[["uno","dos"]],"foo":[{"fizz":"buzz"}]}`,
		custom:         `{"b4r":[["uno","dos"]],"foo":[{"f1zz":"buzz"}]}`,
	},
	{
		input:          `{"firstName":"Higgs","lastName":"Boson","address":{"streetAddress":"Esplanade des Particules 1","city":"Meyrin","canton":"Geneva","countryCode":"Switzerland"},"lineOfBusiness":[{"reasonCode":"SR-042","description":"Standard Model of contemporary particle physics"}],"futureUse":[{"useCode":"QW-011","description":" Validation of LHC, etc."}]}`,
		snake:          `{"address":{"city":"Meyrin","country_code":"Switzerland","canton":"Geneva","street_address":"Esplanade des Particules 1"},"first_name":"Higgs","future_use":[{"description":" Validation of LHC, etc.","use_code":"QW-011"}],"last_name":"Boson","line_of_business":[{"description":"Standard Model of contemporary particle physics","reason_code":"SR-042"}]}`,
		screamingSnake: `{"ADDRESS":{"CITY":"Meyrin","COUNTRY_CODE":"Switzerland","CANTON":"Geneva","STREET_ADDRESS":"Esplanade des Particules 1"},"FIRST_NAME":"Higgs","FUTURE_USE":[{"DESCRIPTION":" Validation of LHC, etc.","USE_CODE":"QW-011"}],"LAST_NAME":"Boson","LINE_OF_BUSINESS":[{"DESCRIPTION":"Standard Model of contemporary particle physics","REASON_CODE":"SR-042"}]}`,
		kebab:          `{"address":{"city":"Meyrin","country-code":"Switzerland","canton":"Geneva","street-address":"Esplanade des Particules 1"},"first-name":"Higgs","future-use":[{"description":" Validation of LHC, etc.","use-code":"QW-011"}],"last-name":"Boson","line-of-business":[{"description":"Standard Model of contemporary particle physics","reason-code":"SR-042"}]}`,
		screamingKebab: `{"ADDRESS":{"CITY":"Meyrin","COUNTRY-CODE":"Switzerland","CANTON":"Geneva","STREET-ADDRESS":"Esplanade des Particules 1"},"FIRST-NAME":"Higgs","FUTURE-USE":[{"DESCRIPTION":" Validation of LHC, etc.","USE-CODE":"QW-011"}],"LAST-NAME":"Boson","LINE-OF-BUSINESS":[{"DESCRIPTION":"Standard Model of contemporary particle physics","REASON-CODE":"SR-042"}]}`,
		camel:          `{"Address":{"City":"Meyrin","CountryCode":"Switzerland","canton":"Geneva","StreetAddress":"Esplanade des Particules 1"},"FirstName":"Higgs","FutureUse":[{"Description":" Validation of LHC, etc.","UseCode":"QW-011"}],"LastName":"Boson","LineOfBusiness":[{"Description":"Standard Model of contemporary particle physics","ReasonCode":"SR-042"}]}`,
		lowerCamel:     `{"address":{"city":"Meyrin","countryCode":"Switzerland","canton":"Geneva","streetAddress":"Esplanade des Particules 1"},"firstName":"Higgs","futureUse":[{"description":" Validation of LHC, etc.","useCode":"QW-011"}],"lastName":"Boson","LineOfBusiness":[{"description":"Standard Model of contemporary particle physics","reasonCode":"SR-042"}]}`,
		custom:         `{"4ddress":{"c1ty":"Meyrin","countryCode":"Switzerland","prov1nce":"Geneva","streetAddress":"Esplanade des Particules 1"},"f1rstN4me":"Higgs","futureUse":[{"descr1pt1on":" Validation of LHC, etc.","useCode":"QW-011"}],"l4stN4me":"Boson","l1neOfBus1ness":[{"descr1pt1on":"Standard Model of contemporary particle physics","re4sonCode":"SR-042"}]}`,
	},
}

func mismatch(t *testing.T, in string, out string) {
	t.Errorf("JSON mismatch: want %v, got %v", in, out)
}

func TestToSnake(t *testing.T) {
	for _, tc := range TestCases {
		out, err := durianfarts.ToSnake(tc.input)
		if nil != err {
			t.Error(err)
		}
		if out != tc.snake {
			mismatch(t, tc.snake, out)
		}
	}
}

func TestToScreamingSnake(t *testing.T) {
	for _, tc := range TestCases {
		out, err := durianfarts.ToScreamingSnake(tc.input)
		if nil != err {
			t.Error(err)
		}
		if out != tc.screamingSnake {
			mismatch(t, tc.screamingSnake, out)
		}
	}
}

func TestToKebab(t *testing.T) {
	for _, tc := range TestCases {
		out, err := durianfarts.ToKebab(tc.input)
		if nil != err {
			t.Error(err)
		}
		if out != tc.kebab {
			mismatch(t, tc.kebab, out)
		}
	}
}

func TestToScreamingKebab(t *testing.T) {
	for _, tc := range TestCases {
		out, err := durianfarts.ToScreamingKebab(tc.input)
		if nil != err {
			t.Error(err)
		}
		if out != tc.screamingKebab {
			mismatch(t, tc.screamingKebab, out)
		}
	}
}

func TestToCamel(t *testing.T) {
	for _, tc := range TestCases {
		out, err := durianfarts.ToCamel(tc.input)
		if nil != err {
			t.Error(err)
		}
		if out != tc.camel {
			mismatch(t, tc.camel, out)
		}
	}
}

func TestToLowerCamel(t *testing.T) {
	for _, tc := range TestCases {
		out, err := durianfarts.ToLowerCamel(tc.input)
		if nil != err {
			t.Error(err)
		}
		if out != tc.lowerCamel {
			mismatch(t, tc.lowerCamel, out)
		}
	}
}

func TestCustom(t *testing.T) {
	for _, tc := range TestCases {
		out, err := durianfarts.ToCustomTransform(tc.input, func(s string) string {
			var o string = s
			o = strings.Replace(o, "a", "4", -1)
			o = strings.Replace(o, "i", "1", -1)
			return o
		})
		if nil != err {
			t.Error(err)
		}
		if out != tc.custom {
			mismatch(t, tc.custom, out)
		}
	}
}

func TestInvalidJSON(t *testing.T) {
	invalidInput := "{,"
	_, err := durianfarts.ToSnake(invalidInput)
	if nil == err {
		t.Error("expected invalid JSON to fail")
	}
}

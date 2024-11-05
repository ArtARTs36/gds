package gds

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSplitWords(t *testing.T) {
	cases := []struct {
		String        string
		ExpectedWords []*SplitWord
	}{
		{
			String:        "",
			ExpectedWords: []*SplitWord{},
		},
		{
			String: "a",
			ExpectedWords: []*SplitWord{
				{
					Word: "a",
				},
			},
		},
		{
			String: "a_",
			ExpectedWords: []*SplitWord{
				{
					Word:           "a",
					SeparatorAfter: "_",
				},
			},
		},
		{
			String: "a-",
			ExpectedWords: []*SplitWord{
				{
					Word:           "a",
					SeparatorAfter: "-",
				},
			},
		},
		{
			String: "a ",
			ExpectedWords: []*SplitWord{
				{
					Word:           "a",
					SeparatorAfter: " ",
				},
			},
		},
		{
			String: "ab_cdf",
			ExpectedWords: []*SplitWord{
				{
					Word:           "ab",
					SeparatorAfter: "_",
				},
				{
					Word: "cdf",
				},
			},
		},
		{
			String: "abCdf",
			ExpectedWords: []*SplitWord{
				{
					Word: "ab",
				},
				{
					Word: "Cdf",
				},
			},
		},
		{
			String: "AbCdf",
			ExpectedWords: []*SplitWord{
				{
					Word: "Ab",
				},
				{
					Word: "Cdf",
				},
			},
		},
		{
			String: "AbCdf_fa_OK",
			ExpectedWords: []*SplitWord{
				{
					Word: "Ab",
				},
				{
					Word:           "Cdf",
					SeparatorAfter: "_",
				},
				{
					Word:           "fa",
					SeparatorAfter: "_",
				},
				{
					Word: "OK",
				},
			},
		},
		{
			String: "goose_db_version",
			ExpectedWords: []*SplitWord{
				{
					Word:           "goose",
					SeparatorAfter: "_",
				},
				{
					Word:           "db",
					SeparatorAfter: "_",
				},
				{
					Word: "version",
				},
			},
		},
		{
			String: "GooseDbVersion",
			ExpectedWords: []*SplitWord{
				{
					Word:           "Goose",
					SeparatorAfter: "",
				},
				{
					Word:           "Db",
					SeparatorAfter: "",
				},
				{
					Word: "Version",
				},
			},
		},
	}

	for i, tCase := range cases {
		t.Run(fmt.Sprintf("%d: %s", i, tCase.String), func(t *testing.T) {
			str := NewString(tCase.String)
			split := str.SplitWords()

			errorMsg := []string{
				"expected:",
			}

			for _, w := range tCase.ExpectedWords {
				errorMsg = append(errorMsg, fmt.Sprintf(
					"(%s, %s)",
					w.Word,
					w.SeparatorAfter,
				))
			}

			errorMsg = append(errorMsg, "\nactual:\n")

			for _, w := range split {
				errorMsg = append(errorMsg, fmt.Sprintf(
					"(%s, %s)",
					w.Word,
					w.SeparatorAfter,
				))
			}

			assert.Equal(t, tCase.ExpectedWords, split, strings.Join(errorMsg, "\n"))
		})
	}
}

func TestStringFixAbbreviations(t *testing.T) {
	cases := []struct {
		String        string
		Abbreviations []string
		Expected      string
	}{
		{
			String:        "",
			Abbreviations: []string{},
			Expected:      "",
		},
		{
			String: "goose_db_version",
			Abbreviations: []string{
				"db",
			},
			Expected: "goose_DB_version",
		},
		{
			String: "GooseDbVersion",
			Abbreviations: []string{
				"db",
			},
			Expected: "GooseDBVersion",
		},
		{
			String: "Id",
			Abbreviations: []string{
				"id",
			},
			Expected: "ID",
		},
	}

	for i, tCase := range cases {
		t.Run(fmt.Sprintf("%d: %s", i, tCase.String), func(t *testing.T) {
			str := NewString(tCase.String)

			abbrSet := map[string]bool{}
			for _, abbreviation := range tCase.Abbreviations {
				abbrSet[abbreviation] = true
			}

			assert.Equal(t, tCase.Expected, str.FixAbbreviations(abbrSet).Value)
		})
	}
}

func TestPluralStringFixAbbreviations(t *testing.T) {
	cases := []struct {
		String        string
		Abbreviations map[string]string
		Expected      string
	}{
		{
			String:        "",
			Abbreviations: map[string]string{},
			Expected:      "",
		},
		{
			String: "goose_db_version",
			Abbreviations: map[string]string{
				"db": "DBs",
			},
			Expected: "goose_DB_versions",
		},
		{
			String: "GooseDbVersion",
			Abbreviations: map[string]string{
				"db": "DBs",
			},
			Expected: "GooseDBVersions",
		},
		{
			String: "Id",
			Abbreviations: map[string]string{
				"id": "IDs",
			},
			Expected: "IDs",
		},
	}

	for i, tCase := range cases {
		t.Run(fmt.Sprintf("%d: %s", i, tCase.String), func(t *testing.T) {
			str := NewString(tCase.String)

			assert.Equal(t, tCase.Expected, str.PluralFixAbbreviations(tCase.Abbreviations).Value)
		})
	}
}

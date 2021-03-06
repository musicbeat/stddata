// Copyright 2014 Musicbeat.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
data loader component for ISO 639.2 language codes.
source is US Library of Congress
http://www.loc.gov/standards/iso639-2/ISO-639-2_utf-8.txt
These are notes from:
http://www.loc.gov/standards/iso639-2/ascii_8bits.html
"These files may be used to download the list of
language codes with their language names, for example into a
database. To read the files, please note that one line of text
contains one entry. An alpha-3 (bibliographic) code, an alpha-3
(terminologic) code (when given), an alpha-2 code (when given),
an English name, and a French name of a language are all separated
by pipe (|) characters. If one of these elements is not applicable
to the entry, the field is left empty, i.e., a pipe (|) character
immediately follows the preceding entry. The Line terminator is
the LF character."
*/
package language

import (
	"encoding/csv"
	"io"
	"net/http"
	"sort"
	"strings"

	"github.com/musicbeat/stddata"
)

// LanguageProvider implements the Provider interfaces.
type LanguageProvider struct {
	loaded          bool
	size            int
	languageIndexes map[string]languageIndex
}

type languageIndex struct {
	languageMap  map[string][]Language
	languageKeys []string
}

// Language is the information on one language in the source data
type Language struct {
	Alpha3bibliographic string
	Alpha3terminologic  string
	Alpha2              string
	EnglishName         string
	FrenchName          string
}

// LanguageResult is the interface{} that is returned from Search
type LanguageResult struct {
	Languages [][]Language
}

var alphaMap map[string][]Language
var englishNameMap map[string][]Language

// Load does the heavy lifting of retrieving the
// Library of Congress' list of languages, a pipe-delimited
// .csv file, and populating maps for searching.
func (p *LanguageProvider) Load() (n int, err error) {
	// initialize the maps:
	p.languageIndexes = make(map[string]languageIndex)
	alphaMap = make(map[string][]Language)
	englishNameMap = make(map[string][]Language)

	res, err := http.Get("http://www.loc.gov/standards/iso639-2/ISO-639-2_utf-8.txt")
	if err != nil {
		return 0, &stddata.ServiceError{err.Error(), http.StatusServiceUnavailable}
	}

	reader := csv.NewReader(res.Body)
	reader.Comma = '|'
	reader.FieldsPerRecord = 5
	reader.TrimLeadingSpace = true

	defer res.Body.Close()

	for {
		// read just one record
		record, err := reader.Read()
		// end-of-file is fitted into err
		if err == io.EOF {
			break
		} else if err != nil {
			return 0, &stddata.ServiceError{err.Error(), http.StatusServiceUnavailable}
		}

		var l Language
		l.Alpha3bibliographic = record[0]
		l.Alpha3terminologic = record[1]
		l.Alpha2 = record[2]
		l.EnglishName = record[3]
		l.FrenchName = record[4]

		// add the language to the maps:
		alphaMap[l.Alpha3bibliographic] = append(alphaMap[l.Alpha3bibliographic], l)
		englishNameMap[l.EnglishName] = append(englishNameMap[l.EnglishName], l)

	}
	p.storeData("alpha", alphaMap)
	p.storeData("name", englishNameMap)
	p.size = len(alphaMap)
	p.loaded = true
	return len(alphaMap), err
}

func (p *LanguageProvider) storeData(s string, m map[string][]Language) {
	// store the map
	var li languageIndex
	li.languageMap = m
	// extract the keys
	li.languageKeys = make([]string, len(m))
	i := 0
	for k, _ := range m {
		li.languageKeys[i] = k
		i++
	}
	// sort the keys
	sort.Strings(li.languageKeys)
	// add to languageIndexes
	p.languageIndexes[s] = li
}

// Search returns a collection as an interface{} and error. The collection
// contains an array of the results to the search. The value
// in index is used to choose the map of Language entities that will be searched.
// If the value in index does not match the name of a map, an error is returned.
// The keys in the map specified by index are searched using a regex-like 'query.*', and
// any matching Languages are returned in the result.
// Search can also "dump" an index. When the value of query is "_dump", the index specified
// is used to supply the entire data set, in the order of the index.
func (p *LanguageProvider) Search(index string, query string) (result interface{}, err error) {
	// make sure the data is loaded
	if p.loaded != true {
		return nil, &stddata.ServiceError{err.Error(), http.StatusServiceUnavailable}
	}
	li, found := p.languageIndexes[index]
	if !found {
		// search cannot be performed
		msg := "No index on " + index
		return nil, &stddata.ServiceError{msg, http.StatusBadRequest}
	}
	result = doSearch(li, query)
	return result, nil
}
func doSearch(li languageIndex, query string) (res LanguageResult) {
	// the "reserved" query term "_dump" is handled by returning all the
	// results in the order of the index.
	dump := query == "_dump"
	// prepare the response. allocate enough space for the response to be the
	// entire data set.
	tmp := make([][]Language, len(li.languageKeys))
	// brute force the sorted list of keys, looking for a match to 'query.*'.
	// add each match to the result array. The results are added in the
	// order of the sorted keys, so the results are sorted.
	i := 0
	for k := range li.languageKeys {
		if dump {
			tmp[i] = li.languageMap[li.languageKeys[k]]
			i++
		} else if len(li.languageKeys[k]) >= len(query) {
			if strings.EqualFold(query, li.languageKeys[k][0:len(query)]) {
				tmp[i] = li.languageMap[li.languageKeys[k]]
				i++
			}
		}
	}
	res.Languages = tmp[0:i]
	return res
}

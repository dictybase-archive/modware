package publication

import "strconv"

var pubColumns = []string{
	"title",
	"volume",
	"series_name",
	"issue",
	"pages",
	"uniquename",
	"name",
	"pubplace",
	"pyear",
}

var pubTestData = [][]string{
	[]string{
		"dictyBase 2015: Expanding data and annotations in a new software environment",
		"12",
		"Genesis",
		"8",
		"765-80",
		"99",
		"journal_article",
		"pubmed",
		"2015",
	},
}

func getPubTestDataRows() [][]string {
	var testData [][]string
	for i := 10; i <= 12; i++ {
		rowData := make([]string, len(pubTestData[0]))
		_ = copy(rowData, pubTestData[0])
		rowData[5] = strconv.Itoa(i)
		testData = append(testData, rowData)
	}
	return testData
}

var selectPubCols = []string{
	"uniquename",
	"series_name",
	"issue",
	"pages",
	"pubplace",
	"pyear",
}

var selectpubTestData = [][]string{
	[]string{
		"99",
		"Genesis",
		"8",
		"765-80",
		"pubmed",
		"2015",
	},
}

var selectPubCols2 = []string{
	"uniquename",
	"series_name",
}

var selectpubTestData2 = [][]string{
	[]string{
		"99",
		"Genesis",
	},
}

func getSelectPubTestDataRows() [][]string {
	var testData [][]string
	for i := 10; i <= 12; i++ {
		rowData := make([]string, len(selectpubTestData2[0]))
		_ = copy(rowData, selectpubTestData2[0])
		rowData[0] = strconv.Itoa(i)
		testData = append(testData, rowData)
	}
	return testData
}

var propTestData = []map[string]string{
	map[string]string{
		"doi":      "10.1002/dvg.22867",
		"abstract": "This is an abstract",
		"status":   "ppublish",
		"month":    "june",
		"issn":     "1526-968X",
	},
}

func getPropsTestDataRows() []map[string]string {
	num := len(getPubTestDataRows())
	testData := make([]map[string]string, num)
	for i := 0; i < num; i++ {
		testData[i] = propTestData[0]
	}
	return testData
}

var selectpropTestData = []map[string]string{
	map[string]string{
		"doi":   "10.1002/dvg.22867",
		"month": "june",
	},
}

var selectpropTestData2 = []map[string]string{
	map[string]string{
		"doi": "10.1002/dvg.22867",
	},
}

func getSelectPropsTestDataRows() []map[string]string {
	num := len(getSelectPubTestDataRows())
	testData := make([]map[string]string, num)
	for i := 0; i < num; i++ {
		testData[i] = selectpropTestData2[0]
	}
	return testData
}

var authorColumns = []string{
	"pubauthor_id",
	"rank",
	"surname",
	"givennames",
}

var selectauthorColumns = []string{
	"pubauthor_id",
	"rank",
}

var selectauthorColumns2 = []string{
	"pubauthor_id",
	"surname",
}

var authorData = [][]string{
	[]string{"23", "3", "Wardroper", "A"},
	[]string{"12", "1", "Quail", "MA"},
	[]string{"13", "1", "Quail", "MA"},
	[]string{"1", "1", "Johnson", "S"},
	[]string{"2", "1", "Procise", "C"},
	[]string{"29", "1", "Lynch", "M"},
}

var selectauthorData = [][]string{
	[]string{"23", "3"},
	[]string{"12", "1"},
}

var selectauthorData2 = [][]string{
	[]string{"23", "Wardroper"},
	[]string{"12", "Quail"},
	[]string{"13", "Quail"},
	[]string{"1", "Johnson"},
	[]string{"2", "Procise"},
	[]string{"29", "Lynch"},
}

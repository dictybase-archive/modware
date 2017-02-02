package publication

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

var propTestData = []map[string]string{
	map[string]string{
		"doi":      "10.1002/dvg.22867",
		"abstract": "This is an abstract",
		"status":   "ppublish",
		"month":    "june",
		"issn":     "1526-968X",
	},
}

var selectpropTestData = []map[string]string{
	map[string]string{
		"doi":   "10.1002/dvg.22867",
		"month": "june",
	},
}

var authorColumns = []string{
	"pubauthor_id",
	"rank",
	"surname",
	"givennames",
}

var authorData = [][]string{
	[]string{"23", "3", "Wardroper", "A"},
	[]string{"12", "1", "Quail", "MA"},
}

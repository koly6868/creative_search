package searcher_templates

type _TemplateModelName struct {
	//# {{ range  .Searchers }}
	_TemplateSearcherKey _TemplateSearcherKeyType
	//# {{ end }}
}

//# {{ range  .Searchers }}
type _TemplateSearcherKeyType string

//# {{ end }}

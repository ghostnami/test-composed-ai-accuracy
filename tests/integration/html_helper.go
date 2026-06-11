// hunk: @@ -37,11 +37,6 @@ func (doc *HTMLDoc) Find(selector string) *goquery.Selection {
	return doc.doc.Find(selector)
}

// AssertHTMLElement check if the element by selector exists or does not exist depending on checkExists
func AssertHTMLElement[T int | bool](t testing.TB, doc *HTMLDoc, selector string, checkExists T) {
	sel := doc.doc.Find(selector)

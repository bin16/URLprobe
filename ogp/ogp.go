package ogp

import (
	"strings"

	"golang.org/x/net/html"
)

// ogp types
const (
	ogpMetaName    = "property"
	ogpMetaContent = "content"

	oUnknown = iota
	oType
	oTitle
	oDescription
	oSiteName
	oURL
	oLocale
	oLocaleAlt
	oImage
	oImageAttr
)

// Parse html.Node as OGP data
func Parse(n *html.Node) Root {
	root := Root{}
	nl := queryHTMLNodes(n, isMeta)
	ml := htmlNodesToOgpMetas(nl)
	for _, m := range ml {
		root.set(m)
	}

	return root
}

// Root is the OGP structure
type Root struct {
	SiteName        string   `json:"siteName"`
	URL             string   `json:"url"`
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	Locale          string   `json:"locale"`
	LocaleAlternate []string `json:"localeAlt"`
	Type            string   `json:"type"`
	rawType         int
	Image           []map[string]string `json:"image"`
}

func (r *Root) set(m ogpMeta) {
	switch m.category {
	case oTitle:
		r.Title = m.String()
	case oDescription:
		r.Description = m.String()
	case oType:
		r.Type = m.String()
	case oURL:
		r.URL = m.String()
	case oLocale:
		r.Locale = m.String()
	case oLocaleAlt:
		r.LocaleAlternate = append(r.LocaleAlternate, m.String())
	case oSiteName:
		r.SiteName = m.String()
	case oImage:
		img := map[string]string{"url": m.String()}
		r.Image = append(r.Image, img)
	case oImageAttr:
		img := r.Image[len(r.Image)-1]
		if img != nil {
			img[m.attrKey()] = m.String()
		}
	}
}

type ogpMeta struct {
	Property string
	Content  string
	index    int
	category int
}

func (m *ogpMeta) attrKey() (key string) {
	kl := strings.Split(m.Property, ":")
	return kl[2]
}

func (m *ogpMeta) String() string {
	return m.Content
}

func htmlNodesToOgpMetas(nl []*html.Node) []ogpMeta {
	ml := []ogpMeta{}
	for i, n := range nl {
		if isMeta(n) {
			if pName, ogpExists := nGetAttr(n, ogpMetaName); ogpExists {
				if pValue, ogpValueExists := ogpGetValue(n, pName); ogpValueExists {
					ml = append(ml, ogpMeta{
						Property: pName,
						Content:  pValue,
						index:    i,
						category: detectProperty(pName),
					})
				}
			}
		}
	}

	return ml
}

func detectProperty(p string) int {
	m := map[string]int{
		"og:title":            oTitle,
		"og:description":      oDescription,
		"og:site_name":        oSiteName,
		"og:url":              oURL,
		"og:locale":           oLocale,
		"og:locale:alternate": oLocaleAlt,

		"og:image":            oImage,
		"og:image:url":        oImage,
		"og:image:secure_url": oImageAttr,
		"og:image:width":      oImageAttr,
		"og:image:height":     oImageAttr,
		"og:image:type":       oImageAttr,
		"og:image:alt":        oImageAttr,
	}

	if o, exists := m[p]; exists {
		return o
	}

	return oUnknown
}

func queryHTMLNodes(root *html.Node, matchFn func(*html.Node) bool) []*html.Node {
	matched := []*html.Node{}
	var find func(*html.Node)
	find = func(n *html.Node) {
		if matchFn(n) {
			matched = append(matched, n)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			find(c)
		}
	}
	find(root)

	return matched
}

func isMeta(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "meta"
}

func nGetAttr(n *html.Node, attrKey string) (value string, exists bool) {
	for _, a := range n.Attr {
		if a.Key == attrKey {
			return a.Val, true
		}
	}

	return "", false
}

func ogpGetValue(n *html.Node, name string) (string, bool) {
	if ogpKey, ok := nGetAttr(n, ogpMetaName); ok && name == ogpKey {
		if ogpValue, ok := nGetAttr(n, ogpMetaContent); ok {
			return ogpValue, true
		}
	}

	return "", false
}

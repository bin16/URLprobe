package ogp

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestParse(t *testing.T) {
	d0 := Root{
		Title:    "Build software better, together",
		SiteName: "GitHub",
		URL:      "https://github.com",
		Image: []map[string]string{
			map[string]string{
				"url":    "https://github.githubassets.com/images/modules/open_graph/github-logo.png",
				"width":  "1200",
				"height": "1200",
			},
			map[string]string{
				"url":    "https://github.githubassets.com/images/modules/open_graph/github-mark.png",
				"width":  "1200",
				"height": "620",
			},
			map[string]string{
				"url":    "https://github.githubassets.com/images/modules/open_graph/github-octocat.png",
				"width":  "1200",
				"height": "620",
			},
		},
	}

	data := Parse(htmlNode())
	if data.Title != d0.Title {
		t.Errorf("og:title: <%s>; want: <%s>", data.Title, d0.Title)
	}
	if data.SiteName != d0.SiteName {
		t.Errorf("og:site_name: <%s>; want: <%s>", data.SiteName, d0.SiteName)
	}
	if data.URL != d0.URL {
		t.Errorf("og:url: <%s>; want: <%s>", data.SiteName, d0.URL)
	}

	for i, img := range d0.Image {
		for k, v := range img {
			if data.Image[i][k] != v {
				t.Errorf("og:image:%s <%s>; want: <%s>", k, data.Image[i][k], v)
			}
		}
	}
}

func htmlNode() *html.Node {
	codes := `
  <!DOCTYPE html>
  <html lang="en">

  <head>
    <meta charset="utf-8">
    <link rel="dns-prefetch" href="https://github.githubassets.com">
    <link rel="dns-prefetch" href="https://avatars0.githubusercontent.com">
    <meta name="viewport" content="width=device-width">
    <title>GitHub</title>
    <meta name="description" content="GitHub is where people build software. More than 50 million people use GitHub to discover, fork, and contribute to over 100 million projects.">
    <link rel="search" type="application/opensearchdescription+xml" href="/opensearch.xml" title="GitHub">
    <link rel="fluid-icon" href="https://github.com/fluidicon.png" title="GitHub">
    <meta property="fb:app_id" content="1401488693436528">
    <meta name="apple-itunes-app" content="app-id=1477376905">
    <meta property="og:url" content="https://github.com">
    <meta property="og:site_name" content="GitHub">
    <meta property="og:title" content="Build software better, together">
    <meta property="og:description"
      content="GitHub is where people build software. More than 50 million people use GitHub to discover, fork, and contribute to over 100 million projects.">
    <meta property="og:image" content="https://github.githubassets.com/images/modules/open_graph/github-logo.png">
    <meta property="og:image:type" content="image/png">
    <meta property="og:image:width" content="1200">
    <meta property="og:image:height" content="1200">
    <meta property="og:image" content="https://github.githubassets.com/images/modules/open_graph/github-mark.png">
    <meta property="og:image:type" content="image/png">
    <meta property="og:image:width" content="1200">
    <meta property="og:image:height" content="620">
    <meta property="og:image" content="https://github.githubassets.com/images/modules/open_graph/github-octocat.png">
    <meta property="og:image:type" content="image/png">
    <meta property="og:image:width" content="1200">
    <meta property="og:image:height" content="620">
    <meta property="twitter:site" content="github">
    <meta property="twitter:site:id" content="13334762">
    <meta property="twitter:creator" content="github">
    <meta property="twitter:creator:id" content="13334762">
    <meta property="twitter:card" content="summary_large_image">
    <meta property="twitter:title" content="GitHub">
    <meta property="twitter:description"
      content="GitHub is where people build software. More than 50 million people use GitHub to discover, fork, and contribute to over 100 million projects.">
    <meta property="twitter:image:src"
      content="https://github.githubassets.com/images/modules/open_graph/github-logo.png">
    <meta property="twitter:image:width" content="1200">
    <meta property="twitter:image:height" content="1200">
    <link rel="assets" href="https://github.githubassets.com/">
    <link rel="sudo-modal" href="/sessions/sudo_modal">
    <meta name="cookie-consent-required" content="false" />
    <meta name="page-subject" content="GitHub">
    <meta name="github-keyboard-shortcuts" content="dashboards" data-pjax-transient="true" />
    <meta name="selected-link" value="/" data-pjax-transient>
    <meta name="analytics-location" content="/dashboard" data-pjax-transient="true" />
    <meta name="hostname" content="github.com">
    <meta name="user-login" content="">
    <meta name="expected-hostname" content="github.com">
    <meta name="enabled-features" content="MARKETPLACE_PENDING_INSTALLATIONS,JS_HTTP_CACHE_HEADERS">
    <meta name="browser-stats-url" content="https://api.github.com/_private/browser/stats">
    <meta name="browser-errors-url" content="https://api.github.com/_private/browser/errors">
    <link rel="mask-icon" href="https://github.githubassets.com/pinned-octocat.svg" color="#000000">
    <link rel="alternate icon" class="js-site-favicon" type="image/png"
      href="https://github.githubassets.com/favicons/favicon.png">
    <link rel="icon" class="js-site-favicon" type="image/svg+xml"
      href="https://github.githubassets.com/favicons/favicon.svg">
    <meta name="theme-color" content="#1e2327">
    <link rel="manifest" href="/manifest.json" crossOrigin="use-credentials">
  </head>
  <body class="logged-in env-production page-responsive full-width">
  </body>

  </html>
`
	n, _ := html.Parse(strings.NewReader(codes))

	return n
}

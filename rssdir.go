// package rssdir generates a static RSS feed from the entries in a directory.
//
// File attributes are used to determine the order and dates/times of the feed,
// as well as other attributes, such as the author, being extracted from the
// owning user and mapped to a name and email based on a configuration table.
// Files are assumed to be plain HTML and are parsed to obtain a title and
// extract. If this is not possible, the file name will be used for the entry
// title with no extract given. The tags from which the extract and title are
// taken can be configured based on either class or element ID. Otherwise, the
// page title and first paragraph element are used.
package rssdir

import (
	"fmt"

	"github.com/gorilla/feeds"
)

// Config contains the configuration for a given generated feed. These are used
// to supplement the information gained from the filesystem (such as file
// owners). Most details in here are cosmetic and the blank configuration is
// always valid.
type Config struct {
	// LinkRoot is the root link which file names should be appended to to
	// create the link to the site. This should include the section of the
	// URI which points to the RSS root directory (e.g https://ejv2.cc/blog/).
	// The trailing slash is optional and unimportant.
	LinkRoot string
	// ExtractLength is the number of characters which will be used to
	// generate the content extract. If zero, defaults to 255.
	ExtractLength uint
	// Authors maps a username to a feed-friendly author.
	// If a file owner is not in this file, his username is used instead.
	Authors map[string]feeds.Author
}

// Generate an RSS feed from the entries in the given directory path, returning
// the generated feed in original gorilla/feeds format, allowing the user to do
// with it as he pleases.
//
// The given directory is not traversed into subdirectories by default. This
// can be reconfigured in the configuration. The only option if walking
// recursively is to flatten the tree - in other words, all files are in the
// same feed.
func Generate(path string) (*feeds.Feed, error) {
	return GenerateWithConfig(path, Config{})
}

func GenerateWithConfig(path string, cfg Config) (*feeds.Feed, error) {
	return nil, fmt.Errorf("not implemented")
}

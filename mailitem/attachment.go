// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package mailitem

import "github.com/adampresley/mailslurper/mailheader"

type Attachment struct {
	Headers  *mailheader.AttachmentHeader
	Contents string
}

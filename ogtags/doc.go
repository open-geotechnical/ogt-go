
// Package `ogtags` handles transformation (and later validation) of ags4 formatted text.multi.csv files
//
// The code consists of two parts
// = the DataDict [FooDD] such as GroupDD, HeadingDD whcih contain the ags4 data dict
// - and a Document, with its lines and structure
//
// The general idea is that at application launch, the data dict DD's are loaded
// and browsable via a www site (an alternate to AGS's site)
// and can also parse on the fly ags documents and convert them
// mainly into json atmo, for research and fun ;-)
package ogtags


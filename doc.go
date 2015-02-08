// Package down is a markdown engine for in Go.
// It is currently in development, which would explain why it doesn't work.
//
// Heres a few random thoughts for me to keep track of.
// - Rewrite all this to be more state machinish
// - Figure out some sort of procedure
//
// Grammar:
// 		line 				::= <composite_string> | <header> | <unordered_list_item>+
//		composite_string 	::= (<link> | <bold> | <italics> | <raw_string> )? <composite_string>
// 		raw_string 			::= \s+
// 		link 				::= "[" <raw_string> "](" <raw_string> ")"
//		bold 				::= "*"<composite_string>"*"
//		italics 			::= "**"<composite_string>"**"
//
//		header 		::= headerOne | header_two | header_three | header_four | header_five | header_six
//		header_one 	::= "#" <composite_string>
//		headerTwo 	::= "##" <composite_string>
//		headerThree ::= "###" <composite_string>
//		headerFour 	::= "####" <composite_string>
//		headerFive 	::= "#####" <composite_string>
//		headerSix 	::= "######" <composite_string>
//
//		list ::= (<unordered_list_item>+ | ordered_list_item+)
// 		unordered_list_item ::= ("* " <list>)
// 		orderered_list_item ::= (/d+ ". " <list>)
package down

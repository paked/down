// Package down implements a markup engine in Go, using a subset of Markdown as it's syntax.
// It is currently in development, which would explain why it doesn't work.
//
// Grammar:
// 		line 				::= <paragraph> | <header> | <list>
//		paragraph 			::= <composite_string>
//		composite_string 	::= (<link> | <bold> | <italics> | <raw_string> )+
// 		raw_string 			::= \s+
// 		link 				::= "[" <raw_string> "](" <raw_string> ")"
//		bold 				::= "*"<composite_string>"*"
//		italics 			::= "**"<composite_string>"**"
//
//		header 		::= <headerOne> | <header_two> | <header_three> | <header_four> | <header_five> | <header_six>
//		header_one 	::= "#" <composite_string>
//		headerTwo 	::= "##" <composite_string>
//		headerThree ::= "###" <composite_string>
//		headerFour 	::= "####" <composite_string>
//		headerFive 	::= "#####" <composite_string>
//		headerSix 	::= "######" <composite_string>
//
//		list ::= (<unordered_list_item>+ | <ordered_list_item>+)
// 		unordered_list_item ::= ("* " (<list> | <composite_string>))
// 		orderered_list_item ::= (/d+ ". " (<list> | <composite_string>))
package down

// Code generated for smslite GSM cleaning; DO NOT EDIT BY HAND.
package smslite

// unicodeCompatibilityReplacementFor maps Unicode compatibility characters whose NFKD form
// can be safely represented with GSM-7/ASCII characters. It intentionally avoids
// dependencies on x/text so the package remains lightweight.
func unicodeCompatibilityReplacementFor(r rune) (string, bool) {
	switch r {
	case '\u00A0':
		return " ", true // NO-BREAK SPACE
	case '\u00AA':
		return "a", true // FEMININE ORDINAL INDICATOR
	case '\u00B2':
		return "2", true // SUPERSCRIPT TWO
	case '\u00B3':
		return "3", true // SUPERSCRIPT THREE
	case '\u00B9':
		return "1", true // SUPERSCRIPT ONE
	case '\u00BA':
		return "o", true // MASCULINE ORDINAL INDICATOR
	case '\u00C0':
		return "A", true // LATIN CAPITAL LETTER A WITH GRAVE
	case '\u00C1':
		return "A", true // LATIN CAPITAL LETTER A WITH ACUTE
	case '\u00C2':
		return "A", true // LATIN CAPITAL LETTER A WITH CIRCUMFLEX
	case '\u00C3':
		return "A", true // LATIN CAPITAL LETTER A WITH TILDE
	case '\u00C4':
		return "A", true // LATIN CAPITAL LETTER A WITH DIAERESIS
	case '\u00C5':
		return "A", true // LATIN CAPITAL LETTER A WITH RING ABOVE
	case '\u00C7':
		return "C", true // LATIN CAPITAL LETTER C WITH CEDILLA
	case '\u00C8':
		return "E", true // LATIN CAPITAL LETTER E WITH GRAVE
	case '\u00C9':
		return "E", true // LATIN CAPITAL LETTER E WITH ACUTE
	case '\u00CA':
		return "E", true // LATIN CAPITAL LETTER E WITH CIRCUMFLEX
	case '\u00CB':
		return "E", true // LATIN CAPITAL LETTER E WITH DIAERESIS
	case '\u00CC':
		return "I", true // LATIN CAPITAL LETTER I WITH GRAVE
	case '\u00CD':
		return "I", true // LATIN CAPITAL LETTER I WITH ACUTE
	case '\u00CE':
		return "I", true // LATIN CAPITAL LETTER I WITH CIRCUMFLEX
	case '\u00CF':
		return "I", true // LATIN CAPITAL LETTER I WITH DIAERESIS
	case '\u00D1':
		return "N", true // LATIN CAPITAL LETTER N WITH TILDE
	case '\u00D2':
		return "O", true // LATIN CAPITAL LETTER O WITH GRAVE
	case '\u00D3':
		return "O", true // LATIN CAPITAL LETTER O WITH ACUTE
	case '\u00D4':
		return "O", true // LATIN CAPITAL LETTER O WITH CIRCUMFLEX
	case '\u00D5':
		return "O", true // LATIN CAPITAL LETTER O WITH TILDE
	case '\u00D6':
		return "O", true // LATIN CAPITAL LETTER O WITH DIAERESIS
	case '\u00D9':
		return "U", true // LATIN CAPITAL LETTER U WITH GRAVE
	case '\u00DA':
		return "U", true // LATIN CAPITAL LETTER U WITH ACUTE
	case '\u00DB':
		return "U", true // LATIN CAPITAL LETTER U WITH CIRCUMFLEX
	case '\u00DC':
		return "U", true // LATIN CAPITAL LETTER U WITH DIAERESIS
	case '\u00DD':
		return "Y", true // LATIN CAPITAL LETTER Y WITH ACUTE
	case '\u00E0':
		return "a", true // LATIN SMALL LETTER A WITH GRAVE
	case '\u00E1':
		return "a", true // LATIN SMALL LETTER A WITH ACUTE
	case '\u00E2':
		return "a", true // LATIN SMALL LETTER A WITH CIRCUMFLEX
	case '\u00E3':
		return "a", true // LATIN SMALL LETTER A WITH TILDE
	case '\u00E4':
		return "a", true // LATIN SMALL LETTER A WITH DIAERESIS
	case '\u00E5':
		return "a", true // LATIN SMALL LETTER A WITH RING ABOVE
	case '\u00E7':
		return "c", true // LATIN SMALL LETTER C WITH CEDILLA
	case '\u00E8':
		return "e", true // LATIN SMALL LETTER E WITH GRAVE
	case '\u00E9':
		return "e", true // LATIN SMALL LETTER E WITH ACUTE
	case '\u00EA':
		return "e", true // LATIN SMALL LETTER E WITH CIRCUMFLEX
	case '\u00EB':
		return "e", true // LATIN SMALL LETTER E WITH DIAERESIS
	case '\u00EC':
		return "i", true // LATIN SMALL LETTER I WITH GRAVE
	case '\u00ED':
		return "i", true // LATIN SMALL LETTER I WITH ACUTE
	case '\u00EE':
		return "i", true // LATIN SMALL LETTER I WITH CIRCUMFLEX
	case '\u00EF':
		return "i", true // LATIN SMALL LETTER I WITH DIAERESIS
	case '\u00F1':
		return "n", true // LATIN SMALL LETTER N WITH TILDE
	case '\u00F2':
		return "o", true // LATIN SMALL LETTER O WITH GRAVE
	case '\u00F3':
		return "o", true // LATIN SMALL LETTER O WITH ACUTE
	case '\u00F4':
		return "o", true // LATIN SMALL LETTER O WITH CIRCUMFLEX
	case '\u00F5':
		return "o", true // LATIN SMALL LETTER O WITH TILDE
	case '\u00F6':
		return "o", true // LATIN SMALL LETTER O WITH DIAERESIS
	case '\u00F9':
		return "u", true // LATIN SMALL LETTER U WITH GRAVE
	case '\u00FA':
		return "u", true // LATIN SMALL LETTER U WITH ACUTE
	case '\u00FB':
		return "u", true // LATIN SMALL LETTER U WITH CIRCUMFLEX
	case '\u00FC':
		return "u", true // LATIN SMALL LETTER U WITH DIAERESIS
	case '\u00FD':
		return "y", true // LATIN SMALL LETTER Y WITH ACUTE
	case '\u00FF':
		return "y", true // LATIN SMALL LETTER Y WITH DIAERESIS
	case '\u0100':
		return "A", true // LATIN CAPITAL LETTER A WITH MACRON
	case '\u0101':
		return "a", true // LATIN SMALL LETTER A WITH MACRON
	case '\u0102':
		return "A", true // LATIN CAPITAL LETTER A WITH BREVE
	case '\u0103':
		return "a", true // LATIN SMALL LETTER A WITH BREVE
	case '\u0104':
		return "A", true // LATIN CAPITAL LETTER A WITH OGONEK
	case '\u0105':
		return "a", true // LATIN SMALL LETTER A WITH OGONEK
	case '\u0106':
		return "C", true // LATIN CAPITAL LETTER C WITH ACUTE
	case '\u0107':
		return "c", true // LATIN SMALL LETTER C WITH ACUTE
	case '\u0108':
		return "C", true // LATIN CAPITAL LETTER C WITH CIRCUMFLEX
	case '\u0109':
		return "c", true // LATIN SMALL LETTER C WITH CIRCUMFLEX
	case '\u010A':
		return "C", true // LATIN CAPITAL LETTER C WITH DOT ABOVE
	case '\u010B':
		return "c", true // LATIN SMALL LETTER C WITH DOT ABOVE
	case '\u010C':
		return "C", true // LATIN CAPITAL LETTER C WITH CARON
	case '\u010D':
		return "c", true // LATIN SMALL LETTER C WITH CARON
	case '\u010E':
		return "D", true // LATIN CAPITAL LETTER D WITH CARON
	case '\u010F':
		return "d", true // LATIN SMALL LETTER D WITH CARON
	case '\u0112':
		return "E", true // LATIN CAPITAL LETTER E WITH MACRON
	case '\u0113':
		return "e", true // LATIN SMALL LETTER E WITH MACRON
	case '\u0114':
		return "E", true // LATIN CAPITAL LETTER E WITH BREVE
	case '\u0115':
		return "e", true // LATIN SMALL LETTER E WITH BREVE
	case '\u0116':
		return "E", true // LATIN CAPITAL LETTER E WITH DOT ABOVE
	case '\u0117':
		return "e", true // LATIN SMALL LETTER E WITH DOT ABOVE
	case '\u0118':
		return "E", true // LATIN CAPITAL LETTER E WITH OGONEK
	case '\u0119':
		return "e", true // LATIN SMALL LETTER E WITH OGONEK
	case '\u011A':
		return "E", true // LATIN CAPITAL LETTER E WITH CARON
	case '\u011B':
		return "e", true // LATIN SMALL LETTER E WITH CARON
	case '\u011C':
		return "G", true // LATIN CAPITAL LETTER G WITH CIRCUMFLEX
	case '\u011D':
		return "g", true // LATIN SMALL LETTER G WITH CIRCUMFLEX
	case '\u011E':
		return "G", true // LATIN CAPITAL LETTER G WITH BREVE
	case '\u011F':
		return "g", true // LATIN SMALL LETTER G WITH BREVE
	case '\u0120':
		return "G", true // LATIN CAPITAL LETTER G WITH DOT ABOVE
	case '\u0121':
		return "g", true // LATIN SMALL LETTER G WITH DOT ABOVE
	case '\u0122':
		return "G", true // LATIN CAPITAL LETTER G WITH CEDILLA
	case '\u0123':
		return "g", true // LATIN SMALL LETTER G WITH CEDILLA
	case '\u0124':
		return "H", true // LATIN CAPITAL LETTER H WITH CIRCUMFLEX
	case '\u0125':
		return "h", true // LATIN SMALL LETTER H WITH CIRCUMFLEX
	case '\u0128':
		return "I", true // LATIN CAPITAL LETTER I WITH TILDE
	case '\u0129':
		return "i", true // LATIN SMALL LETTER I WITH TILDE
	case '\u012A':
		return "I", true // LATIN CAPITAL LETTER I WITH MACRON
	case '\u012B':
		return "i", true // LATIN SMALL LETTER I WITH MACRON
	case '\u012C':
		return "I", true // LATIN CAPITAL LETTER I WITH BREVE
	case '\u012D':
		return "i", true // LATIN SMALL LETTER I WITH BREVE
	case '\u012E':
		return "I", true // LATIN CAPITAL LETTER I WITH OGONEK
	case '\u012F':
		return "i", true // LATIN SMALL LETTER I WITH OGONEK
	case '\u0130':
		return "I", true // LATIN CAPITAL LETTER I WITH DOT ABOVE
	case '\u0132':
		return "IJ", true // LATIN CAPITAL LIGATURE IJ
	case '\u0133':
		return "ij", true // LATIN SMALL LIGATURE IJ
	case '\u0134':
		return "J", true // LATIN CAPITAL LETTER J WITH CIRCUMFLEX
	case '\u0135':
		return "j", true // LATIN SMALL LETTER J WITH CIRCUMFLEX
	case '\u0136':
		return "K", true // LATIN CAPITAL LETTER K WITH CEDILLA
	case '\u0137':
		return "k", true // LATIN SMALL LETTER K WITH CEDILLA
	case '\u0139':
		return "L", true // LATIN CAPITAL LETTER L WITH ACUTE
	case '\u013A':
		return "l", true // LATIN SMALL LETTER L WITH ACUTE
	case '\u013B':
		return "L", true // LATIN CAPITAL LETTER L WITH CEDILLA
	case '\u013C':
		return "l", true // LATIN SMALL LETTER L WITH CEDILLA
	case '\u013D':
		return "L", true // LATIN CAPITAL LETTER L WITH CARON
	case '\u013E':
		return "l", true // LATIN SMALL LETTER L WITH CARON
	case '\u0143':
		return "N", true // LATIN CAPITAL LETTER N WITH ACUTE
	case '\u0144':
		return "n", true // LATIN SMALL LETTER N WITH ACUTE
	case '\u0145':
		return "N", true // LATIN CAPITAL LETTER N WITH CEDILLA
	case '\u0146':
		return "n", true // LATIN SMALL LETTER N WITH CEDILLA
	case '\u0147':
		return "N", true // LATIN CAPITAL LETTER N WITH CARON
	case '\u0148':
		return "n", true // LATIN SMALL LETTER N WITH CARON
	case '\u014C':
		return "O", true // LATIN CAPITAL LETTER O WITH MACRON
	case '\u014D':
		return "o", true // LATIN SMALL LETTER O WITH MACRON
	case '\u014E':
		return "O", true // LATIN CAPITAL LETTER O WITH BREVE
	case '\u014F':
		return "o", true // LATIN SMALL LETTER O WITH BREVE
	case '\u0150':
		return "O", true // LATIN CAPITAL LETTER O WITH DOUBLE ACUTE
	case '\u0151':
		return "o", true // LATIN SMALL LETTER O WITH DOUBLE ACUTE
	case '\u0154':
		return "R", true // LATIN CAPITAL LETTER R WITH ACUTE
	case '\u0155':
		return "r", true // LATIN SMALL LETTER R WITH ACUTE
	case '\u0156':
		return "R", true // LATIN CAPITAL LETTER R WITH CEDILLA
	case '\u0157':
		return "r", true // LATIN SMALL LETTER R WITH CEDILLA
	case '\u0158':
		return "R", true // LATIN CAPITAL LETTER R WITH CARON
	case '\u0159':
		return "r", true // LATIN SMALL LETTER R WITH CARON
	case '\u015A':
		return "S", true // LATIN CAPITAL LETTER S WITH ACUTE
	case '\u015B':
		return "s", true // LATIN SMALL LETTER S WITH ACUTE
	case '\u015C':
		return "S", true // LATIN CAPITAL LETTER S WITH CIRCUMFLEX
	case '\u015D':
		return "s", true // LATIN SMALL LETTER S WITH CIRCUMFLEX
	case '\u015E':
		return "S", true // LATIN CAPITAL LETTER S WITH CEDILLA
	case '\u015F':
		return "s", true // LATIN SMALL LETTER S WITH CEDILLA
	case '\u0160':
		return "S", true // LATIN CAPITAL LETTER S WITH CARON
	case '\u0161':
		return "s", true // LATIN SMALL LETTER S WITH CARON
	case '\u0162':
		return "T", true // LATIN CAPITAL LETTER T WITH CEDILLA
	case '\u0163':
		return "t", true // LATIN SMALL LETTER T WITH CEDILLA
	case '\u0164':
		return "T", true // LATIN CAPITAL LETTER T WITH CARON
	case '\u0165':
		return "t", true // LATIN SMALL LETTER T WITH CARON
	case '\u0168':
		return "U", true // LATIN CAPITAL LETTER U WITH TILDE
	case '\u0169':
		return "u", true // LATIN SMALL LETTER U WITH TILDE
	case '\u016A':
		return "U", true // LATIN CAPITAL LETTER U WITH MACRON
	case '\u016B':
		return "u", true // LATIN SMALL LETTER U WITH MACRON
	case '\u016C':
		return "U", true // LATIN CAPITAL LETTER U WITH BREVE
	case '\u016D':
		return "u", true // LATIN SMALL LETTER U WITH BREVE
	case '\u016E':
		return "U", true // LATIN CAPITAL LETTER U WITH RING ABOVE
	case '\u016F':
		return "u", true // LATIN SMALL LETTER U WITH RING ABOVE
	case '\u0170':
		return "U", true // LATIN CAPITAL LETTER U WITH DOUBLE ACUTE
	case '\u0171':
		return "u", true // LATIN SMALL LETTER U WITH DOUBLE ACUTE
	case '\u0172':
		return "U", true // LATIN CAPITAL LETTER U WITH OGONEK
	case '\u0173':
		return "u", true // LATIN SMALL LETTER U WITH OGONEK
	case '\u0174':
		return "W", true // LATIN CAPITAL LETTER W WITH CIRCUMFLEX
	case '\u0175':
		return "w", true // LATIN SMALL LETTER W WITH CIRCUMFLEX
	case '\u0176':
		return "Y", true // LATIN CAPITAL LETTER Y WITH CIRCUMFLEX
	case '\u0177':
		return "y", true // LATIN SMALL LETTER Y WITH CIRCUMFLEX
	case '\u0178':
		return "Y", true // LATIN CAPITAL LETTER Y WITH DIAERESIS
	case '\u0179':
		return "Z", true // LATIN CAPITAL LETTER Z WITH ACUTE
	case '\u017A':
		return "z", true // LATIN SMALL LETTER Z WITH ACUTE
	case '\u017B':
		return "Z", true // LATIN CAPITAL LETTER Z WITH DOT ABOVE
	case '\u017C':
		return "z", true // LATIN SMALL LETTER Z WITH DOT ABOVE
	case '\u017D':
		return "Z", true // LATIN CAPITAL LETTER Z WITH CARON
	case '\u017E':
		return "z", true // LATIN SMALL LETTER Z WITH CARON
	case '\u017F':
		return "s", true // LATIN SMALL LETTER LONG S
	case '\u01A0':
		return "O", true // LATIN CAPITAL LETTER O WITH HORN
	case '\u01A1':
		return "o", true // LATIN SMALL LETTER O WITH HORN
	case '\u01AF':
		return "U", true // LATIN CAPITAL LETTER U WITH HORN
	case '\u01B0':
		return "u", true // LATIN SMALL LETTER U WITH HORN
	case '\u01C4':
		return "DZ", true // LATIN CAPITAL LETTER DZ WITH CARON
	case '\u01C5':
		return "Dz", true // LATIN CAPITAL LETTER D WITH SMALL LETTER Z WITH CARON
	case '\u01C6':
		return "dz", true // LATIN SMALL LETTER DZ WITH CARON
	case '\u01C7':
		return "LJ", true // LATIN CAPITAL LETTER LJ
	case '\u01C8':
		return "Lj", true // LATIN CAPITAL LETTER L WITH SMALL LETTER J
	case '\u01C9':
		return "lj", true // LATIN SMALL LETTER LJ
	case '\u01CA':
		return "NJ", true // LATIN CAPITAL LETTER NJ
	case '\u01CB':
		return "Nj", true // LATIN CAPITAL LETTER N WITH SMALL LETTER J
	case '\u01CC':
		return "nj", true // LATIN SMALL LETTER NJ
	case '\u01CD':
		return "A", true // LATIN CAPITAL LETTER A WITH CARON
	case '\u01CE':
		return "a", true // LATIN SMALL LETTER A WITH CARON
	case '\u01CF':
		return "I", true // LATIN CAPITAL LETTER I WITH CARON
	case '\u01D0':
		return "i", true // LATIN SMALL LETTER I WITH CARON
	case '\u01D1':
		return "O", true // LATIN CAPITAL LETTER O WITH CARON
	case '\u01D2':
		return "o", true // LATIN SMALL LETTER O WITH CARON
	case '\u01D3':
		return "U", true // LATIN CAPITAL LETTER U WITH CARON
	case '\u01D4':
		return "u", true // LATIN SMALL LETTER U WITH CARON
	case '\u01D5':
		return "U", true // LATIN CAPITAL LETTER U WITH DIAERESIS AND MACRON
	case '\u01D6':
		return "u", true // LATIN SMALL LETTER U WITH DIAERESIS AND MACRON
	case '\u01D7':
		return "U", true // LATIN CAPITAL LETTER U WITH DIAERESIS AND ACUTE
	case '\u01D8':
		return "u", true // LATIN SMALL LETTER U WITH DIAERESIS AND ACUTE
	case '\u01D9':
		return "U", true // LATIN CAPITAL LETTER U WITH DIAERESIS AND CARON
	case '\u01DA':
		return "u", true // LATIN SMALL LETTER U WITH DIAERESIS AND CARON
	case '\u01DB':
		return "U", true // LATIN CAPITAL LETTER U WITH DIAERESIS AND GRAVE
	case '\u01DC':
		return "u", true // LATIN SMALL LETTER U WITH DIAERESIS AND GRAVE
	case '\u01DE':
		return "A", true // LATIN CAPITAL LETTER A WITH DIAERESIS AND MACRON
	case '\u01DF':
		return "a", true // LATIN SMALL LETTER A WITH DIAERESIS AND MACRON
	case '\u01E0':
		return "A", true // LATIN CAPITAL LETTER A WITH DOT ABOVE AND MACRON
	case '\u01E1':
		return "a", true // LATIN SMALL LETTER A WITH DOT ABOVE AND MACRON
	case '\u01E6':
		return "G", true // LATIN CAPITAL LETTER G WITH CARON
	case '\u01E7':
		return "g", true // LATIN SMALL LETTER G WITH CARON
	case '\u01E8':
		return "K", true // LATIN CAPITAL LETTER K WITH CARON
	case '\u01E9':
		return "k", true // LATIN SMALL LETTER K WITH CARON
	case '\u01EA':
		return "O", true // LATIN CAPITAL LETTER O WITH OGONEK
	case '\u01EB':
		return "o", true // LATIN SMALL LETTER O WITH OGONEK
	case '\u01EC':
		return "O", true // LATIN CAPITAL LETTER O WITH OGONEK AND MACRON
	case '\u01ED':
		return "o", true // LATIN SMALL LETTER O WITH OGONEK AND MACRON
	case '\u01F0':
		return "j", true // LATIN SMALL LETTER J WITH CARON
	case '\u01F1':
		return "DZ", true // LATIN CAPITAL LETTER DZ
	case '\u01F2':
		return "Dz", true // LATIN CAPITAL LETTER D WITH SMALL LETTER Z
	case '\u01F3':
		return "dz", true // LATIN SMALL LETTER DZ
	case '\u01F4':
		return "G", true // LATIN CAPITAL LETTER G WITH ACUTE
	case '\u01F5':
		return "g", true // LATIN SMALL LETTER G WITH ACUTE
	case '\u01F8':
		return "N", true // LATIN CAPITAL LETTER N WITH GRAVE
	case '\u01F9':
		return "n", true // LATIN SMALL LETTER N WITH GRAVE
	case '\u01FA':
		return "A", true // LATIN CAPITAL LETTER A WITH RING ABOVE AND ACUTE
	case '\u01FB':
		return "a", true // LATIN SMALL LETTER A WITH RING ABOVE AND ACUTE
	case '\u0200':
		return "A", true // LATIN CAPITAL LETTER A WITH DOUBLE GRAVE
	case '\u0201':
		return "a", true // LATIN SMALL LETTER A WITH DOUBLE GRAVE
	case '\u0202':
		return "A", true // LATIN CAPITAL LETTER A WITH INVERTED BREVE
	case '\u0203':
		return "a", true // LATIN SMALL LETTER A WITH INVERTED BREVE
	case '\u0204':
		return "E", true // LATIN CAPITAL LETTER E WITH DOUBLE GRAVE
	case '\u0205':
		return "e", true // LATIN SMALL LETTER E WITH DOUBLE GRAVE
	case '\u0206':
		return "E", true // LATIN CAPITAL LETTER E WITH INVERTED BREVE
	case '\u0207':
		return "e", true // LATIN SMALL LETTER E WITH INVERTED BREVE
	case '\u0208':
		return "I", true // LATIN CAPITAL LETTER I WITH DOUBLE GRAVE
	case '\u0209':
		return "i", true // LATIN SMALL LETTER I WITH DOUBLE GRAVE
	case '\u020A':
		return "I", true // LATIN CAPITAL LETTER I WITH INVERTED BREVE
	case '\u020B':
		return "i", true // LATIN SMALL LETTER I WITH INVERTED BREVE
	case '\u020C':
		return "O", true // LATIN CAPITAL LETTER O WITH DOUBLE GRAVE
	case '\u020D':
		return "o", true // LATIN SMALL LETTER O WITH DOUBLE GRAVE
	case '\u020E':
		return "O", true // LATIN CAPITAL LETTER O WITH INVERTED BREVE
	case '\u020F':
		return "o", true // LATIN SMALL LETTER O WITH INVERTED BREVE
	case '\u0210':
		return "R", true // LATIN CAPITAL LETTER R WITH DOUBLE GRAVE
	case '\u0211':
		return "r", true // LATIN SMALL LETTER R WITH DOUBLE GRAVE
	case '\u0212':
		return "R", true // LATIN CAPITAL LETTER R WITH INVERTED BREVE
	case '\u0213':
		return "r", true // LATIN SMALL LETTER R WITH INVERTED BREVE
	case '\u0214':
		return "U", true // LATIN CAPITAL LETTER U WITH DOUBLE GRAVE
	case '\u0215':
		return "u", true // LATIN SMALL LETTER U WITH DOUBLE GRAVE
	case '\u0216':
		return "U", true // LATIN CAPITAL LETTER U WITH INVERTED BREVE
	case '\u0217':
		return "u", true // LATIN SMALL LETTER U WITH INVERTED BREVE
	case '\u0218':
		return "S", true // LATIN CAPITAL LETTER S WITH COMMA BELOW
	case '\u0219':
		return "s", true // LATIN SMALL LETTER S WITH COMMA BELOW
	case '\u021A':
		return "T", true // LATIN CAPITAL LETTER T WITH COMMA BELOW
	case '\u021B':
		return "t", true // LATIN SMALL LETTER T WITH COMMA BELOW
	case '\u021E':
		return "H", true // LATIN CAPITAL LETTER H WITH CARON
	case '\u021F':
		return "h", true // LATIN SMALL LETTER H WITH CARON
	case '\u0226':
		return "A", true // LATIN CAPITAL LETTER A WITH DOT ABOVE
	case '\u0227':
		return "a", true // LATIN SMALL LETTER A WITH DOT ABOVE
	case '\u0228':
		return "E", true // LATIN CAPITAL LETTER E WITH CEDILLA
	case '\u0229':
		return "e", true // LATIN SMALL LETTER E WITH CEDILLA
	case '\u022A':
		return "O", true // LATIN CAPITAL LETTER O WITH DIAERESIS AND MACRON
	case '\u022B':
		return "o", true // LATIN SMALL LETTER O WITH DIAERESIS AND MACRON
	case '\u022C':
		return "O", true // LATIN CAPITAL LETTER O WITH TILDE AND MACRON
	case '\u022D':
		return "o", true // LATIN SMALL LETTER O WITH TILDE AND MACRON
	case '\u022E':
		return "O", true // LATIN CAPITAL LETTER O WITH DOT ABOVE
	case '\u022F':
		return "o", true // LATIN SMALL LETTER O WITH DOT ABOVE
	case '\u0230':
		return "O", true // LATIN CAPITAL LETTER O WITH DOT ABOVE AND MACRON
	case '\u0231':
		return "o", true // LATIN SMALL LETTER O WITH DOT ABOVE AND MACRON
	case '\u0232':
		return "Y", true // LATIN CAPITAL LETTER Y WITH MACRON
	case '\u0233':
		return "y", true // LATIN SMALL LETTER Y WITH MACRON
	case '\u02B0':
		return "h", true // MODIFIER LETTER SMALL H
	case '\u02B2':
		return "j", true // MODIFIER LETTER SMALL J
	case '\u02B3':
		return "r", true // MODIFIER LETTER SMALL R
	case '\u02B7':
		return "w", true // MODIFIER LETTER SMALL W
	case '\u02B8':
		return "y", true // MODIFIER LETTER SMALL Y
	case '\u02E1':
		return "l", true // MODIFIER LETTER SMALL L
	case '\u02E2':
		return "s", true // MODIFIER LETTER SMALL S
	case '\u02E3':
		return "x", true // MODIFIER LETTER SMALL X
	case '\u037E':
		return ";", true // GREEK QUESTION MARK
	case '\u1D2C':
		return "A", true // MODIFIER LETTER CAPITAL A
	case '\u1D2E':
		return "B", true // MODIFIER LETTER CAPITAL B
	case '\u1D30':
		return "D", true // MODIFIER LETTER CAPITAL D
	case '\u1D31':
		return "E", true // MODIFIER LETTER CAPITAL E
	case '\u1D33':
		return "G", true // MODIFIER LETTER CAPITAL G
	case '\u1D34':
		return "H", true // MODIFIER LETTER CAPITAL H
	case '\u1D35':
		return "I", true // MODIFIER LETTER CAPITAL I
	case '\u1D36':
		return "J", true // MODIFIER LETTER CAPITAL J
	case '\u1D37':
		return "K", true // MODIFIER LETTER CAPITAL K
	case '\u1D38':
		return "L", true // MODIFIER LETTER CAPITAL L
	case '\u1D39':
		return "M", true // MODIFIER LETTER CAPITAL M
	case '\u1D3A':
		return "N", true // MODIFIER LETTER CAPITAL N
	case '\u1D3C':
		return "O", true // MODIFIER LETTER CAPITAL O
	case '\u1D3E':
		return "P", true // MODIFIER LETTER CAPITAL P
	case '\u1D3F':
		return "R", true // MODIFIER LETTER CAPITAL R
	case '\u1D40':
		return "T", true // MODIFIER LETTER CAPITAL T
	case '\u1D41':
		return "U", true // MODIFIER LETTER CAPITAL U
	case '\u1D42':
		return "W", true // MODIFIER LETTER CAPITAL W
	case '\u1D43':
		return "a", true // MODIFIER LETTER SMALL A
	case '\u1D47':
		return "b", true // MODIFIER LETTER SMALL B
	case '\u1D48':
		return "d", true // MODIFIER LETTER SMALL D
	case '\u1D49':
		return "e", true // MODIFIER LETTER SMALL E
	case '\u1D4D':
		return "g", true // MODIFIER LETTER SMALL G
	case '\u1D4F':
		return "k", true // MODIFIER LETTER SMALL K
	case '\u1D50':
		return "m", true // MODIFIER LETTER SMALL M
	case '\u1D52':
		return "o", true // MODIFIER LETTER SMALL O
	case '\u1D56':
		return "p", true // MODIFIER LETTER SMALL P
	case '\u1D57':
		return "t", true // MODIFIER LETTER SMALL T
	case '\u1D58':
		return "u", true // MODIFIER LETTER SMALL U
	case '\u1D5B':
		return "v", true // MODIFIER LETTER SMALL V
	case '\u1D62':
		return "i", true // LATIN SUBSCRIPT SMALL LETTER I
	case '\u1D63':
		return "r", true // LATIN SUBSCRIPT SMALL LETTER R
	case '\u1D64':
		return "u", true // LATIN SUBSCRIPT SMALL LETTER U
	case '\u1D65':
		return "v", true // LATIN SUBSCRIPT SMALL LETTER V
	case '\u1D9C':
		return "c", true // MODIFIER LETTER SMALL C
	case '\u1DA0':
		return "f", true // MODIFIER LETTER SMALL F
	case '\u1DBB':
		return "z", true // MODIFIER LETTER SMALL Z
	case '\u1E00':
		return "A", true // LATIN CAPITAL LETTER A WITH RING BELOW
	case '\u1E01':
		return "a", true // LATIN SMALL LETTER A WITH RING BELOW
	case '\u1E02':
		return "B", true // LATIN CAPITAL LETTER B WITH DOT ABOVE
	case '\u1E03':
		return "b", true // LATIN SMALL LETTER B WITH DOT ABOVE
	case '\u1E04':
		return "B", true // LATIN CAPITAL LETTER B WITH DOT BELOW
	case '\u1E05':
		return "b", true // LATIN SMALL LETTER B WITH DOT BELOW
	case '\u1E06':
		return "B", true // LATIN CAPITAL LETTER B WITH LINE BELOW
	case '\u1E07':
		return "b", true // LATIN SMALL LETTER B WITH LINE BELOW
	case '\u1E08':
		return "C", true // LATIN CAPITAL LETTER C WITH CEDILLA AND ACUTE
	case '\u1E09':
		return "c", true // LATIN SMALL LETTER C WITH CEDILLA AND ACUTE
	case '\u1E0A':
		return "D", true // LATIN CAPITAL LETTER D WITH DOT ABOVE
	case '\u1E0B':
		return "d", true // LATIN SMALL LETTER D WITH DOT ABOVE
	case '\u1E0C':
		return "D", true // LATIN CAPITAL LETTER D WITH DOT BELOW
	case '\u1E0D':
		return "d", true // LATIN SMALL LETTER D WITH DOT BELOW
	case '\u1E0E':
		return "D", true // LATIN CAPITAL LETTER D WITH LINE BELOW
	case '\u1E0F':
		return "d", true // LATIN SMALL LETTER D WITH LINE BELOW
	case '\u1E10':
		return "D", true // LATIN CAPITAL LETTER D WITH CEDILLA
	case '\u1E11':
		return "d", true // LATIN SMALL LETTER D WITH CEDILLA
	case '\u1E12':
		return "D", true // LATIN CAPITAL LETTER D WITH CIRCUMFLEX BELOW
	case '\u1E13':
		return "d", true // LATIN SMALL LETTER D WITH CIRCUMFLEX BELOW
	case '\u1E14':
		return "E", true // LATIN CAPITAL LETTER E WITH MACRON AND GRAVE
	case '\u1E15':
		return "e", true // LATIN SMALL LETTER E WITH MACRON AND GRAVE
	case '\u1E16':
		return "E", true // LATIN CAPITAL LETTER E WITH MACRON AND ACUTE
	case '\u1E17':
		return "e", true // LATIN SMALL LETTER E WITH MACRON AND ACUTE
	case '\u1E18':
		return "E", true // LATIN CAPITAL LETTER E WITH CIRCUMFLEX BELOW
	case '\u1E19':
		return "e", true // LATIN SMALL LETTER E WITH CIRCUMFLEX BELOW
	case '\u1E1A':
		return "E", true // LATIN CAPITAL LETTER E WITH TILDE BELOW
	case '\u1E1B':
		return "e", true // LATIN SMALL LETTER E WITH TILDE BELOW
	case '\u1E1C':
		return "E", true // LATIN CAPITAL LETTER E WITH CEDILLA AND BREVE
	case '\u1E1D':
		return "e", true // LATIN SMALL LETTER E WITH CEDILLA AND BREVE
	case '\u1E1E':
		return "F", true // LATIN CAPITAL LETTER F WITH DOT ABOVE
	case '\u1E1F':
		return "f", true // LATIN SMALL LETTER F WITH DOT ABOVE
	case '\u1E20':
		return "G", true // LATIN CAPITAL LETTER G WITH MACRON
	case '\u1E21':
		return "g", true // LATIN SMALL LETTER G WITH MACRON
	case '\u1E22':
		return "H", true // LATIN CAPITAL LETTER H WITH DOT ABOVE
	case '\u1E23':
		return "h", true // LATIN SMALL LETTER H WITH DOT ABOVE
	case '\u1E24':
		return "H", true // LATIN CAPITAL LETTER H WITH DOT BELOW
	case '\u1E25':
		return "h", true // LATIN SMALL LETTER H WITH DOT BELOW
	case '\u1E26':
		return "H", true // LATIN CAPITAL LETTER H WITH DIAERESIS
	case '\u1E27':
		return "h", true // LATIN SMALL LETTER H WITH DIAERESIS
	case '\u1E28':
		return "H", true // LATIN CAPITAL LETTER H WITH CEDILLA
	case '\u1E29':
		return "h", true // LATIN SMALL LETTER H WITH CEDILLA
	case '\u1E2A':
		return "H", true // LATIN CAPITAL LETTER H WITH BREVE BELOW
	case '\u1E2B':
		return "h", true // LATIN SMALL LETTER H WITH BREVE BELOW
	case '\u1E2C':
		return "I", true // LATIN CAPITAL LETTER I WITH TILDE BELOW
	case '\u1E2D':
		return "i", true // LATIN SMALL LETTER I WITH TILDE BELOW
	case '\u1E2E':
		return "I", true // LATIN CAPITAL LETTER I WITH DIAERESIS AND ACUTE
	case '\u1E2F':
		return "i", true // LATIN SMALL LETTER I WITH DIAERESIS AND ACUTE
	case '\u1E30':
		return "K", true // LATIN CAPITAL LETTER K WITH ACUTE
	case '\u1E31':
		return "k", true // LATIN SMALL LETTER K WITH ACUTE
	case '\u1E32':
		return "K", true // LATIN CAPITAL LETTER K WITH DOT BELOW
	case '\u1E33':
		return "k", true // LATIN SMALL LETTER K WITH DOT BELOW
	case '\u1E34':
		return "K", true // LATIN CAPITAL LETTER K WITH LINE BELOW
	case '\u1E35':
		return "k", true // LATIN SMALL LETTER K WITH LINE BELOW
	case '\u1E36':
		return "L", true // LATIN CAPITAL LETTER L WITH DOT BELOW
	case '\u1E37':
		return "l", true // LATIN SMALL LETTER L WITH DOT BELOW
	case '\u1E38':
		return "L", true // LATIN CAPITAL LETTER L WITH DOT BELOW AND MACRON
	case '\u1E39':
		return "l", true // LATIN SMALL LETTER L WITH DOT BELOW AND MACRON
	case '\u1E3A':
		return "L", true // LATIN CAPITAL LETTER L WITH LINE BELOW
	case '\u1E3B':
		return "l", true // LATIN SMALL LETTER L WITH LINE BELOW
	case '\u1E3C':
		return "L", true // LATIN CAPITAL LETTER L WITH CIRCUMFLEX BELOW
	case '\u1E3D':
		return "l", true // LATIN SMALL LETTER L WITH CIRCUMFLEX BELOW
	case '\u1E3E':
		return "M", true // LATIN CAPITAL LETTER M WITH ACUTE
	case '\u1E3F':
		return "m", true // LATIN SMALL LETTER M WITH ACUTE
	case '\u1E40':
		return "M", true // LATIN CAPITAL LETTER M WITH DOT ABOVE
	case '\u1E41':
		return "m", true // LATIN SMALL LETTER M WITH DOT ABOVE
	case '\u1E42':
		return "M", true // LATIN CAPITAL LETTER M WITH DOT BELOW
	case '\u1E43':
		return "m", true // LATIN SMALL LETTER M WITH DOT BELOW
	case '\u1E44':
		return "N", true // LATIN CAPITAL LETTER N WITH DOT ABOVE
	case '\u1E45':
		return "n", true // LATIN SMALL LETTER N WITH DOT ABOVE
	case '\u1E46':
		return "N", true // LATIN CAPITAL LETTER N WITH DOT BELOW
	case '\u1E47':
		return "n", true // LATIN SMALL LETTER N WITH DOT BELOW
	case '\u1E48':
		return "N", true // LATIN CAPITAL LETTER N WITH LINE BELOW
	case '\u1E49':
		return "n", true // LATIN SMALL LETTER N WITH LINE BELOW
	case '\u1E4A':
		return "N", true // LATIN CAPITAL LETTER N WITH CIRCUMFLEX BELOW
	case '\u1E4B':
		return "n", true // LATIN SMALL LETTER N WITH CIRCUMFLEX BELOW
	case '\u1E4C':
		return "O", true // LATIN CAPITAL LETTER O WITH TILDE AND ACUTE
	case '\u1E4D':
		return "o", true // LATIN SMALL LETTER O WITH TILDE AND ACUTE
	case '\u1E4E':
		return "O", true // LATIN CAPITAL LETTER O WITH TILDE AND DIAERESIS
	case '\u1E4F':
		return "o", true // LATIN SMALL LETTER O WITH TILDE AND DIAERESIS
	case '\u1E50':
		return "O", true // LATIN CAPITAL LETTER O WITH MACRON AND GRAVE
	case '\u1E51':
		return "o", true // LATIN SMALL LETTER O WITH MACRON AND GRAVE
	case '\u1E52':
		return "O", true // LATIN CAPITAL LETTER O WITH MACRON AND ACUTE
	case '\u1E53':
		return "o", true // LATIN SMALL LETTER O WITH MACRON AND ACUTE
	case '\u1E54':
		return "P", true // LATIN CAPITAL LETTER P WITH ACUTE
	case '\u1E55':
		return "p", true // LATIN SMALL LETTER P WITH ACUTE
	case '\u1E56':
		return "P", true // LATIN CAPITAL LETTER P WITH DOT ABOVE
	case '\u1E57':
		return "p", true // LATIN SMALL LETTER P WITH DOT ABOVE
	case '\u1E58':
		return "R", true // LATIN CAPITAL LETTER R WITH DOT ABOVE
	case '\u1E59':
		return "r", true // LATIN SMALL LETTER R WITH DOT ABOVE
	case '\u1E5A':
		return "R", true // LATIN CAPITAL LETTER R WITH DOT BELOW
	case '\u1E5B':
		return "r", true // LATIN SMALL LETTER R WITH DOT BELOW
	case '\u1E5C':
		return "R", true // LATIN CAPITAL LETTER R WITH DOT BELOW AND MACRON
	case '\u1E5D':
		return "r", true // LATIN SMALL LETTER R WITH DOT BELOW AND MACRON
	case '\u1E5E':
		return "R", true // LATIN CAPITAL LETTER R WITH LINE BELOW
	case '\u1E5F':
		return "r", true // LATIN SMALL LETTER R WITH LINE BELOW
	case '\u1E60':
		return "S", true // LATIN CAPITAL LETTER S WITH DOT ABOVE
	case '\u1E61':
		return "s", true // LATIN SMALL LETTER S WITH DOT ABOVE
	case '\u1E62':
		return "S", true // LATIN CAPITAL LETTER S WITH DOT BELOW
	case '\u1E63':
		return "s", true // LATIN SMALL LETTER S WITH DOT BELOW
	case '\u1E64':
		return "S", true // LATIN CAPITAL LETTER S WITH ACUTE AND DOT ABOVE
	case '\u1E65':
		return "s", true // LATIN SMALL LETTER S WITH ACUTE AND DOT ABOVE
	case '\u1E66':
		return "S", true // LATIN CAPITAL LETTER S WITH CARON AND DOT ABOVE
	case '\u1E67':
		return "s", true // LATIN SMALL LETTER S WITH CARON AND DOT ABOVE
	case '\u1E68':
		return "S", true // LATIN CAPITAL LETTER S WITH DOT BELOW AND DOT ABOVE
	case '\u1E69':
		return "s", true // LATIN SMALL LETTER S WITH DOT BELOW AND DOT ABOVE
	case '\u1E6A':
		return "T", true // LATIN CAPITAL LETTER T WITH DOT ABOVE
	case '\u1E6B':
		return "t", true // LATIN SMALL LETTER T WITH DOT ABOVE
	case '\u1E6C':
		return "T", true // LATIN CAPITAL LETTER T WITH DOT BELOW
	case '\u1E6D':
		return "t", true // LATIN SMALL LETTER T WITH DOT BELOW
	case '\u1E6E':
		return "T", true // LATIN CAPITAL LETTER T WITH LINE BELOW
	case '\u1E6F':
		return "t", true // LATIN SMALL LETTER T WITH LINE BELOW
	case '\u1E70':
		return "T", true // LATIN CAPITAL LETTER T WITH CIRCUMFLEX BELOW
	case '\u1E71':
		return "t", true // LATIN SMALL LETTER T WITH CIRCUMFLEX BELOW
	case '\u1E72':
		return "U", true // LATIN CAPITAL LETTER U WITH DIAERESIS BELOW
	case '\u1E73':
		return "u", true // LATIN SMALL LETTER U WITH DIAERESIS BELOW
	case '\u1E74':
		return "U", true // LATIN CAPITAL LETTER U WITH TILDE BELOW
	case '\u1E75':
		return "u", true // LATIN SMALL LETTER U WITH TILDE BELOW
	case '\u1E76':
		return "U", true // LATIN CAPITAL LETTER U WITH CIRCUMFLEX BELOW
	case '\u1E77':
		return "u", true // LATIN SMALL LETTER U WITH CIRCUMFLEX BELOW
	case '\u1E78':
		return "U", true // LATIN CAPITAL LETTER U WITH TILDE AND ACUTE
	case '\u1E79':
		return "u", true // LATIN SMALL LETTER U WITH TILDE AND ACUTE
	case '\u1E7A':
		return "U", true // LATIN CAPITAL LETTER U WITH MACRON AND DIAERESIS
	case '\u1E7B':
		return "u", true // LATIN SMALL LETTER U WITH MACRON AND DIAERESIS
	case '\u1E7C':
		return "V", true // LATIN CAPITAL LETTER V WITH TILDE
	case '\u1E7D':
		return "v", true // LATIN SMALL LETTER V WITH TILDE
	case '\u1E7E':
		return "V", true // LATIN CAPITAL LETTER V WITH DOT BELOW
	case '\u1E7F':
		return "v", true // LATIN SMALL LETTER V WITH DOT BELOW
	case '\u1E80':
		return "W", true // LATIN CAPITAL LETTER W WITH GRAVE
	case '\u1E81':
		return "w", true // LATIN SMALL LETTER W WITH GRAVE
	case '\u1E82':
		return "W", true // LATIN CAPITAL LETTER W WITH ACUTE
	case '\u1E83':
		return "w", true // LATIN SMALL LETTER W WITH ACUTE
	case '\u1E84':
		return "W", true // LATIN CAPITAL LETTER W WITH DIAERESIS
	case '\u1E85':
		return "w", true // LATIN SMALL LETTER W WITH DIAERESIS
	case '\u1E86':
		return "W", true // LATIN CAPITAL LETTER W WITH DOT ABOVE
	case '\u1E87':
		return "w", true // LATIN SMALL LETTER W WITH DOT ABOVE
	case '\u1E88':
		return "W", true // LATIN CAPITAL LETTER W WITH DOT BELOW
	case '\u1E89':
		return "w", true // LATIN SMALL LETTER W WITH DOT BELOW
	case '\u1E8A':
		return "X", true // LATIN CAPITAL LETTER X WITH DOT ABOVE
	case '\u1E8B':
		return "x", true // LATIN SMALL LETTER X WITH DOT ABOVE
	case '\u1E8C':
		return "X", true // LATIN CAPITAL LETTER X WITH DIAERESIS
	case '\u1E8D':
		return "x", true // LATIN SMALL LETTER X WITH DIAERESIS
	case '\u1E8E':
		return "Y", true // LATIN CAPITAL LETTER Y WITH DOT ABOVE
	case '\u1E8F':
		return "y", true // LATIN SMALL LETTER Y WITH DOT ABOVE
	case '\u1E90':
		return "Z", true // LATIN CAPITAL LETTER Z WITH CIRCUMFLEX
	case '\u1E91':
		return "z", true // LATIN SMALL LETTER Z WITH CIRCUMFLEX
	case '\u1E92':
		return "Z", true // LATIN CAPITAL LETTER Z WITH DOT BELOW
	case '\u1E93':
		return "z", true // LATIN SMALL LETTER Z WITH DOT BELOW
	case '\u1E94':
		return "Z", true // LATIN CAPITAL LETTER Z WITH LINE BELOW
	case '\u1E95':
		return "z", true // LATIN SMALL LETTER Z WITH LINE BELOW
	case '\u1E96':
		return "h", true // LATIN SMALL LETTER H WITH LINE BELOW
	case '\u1E97':
		return "t", true // LATIN SMALL LETTER T WITH DIAERESIS
	case '\u1E98':
		return "w", true // LATIN SMALL LETTER W WITH RING ABOVE
	case '\u1E99':
		return "y", true // LATIN SMALL LETTER Y WITH RING ABOVE
	case '\u1E9B':
		return "s", true // LATIN SMALL LETTER LONG S WITH DOT ABOVE
	case '\u1EA0':
		return "A", true // LATIN CAPITAL LETTER A WITH DOT BELOW
	case '\u1EA1':
		return "a", true // LATIN SMALL LETTER A WITH DOT BELOW
	case '\u1EA2':
		return "A", true // LATIN CAPITAL LETTER A WITH HOOK ABOVE
	case '\u1EA3':
		return "a", true // LATIN SMALL LETTER A WITH HOOK ABOVE
	case '\u1EA4':
		return "A", true // LATIN CAPITAL LETTER A WITH CIRCUMFLEX AND ACUTE
	case '\u1EA5':
		return "a", true // LATIN SMALL LETTER A WITH CIRCUMFLEX AND ACUTE
	case '\u1EA6':
		return "A", true // LATIN CAPITAL LETTER A WITH CIRCUMFLEX AND GRAVE
	case '\u1EA7':
		return "a", true // LATIN SMALL LETTER A WITH CIRCUMFLEX AND GRAVE
	case '\u1EA8':
		return "A", true // LATIN CAPITAL LETTER A WITH CIRCUMFLEX AND HOOK ABOVE
	case '\u1EA9':
		return "a", true // LATIN SMALL LETTER A WITH CIRCUMFLEX AND HOOK ABOVE
	case '\u1EAA':
		return "A", true // LATIN CAPITAL LETTER A WITH CIRCUMFLEX AND TILDE
	case '\u1EAB':
		return "a", true // LATIN SMALL LETTER A WITH CIRCUMFLEX AND TILDE
	case '\u1EAC':
		return "A", true // LATIN CAPITAL LETTER A WITH CIRCUMFLEX AND DOT BELOW
	case '\u1EAD':
		return "a", true // LATIN SMALL LETTER A WITH CIRCUMFLEX AND DOT BELOW
	case '\u1EAE':
		return "A", true // LATIN CAPITAL LETTER A WITH BREVE AND ACUTE
	case '\u1EAF':
		return "a", true // LATIN SMALL LETTER A WITH BREVE AND ACUTE
	case '\u1EB0':
		return "A", true // LATIN CAPITAL LETTER A WITH BREVE AND GRAVE
	case '\u1EB1':
		return "a", true // LATIN SMALL LETTER A WITH BREVE AND GRAVE
	case '\u1EB2':
		return "A", true // LATIN CAPITAL LETTER A WITH BREVE AND HOOK ABOVE
	case '\u1EB3':
		return "a", true // LATIN SMALL LETTER A WITH BREVE AND HOOK ABOVE
	case '\u1EB4':
		return "A", true // LATIN CAPITAL LETTER A WITH BREVE AND TILDE
	case '\u1EB5':
		return "a", true // LATIN SMALL LETTER A WITH BREVE AND TILDE
	case '\u1EB6':
		return "A", true // LATIN CAPITAL LETTER A WITH BREVE AND DOT BELOW
	case '\u1EB7':
		return "a", true // LATIN SMALL LETTER A WITH BREVE AND DOT BELOW
	case '\u1EB8':
		return "E", true // LATIN CAPITAL LETTER E WITH DOT BELOW
	case '\u1EB9':
		return "e", true // LATIN SMALL LETTER E WITH DOT BELOW
	case '\u1EBA':
		return "E", true // LATIN CAPITAL LETTER E WITH HOOK ABOVE
	case '\u1EBB':
		return "e", true // LATIN SMALL LETTER E WITH HOOK ABOVE
	case '\u1EBC':
		return "E", true // LATIN CAPITAL LETTER E WITH TILDE
	case '\u1EBD':
		return "e", true // LATIN SMALL LETTER E WITH TILDE
	case '\u1EBE':
		return "E", true // LATIN CAPITAL LETTER E WITH CIRCUMFLEX AND ACUTE
	case '\u1EBF':
		return "e", true // LATIN SMALL LETTER E WITH CIRCUMFLEX AND ACUTE
	case '\u1EC0':
		return "E", true // LATIN CAPITAL LETTER E WITH CIRCUMFLEX AND GRAVE
	case '\u1EC1':
		return "e", true // LATIN SMALL LETTER E WITH CIRCUMFLEX AND GRAVE
	case '\u1EC2':
		return "E", true // LATIN CAPITAL LETTER E WITH CIRCUMFLEX AND HOOK ABOVE
	case '\u1EC3':
		return "e", true // LATIN SMALL LETTER E WITH CIRCUMFLEX AND HOOK ABOVE
	case '\u1EC4':
		return "E", true // LATIN CAPITAL LETTER E WITH CIRCUMFLEX AND TILDE
	case '\u1EC5':
		return "e", true // LATIN SMALL LETTER E WITH CIRCUMFLEX AND TILDE
	case '\u1EC6':
		return "E", true // LATIN CAPITAL LETTER E WITH CIRCUMFLEX AND DOT BELOW
	case '\u1EC7':
		return "e", true // LATIN SMALL LETTER E WITH CIRCUMFLEX AND DOT BELOW
	case '\u1EC8':
		return "I", true // LATIN CAPITAL LETTER I WITH HOOK ABOVE
	case '\u1EC9':
		return "i", true // LATIN SMALL LETTER I WITH HOOK ABOVE
	case '\u1ECA':
		return "I", true // LATIN CAPITAL LETTER I WITH DOT BELOW
	case '\u1ECB':
		return "i", true // LATIN SMALL LETTER I WITH DOT BELOW
	case '\u1ECC':
		return "O", true // LATIN CAPITAL LETTER O WITH DOT BELOW
	case '\u1ECD':
		return "o", true // LATIN SMALL LETTER O WITH DOT BELOW
	case '\u1ECE':
		return "O", true // LATIN CAPITAL LETTER O WITH HOOK ABOVE
	case '\u1ECF':
		return "o", true // LATIN SMALL LETTER O WITH HOOK ABOVE
	case '\u1ED0':
		return "O", true // LATIN CAPITAL LETTER O WITH CIRCUMFLEX AND ACUTE
	case '\u1ED1':
		return "o", true // LATIN SMALL LETTER O WITH CIRCUMFLEX AND ACUTE
	case '\u1ED2':
		return "O", true // LATIN CAPITAL LETTER O WITH CIRCUMFLEX AND GRAVE
	case '\u1ED3':
		return "o", true // LATIN SMALL LETTER O WITH CIRCUMFLEX AND GRAVE
	case '\u1ED4':
		return "O", true // LATIN CAPITAL LETTER O WITH CIRCUMFLEX AND HOOK ABOVE
	case '\u1ED5':
		return "o", true // LATIN SMALL LETTER O WITH CIRCUMFLEX AND HOOK ABOVE
	case '\u1ED6':
		return "O", true // LATIN CAPITAL LETTER O WITH CIRCUMFLEX AND TILDE
	case '\u1ED7':
		return "o", true // LATIN SMALL LETTER O WITH CIRCUMFLEX AND TILDE
	case '\u1ED8':
		return "O", true // LATIN CAPITAL LETTER O WITH CIRCUMFLEX AND DOT BELOW
	case '\u1ED9':
		return "o", true // LATIN SMALL LETTER O WITH CIRCUMFLEX AND DOT BELOW
	case '\u1EDA':
		return "O", true // LATIN CAPITAL LETTER O WITH HORN AND ACUTE
	case '\u1EDB':
		return "o", true // LATIN SMALL LETTER O WITH HORN AND ACUTE
	case '\u1EDC':
		return "O", true // LATIN CAPITAL LETTER O WITH HORN AND GRAVE
	case '\u1EDD':
		return "o", true // LATIN SMALL LETTER O WITH HORN AND GRAVE
	case '\u1EDE':
		return "O", true // LATIN CAPITAL LETTER O WITH HORN AND HOOK ABOVE
	case '\u1EDF':
		return "o", true // LATIN SMALL LETTER O WITH HORN AND HOOK ABOVE
	case '\u1EE0':
		return "O", true // LATIN CAPITAL LETTER O WITH HORN AND TILDE
	case '\u1EE1':
		return "o", true // LATIN SMALL LETTER O WITH HORN AND TILDE
	case '\u1EE2':
		return "O", true // LATIN CAPITAL LETTER O WITH HORN AND DOT BELOW
	case '\u1EE3':
		return "o", true // LATIN SMALL LETTER O WITH HORN AND DOT BELOW
	case '\u1EE4':
		return "U", true // LATIN CAPITAL LETTER U WITH DOT BELOW
	case '\u1EE5':
		return "u", true // LATIN SMALL LETTER U WITH DOT BELOW
	case '\u1EE6':
		return "U", true // LATIN CAPITAL LETTER U WITH HOOK ABOVE
	case '\u1EE7':
		return "u", true // LATIN SMALL LETTER U WITH HOOK ABOVE
	case '\u1EE8':
		return "U", true // LATIN CAPITAL LETTER U WITH HORN AND ACUTE
	case '\u1EE9':
		return "u", true // LATIN SMALL LETTER U WITH HORN AND ACUTE
	case '\u1EEA':
		return "U", true // LATIN CAPITAL LETTER U WITH HORN AND GRAVE
	case '\u1EEB':
		return "u", true // LATIN SMALL LETTER U WITH HORN AND GRAVE
	case '\u1EEC':
		return "U", true // LATIN CAPITAL LETTER U WITH HORN AND HOOK ABOVE
	case '\u1EED':
		return "u", true // LATIN SMALL LETTER U WITH HORN AND HOOK ABOVE
	case '\u1EEE':
		return "U", true // LATIN CAPITAL LETTER U WITH HORN AND TILDE
	case '\u1EEF':
		return "u", true // LATIN SMALL LETTER U WITH HORN AND TILDE
	case '\u1EF0':
		return "U", true // LATIN CAPITAL LETTER U WITH HORN AND DOT BELOW
	case '\u1EF1':
		return "u", true // LATIN SMALL LETTER U WITH HORN AND DOT BELOW
	case '\u1EF2':
		return "Y", true // LATIN CAPITAL LETTER Y WITH GRAVE
	case '\u1EF3':
		return "y", true // LATIN SMALL LETTER Y WITH GRAVE
	case '\u1EF4':
		return "Y", true // LATIN CAPITAL LETTER Y WITH DOT BELOW
	case '\u1EF5':
		return "y", true // LATIN SMALL LETTER Y WITH DOT BELOW
	case '\u1EF6':
		return "Y", true // LATIN CAPITAL LETTER Y WITH HOOK ABOVE
	case '\u1EF7':
		return "y", true // LATIN SMALL LETTER Y WITH HOOK ABOVE
	case '\u1EF8':
		return "Y", true // LATIN CAPITAL LETTER Y WITH TILDE
	case '\u1EF9':
		return "y", true // LATIN SMALL LETTER Y WITH TILDE
	case '\u2000':
		return " ", true // EN QUAD
	case '\u2001':
		return " ", true // EM QUAD
	case '\u2002':
		return " ", true // EN SPACE
	case '\u2003':
		return " ", true // EM SPACE
	case '\u2004':
		return " ", true // THREE-PER-EM SPACE
	case '\u2005':
		return " ", true // FOUR-PER-EM SPACE
	case '\u2006':
		return " ", true // SIX-PER-EM SPACE
	case '\u2007':
		return " ", true // FIGURE SPACE
	case '\u2008':
		return " ", true // PUNCTUATION SPACE
	case '\u2009':
		return " ", true // THIN SPACE
	case '\u200A':
		return " ", true // HAIR SPACE
	case '\u2024':
		return ".", true // ONE DOT LEADER
	case '\u2025':
		return "..", true // TWO DOT LEADER
	case '\u2026':
		return "...", true // HORIZONTAL ELLIPSIS
	case '\u202F':
		return " ", true // NARROW NO-BREAK SPACE
	case '\u203C':
		return "!!", true // DOUBLE EXCLAMATION MARK
	case '\u2047':
		return "??", true // DOUBLE QUESTION MARK
	case '\u2048':
		return "?!", true // QUESTION EXCLAMATION MARK
	case '\u2049':
		return "!?", true // EXCLAMATION QUESTION MARK
	case '\u205F':
		return " ", true // MEDIUM MATHEMATICAL SPACE
	case '\u2070':
		return "0", true // SUPERSCRIPT ZERO
	case '\u2071':
		return "i", true // SUPERSCRIPT LATIN SMALL LETTER I
	case '\u2074':
		return "4", true // SUPERSCRIPT FOUR
	case '\u2075':
		return "5", true // SUPERSCRIPT FIVE
	case '\u2076':
		return "6", true // SUPERSCRIPT SIX
	case '\u2077':
		return "7", true // SUPERSCRIPT SEVEN
	case '\u2078':
		return "8", true // SUPERSCRIPT EIGHT
	case '\u2079':
		return "9", true // SUPERSCRIPT NINE
	case '\u207A':
		return "+", true // SUPERSCRIPT PLUS SIGN
	case '\u207C':
		return "=", true // SUPERSCRIPT EQUALS SIGN
	case '\u207D':
		return "(", true // SUPERSCRIPT LEFT PARENTHESIS
	case '\u207E':
		return ")", true // SUPERSCRIPT RIGHT PARENTHESIS
	case '\u207F':
		return "n", true // SUPERSCRIPT LATIN SMALL LETTER N
	case '\u2080':
		return "0", true // SUBSCRIPT ZERO
	case '\u2081':
		return "1", true // SUBSCRIPT ONE
	case '\u2082':
		return "2", true // SUBSCRIPT TWO
	case '\u2083':
		return "3", true // SUBSCRIPT THREE
	case '\u2084':
		return "4", true // SUBSCRIPT FOUR
	case '\u2085':
		return "5", true // SUBSCRIPT FIVE
	case '\u2086':
		return "6", true // SUBSCRIPT SIX
	case '\u2087':
		return "7", true // SUBSCRIPT SEVEN
	case '\u2088':
		return "8", true // SUBSCRIPT EIGHT
	case '\u2089':
		return "9", true // SUBSCRIPT NINE
	case '\u208A':
		return "+", true // SUBSCRIPT PLUS SIGN
	case '\u208C':
		return "=", true // SUBSCRIPT EQUALS SIGN
	case '\u208D':
		return "(", true // SUBSCRIPT LEFT PARENTHESIS
	case '\u208E':
		return ")", true // SUBSCRIPT RIGHT PARENTHESIS
	case '\u2090':
		return "a", true // LATIN SUBSCRIPT SMALL LETTER A
	case '\u2091':
		return "e", true // LATIN SUBSCRIPT SMALL LETTER E
	case '\u2092':
		return "o", true // LATIN SUBSCRIPT SMALL LETTER O
	case '\u2093':
		return "x", true // LATIN SUBSCRIPT SMALL LETTER X
	case '\u2095':
		return "h", true // LATIN SUBSCRIPT SMALL LETTER H
	case '\u2096':
		return "k", true // LATIN SUBSCRIPT SMALL LETTER K
	case '\u2097':
		return "l", true // LATIN SUBSCRIPT SMALL LETTER L
	case '\u2098':
		return "m", true // LATIN SUBSCRIPT SMALL LETTER M
	case '\u2099':
		return "n", true // LATIN SUBSCRIPT SMALL LETTER N
	case '\u209A':
		return "p", true // LATIN SUBSCRIPT SMALL LETTER P
	case '\u209B':
		return "s", true // LATIN SUBSCRIPT SMALL LETTER S
	case '\u209C':
		return "t", true // LATIN SUBSCRIPT SMALL LETTER T
	case '\u20A8':
		return "Rs", true // RUPEE SIGN
	case '\u2100':
		return "a/c", true // ACCOUNT OF
	case '\u2101':
		return "a/s", true // ADDRESSED TO THE SUBJECT
	case '\u2102':
		return "C", true // DOUBLE-STRUCK CAPITAL C
	case '\u2105':
		return "c/o", true // CARE OF
	case '\u2106':
		return "c/u", true // CADA UNA
	case '\u210A':
		return "g", true // SCRIPT SMALL G
	case '\u210B':
		return "H", true // SCRIPT CAPITAL H
	case '\u210C':
		return "H", true // BLACK-LETTER CAPITAL H
	case '\u210D':
		return "H", true // DOUBLE-STRUCK CAPITAL H
	case '\u210E':
		return "h", true // PLANCK CONSTANT
	case '\u2110':
		return "I", true // SCRIPT CAPITAL I
	case '\u2111':
		return "I", true // BLACK-LETTER CAPITAL I
	case '\u2112':
		return "L", true // SCRIPT CAPITAL L
	case '\u2113':
		return "l", true // SCRIPT SMALL L
	case '\u2115':
		return "N", true // DOUBLE-STRUCK CAPITAL N
	case '\u2116':
		return "No", true // NUMERO SIGN
	case '\u2119':
		return "P", true // DOUBLE-STRUCK CAPITAL P
	case '\u211A':
		return "Q", true // DOUBLE-STRUCK CAPITAL Q
	case '\u211B':
		return "R", true // SCRIPT CAPITAL R
	case '\u211C':
		return "R", true // BLACK-LETTER CAPITAL R
	case '\u211D':
		return "R", true // DOUBLE-STRUCK CAPITAL R
	case '\u2120':
		return "SM", true // SERVICE MARK
	case '\u2121':
		return "TEL", true // TELEPHONE SIGN
	case '\u2122':
		return "TM", true // TRADE MARK SIGN
	case '\u2124':
		return "Z", true // DOUBLE-STRUCK CAPITAL Z
	case '\u2128':
		return "Z", true // BLACK-LETTER CAPITAL Z
	case '\u212A':
		return "K", true // KELVIN SIGN
	case '\u212B':
		return "A", true // ANGSTROM SIGN
	case '\u212C':
		return "B", true // SCRIPT CAPITAL B
	case '\u212D':
		return "C", true // BLACK-LETTER CAPITAL C
	case '\u212F':
		return "e", true // SCRIPT SMALL E
	case '\u2130':
		return "E", true // SCRIPT CAPITAL E
	case '\u2131':
		return "F", true // SCRIPT CAPITAL F
	case '\u2133':
		return "M", true // SCRIPT CAPITAL M
	case '\u2134':
		return "o", true // SCRIPT SMALL O
	case '\u2139':
		return "i", true // INFORMATION SOURCE
	case '\u213B':
		return "FAX", true // FACSIMILE SIGN
	case '\u2145':
		return "D", true // DOUBLE-STRUCK ITALIC CAPITAL D
	case '\u2146':
		return "d", true // DOUBLE-STRUCK ITALIC SMALL D
	case '\u2147':
		return "e", true // DOUBLE-STRUCK ITALIC SMALL E
	case '\u2148':
		return "i", true // DOUBLE-STRUCK ITALIC SMALL I
	case '\u2149':
		return "j", true // DOUBLE-STRUCK ITALIC SMALL J
	case '\u2160':
		return "I", true // ROMAN NUMERAL ONE
	case '\u2161':
		return "II", true // ROMAN NUMERAL TWO
	case '\u2162':
		return "III", true // ROMAN NUMERAL THREE
	case '\u2163':
		return "IV", true // ROMAN NUMERAL FOUR
	case '\u2164':
		return "V", true // ROMAN NUMERAL FIVE
	case '\u2165':
		return "VI", true // ROMAN NUMERAL SIX
	case '\u2166':
		return "VII", true // ROMAN NUMERAL SEVEN
	case '\u2167':
		return "VIII", true // ROMAN NUMERAL EIGHT
	case '\u2168':
		return "IX", true // ROMAN NUMERAL NINE
	case '\u2169':
		return "X", true // ROMAN NUMERAL TEN
	case '\u216A':
		return "XI", true // ROMAN NUMERAL ELEVEN
	case '\u216B':
		return "XII", true // ROMAN NUMERAL TWELVE
	case '\u216C':
		return "L", true // ROMAN NUMERAL FIFTY
	case '\u216D':
		return "C", true // ROMAN NUMERAL ONE HUNDRED
	case '\u216E':
		return "D", true // ROMAN NUMERAL FIVE HUNDRED
	case '\u216F':
		return "M", true // ROMAN NUMERAL ONE THOUSAND
	case '\u2170':
		return "i", true // SMALL ROMAN NUMERAL ONE
	case '\u2171':
		return "ii", true // SMALL ROMAN NUMERAL TWO
	case '\u2172':
		return "iii", true // SMALL ROMAN NUMERAL THREE
	case '\u2173':
		return "iv", true // SMALL ROMAN NUMERAL FOUR
	case '\u2174':
		return "v", true // SMALL ROMAN NUMERAL FIVE
	case '\u2175':
		return "vi", true // SMALL ROMAN NUMERAL SIX
	case '\u2176':
		return "vii", true // SMALL ROMAN NUMERAL SEVEN
	case '\u2177':
		return "viii", true // SMALL ROMAN NUMERAL EIGHT
	case '\u2178':
		return "ix", true // SMALL ROMAN NUMERAL NINE
	case '\u2179':
		return "x", true // SMALL ROMAN NUMERAL TEN
	case '\u217A':
		return "xi", true // SMALL ROMAN NUMERAL ELEVEN
	case '\u217B':
		return "xii", true // SMALL ROMAN NUMERAL TWELVE
	case '\u217C':
		return "l", true // SMALL ROMAN NUMERAL FIFTY
	case '\u217D':
		return "c", true // SMALL ROMAN NUMERAL ONE HUNDRED
	case '\u217E':
		return "d", true // SMALL ROMAN NUMERAL FIVE HUNDRED
	case '\u217F':
		return "m", true // SMALL ROMAN NUMERAL ONE THOUSAND
	case '\u2260':
		return "=", true // NOT EQUAL TO
	case '\u226E':
		return "<", true // NOT LESS-THAN
	case '\u226F':
		return ">", true // NOT GREATER-THAN
	case '\u2460':
		return "1", true // CIRCLED DIGIT ONE
	case '\u2461':
		return "2", true // CIRCLED DIGIT TWO
	case '\u2462':
		return "3", true // CIRCLED DIGIT THREE
	case '\u2463':
		return "4", true // CIRCLED DIGIT FOUR
	case '\u2464':
		return "5", true // CIRCLED DIGIT FIVE
	case '\u2465':
		return "6", true // CIRCLED DIGIT SIX
	case '\u2466':
		return "7", true // CIRCLED DIGIT SEVEN
	case '\u2467':
		return "8", true // CIRCLED DIGIT EIGHT
	case '\u2468':
		return "9", true // CIRCLED DIGIT NINE
	case '\u2469':
		return "10", true // CIRCLED NUMBER TEN
	case '\u246A':
		return "11", true // CIRCLED NUMBER ELEVEN
	case '\u246B':
		return "12", true // CIRCLED NUMBER TWELVE
	case '\u246C':
		return "13", true // CIRCLED NUMBER THIRTEEN
	case '\u246D':
		return "14", true // CIRCLED NUMBER FOURTEEN
	case '\u246E':
		return "15", true // CIRCLED NUMBER FIFTEEN
	case '\u246F':
		return "16", true // CIRCLED NUMBER SIXTEEN
	case '\u2470':
		return "17", true // CIRCLED NUMBER SEVENTEEN
	case '\u2471':
		return "18", true // CIRCLED NUMBER EIGHTEEN
	case '\u2472':
		return "19", true // CIRCLED NUMBER NINETEEN
	case '\u2473':
		return "20", true // CIRCLED NUMBER TWENTY
	case '\u2474':
		return "(1)", true // PARENTHESIZED DIGIT ONE
	case '\u2475':
		return "(2)", true // PARENTHESIZED DIGIT TWO
	case '\u2476':
		return "(3)", true // PARENTHESIZED DIGIT THREE
	case '\u2477':
		return "(4)", true // PARENTHESIZED DIGIT FOUR
	case '\u2478':
		return "(5)", true // PARENTHESIZED DIGIT FIVE
	case '\u2479':
		return "(6)", true // PARENTHESIZED DIGIT SIX
	case '\u247A':
		return "(7)", true // PARENTHESIZED DIGIT SEVEN
	case '\u247B':
		return "(8)", true // PARENTHESIZED DIGIT EIGHT
	case '\u247C':
		return "(9)", true // PARENTHESIZED DIGIT NINE
	case '\u247D':
		return "(10)", true // PARENTHESIZED NUMBER TEN
	case '\u247E':
		return "(11)", true // PARENTHESIZED NUMBER ELEVEN
	case '\u247F':
		return "(12)", true // PARENTHESIZED NUMBER TWELVE
	case '\u2480':
		return "(13)", true // PARENTHESIZED NUMBER THIRTEEN
	case '\u2481':
		return "(14)", true // PARENTHESIZED NUMBER FOURTEEN
	case '\u2482':
		return "(15)", true // PARENTHESIZED NUMBER FIFTEEN
	case '\u2483':
		return "(16)", true // PARENTHESIZED NUMBER SIXTEEN
	case '\u2484':
		return "(17)", true // PARENTHESIZED NUMBER SEVENTEEN
	case '\u2485':
		return "(18)", true // PARENTHESIZED NUMBER EIGHTEEN
	case '\u2486':
		return "(19)", true // PARENTHESIZED NUMBER NINETEEN
	case '\u2487':
		return "(20)", true // PARENTHESIZED NUMBER TWENTY
	case '\u2488':
		return "1.", true // DIGIT ONE FULL STOP
	case '\u2489':
		return "2.", true // DIGIT TWO FULL STOP
	case '\u248A':
		return "3.", true // DIGIT THREE FULL STOP
	case '\u248B':
		return "4.", true // DIGIT FOUR FULL STOP
	case '\u248C':
		return "5.", true // DIGIT FIVE FULL STOP
	case '\u248D':
		return "6.", true // DIGIT SIX FULL STOP
	case '\u248E':
		return "7.", true // DIGIT SEVEN FULL STOP
	case '\u248F':
		return "8.", true // DIGIT EIGHT FULL STOP
	case '\u2490':
		return "9.", true // DIGIT NINE FULL STOP
	case '\u2491':
		return "10.", true // NUMBER TEN FULL STOP
	case '\u2492':
		return "11.", true // NUMBER ELEVEN FULL STOP
	case '\u2493':
		return "12.", true // NUMBER TWELVE FULL STOP
	case '\u2494':
		return "13.", true // NUMBER THIRTEEN FULL STOP
	case '\u2495':
		return "14.", true // NUMBER FOURTEEN FULL STOP
	case '\u2496':
		return "15.", true // NUMBER FIFTEEN FULL STOP
	case '\u2497':
		return "16.", true // NUMBER SIXTEEN FULL STOP
	case '\u2498':
		return "17.", true // NUMBER SEVENTEEN FULL STOP
	case '\u2499':
		return "18.", true // NUMBER EIGHTEEN FULL STOP
	case '\u249A':
		return "19.", true // NUMBER NINETEEN FULL STOP
	case '\u249B':
		return "20.", true // NUMBER TWENTY FULL STOP
	case '\u249C':
		return "(a)", true // PARENTHESIZED LATIN SMALL LETTER A
	case '\u249D':
		return "(b)", true // PARENTHESIZED LATIN SMALL LETTER B
	case '\u249E':
		return "(c)", true // PARENTHESIZED LATIN SMALL LETTER C
	case '\u249F':
		return "(d)", true // PARENTHESIZED LATIN SMALL LETTER D
	case '\u24A0':
		return "(e)", true // PARENTHESIZED LATIN SMALL LETTER E
	case '\u24A1':
		return "(f)", true // PARENTHESIZED LATIN SMALL LETTER F
	case '\u24A2':
		return "(g)", true // PARENTHESIZED LATIN SMALL LETTER G
	case '\u24A3':
		return "(h)", true // PARENTHESIZED LATIN SMALL LETTER H
	case '\u24A4':
		return "(i)", true // PARENTHESIZED LATIN SMALL LETTER I
	case '\u24A5':
		return "(j)", true // PARENTHESIZED LATIN SMALL LETTER J
	case '\u24A6':
		return "(k)", true // PARENTHESIZED LATIN SMALL LETTER K
	case '\u24A7':
		return "(l)", true // PARENTHESIZED LATIN SMALL LETTER L
	case '\u24A8':
		return "(m)", true // PARENTHESIZED LATIN SMALL LETTER M
	case '\u24A9':
		return "(n)", true // PARENTHESIZED LATIN SMALL LETTER N
	case '\u24AA':
		return "(o)", true // PARENTHESIZED LATIN SMALL LETTER O
	case '\u24AB':
		return "(p)", true // PARENTHESIZED LATIN SMALL LETTER P
	case '\u24AC':
		return "(q)", true // PARENTHESIZED LATIN SMALL LETTER Q
	case '\u24AD':
		return "(r)", true // PARENTHESIZED LATIN SMALL LETTER R
	case '\u24AE':
		return "(s)", true // PARENTHESIZED LATIN SMALL LETTER S
	case '\u24AF':
		return "(t)", true // PARENTHESIZED LATIN SMALL LETTER T
	case '\u24B0':
		return "(u)", true // PARENTHESIZED LATIN SMALL LETTER U
	case '\u24B1':
		return "(v)", true // PARENTHESIZED LATIN SMALL LETTER V
	case '\u24B2':
		return "(w)", true // PARENTHESIZED LATIN SMALL LETTER W
	case '\u24B3':
		return "(x)", true // PARENTHESIZED LATIN SMALL LETTER X
	case '\u24B4':
		return "(y)", true // PARENTHESIZED LATIN SMALL LETTER Y
	case '\u24B5':
		return "(z)", true // PARENTHESIZED LATIN SMALL LETTER Z
	case '\u24B6':
		return "A", true // CIRCLED LATIN CAPITAL LETTER A
	case '\u24B7':
		return "B", true // CIRCLED LATIN CAPITAL LETTER B
	case '\u24B8':
		return "C", true // CIRCLED LATIN CAPITAL LETTER C
	case '\u24B9':
		return "D", true // CIRCLED LATIN CAPITAL LETTER D
	case '\u24BA':
		return "E", true // CIRCLED LATIN CAPITAL LETTER E
	case '\u24BB':
		return "F", true // CIRCLED LATIN CAPITAL LETTER F
	case '\u24BC':
		return "G", true // CIRCLED LATIN CAPITAL LETTER G
	case '\u24BD':
		return "H", true // CIRCLED LATIN CAPITAL LETTER H
	case '\u24BE':
		return "I", true // CIRCLED LATIN CAPITAL LETTER I
	case '\u24BF':
		return "J", true // CIRCLED LATIN CAPITAL LETTER J
	case '\u24C0':
		return "K", true // CIRCLED LATIN CAPITAL LETTER K
	case '\u24C1':
		return "L", true // CIRCLED LATIN CAPITAL LETTER L
	case '\u24C2':
		return "M", true // CIRCLED LATIN CAPITAL LETTER M
	case '\u24C3':
		return "N", true // CIRCLED LATIN CAPITAL LETTER N
	case '\u24C4':
		return "O", true // CIRCLED LATIN CAPITAL LETTER O
	case '\u24C5':
		return "P", true // CIRCLED LATIN CAPITAL LETTER P
	case '\u24C6':
		return "Q", true // CIRCLED LATIN CAPITAL LETTER Q
	case '\u24C7':
		return "R", true // CIRCLED LATIN CAPITAL LETTER R
	case '\u24C8':
		return "S", true // CIRCLED LATIN CAPITAL LETTER S
	case '\u24C9':
		return "T", true // CIRCLED LATIN CAPITAL LETTER T
	case '\u24CA':
		return "U", true // CIRCLED LATIN CAPITAL LETTER U
	case '\u24CB':
		return "V", true // CIRCLED LATIN CAPITAL LETTER V
	case '\u24CC':
		return "W", true // CIRCLED LATIN CAPITAL LETTER W
	case '\u24CD':
		return "X", true // CIRCLED LATIN CAPITAL LETTER X
	case '\u24CE':
		return "Y", true // CIRCLED LATIN CAPITAL LETTER Y
	case '\u24CF':
		return "Z", true // CIRCLED LATIN CAPITAL LETTER Z
	case '\u24D0':
		return "a", true // CIRCLED LATIN SMALL LETTER A
	case '\u24D1':
		return "b", true // CIRCLED LATIN SMALL LETTER B
	case '\u24D2':
		return "c", true // CIRCLED LATIN SMALL LETTER C
	case '\u24D3':
		return "d", true // CIRCLED LATIN SMALL LETTER D
	case '\u24D4':
		return "e", true // CIRCLED LATIN SMALL LETTER E
	case '\u24D5':
		return "f", true // CIRCLED LATIN SMALL LETTER F
	case '\u24D6':
		return "g", true // CIRCLED LATIN SMALL LETTER G
	case '\u24D7':
		return "h", true // CIRCLED LATIN SMALL LETTER H
	case '\u24D8':
		return "i", true // CIRCLED LATIN SMALL LETTER I
	case '\u24D9':
		return "j", true // CIRCLED LATIN SMALL LETTER J
	case '\u24DA':
		return "k", true // CIRCLED LATIN SMALL LETTER K
	case '\u24DB':
		return "l", true // CIRCLED LATIN SMALL LETTER L
	case '\u24DC':
		return "m", true // CIRCLED LATIN SMALL LETTER M
	case '\u24DD':
		return "n", true // CIRCLED LATIN SMALL LETTER N
	case '\u24DE':
		return "o", true // CIRCLED LATIN SMALL LETTER O
	case '\u24DF':
		return "p", true // CIRCLED LATIN SMALL LETTER P
	case '\u24E0':
		return "q", true // CIRCLED LATIN SMALL LETTER Q
	case '\u24E1':
		return "r", true // CIRCLED LATIN SMALL LETTER R
	case '\u24E2':
		return "s", true // CIRCLED LATIN SMALL LETTER S
	case '\u24E3':
		return "t", true // CIRCLED LATIN SMALL LETTER T
	case '\u24E4':
		return "u", true // CIRCLED LATIN SMALL LETTER U
	case '\u24E5':
		return "v", true // CIRCLED LATIN SMALL LETTER V
	case '\u24E6':
		return "w", true // CIRCLED LATIN SMALL LETTER W
	case '\u24E7':
		return "x", true // CIRCLED LATIN SMALL LETTER X
	case '\u24E8':
		return "y", true // CIRCLED LATIN SMALL LETTER Y
	case '\u24E9':
		return "z", true // CIRCLED LATIN SMALL LETTER Z
	case '\u24EA':
		return "0", true // CIRCLED DIGIT ZERO
	case '\u2A74':
		return "::=", true // DOUBLE COLON EQUAL
	case '\u2A75':
		return "==", true // TWO CONSECUTIVE EQUALS SIGNS
	case '\u2A76':
		return "===", true // THREE CONSECUTIVE EQUALS SIGNS
	case '\u2C7C':
		return "j", true // LATIN SUBSCRIPT SMALL LETTER J
	case '\u2C7D':
		return "V", true // MODIFIER LETTER CAPITAL V
	case '\u3000':
		return " ", true // IDEOGRAPHIC SPACE
	case '\u3250':
		return "PTE", true // PARTNERSHIP SIGN
	case '\u3251':
		return "21", true // CIRCLED NUMBER TWENTY ONE
	case '\u3252':
		return "22", true // CIRCLED NUMBER TWENTY TWO
	case '\u3253':
		return "23", true // CIRCLED NUMBER TWENTY THREE
	case '\u3254':
		return "24", true // CIRCLED NUMBER TWENTY FOUR
	case '\u3255':
		return "25", true // CIRCLED NUMBER TWENTY FIVE
	case '\u3256':
		return "26", true // CIRCLED NUMBER TWENTY SIX
	case '\u3257':
		return "27", true // CIRCLED NUMBER TWENTY SEVEN
	case '\u3258':
		return "28", true // CIRCLED NUMBER TWENTY EIGHT
	case '\u3259':
		return "29", true // CIRCLED NUMBER TWENTY NINE
	case '\u325A':
		return "30", true // CIRCLED NUMBER THIRTY
	case '\u325B':
		return "31", true // CIRCLED NUMBER THIRTY ONE
	case '\u325C':
		return "32", true // CIRCLED NUMBER THIRTY TWO
	case '\u325D':
		return "33", true // CIRCLED NUMBER THIRTY THREE
	case '\u325E':
		return "34", true // CIRCLED NUMBER THIRTY FOUR
	case '\u325F':
		return "35", true // CIRCLED NUMBER THIRTY FIVE
	case '\u32B1':
		return "36", true // CIRCLED NUMBER THIRTY SIX
	case '\u32B2':
		return "37", true // CIRCLED NUMBER THIRTY SEVEN
	case '\u32B3':
		return "38", true // CIRCLED NUMBER THIRTY EIGHT
	case '\u32B4':
		return "39", true // CIRCLED NUMBER THIRTY NINE
	case '\u32B5':
		return "40", true // CIRCLED NUMBER FORTY
	case '\u32B6':
		return "41", true // CIRCLED NUMBER FORTY ONE
	case '\u32B7':
		return "42", true // CIRCLED NUMBER FORTY TWO
	case '\u32B8':
		return "43", true // CIRCLED NUMBER FORTY THREE
	case '\u32B9':
		return "44", true // CIRCLED NUMBER FORTY FOUR
	case '\u32BA':
		return "45", true // CIRCLED NUMBER FORTY FIVE
	case '\u32BB':
		return "46", true // CIRCLED NUMBER FORTY SIX
	case '\u32BC':
		return "47", true // CIRCLED NUMBER FORTY SEVEN
	case '\u32BD':
		return "48", true // CIRCLED NUMBER FORTY EIGHT
	case '\u32BE':
		return "49", true // CIRCLED NUMBER FORTY NINE
	case '\u32BF':
		return "50", true // CIRCLED NUMBER FIFTY
	case '\u32CC':
		return "Hg", true // SQUARE HG
	case '\u32CD':
		return "erg", true // SQUARE ERG
	case '\u32CE':
		return "eV", true // SQUARE EV
	case '\u32CF':
		return "LTD", true // LIMITED LIABILITY SIGN
	case '\u3371':
		return "hPa", true // SQUARE HPA
	case '\u3372':
		return "da", true // SQUARE DA
	case '\u3373':
		return "AU", true // SQUARE AU
	case '\u3374':
		return "bar", true // SQUARE BAR
	case '\u3375':
		return "oV", true // SQUARE OV
	case '\u3376':
		return "pc", true // SQUARE PC
	case '\u3377':
		return "dm", true // SQUARE DM
	case '\u3378':
		return "dm2", true // SQUARE DM SQUARED
	case '\u3379':
		return "dm3", true // SQUARE DM CUBED
	case '\u337A':
		return "IU", true // SQUARE IU
	case '\u3380':
		return "pA", true // SQUARE PA AMPS
	case '\u3381':
		return "nA", true // SQUARE NA
	case '\u3383':
		return "mA", true // SQUARE MA
	case '\u3384':
		return "kA", true // SQUARE KA
	case '\u3385':
		return "KB", true // SQUARE KB
	case '\u3386':
		return "MB", true // SQUARE MB
	case '\u3387':
		return "GB", true // SQUARE GB
	case '\u3388':
		return "cal", true // SQUARE CAL
	case '\u3389':
		return "kcal", true // SQUARE KCAL
	case '\u338A':
		return "pF", true // SQUARE PF
	case '\u338B':
		return "nF", true // SQUARE NF
	case '\u338E':
		return "mg", true // SQUARE MG
	case '\u338F':
		return "kg", true // SQUARE KG
	case '\u3390':
		return "Hz", true // SQUARE HZ
	case '\u3391':
		return "kHz", true // SQUARE KHZ
	case '\u3392':
		return "MHz", true // SQUARE MHZ
	case '\u3393':
		return "GHz", true // SQUARE GHZ
	case '\u3394':
		return "THz", true // SQUARE THZ
	case '\u3396':
		return "ml", true // SQUARE ML
	case '\u3397':
		return "dl", true // SQUARE DL
	case '\u3398':
		return "kl", true // SQUARE KL
	case '\u3399':
		return "fm", true // SQUARE FM
	case '\u339A':
		return "nm", true // SQUARE NM
	case '\u339C':
		return "mm", true // SQUARE MM
	case '\u339D':
		return "cm", true // SQUARE CM
	case '\u339E':
		return "km", true // SQUARE KM
	case '\u339F':
		return "mm2", true // SQUARE MM SQUARED
	case '\u33A0':
		return "cm2", true // SQUARE CM SQUARED
	case '\u33A1':
		return "m2", true // SQUARE M SQUARED
	case '\u33A2':
		return "km2", true // SQUARE KM SQUARED
	case '\u33A3':
		return "mm3", true // SQUARE MM CUBED
	case '\u33A4':
		return "cm3", true // SQUARE CM CUBED
	case '\u33A5':
		return "m3", true // SQUARE M CUBED
	case '\u33A6':
		return "km3", true // SQUARE KM CUBED
	case '\u33A9':
		return "Pa", true // SQUARE PA
	case '\u33AA':
		return "kPa", true // SQUARE KPA
	case '\u33AB':
		return "MPa", true // SQUARE MPA
	case '\u33AC':
		return "GPa", true // SQUARE GPA
	case '\u33AD':
		return "rad", true // SQUARE RAD
	case '\u33B0':
		return "ps", true // SQUARE PS
	case '\u33B1':
		return "ns", true // SQUARE NS
	case '\u33B3':
		return "ms", true // SQUARE MS
	case '\u33B4':
		return "pV", true // SQUARE PV
	case '\u33B5':
		return "nV", true // SQUARE NV
	case '\u33B7':
		return "mV", true // SQUARE MV
	case '\u33B8':
		return "kV", true // SQUARE KV
	case '\u33B9':
		return "MV", true // SQUARE MV MEGA
	case '\u33BA':
		return "pW", true // SQUARE PW
	case '\u33BB':
		return "nW", true // SQUARE NW
	case '\u33BD':
		return "mW", true // SQUARE MW
	case '\u33BE':
		return "kW", true // SQUARE KW
	case '\u33BF':
		return "MW", true // SQUARE MW MEGA
	case '\u33C2':
		return "a.m.", true // SQUARE AM
	case '\u33C3':
		return "Bq", true // SQUARE BQ
	case '\u33C4':
		return "cc", true // SQUARE CC
	case '\u33C5':
		return "cd", true // SQUARE CD
	case '\u33C7':
		return "Co.", true // SQUARE CO
	case '\u33C8':
		return "dB", true // SQUARE DB
	case '\u33C9':
		return "Gy", true // SQUARE GY
	case '\u33CA':
		return "ha", true // SQUARE HA
	case '\u33CB':
		return "HP", true // SQUARE HP
	case '\u33CC':
		return "in", true // SQUARE IN
	case '\u33CD':
		return "KK", true // SQUARE KK
	case '\u33CE':
		return "KM", true // SQUARE KM CAPITAL
	case '\u33CF':
		return "kt", true // SQUARE KT
	case '\u33D0':
		return "lm", true // SQUARE LM
	case '\u33D1':
		return "ln", true // SQUARE LN
	case '\u33D2':
		return "log", true // SQUARE LOG
	case '\u33D3':
		return "lx", true // SQUARE LX
	case '\u33D4':
		return "mb", true // SQUARE MB SMALL
	case '\u33D5':
		return "mil", true // SQUARE MIL
	case '\u33D6':
		return "mol", true // SQUARE MOL
	case '\u33D7':
		return "PH", true // SQUARE PH
	case '\u33D8':
		return "p.m.", true // SQUARE PM
	case '\u33D9':
		return "PPM", true // SQUARE PPM
	case '\u33DA':
		return "PR", true // SQUARE PR
	case '\u33DB':
		return "sr", true // SQUARE SR
	case '\u33DC':
		return "Sv", true // SQUARE SV
	case '\u33DD':
		return "Wb", true // SQUARE WB
	case '\u33FF':
		return "gal", true // SQUARE GAL
	case '\uA7F2':
		return "C", true // MODIFIER LETTER CAPITAL C
	case '\uA7F3':
		return "F", true // MODIFIER LETTER CAPITAL F
	case '\uA7F4':
		return "Q", true // MODIFIER LETTER CAPITAL Q
	case '\uFB00':
		return "ff", true // LATIN SMALL LIGATURE FF
	case '\uFB01':
		return "fi", true // LATIN SMALL LIGATURE FI
	case '\uFB02':
		return "fl", true // LATIN SMALL LIGATURE FL
	case '\uFB03':
		return "ffi", true // LATIN SMALL LIGATURE FFI
	case '\uFB04':
		return "ffl", true // LATIN SMALL LIGATURE FFL
	case '\uFB05':
		return "st", true // LATIN SMALL LIGATURE LONG S T
	case '\uFB06':
		return "st", true // LATIN SMALL LIGATURE ST
	case '\uFB29':
		return "+", true // HEBREW LETTER ALTERNATIVE PLUS SIGN
	case '\uFE10':
		return ",", true // PRESENTATION FORM FOR VERTICAL COMMA
	case '\uFE13':
		return ":", true // PRESENTATION FORM FOR VERTICAL COLON
	case '\uFE14':
		return ";", true // PRESENTATION FORM FOR VERTICAL SEMICOLON
	case '\uFE15':
		return "!", true // PRESENTATION FORM FOR VERTICAL EXCLAMATION MARK
	case '\uFE16':
		return "?", true // PRESENTATION FORM FOR VERTICAL QUESTION MARK
	case '\uFE19':
		return "...", true // PRESENTATION FORM FOR VERTICAL HORIZONTAL ELLIPSIS
	case '\uFE30':
		return "..", true // PRESENTATION FORM FOR VERTICAL TWO DOT LEADER
	case '\uFE33':
		return "_", true // PRESENTATION FORM FOR VERTICAL LOW LINE
	case '\uFE34':
		return "_", true // PRESENTATION FORM FOR VERTICAL WAVY LOW LINE
	case '\uFE35':
		return "(", true // PRESENTATION FORM FOR VERTICAL LEFT PARENTHESIS
	case '\uFE36':
		return ")", true // PRESENTATION FORM FOR VERTICAL RIGHT PARENTHESIS
	case '\uFE37':
		return "{", true // PRESENTATION FORM FOR VERTICAL LEFT CURLY BRACKET
	case '\uFE38':
		return "}", true // PRESENTATION FORM FOR VERTICAL RIGHT CURLY BRACKET
	case '\uFE47':
		return "[", true // PRESENTATION FORM FOR VERTICAL LEFT SQUARE BRACKET
	case '\uFE48':
		return "]", true // PRESENTATION FORM FOR VERTICAL RIGHT SQUARE BRACKET
	case '\uFE4D':
		return "_", true // DASHED LOW LINE
	case '\uFE4E':
		return "_", true // CENTRELINE LOW LINE
	case '\uFE4F':
		return "_", true // WAVY LOW LINE
	case '\uFE50':
		return ",", true // SMALL COMMA
	case '\uFE52':
		return ".", true // SMALL FULL STOP
	case '\uFE54':
		return ";", true // SMALL SEMICOLON
	case '\uFE55':
		return ":", true // SMALL COLON
	case '\uFE56':
		return "?", true // SMALL QUESTION MARK
	case '\uFE57':
		return "!", true // SMALL EXCLAMATION MARK
	case '\uFE59':
		return "(", true // SMALL LEFT PARENTHESIS
	case '\uFE5A':
		return ")", true // SMALL RIGHT PARENTHESIS
	case '\uFE5B':
		return "{", true // SMALL LEFT CURLY BRACKET
	case '\uFE5C':
		return "}", true // SMALL RIGHT CURLY BRACKET
	case '\uFE5F':
		return "#", true // SMALL NUMBER SIGN
	case '\uFE60':
		return "&", true // SMALL AMPERSAND
	case '\uFE61':
		return "*", true // SMALL ASTERISK
	case '\uFE62':
		return "+", true // SMALL PLUS SIGN
	case '\uFE63':
		return "-", true // SMALL HYPHEN-MINUS
	case '\uFE64':
		return "<", true // SMALL LESS-THAN SIGN
	case '\uFE65':
		return ">", true // SMALL GREATER-THAN SIGN
	case '\uFE66':
		return "=", true // SMALL EQUALS SIGN
	case '\uFE68':
		return "\\", true // SMALL REVERSE SOLIDUS
	case '\uFE69':
		return "$", true // SMALL DOLLAR SIGN
	case '\uFE6A':
		return "%", true // SMALL PERCENT SIGN
	case '\uFE6B':
		return "@", true // SMALL COMMERCIAL AT
	case '\uFF01':
		return "!", true // FULLWIDTH EXCLAMATION MARK
	case '\uFF02':
		return "\"", true // FULLWIDTH QUOTATION MARK
	case '\uFF03':
		return "#", true // FULLWIDTH NUMBER SIGN
	case '\uFF04':
		return "$", true // FULLWIDTH DOLLAR SIGN
	case '\uFF05':
		return "%", true // FULLWIDTH PERCENT SIGN
	case '\uFF06':
		return "&", true // FULLWIDTH AMPERSAND
	case '\uFF07':
		return "'", true // FULLWIDTH APOSTROPHE
	case '\uFF08':
		return "(", true // FULLWIDTH LEFT PARENTHESIS
	case '\uFF09':
		return ")", true // FULLWIDTH RIGHT PARENTHESIS
	case '\uFF0A':
		return "*", true // FULLWIDTH ASTERISK
	case '\uFF0B':
		return "+", true // FULLWIDTH PLUS SIGN
	case '\uFF0C':
		return ",", true // FULLWIDTH COMMA
	case '\uFF0D':
		return "-", true // FULLWIDTH HYPHEN-MINUS
	case '\uFF0E':
		return ".", true // FULLWIDTH FULL STOP
	case '\uFF0F':
		return "/", true // FULLWIDTH SOLIDUS
	case '\uFF10':
		return "0", true // FULLWIDTH DIGIT ZERO
	case '\uFF11':
		return "1", true // FULLWIDTH DIGIT ONE
	case '\uFF12':
		return "2", true // FULLWIDTH DIGIT TWO
	case '\uFF13':
		return "3", true // FULLWIDTH DIGIT THREE
	case '\uFF14':
		return "4", true // FULLWIDTH DIGIT FOUR
	case '\uFF15':
		return "5", true // FULLWIDTH DIGIT FIVE
	case '\uFF16':
		return "6", true // FULLWIDTH DIGIT SIX
	case '\uFF17':
		return "7", true // FULLWIDTH DIGIT SEVEN
	case '\uFF18':
		return "8", true // FULLWIDTH DIGIT EIGHT
	case '\uFF19':
		return "9", true // FULLWIDTH DIGIT NINE
	case '\uFF1A':
		return ":", true // FULLWIDTH COLON
	case '\uFF1B':
		return ";", true // FULLWIDTH SEMICOLON
	case '\uFF1C':
		return "<", true // FULLWIDTH LESS-THAN SIGN
	case '\uFF1D':
		return "=", true // FULLWIDTH EQUALS SIGN
	case '\uFF1E':
		return ">", true // FULLWIDTH GREATER-THAN SIGN
	case '\uFF1F':
		return "?", true // FULLWIDTH QUESTION MARK
	case '\uFF20':
		return "@", true // FULLWIDTH COMMERCIAL AT
	case '\uFF21':
		return "A", true // FULLWIDTH LATIN CAPITAL LETTER A
	case '\uFF22':
		return "B", true // FULLWIDTH LATIN CAPITAL LETTER B
	case '\uFF23':
		return "C", true // FULLWIDTH LATIN CAPITAL LETTER C
	case '\uFF24':
		return "D", true // FULLWIDTH LATIN CAPITAL LETTER D
	case '\uFF25':
		return "E", true // FULLWIDTH LATIN CAPITAL LETTER E
	case '\uFF26':
		return "F", true // FULLWIDTH LATIN CAPITAL LETTER F
	case '\uFF27':
		return "G", true // FULLWIDTH LATIN CAPITAL LETTER G
	case '\uFF28':
		return "H", true // FULLWIDTH LATIN CAPITAL LETTER H
	case '\uFF29':
		return "I", true // FULLWIDTH LATIN CAPITAL LETTER I
	case '\uFF2A':
		return "J", true // FULLWIDTH LATIN CAPITAL LETTER J
	case '\uFF2B':
		return "K", true // FULLWIDTH LATIN CAPITAL LETTER K
	case '\uFF2C':
		return "L", true // FULLWIDTH LATIN CAPITAL LETTER L
	case '\uFF2D':
		return "M", true // FULLWIDTH LATIN CAPITAL LETTER M
	case '\uFF2E':
		return "N", true // FULLWIDTH LATIN CAPITAL LETTER N
	case '\uFF2F':
		return "O", true // FULLWIDTH LATIN CAPITAL LETTER O
	case '\uFF30':
		return "P", true // FULLWIDTH LATIN CAPITAL LETTER P
	case '\uFF31':
		return "Q", true // FULLWIDTH LATIN CAPITAL LETTER Q
	case '\uFF32':
		return "R", true // FULLWIDTH LATIN CAPITAL LETTER R
	case '\uFF33':
		return "S", true // FULLWIDTH LATIN CAPITAL LETTER S
	case '\uFF34':
		return "T", true // FULLWIDTH LATIN CAPITAL LETTER T
	case '\uFF35':
		return "U", true // FULLWIDTH LATIN CAPITAL LETTER U
	case '\uFF36':
		return "V", true // FULLWIDTH LATIN CAPITAL LETTER V
	case '\uFF37':
		return "W", true // FULLWIDTH LATIN CAPITAL LETTER W
	case '\uFF38':
		return "X", true // FULLWIDTH LATIN CAPITAL LETTER X
	case '\uFF39':
		return "Y", true // FULLWIDTH LATIN CAPITAL LETTER Y
	case '\uFF3A':
		return "Z", true // FULLWIDTH LATIN CAPITAL LETTER Z
	case '\uFF3B':
		return "[", true // FULLWIDTH LEFT SQUARE BRACKET
	case '\uFF3C':
		return "\\", true // FULLWIDTH REVERSE SOLIDUS
	case '\uFF3D':
		return "]", true // FULLWIDTH RIGHT SQUARE BRACKET
	case '\uFF3E':
		return "^", true // FULLWIDTH CIRCUMFLEX ACCENT
	case '\uFF3F':
		return "_", true // FULLWIDTH LOW LINE
	case '\uFF41':
		return "a", true // FULLWIDTH LATIN SMALL LETTER A
	case '\uFF42':
		return "b", true // FULLWIDTH LATIN SMALL LETTER B
	case '\uFF43':
		return "c", true // FULLWIDTH LATIN SMALL LETTER C
	case '\uFF44':
		return "d", true // FULLWIDTH LATIN SMALL LETTER D
	case '\uFF45':
		return "e", true // FULLWIDTH LATIN SMALL LETTER E
	case '\uFF46':
		return "f", true // FULLWIDTH LATIN SMALL LETTER F
	case '\uFF47':
		return "g", true // FULLWIDTH LATIN SMALL LETTER G
	case '\uFF48':
		return "h", true // FULLWIDTH LATIN SMALL LETTER H
	case '\uFF49':
		return "i", true // FULLWIDTH LATIN SMALL LETTER I
	case '\uFF4A':
		return "j", true // FULLWIDTH LATIN SMALL LETTER J
	case '\uFF4B':
		return "k", true // FULLWIDTH LATIN SMALL LETTER K
	case '\uFF4C':
		return "l", true // FULLWIDTH LATIN SMALL LETTER L
	case '\uFF4D':
		return "m", true // FULLWIDTH LATIN SMALL LETTER M
	case '\uFF4E':
		return "n", true // FULLWIDTH LATIN SMALL LETTER N
	case '\uFF4F':
		return "o", true // FULLWIDTH LATIN SMALL LETTER O
	case '\uFF50':
		return "p", true // FULLWIDTH LATIN SMALL LETTER P
	case '\uFF51':
		return "q", true // FULLWIDTH LATIN SMALL LETTER Q
	case '\uFF52':
		return "r", true // FULLWIDTH LATIN SMALL LETTER R
	case '\uFF53':
		return "s", true // FULLWIDTH LATIN SMALL LETTER S
	case '\uFF54':
		return "t", true // FULLWIDTH LATIN SMALL LETTER T
	case '\uFF55':
		return "u", true // FULLWIDTH LATIN SMALL LETTER U
	case '\uFF56':
		return "v", true // FULLWIDTH LATIN SMALL LETTER V
	case '\uFF57':
		return "w", true // FULLWIDTH LATIN SMALL LETTER W
	case '\uFF58':
		return "x", true // FULLWIDTH LATIN SMALL LETTER X
	case '\uFF59':
		return "y", true // FULLWIDTH LATIN SMALL LETTER Y
	case '\uFF5A':
		return "z", true // FULLWIDTH LATIN SMALL LETTER Z
	case '\uFF5B':
		return "{", true // FULLWIDTH LEFT CURLY BRACKET
	case '\uFF5C':
		return "|", true // FULLWIDTH VERTICAL LINE
	case '\uFF5D':
		return "}", true // FULLWIDTH RIGHT CURLY BRACKET
	case '\uFF5E':
		return "~", true // FULLWIDTH TILDE
	case '\U000107A5':
		return "q", true // MODIFIER LETTER SMALL Q
	case '\U0001D400':
		return "A", true // MATHEMATICAL BOLD CAPITAL A
	case '\U0001D401':
		return "B", true // MATHEMATICAL BOLD CAPITAL B
	case '\U0001D402':
		return "C", true // MATHEMATICAL BOLD CAPITAL C
	case '\U0001D403':
		return "D", true // MATHEMATICAL BOLD CAPITAL D
	case '\U0001D404':
		return "E", true // MATHEMATICAL BOLD CAPITAL E
	case '\U0001D405':
		return "F", true // MATHEMATICAL BOLD CAPITAL F
	case '\U0001D406':
		return "G", true // MATHEMATICAL BOLD CAPITAL G
	case '\U0001D407':
		return "H", true // MATHEMATICAL BOLD CAPITAL H
	case '\U0001D408':
		return "I", true // MATHEMATICAL BOLD CAPITAL I
	case '\U0001D409':
		return "J", true // MATHEMATICAL BOLD CAPITAL J
	case '\U0001D40A':
		return "K", true // MATHEMATICAL BOLD CAPITAL K
	case '\U0001D40B':
		return "L", true // MATHEMATICAL BOLD CAPITAL L
	case '\U0001D40C':
		return "M", true // MATHEMATICAL BOLD CAPITAL M
	case '\U0001D40D':
		return "N", true // MATHEMATICAL BOLD CAPITAL N
	case '\U0001D40E':
		return "O", true // MATHEMATICAL BOLD CAPITAL O
	case '\U0001D40F':
		return "P", true // MATHEMATICAL BOLD CAPITAL P
	case '\U0001D410':
		return "Q", true // MATHEMATICAL BOLD CAPITAL Q
	case '\U0001D411':
		return "R", true // MATHEMATICAL BOLD CAPITAL R
	case '\U0001D412':
		return "S", true // MATHEMATICAL BOLD CAPITAL S
	case '\U0001D413':
		return "T", true // MATHEMATICAL BOLD CAPITAL T
	case '\U0001D414':
		return "U", true // MATHEMATICAL BOLD CAPITAL U
	case '\U0001D415':
		return "V", true // MATHEMATICAL BOLD CAPITAL V
	case '\U0001D416':
		return "W", true // MATHEMATICAL BOLD CAPITAL W
	case '\U0001D417':
		return "X", true // MATHEMATICAL BOLD CAPITAL X
	case '\U0001D418':
		return "Y", true // MATHEMATICAL BOLD CAPITAL Y
	case '\U0001D419':
		return "Z", true // MATHEMATICAL BOLD CAPITAL Z
	case '\U0001D41A':
		return "a", true // MATHEMATICAL BOLD SMALL A
	case '\U0001D41B':
		return "b", true // MATHEMATICAL BOLD SMALL B
	case '\U0001D41C':
		return "c", true // MATHEMATICAL BOLD SMALL C
	case '\U0001D41D':
		return "d", true // MATHEMATICAL BOLD SMALL D
	case '\U0001D41E':
		return "e", true // MATHEMATICAL BOLD SMALL E
	case '\U0001D41F':
		return "f", true // MATHEMATICAL BOLD SMALL F
	case '\U0001D420':
		return "g", true // MATHEMATICAL BOLD SMALL G
	case '\U0001D421':
		return "h", true // MATHEMATICAL BOLD SMALL H
	case '\U0001D422':
		return "i", true // MATHEMATICAL BOLD SMALL I
	case '\U0001D423':
		return "j", true // MATHEMATICAL BOLD SMALL J
	case '\U0001D424':
		return "k", true // MATHEMATICAL BOLD SMALL K
	case '\U0001D425':
		return "l", true // MATHEMATICAL BOLD SMALL L
	case '\U0001D426':
		return "m", true // MATHEMATICAL BOLD SMALL M
	case '\U0001D427':
		return "n", true // MATHEMATICAL BOLD SMALL N
	case '\U0001D428':
		return "o", true // MATHEMATICAL BOLD SMALL O
	case '\U0001D429':
		return "p", true // MATHEMATICAL BOLD SMALL P
	case '\U0001D42A':
		return "q", true // MATHEMATICAL BOLD SMALL Q
	case '\U0001D42B':
		return "r", true // MATHEMATICAL BOLD SMALL R
	case '\U0001D42C':
		return "s", true // MATHEMATICAL BOLD SMALL S
	case '\U0001D42D':
		return "t", true // MATHEMATICAL BOLD SMALL T
	case '\U0001D42E':
		return "u", true // MATHEMATICAL BOLD SMALL U
	case '\U0001D42F':
		return "v", true // MATHEMATICAL BOLD SMALL V
	case '\U0001D430':
		return "w", true // MATHEMATICAL BOLD SMALL W
	case '\U0001D431':
		return "x", true // MATHEMATICAL BOLD SMALL X
	case '\U0001D432':
		return "y", true // MATHEMATICAL BOLD SMALL Y
	case '\U0001D433':
		return "z", true // MATHEMATICAL BOLD SMALL Z
	case '\U0001D434':
		return "A", true // MATHEMATICAL ITALIC CAPITAL A
	case '\U0001D435':
		return "B", true // MATHEMATICAL ITALIC CAPITAL B
	case '\U0001D436':
		return "C", true // MATHEMATICAL ITALIC CAPITAL C
	case '\U0001D437':
		return "D", true // MATHEMATICAL ITALIC CAPITAL D
	case '\U0001D438':
		return "E", true // MATHEMATICAL ITALIC CAPITAL E
	case '\U0001D439':
		return "F", true // MATHEMATICAL ITALIC CAPITAL F
	case '\U0001D43A':
		return "G", true // MATHEMATICAL ITALIC CAPITAL G
	case '\U0001D43B':
		return "H", true // MATHEMATICAL ITALIC CAPITAL H
	case '\U0001D43C':
		return "I", true // MATHEMATICAL ITALIC CAPITAL I
	case '\U0001D43D':
		return "J", true // MATHEMATICAL ITALIC CAPITAL J
	case '\U0001D43E':
		return "K", true // MATHEMATICAL ITALIC CAPITAL K
	case '\U0001D43F':
		return "L", true // MATHEMATICAL ITALIC CAPITAL L
	case '\U0001D440':
		return "M", true // MATHEMATICAL ITALIC CAPITAL M
	case '\U0001D441':
		return "N", true // MATHEMATICAL ITALIC CAPITAL N
	case '\U0001D442':
		return "O", true // MATHEMATICAL ITALIC CAPITAL O
	case '\U0001D443':
		return "P", true // MATHEMATICAL ITALIC CAPITAL P
	case '\U0001D444':
		return "Q", true // MATHEMATICAL ITALIC CAPITAL Q
	case '\U0001D445':
		return "R", true // MATHEMATICAL ITALIC CAPITAL R
	case '\U0001D446':
		return "S", true // MATHEMATICAL ITALIC CAPITAL S
	case '\U0001D447':
		return "T", true // MATHEMATICAL ITALIC CAPITAL T
	case '\U0001D448':
		return "U", true // MATHEMATICAL ITALIC CAPITAL U
	case '\U0001D449':
		return "V", true // MATHEMATICAL ITALIC CAPITAL V
	case '\U0001D44A':
		return "W", true // MATHEMATICAL ITALIC CAPITAL W
	case '\U0001D44B':
		return "X", true // MATHEMATICAL ITALIC CAPITAL X
	case '\U0001D44C':
		return "Y", true // MATHEMATICAL ITALIC CAPITAL Y
	case '\U0001D44D':
		return "Z", true // MATHEMATICAL ITALIC CAPITAL Z
	case '\U0001D44E':
		return "a", true // MATHEMATICAL ITALIC SMALL A
	case '\U0001D44F':
		return "b", true // MATHEMATICAL ITALIC SMALL B
	case '\U0001D450':
		return "c", true // MATHEMATICAL ITALIC SMALL C
	case '\U0001D451':
		return "d", true // MATHEMATICAL ITALIC SMALL D
	case '\U0001D452':
		return "e", true // MATHEMATICAL ITALIC SMALL E
	case '\U0001D453':
		return "f", true // MATHEMATICAL ITALIC SMALL F
	case '\U0001D454':
		return "g", true // MATHEMATICAL ITALIC SMALL G
	case '\U0001D456':
		return "i", true // MATHEMATICAL ITALIC SMALL I
	case '\U0001D457':
		return "j", true // MATHEMATICAL ITALIC SMALL J
	case '\U0001D458':
		return "k", true // MATHEMATICAL ITALIC SMALL K
	case '\U0001D459':
		return "l", true // MATHEMATICAL ITALIC SMALL L
	case '\U0001D45A':
		return "m", true // MATHEMATICAL ITALIC SMALL M
	case '\U0001D45B':
		return "n", true // MATHEMATICAL ITALIC SMALL N
	case '\U0001D45C':
		return "o", true // MATHEMATICAL ITALIC SMALL O
	case '\U0001D45D':
		return "p", true // MATHEMATICAL ITALIC SMALL P
	case '\U0001D45E':
		return "q", true // MATHEMATICAL ITALIC SMALL Q
	case '\U0001D45F':
		return "r", true // MATHEMATICAL ITALIC SMALL R
	case '\U0001D460':
		return "s", true // MATHEMATICAL ITALIC SMALL S
	case '\U0001D461':
		return "t", true // MATHEMATICAL ITALIC SMALL T
	case '\U0001D462':
		return "u", true // MATHEMATICAL ITALIC SMALL U
	case '\U0001D463':
		return "v", true // MATHEMATICAL ITALIC SMALL V
	case '\U0001D464':
		return "w", true // MATHEMATICAL ITALIC SMALL W
	case '\U0001D465':
		return "x", true // MATHEMATICAL ITALIC SMALL X
	case '\U0001D466':
		return "y", true // MATHEMATICAL ITALIC SMALL Y
	case '\U0001D467':
		return "z", true // MATHEMATICAL ITALIC SMALL Z
	case '\U0001D468':
		return "A", true // MATHEMATICAL BOLD ITALIC CAPITAL A
	case '\U0001D469':
		return "B", true // MATHEMATICAL BOLD ITALIC CAPITAL B
	case '\U0001D46A':
		return "C", true // MATHEMATICAL BOLD ITALIC CAPITAL C
	case '\U0001D46B':
		return "D", true // MATHEMATICAL BOLD ITALIC CAPITAL D
	case '\U0001D46C':
		return "E", true // MATHEMATICAL BOLD ITALIC CAPITAL E
	case '\U0001D46D':
		return "F", true // MATHEMATICAL BOLD ITALIC CAPITAL F
	case '\U0001D46E':
		return "G", true // MATHEMATICAL BOLD ITALIC CAPITAL G
	case '\U0001D46F':
		return "H", true // MATHEMATICAL BOLD ITALIC CAPITAL H
	case '\U0001D470':
		return "I", true // MATHEMATICAL BOLD ITALIC CAPITAL I
	case '\U0001D471':
		return "J", true // MATHEMATICAL BOLD ITALIC CAPITAL J
	case '\U0001D472':
		return "K", true // MATHEMATICAL BOLD ITALIC CAPITAL K
	case '\U0001D473':
		return "L", true // MATHEMATICAL BOLD ITALIC CAPITAL L
	case '\U0001D474':
		return "M", true // MATHEMATICAL BOLD ITALIC CAPITAL M
	case '\U0001D475':
		return "N", true // MATHEMATICAL BOLD ITALIC CAPITAL N
	case '\U0001D476':
		return "O", true // MATHEMATICAL BOLD ITALIC CAPITAL O
	case '\U0001D477':
		return "P", true // MATHEMATICAL BOLD ITALIC CAPITAL P
	case '\U0001D478':
		return "Q", true // MATHEMATICAL BOLD ITALIC CAPITAL Q
	case '\U0001D479':
		return "R", true // MATHEMATICAL BOLD ITALIC CAPITAL R
	case '\U0001D47A':
		return "S", true // MATHEMATICAL BOLD ITALIC CAPITAL S
	case '\U0001D47B':
		return "T", true // MATHEMATICAL BOLD ITALIC CAPITAL T
	case '\U0001D47C':
		return "U", true // MATHEMATICAL BOLD ITALIC CAPITAL U
	case '\U0001D47D':
		return "V", true // MATHEMATICAL BOLD ITALIC CAPITAL V
	case '\U0001D47E':
		return "W", true // MATHEMATICAL BOLD ITALIC CAPITAL W
	case '\U0001D47F':
		return "X", true // MATHEMATICAL BOLD ITALIC CAPITAL X
	case '\U0001D480':
		return "Y", true // MATHEMATICAL BOLD ITALIC CAPITAL Y
	case '\U0001D481':
		return "Z", true // MATHEMATICAL BOLD ITALIC CAPITAL Z
	case '\U0001D482':
		return "a", true // MATHEMATICAL BOLD ITALIC SMALL A
	case '\U0001D483':
		return "b", true // MATHEMATICAL BOLD ITALIC SMALL B
	case '\U0001D484':
		return "c", true // MATHEMATICAL BOLD ITALIC SMALL C
	case '\U0001D485':
		return "d", true // MATHEMATICAL BOLD ITALIC SMALL D
	case '\U0001D486':
		return "e", true // MATHEMATICAL BOLD ITALIC SMALL E
	case '\U0001D487':
		return "f", true // MATHEMATICAL BOLD ITALIC SMALL F
	case '\U0001D488':
		return "g", true // MATHEMATICAL BOLD ITALIC SMALL G
	case '\U0001D489':
		return "h", true // MATHEMATICAL BOLD ITALIC SMALL H
	case '\U0001D48A':
		return "i", true // MATHEMATICAL BOLD ITALIC SMALL I
	case '\U0001D48B':
		return "j", true // MATHEMATICAL BOLD ITALIC SMALL J
	case '\U0001D48C':
		return "k", true // MATHEMATICAL BOLD ITALIC SMALL K
	case '\U0001D48D':
		return "l", true // MATHEMATICAL BOLD ITALIC SMALL L
	case '\U0001D48E':
		return "m", true // MATHEMATICAL BOLD ITALIC SMALL M
	case '\U0001D48F':
		return "n", true // MATHEMATICAL BOLD ITALIC SMALL N
	case '\U0001D490':
		return "o", true // MATHEMATICAL BOLD ITALIC SMALL O
	case '\U0001D491':
		return "p", true // MATHEMATICAL BOLD ITALIC SMALL P
	case '\U0001D492':
		return "q", true // MATHEMATICAL BOLD ITALIC SMALL Q
	case '\U0001D493':
		return "r", true // MATHEMATICAL BOLD ITALIC SMALL R
	case '\U0001D494':
		return "s", true // MATHEMATICAL BOLD ITALIC SMALL S
	case '\U0001D495':
		return "t", true // MATHEMATICAL BOLD ITALIC SMALL T
	case '\U0001D496':
		return "u", true // MATHEMATICAL BOLD ITALIC SMALL U
	case '\U0001D497':
		return "v", true // MATHEMATICAL BOLD ITALIC SMALL V
	case '\U0001D498':
		return "w", true // MATHEMATICAL BOLD ITALIC SMALL W
	case '\U0001D499':
		return "x", true // MATHEMATICAL BOLD ITALIC SMALL X
	case '\U0001D49A':
		return "y", true // MATHEMATICAL BOLD ITALIC SMALL Y
	case '\U0001D49B':
		return "z", true // MATHEMATICAL BOLD ITALIC SMALL Z
	case '\U0001D49C':
		return "A", true // MATHEMATICAL SCRIPT CAPITAL A
	case '\U0001D49E':
		return "C", true // MATHEMATICAL SCRIPT CAPITAL C
	case '\U0001D49F':
		return "D", true // MATHEMATICAL SCRIPT CAPITAL D
	case '\U0001D4A2':
		return "G", true // MATHEMATICAL SCRIPT CAPITAL G
	case '\U0001D4A5':
		return "J", true // MATHEMATICAL SCRIPT CAPITAL J
	case '\U0001D4A6':
		return "K", true // MATHEMATICAL SCRIPT CAPITAL K
	case '\U0001D4A9':
		return "N", true // MATHEMATICAL SCRIPT CAPITAL N
	case '\U0001D4AA':
		return "O", true // MATHEMATICAL SCRIPT CAPITAL O
	case '\U0001D4AB':
		return "P", true // MATHEMATICAL SCRIPT CAPITAL P
	case '\U0001D4AC':
		return "Q", true // MATHEMATICAL SCRIPT CAPITAL Q
	case '\U0001D4AE':
		return "S", true // MATHEMATICAL SCRIPT CAPITAL S
	case '\U0001D4AF':
		return "T", true // MATHEMATICAL SCRIPT CAPITAL T
	case '\U0001D4B0':
		return "U", true // MATHEMATICAL SCRIPT CAPITAL U
	case '\U0001D4B1':
		return "V", true // MATHEMATICAL SCRIPT CAPITAL V
	case '\U0001D4B2':
		return "W", true // MATHEMATICAL SCRIPT CAPITAL W
	case '\U0001D4B3':
		return "X", true // MATHEMATICAL SCRIPT CAPITAL X
	case '\U0001D4B4':
		return "Y", true // MATHEMATICAL SCRIPT CAPITAL Y
	case '\U0001D4B5':
		return "Z", true // MATHEMATICAL SCRIPT CAPITAL Z
	case '\U0001D4B6':
		return "a", true // MATHEMATICAL SCRIPT SMALL A
	case '\U0001D4B7':
		return "b", true // MATHEMATICAL SCRIPT SMALL B
	case '\U0001D4B8':
		return "c", true // MATHEMATICAL SCRIPT SMALL C
	case '\U0001D4B9':
		return "d", true // MATHEMATICAL SCRIPT SMALL D
	case '\U0001D4BB':
		return "f", true // MATHEMATICAL SCRIPT SMALL F
	case '\U0001D4BD':
		return "h", true // MATHEMATICAL SCRIPT SMALL H
	case '\U0001D4BE':
		return "i", true // MATHEMATICAL SCRIPT SMALL I
	case '\U0001D4BF':
		return "j", true // MATHEMATICAL SCRIPT SMALL J
	case '\U0001D4C0':
		return "k", true // MATHEMATICAL SCRIPT SMALL K
	case '\U0001D4C1':
		return "l", true // MATHEMATICAL SCRIPT SMALL L
	case '\U0001D4C2':
		return "m", true // MATHEMATICAL SCRIPT SMALL M
	case '\U0001D4C3':
		return "n", true // MATHEMATICAL SCRIPT SMALL N
	case '\U0001D4C5':
		return "p", true // MATHEMATICAL SCRIPT SMALL P
	case '\U0001D4C6':
		return "q", true // MATHEMATICAL SCRIPT SMALL Q
	case '\U0001D4C7':
		return "r", true // MATHEMATICAL SCRIPT SMALL R
	case '\U0001D4C8':
		return "s", true // MATHEMATICAL SCRIPT SMALL S
	case '\U0001D4C9':
		return "t", true // MATHEMATICAL SCRIPT SMALL T
	case '\U0001D4CA':
		return "u", true // MATHEMATICAL SCRIPT SMALL U
	case '\U0001D4CB':
		return "v", true // MATHEMATICAL SCRIPT SMALL V
	case '\U0001D4CC':
		return "w", true // MATHEMATICAL SCRIPT SMALL W
	case '\U0001D4CD':
		return "x", true // MATHEMATICAL SCRIPT SMALL X
	case '\U0001D4CE':
		return "y", true // MATHEMATICAL SCRIPT SMALL Y
	case '\U0001D4CF':
		return "z", true // MATHEMATICAL SCRIPT SMALL Z
	case '\U0001D4D0':
		return "A", true // MATHEMATICAL BOLD SCRIPT CAPITAL A
	case '\U0001D4D1':
		return "B", true // MATHEMATICAL BOLD SCRIPT CAPITAL B
	case '\U0001D4D2':
		return "C", true // MATHEMATICAL BOLD SCRIPT CAPITAL C
	case '\U0001D4D3':
		return "D", true // MATHEMATICAL BOLD SCRIPT CAPITAL D
	case '\U0001D4D4':
		return "E", true // MATHEMATICAL BOLD SCRIPT CAPITAL E
	case '\U0001D4D5':
		return "F", true // MATHEMATICAL BOLD SCRIPT CAPITAL F
	case '\U0001D4D6':
		return "G", true // MATHEMATICAL BOLD SCRIPT CAPITAL G
	case '\U0001D4D7':
		return "H", true // MATHEMATICAL BOLD SCRIPT CAPITAL H
	case '\U0001D4D8':
		return "I", true // MATHEMATICAL BOLD SCRIPT CAPITAL I
	case '\U0001D4D9':
		return "J", true // MATHEMATICAL BOLD SCRIPT CAPITAL J
	case '\U0001D4DA':
		return "K", true // MATHEMATICAL BOLD SCRIPT CAPITAL K
	case '\U0001D4DB':
		return "L", true // MATHEMATICAL BOLD SCRIPT CAPITAL L
	case '\U0001D4DC':
		return "M", true // MATHEMATICAL BOLD SCRIPT CAPITAL M
	case '\U0001D4DD':
		return "N", true // MATHEMATICAL BOLD SCRIPT CAPITAL N
	case '\U0001D4DE':
		return "O", true // MATHEMATICAL BOLD SCRIPT CAPITAL O
	case '\U0001D4DF':
		return "P", true // MATHEMATICAL BOLD SCRIPT CAPITAL P
	case '\U0001D4E0':
		return "Q", true // MATHEMATICAL BOLD SCRIPT CAPITAL Q
	case '\U0001D4E1':
		return "R", true // MATHEMATICAL BOLD SCRIPT CAPITAL R
	case '\U0001D4E2':
		return "S", true // MATHEMATICAL BOLD SCRIPT CAPITAL S
	case '\U0001D4E3':
		return "T", true // MATHEMATICAL BOLD SCRIPT CAPITAL T
	case '\U0001D4E4':
		return "U", true // MATHEMATICAL BOLD SCRIPT CAPITAL U
	case '\U0001D4E5':
		return "V", true // MATHEMATICAL BOLD SCRIPT CAPITAL V
	case '\U0001D4E6':
		return "W", true // MATHEMATICAL BOLD SCRIPT CAPITAL W
	case '\U0001D4E7':
		return "X", true // MATHEMATICAL BOLD SCRIPT CAPITAL X
	case '\U0001D4E8':
		return "Y", true // MATHEMATICAL BOLD SCRIPT CAPITAL Y
	case '\U0001D4E9':
		return "Z", true // MATHEMATICAL BOLD SCRIPT CAPITAL Z
	case '\U0001D4EA':
		return "a", true // MATHEMATICAL BOLD SCRIPT SMALL A
	case '\U0001D4EB':
		return "b", true // MATHEMATICAL BOLD SCRIPT SMALL B
	case '\U0001D4EC':
		return "c", true // MATHEMATICAL BOLD SCRIPT SMALL C
	case '\U0001D4ED':
		return "d", true // MATHEMATICAL BOLD SCRIPT SMALL D
	case '\U0001D4EE':
		return "e", true // MATHEMATICAL BOLD SCRIPT SMALL E
	case '\U0001D4EF':
		return "f", true // MATHEMATICAL BOLD SCRIPT SMALL F
	case '\U0001D4F0':
		return "g", true // MATHEMATICAL BOLD SCRIPT SMALL G
	case '\U0001D4F1':
		return "h", true // MATHEMATICAL BOLD SCRIPT SMALL H
	case '\U0001D4F2':
		return "i", true // MATHEMATICAL BOLD SCRIPT SMALL I
	case '\U0001D4F3':
		return "j", true // MATHEMATICAL BOLD SCRIPT SMALL J
	case '\U0001D4F4':
		return "k", true // MATHEMATICAL BOLD SCRIPT SMALL K
	case '\U0001D4F5':
		return "l", true // MATHEMATICAL BOLD SCRIPT SMALL L
	case '\U0001D4F6':
		return "m", true // MATHEMATICAL BOLD SCRIPT SMALL M
	case '\U0001D4F7':
		return "n", true // MATHEMATICAL BOLD SCRIPT SMALL N
	case '\U0001D4F8':
		return "o", true // MATHEMATICAL BOLD SCRIPT SMALL O
	case '\U0001D4F9':
		return "p", true // MATHEMATICAL BOLD SCRIPT SMALL P
	case '\U0001D4FA':
		return "q", true // MATHEMATICAL BOLD SCRIPT SMALL Q
	case '\U0001D4FB':
		return "r", true // MATHEMATICAL BOLD SCRIPT SMALL R
	case '\U0001D4FC':
		return "s", true // MATHEMATICAL BOLD SCRIPT SMALL S
	case '\U0001D4FD':
		return "t", true // MATHEMATICAL BOLD SCRIPT SMALL T
	case '\U0001D4FE':
		return "u", true // MATHEMATICAL BOLD SCRIPT SMALL U
	case '\U0001D4FF':
		return "v", true // MATHEMATICAL BOLD SCRIPT SMALL V
	case '\U0001D500':
		return "w", true // MATHEMATICAL BOLD SCRIPT SMALL W
	case '\U0001D501':
		return "x", true // MATHEMATICAL BOLD SCRIPT SMALL X
	case '\U0001D502':
		return "y", true // MATHEMATICAL BOLD SCRIPT SMALL Y
	case '\U0001D503':
		return "z", true // MATHEMATICAL BOLD SCRIPT SMALL Z
	case '\U0001D504':
		return "A", true // MATHEMATICAL FRAKTUR CAPITAL A
	case '\U0001D505':
		return "B", true // MATHEMATICAL FRAKTUR CAPITAL B
	case '\U0001D507':
		return "D", true // MATHEMATICAL FRAKTUR CAPITAL D
	case '\U0001D508':
		return "E", true // MATHEMATICAL FRAKTUR CAPITAL E
	case '\U0001D509':
		return "F", true // MATHEMATICAL FRAKTUR CAPITAL F
	case '\U0001D50A':
		return "G", true // MATHEMATICAL FRAKTUR CAPITAL G
	case '\U0001D50D':
		return "J", true // MATHEMATICAL FRAKTUR CAPITAL J
	case '\U0001D50E':
		return "K", true // MATHEMATICAL FRAKTUR CAPITAL K
	case '\U0001D50F':
		return "L", true // MATHEMATICAL FRAKTUR CAPITAL L
	case '\U0001D510':
		return "M", true // MATHEMATICAL FRAKTUR CAPITAL M
	case '\U0001D511':
		return "N", true // MATHEMATICAL FRAKTUR CAPITAL N
	case '\U0001D512':
		return "O", true // MATHEMATICAL FRAKTUR CAPITAL O
	case '\U0001D513':
		return "P", true // MATHEMATICAL FRAKTUR CAPITAL P
	case '\U0001D514':
		return "Q", true // MATHEMATICAL FRAKTUR CAPITAL Q
	case '\U0001D516':
		return "S", true // MATHEMATICAL FRAKTUR CAPITAL S
	case '\U0001D517':
		return "T", true // MATHEMATICAL FRAKTUR CAPITAL T
	case '\U0001D518':
		return "U", true // MATHEMATICAL FRAKTUR CAPITAL U
	case '\U0001D519':
		return "V", true // MATHEMATICAL FRAKTUR CAPITAL V
	case '\U0001D51A':
		return "W", true // MATHEMATICAL FRAKTUR CAPITAL W
	case '\U0001D51B':
		return "X", true // MATHEMATICAL FRAKTUR CAPITAL X
	case '\U0001D51C':
		return "Y", true // MATHEMATICAL FRAKTUR CAPITAL Y
	case '\U0001D51E':
		return "a", true // MATHEMATICAL FRAKTUR SMALL A
	case '\U0001D51F':
		return "b", true // MATHEMATICAL FRAKTUR SMALL B
	case '\U0001D520':
		return "c", true // MATHEMATICAL FRAKTUR SMALL C
	case '\U0001D521':
		return "d", true // MATHEMATICAL FRAKTUR SMALL D
	case '\U0001D522':
		return "e", true // MATHEMATICAL FRAKTUR SMALL E
	case '\U0001D523':
		return "f", true // MATHEMATICAL FRAKTUR SMALL F
	case '\U0001D524':
		return "g", true // MATHEMATICAL FRAKTUR SMALL G
	case '\U0001D525':
		return "h", true // MATHEMATICAL FRAKTUR SMALL H
	case '\U0001D526':
		return "i", true // MATHEMATICAL FRAKTUR SMALL I
	case '\U0001D527':
		return "j", true // MATHEMATICAL FRAKTUR SMALL J
	case '\U0001D528':
		return "k", true // MATHEMATICAL FRAKTUR SMALL K
	case '\U0001D529':
		return "l", true // MATHEMATICAL FRAKTUR SMALL L
	case '\U0001D52A':
		return "m", true // MATHEMATICAL FRAKTUR SMALL M
	case '\U0001D52B':
		return "n", true // MATHEMATICAL FRAKTUR SMALL N
	case '\U0001D52C':
		return "o", true // MATHEMATICAL FRAKTUR SMALL O
	case '\U0001D52D':
		return "p", true // MATHEMATICAL FRAKTUR SMALL P
	case '\U0001D52E':
		return "q", true // MATHEMATICAL FRAKTUR SMALL Q
	case '\U0001D52F':
		return "r", true // MATHEMATICAL FRAKTUR SMALL R
	case '\U0001D530':
		return "s", true // MATHEMATICAL FRAKTUR SMALL S
	case '\U0001D531':
		return "t", true // MATHEMATICAL FRAKTUR SMALL T
	case '\U0001D532':
		return "u", true // MATHEMATICAL FRAKTUR SMALL U
	case '\U0001D533':
		return "v", true // MATHEMATICAL FRAKTUR SMALL V
	case '\U0001D534':
		return "w", true // MATHEMATICAL FRAKTUR SMALL W
	case '\U0001D535':
		return "x", true // MATHEMATICAL FRAKTUR SMALL X
	case '\U0001D536':
		return "y", true // MATHEMATICAL FRAKTUR SMALL Y
	case '\U0001D537':
		return "z", true // MATHEMATICAL FRAKTUR SMALL Z
	case '\U0001D538':
		return "A", true // MATHEMATICAL DOUBLE-STRUCK CAPITAL A
	case '\U0001D539':
		return "B", true // MATHEMATICAL DOUBLE-STRUCK CAPITAL B
	case '\U0001D53B':
		return "D", true // MATHEMATICAL DOUBLE-STRUCK CAPITAL D
	case '\U0001D53C':
		return "E", true // MATHEMATICAL DOUBLE-STRUCK CAPITAL E
	case '\U0001D53D':
		return "F", true // MATHEMATICAL DOUBLE-STRUCK CAPITAL F
	case '\U0001D53E':
		return "G", true // MATHEMATICAL DOUBLE-STRUCK CAPITAL G
	case '\U0001D540':
		return "I", true // MATHEMATICAL DOUBLE-STRUCK CAPITAL I
	case '\U0001D541':
		return "J", true // MATHEMATICAL DOUBLE-STRUCK CAPITAL J
	case '\U0001D542':
		return "K", true // MATHEMATICAL DOUBLE-STRUCK CAPITAL K
	case '\U0001D543':
		return "L", true // MATHEMATICAL DOUBLE-STRUCK CAPITAL L
	case '\U0001D544':
		return "M", true // MATHEMATICAL DOUBLE-STRUCK CAPITAL M
	case '\U0001D546':
		return "O", true // MATHEMATICAL DOUBLE-STRUCK CAPITAL O
	case '\U0001D54A':
		return "S", true // MATHEMATICAL DOUBLE-STRUCK CAPITAL S
	case '\U0001D54B':
		return "T", true // MATHEMATICAL DOUBLE-STRUCK CAPITAL T
	case '\U0001D54C':
		return "U", true // MATHEMATICAL DOUBLE-STRUCK CAPITAL U
	case '\U0001D54D':
		return "V", true // MATHEMATICAL DOUBLE-STRUCK CAPITAL V
	case '\U0001D54E':
		return "W", true // MATHEMATICAL DOUBLE-STRUCK CAPITAL W
	case '\U0001D54F':
		return "X", true // MATHEMATICAL DOUBLE-STRUCK CAPITAL X
	case '\U0001D550':
		return "Y", true // MATHEMATICAL DOUBLE-STRUCK CAPITAL Y
	case '\U0001D552':
		return "a", true // MATHEMATICAL DOUBLE-STRUCK SMALL A
	case '\U0001D553':
		return "b", true // MATHEMATICAL DOUBLE-STRUCK SMALL B
	case '\U0001D554':
		return "c", true // MATHEMATICAL DOUBLE-STRUCK SMALL C
	case '\U0001D555':
		return "d", true // MATHEMATICAL DOUBLE-STRUCK SMALL D
	case '\U0001D556':
		return "e", true // MATHEMATICAL DOUBLE-STRUCK SMALL E
	case '\U0001D557':
		return "f", true // MATHEMATICAL DOUBLE-STRUCK SMALL F
	case '\U0001D558':
		return "g", true // MATHEMATICAL DOUBLE-STRUCK SMALL G
	case '\U0001D559':
		return "h", true // MATHEMATICAL DOUBLE-STRUCK SMALL H
	case '\U0001D55A':
		return "i", true // MATHEMATICAL DOUBLE-STRUCK SMALL I
	case '\U0001D55B':
		return "j", true // MATHEMATICAL DOUBLE-STRUCK SMALL J
	case '\U0001D55C':
		return "k", true // MATHEMATICAL DOUBLE-STRUCK SMALL K
	case '\U0001D55D':
		return "l", true // MATHEMATICAL DOUBLE-STRUCK SMALL L
	case '\U0001D55E':
		return "m", true // MATHEMATICAL DOUBLE-STRUCK SMALL M
	case '\U0001D55F':
		return "n", true // MATHEMATICAL DOUBLE-STRUCK SMALL N
	case '\U0001D560':
		return "o", true // MATHEMATICAL DOUBLE-STRUCK SMALL O
	case '\U0001D561':
		return "p", true // MATHEMATICAL DOUBLE-STRUCK SMALL P
	case '\U0001D562':
		return "q", true // MATHEMATICAL DOUBLE-STRUCK SMALL Q
	case '\U0001D563':
		return "r", true // MATHEMATICAL DOUBLE-STRUCK SMALL R
	case '\U0001D564':
		return "s", true // MATHEMATICAL DOUBLE-STRUCK SMALL S
	case '\U0001D565':
		return "t", true // MATHEMATICAL DOUBLE-STRUCK SMALL T
	case '\U0001D566':
		return "u", true // MATHEMATICAL DOUBLE-STRUCK SMALL U
	case '\U0001D567':
		return "v", true // MATHEMATICAL DOUBLE-STRUCK SMALL V
	case '\U0001D568':
		return "w", true // MATHEMATICAL DOUBLE-STRUCK SMALL W
	case '\U0001D569':
		return "x", true // MATHEMATICAL DOUBLE-STRUCK SMALL X
	case '\U0001D56A':
		return "y", true // MATHEMATICAL DOUBLE-STRUCK SMALL Y
	case '\U0001D56B':
		return "z", true // MATHEMATICAL DOUBLE-STRUCK SMALL Z
	case '\U0001D56C':
		return "A", true // MATHEMATICAL BOLD FRAKTUR CAPITAL A
	case '\U0001D56D':
		return "B", true // MATHEMATICAL BOLD FRAKTUR CAPITAL B
	case '\U0001D56E':
		return "C", true // MATHEMATICAL BOLD FRAKTUR CAPITAL C
	case '\U0001D56F':
		return "D", true // MATHEMATICAL BOLD FRAKTUR CAPITAL D
	case '\U0001D570':
		return "E", true // MATHEMATICAL BOLD FRAKTUR CAPITAL E
	case '\U0001D571':
		return "F", true // MATHEMATICAL BOLD FRAKTUR CAPITAL F
	case '\U0001D572':
		return "G", true // MATHEMATICAL BOLD FRAKTUR CAPITAL G
	case '\U0001D573':
		return "H", true // MATHEMATICAL BOLD FRAKTUR CAPITAL H
	case '\U0001D574':
		return "I", true // MATHEMATICAL BOLD FRAKTUR CAPITAL I
	case '\U0001D575':
		return "J", true // MATHEMATICAL BOLD FRAKTUR CAPITAL J
	case '\U0001D576':
		return "K", true // MATHEMATICAL BOLD FRAKTUR CAPITAL K
	case '\U0001D577':
		return "L", true // MATHEMATICAL BOLD FRAKTUR CAPITAL L
	case '\U0001D578':
		return "M", true // MATHEMATICAL BOLD FRAKTUR CAPITAL M
	case '\U0001D579':
		return "N", true // MATHEMATICAL BOLD FRAKTUR CAPITAL N
	case '\U0001D57A':
		return "O", true // MATHEMATICAL BOLD FRAKTUR CAPITAL O
	case '\U0001D57B':
		return "P", true // MATHEMATICAL BOLD FRAKTUR CAPITAL P
	case '\U0001D57C':
		return "Q", true // MATHEMATICAL BOLD FRAKTUR CAPITAL Q
	case '\U0001D57D':
		return "R", true // MATHEMATICAL BOLD FRAKTUR CAPITAL R
	case '\U0001D57E':
		return "S", true // MATHEMATICAL BOLD FRAKTUR CAPITAL S
	case '\U0001D57F':
		return "T", true // MATHEMATICAL BOLD FRAKTUR CAPITAL T
	case '\U0001D580':
		return "U", true // MATHEMATICAL BOLD FRAKTUR CAPITAL U
	case '\U0001D581':
		return "V", true // MATHEMATICAL BOLD FRAKTUR CAPITAL V
	case '\U0001D582':
		return "W", true // MATHEMATICAL BOLD FRAKTUR CAPITAL W
	case '\U0001D583':
		return "X", true // MATHEMATICAL BOLD FRAKTUR CAPITAL X
	case '\U0001D584':
		return "Y", true // MATHEMATICAL BOLD FRAKTUR CAPITAL Y
	case '\U0001D585':
		return "Z", true // MATHEMATICAL BOLD FRAKTUR CAPITAL Z
	case '\U0001D586':
		return "a", true // MATHEMATICAL BOLD FRAKTUR SMALL A
	case '\U0001D587':
		return "b", true // MATHEMATICAL BOLD FRAKTUR SMALL B
	case '\U0001D588':
		return "c", true // MATHEMATICAL BOLD FRAKTUR SMALL C
	case '\U0001D589':
		return "d", true // MATHEMATICAL BOLD FRAKTUR SMALL D
	case '\U0001D58A':
		return "e", true // MATHEMATICAL BOLD FRAKTUR SMALL E
	case '\U0001D58B':
		return "f", true // MATHEMATICAL BOLD FRAKTUR SMALL F
	case '\U0001D58C':
		return "g", true // MATHEMATICAL BOLD FRAKTUR SMALL G
	case '\U0001D58D':
		return "h", true // MATHEMATICAL BOLD FRAKTUR SMALL H
	case '\U0001D58E':
		return "i", true // MATHEMATICAL BOLD FRAKTUR SMALL I
	case '\U0001D58F':
		return "j", true // MATHEMATICAL BOLD FRAKTUR SMALL J
	case '\U0001D590':
		return "k", true // MATHEMATICAL BOLD FRAKTUR SMALL K
	case '\U0001D591':
		return "l", true // MATHEMATICAL BOLD FRAKTUR SMALL L
	case '\U0001D592':
		return "m", true // MATHEMATICAL BOLD FRAKTUR SMALL M
	case '\U0001D593':
		return "n", true // MATHEMATICAL BOLD FRAKTUR SMALL N
	case '\U0001D594':
		return "o", true // MATHEMATICAL BOLD FRAKTUR SMALL O
	case '\U0001D595':
		return "p", true // MATHEMATICAL BOLD FRAKTUR SMALL P
	case '\U0001D596':
		return "q", true // MATHEMATICAL BOLD FRAKTUR SMALL Q
	case '\U0001D597':
		return "r", true // MATHEMATICAL BOLD FRAKTUR SMALL R
	case '\U0001D598':
		return "s", true // MATHEMATICAL BOLD FRAKTUR SMALL S
	case '\U0001D599':
		return "t", true // MATHEMATICAL BOLD FRAKTUR SMALL T
	case '\U0001D59A':
		return "u", true // MATHEMATICAL BOLD FRAKTUR SMALL U
	case '\U0001D59B':
		return "v", true // MATHEMATICAL BOLD FRAKTUR SMALL V
	case '\U0001D59C':
		return "w", true // MATHEMATICAL BOLD FRAKTUR SMALL W
	case '\U0001D59D':
		return "x", true // MATHEMATICAL BOLD FRAKTUR SMALL X
	case '\U0001D59E':
		return "y", true // MATHEMATICAL BOLD FRAKTUR SMALL Y
	case '\U0001D59F':
		return "z", true // MATHEMATICAL BOLD FRAKTUR SMALL Z
	case '\U0001D5A0':
		return "A", true // MATHEMATICAL SANS-SERIF CAPITAL A
	case '\U0001D5A1':
		return "B", true // MATHEMATICAL SANS-SERIF CAPITAL B
	case '\U0001D5A2':
		return "C", true // MATHEMATICAL SANS-SERIF CAPITAL C
	case '\U0001D5A3':
		return "D", true // MATHEMATICAL SANS-SERIF CAPITAL D
	case '\U0001D5A4':
		return "E", true // MATHEMATICAL SANS-SERIF CAPITAL E
	case '\U0001D5A5':
		return "F", true // MATHEMATICAL SANS-SERIF CAPITAL F
	case '\U0001D5A6':
		return "G", true // MATHEMATICAL SANS-SERIF CAPITAL G
	case '\U0001D5A7':
		return "H", true // MATHEMATICAL SANS-SERIF CAPITAL H
	case '\U0001D5A8':
		return "I", true // MATHEMATICAL SANS-SERIF CAPITAL I
	case '\U0001D5A9':
		return "J", true // MATHEMATICAL SANS-SERIF CAPITAL J
	case '\U0001D5AA':
		return "K", true // MATHEMATICAL SANS-SERIF CAPITAL K
	case '\U0001D5AB':
		return "L", true // MATHEMATICAL SANS-SERIF CAPITAL L
	case '\U0001D5AC':
		return "M", true // MATHEMATICAL SANS-SERIF CAPITAL M
	case '\U0001D5AD':
		return "N", true // MATHEMATICAL SANS-SERIF CAPITAL N
	case '\U0001D5AE':
		return "O", true // MATHEMATICAL SANS-SERIF CAPITAL O
	case '\U0001D5AF':
		return "P", true // MATHEMATICAL SANS-SERIF CAPITAL P
	case '\U0001D5B0':
		return "Q", true // MATHEMATICAL SANS-SERIF CAPITAL Q
	case '\U0001D5B1':
		return "R", true // MATHEMATICAL SANS-SERIF CAPITAL R
	case '\U0001D5B2':
		return "S", true // MATHEMATICAL SANS-SERIF CAPITAL S
	case '\U0001D5B3':
		return "T", true // MATHEMATICAL SANS-SERIF CAPITAL T
	case '\U0001D5B4':
		return "U", true // MATHEMATICAL SANS-SERIF CAPITAL U
	case '\U0001D5B5':
		return "V", true // MATHEMATICAL SANS-SERIF CAPITAL V
	case '\U0001D5B6':
		return "W", true // MATHEMATICAL SANS-SERIF CAPITAL W
	case '\U0001D5B7':
		return "X", true // MATHEMATICAL SANS-SERIF CAPITAL X
	case '\U0001D5B8':
		return "Y", true // MATHEMATICAL SANS-SERIF CAPITAL Y
	case '\U0001D5B9':
		return "Z", true // MATHEMATICAL SANS-SERIF CAPITAL Z
	case '\U0001D5BA':
		return "a", true // MATHEMATICAL SANS-SERIF SMALL A
	case '\U0001D5BB':
		return "b", true // MATHEMATICAL SANS-SERIF SMALL B
	case '\U0001D5BC':
		return "c", true // MATHEMATICAL SANS-SERIF SMALL C
	case '\U0001D5BD':
		return "d", true // MATHEMATICAL SANS-SERIF SMALL D
	case '\U0001D5BE':
		return "e", true // MATHEMATICAL SANS-SERIF SMALL E
	case '\U0001D5BF':
		return "f", true // MATHEMATICAL SANS-SERIF SMALL F
	case '\U0001D5C0':
		return "g", true // MATHEMATICAL SANS-SERIF SMALL G
	case '\U0001D5C1':
		return "h", true // MATHEMATICAL SANS-SERIF SMALL H
	case '\U0001D5C2':
		return "i", true // MATHEMATICAL SANS-SERIF SMALL I
	case '\U0001D5C3':
		return "j", true // MATHEMATICAL SANS-SERIF SMALL J
	case '\U0001D5C4':
		return "k", true // MATHEMATICAL SANS-SERIF SMALL K
	case '\U0001D5C5':
		return "l", true // MATHEMATICAL SANS-SERIF SMALL L
	case '\U0001D5C6':
		return "m", true // MATHEMATICAL SANS-SERIF SMALL M
	case '\U0001D5C7':
		return "n", true // MATHEMATICAL SANS-SERIF SMALL N
	case '\U0001D5C8':
		return "o", true // MATHEMATICAL SANS-SERIF SMALL O
	case '\U0001D5C9':
		return "p", true // MATHEMATICAL SANS-SERIF SMALL P
	case '\U0001D5CA':
		return "q", true // MATHEMATICAL SANS-SERIF SMALL Q
	case '\U0001D5CB':
		return "r", true // MATHEMATICAL SANS-SERIF SMALL R
	case '\U0001D5CC':
		return "s", true // MATHEMATICAL SANS-SERIF SMALL S
	case '\U0001D5CD':
		return "t", true // MATHEMATICAL SANS-SERIF SMALL T
	case '\U0001D5CE':
		return "u", true // MATHEMATICAL SANS-SERIF SMALL U
	case '\U0001D5CF':
		return "v", true // MATHEMATICAL SANS-SERIF SMALL V
	case '\U0001D5D0':
		return "w", true // MATHEMATICAL SANS-SERIF SMALL W
	case '\U0001D5D1':
		return "x", true // MATHEMATICAL SANS-SERIF SMALL X
	case '\U0001D5D2':
		return "y", true // MATHEMATICAL SANS-SERIF SMALL Y
	case '\U0001D5D3':
		return "z", true // MATHEMATICAL SANS-SERIF SMALL Z
	case '\U0001D5D4':
		return "A", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL A
	case '\U0001D5D5':
		return "B", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL B
	case '\U0001D5D6':
		return "C", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL C
	case '\U0001D5D7':
		return "D", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL D
	case '\U0001D5D8':
		return "E", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL E
	case '\U0001D5D9':
		return "F", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL F
	case '\U0001D5DA':
		return "G", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL G
	case '\U0001D5DB':
		return "H", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL H
	case '\U0001D5DC':
		return "I", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL I
	case '\U0001D5DD':
		return "J", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL J
	case '\U0001D5DE':
		return "K", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL K
	case '\U0001D5DF':
		return "L", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL L
	case '\U0001D5E0':
		return "M", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL M
	case '\U0001D5E1':
		return "N", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL N
	case '\U0001D5E2':
		return "O", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL O
	case '\U0001D5E3':
		return "P", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL P
	case '\U0001D5E4':
		return "Q", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL Q
	case '\U0001D5E5':
		return "R", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL R
	case '\U0001D5E6':
		return "S", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL S
	case '\U0001D5E7':
		return "T", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL T
	case '\U0001D5E8':
		return "U", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL U
	case '\U0001D5E9':
		return "V", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL V
	case '\U0001D5EA':
		return "W", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL W
	case '\U0001D5EB':
		return "X", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL X
	case '\U0001D5EC':
		return "Y", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL Y
	case '\U0001D5ED':
		return "Z", true // MATHEMATICAL SANS-SERIF BOLD CAPITAL Z
	case '\U0001D5EE':
		return "a", true // MATHEMATICAL SANS-SERIF BOLD SMALL A
	case '\U0001D5EF':
		return "b", true // MATHEMATICAL SANS-SERIF BOLD SMALL B
	case '\U0001D5F0':
		return "c", true // MATHEMATICAL SANS-SERIF BOLD SMALL C
	case '\U0001D5F1':
		return "d", true // MATHEMATICAL SANS-SERIF BOLD SMALL D
	case '\U0001D5F2':
		return "e", true // MATHEMATICAL SANS-SERIF BOLD SMALL E
	case '\U0001D5F3':
		return "f", true // MATHEMATICAL SANS-SERIF BOLD SMALL F
	case '\U0001D5F4':
		return "g", true // MATHEMATICAL SANS-SERIF BOLD SMALL G
	case '\U0001D5F5':
		return "h", true // MATHEMATICAL SANS-SERIF BOLD SMALL H
	case '\U0001D5F6':
		return "i", true // MATHEMATICAL SANS-SERIF BOLD SMALL I
	case '\U0001D5F7':
		return "j", true // MATHEMATICAL SANS-SERIF BOLD SMALL J
	case '\U0001D5F8':
		return "k", true // MATHEMATICAL SANS-SERIF BOLD SMALL K
	case '\U0001D5F9':
		return "l", true // MATHEMATICAL SANS-SERIF BOLD SMALL L
	case '\U0001D5FA':
		return "m", true // MATHEMATICAL SANS-SERIF BOLD SMALL M
	case '\U0001D5FB':
		return "n", true // MATHEMATICAL SANS-SERIF BOLD SMALL N
	case '\U0001D5FC':
		return "o", true // MATHEMATICAL SANS-SERIF BOLD SMALL O
	case '\U0001D5FD':
		return "p", true // MATHEMATICAL SANS-SERIF BOLD SMALL P
	case '\U0001D5FE':
		return "q", true // MATHEMATICAL SANS-SERIF BOLD SMALL Q
	case '\U0001D5FF':
		return "r", true // MATHEMATICAL SANS-SERIF BOLD SMALL R
	case '\U0001D600':
		return "s", true // MATHEMATICAL SANS-SERIF BOLD SMALL S
	case '\U0001D601':
		return "t", true // MATHEMATICAL SANS-SERIF BOLD SMALL T
	case '\U0001D602':
		return "u", true // MATHEMATICAL SANS-SERIF BOLD SMALL U
	case '\U0001D603':
		return "v", true // MATHEMATICAL SANS-SERIF BOLD SMALL V
	case '\U0001D604':
		return "w", true // MATHEMATICAL SANS-SERIF BOLD SMALL W
	case '\U0001D605':
		return "x", true // MATHEMATICAL SANS-SERIF BOLD SMALL X
	case '\U0001D606':
		return "y", true // MATHEMATICAL SANS-SERIF BOLD SMALL Y
	case '\U0001D607':
		return "z", true // MATHEMATICAL SANS-SERIF BOLD SMALL Z
	case '\U0001D608':
		return "A", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL A
	case '\U0001D609':
		return "B", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL B
	case '\U0001D60A':
		return "C", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL C
	case '\U0001D60B':
		return "D", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL D
	case '\U0001D60C':
		return "E", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL E
	case '\U0001D60D':
		return "F", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL F
	case '\U0001D60E':
		return "G", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL G
	case '\U0001D60F':
		return "H", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL H
	case '\U0001D610':
		return "I", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL I
	case '\U0001D611':
		return "J", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL J
	case '\U0001D612':
		return "K", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL K
	case '\U0001D613':
		return "L", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL L
	case '\U0001D614':
		return "M", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL M
	case '\U0001D615':
		return "N", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL N
	case '\U0001D616':
		return "O", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL O
	case '\U0001D617':
		return "P", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL P
	case '\U0001D618':
		return "Q", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL Q
	case '\U0001D619':
		return "R", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL R
	case '\U0001D61A':
		return "S", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL S
	case '\U0001D61B':
		return "T", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL T
	case '\U0001D61C':
		return "U", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL U
	case '\U0001D61D':
		return "V", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL V
	case '\U0001D61E':
		return "W", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL W
	case '\U0001D61F':
		return "X", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL X
	case '\U0001D620':
		return "Y", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL Y
	case '\U0001D621':
		return "Z", true // MATHEMATICAL SANS-SERIF ITALIC CAPITAL Z
	case '\U0001D622':
		return "a", true // MATHEMATICAL SANS-SERIF ITALIC SMALL A
	case '\U0001D623':
		return "b", true // MATHEMATICAL SANS-SERIF ITALIC SMALL B
	case '\U0001D624':
		return "c", true // MATHEMATICAL SANS-SERIF ITALIC SMALL C
	case '\U0001D625':
		return "d", true // MATHEMATICAL SANS-SERIF ITALIC SMALL D
	case '\U0001D626':
		return "e", true // MATHEMATICAL SANS-SERIF ITALIC SMALL E
	case '\U0001D627':
		return "f", true // MATHEMATICAL SANS-SERIF ITALIC SMALL F
	case '\U0001D628':
		return "g", true // MATHEMATICAL SANS-SERIF ITALIC SMALL G
	case '\U0001D629':
		return "h", true // MATHEMATICAL SANS-SERIF ITALIC SMALL H
	case '\U0001D62A':
		return "i", true // MATHEMATICAL SANS-SERIF ITALIC SMALL I
	case '\U0001D62B':
		return "j", true // MATHEMATICAL SANS-SERIF ITALIC SMALL J
	case '\U0001D62C':
		return "k", true // MATHEMATICAL SANS-SERIF ITALIC SMALL K
	case '\U0001D62D':
		return "l", true // MATHEMATICAL SANS-SERIF ITALIC SMALL L
	case '\U0001D62E':
		return "m", true // MATHEMATICAL SANS-SERIF ITALIC SMALL M
	case '\U0001D62F':
		return "n", true // MATHEMATICAL SANS-SERIF ITALIC SMALL N
	case '\U0001D630':
		return "o", true // MATHEMATICAL SANS-SERIF ITALIC SMALL O
	case '\U0001D631':
		return "p", true // MATHEMATICAL SANS-SERIF ITALIC SMALL P
	case '\U0001D632':
		return "q", true // MATHEMATICAL SANS-SERIF ITALIC SMALL Q
	case '\U0001D633':
		return "r", true // MATHEMATICAL SANS-SERIF ITALIC SMALL R
	case '\U0001D634':
		return "s", true // MATHEMATICAL SANS-SERIF ITALIC SMALL S
	case '\U0001D635':
		return "t", true // MATHEMATICAL SANS-SERIF ITALIC SMALL T
	case '\U0001D636':
		return "u", true // MATHEMATICAL SANS-SERIF ITALIC SMALL U
	case '\U0001D637':
		return "v", true // MATHEMATICAL SANS-SERIF ITALIC SMALL V
	case '\U0001D638':
		return "w", true // MATHEMATICAL SANS-SERIF ITALIC SMALL W
	case '\U0001D639':
		return "x", true // MATHEMATICAL SANS-SERIF ITALIC SMALL X
	case '\U0001D63A':
		return "y", true // MATHEMATICAL SANS-SERIF ITALIC SMALL Y
	case '\U0001D63B':
		return "z", true // MATHEMATICAL SANS-SERIF ITALIC SMALL Z
	case '\U0001D63C':
		return "A", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL A
	case '\U0001D63D':
		return "B", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL B
	case '\U0001D63E':
		return "C", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL C
	case '\U0001D63F':
		return "D", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL D
	case '\U0001D640':
		return "E", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL E
	case '\U0001D641':
		return "F", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL F
	case '\U0001D642':
		return "G", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL G
	case '\U0001D643':
		return "H", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL H
	case '\U0001D644':
		return "I", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL I
	case '\U0001D645':
		return "J", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL J
	case '\U0001D646':
		return "K", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL K
	case '\U0001D647':
		return "L", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL L
	case '\U0001D648':
		return "M", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL M
	case '\U0001D649':
		return "N", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL N
	case '\U0001D64A':
		return "O", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL O
	case '\U0001D64B':
		return "P", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL P
	case '\U0001D64C':
		return "Q", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL Q
	case '\U0001D64D':
		return "R", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL R
	case '\U0001D64E':
		return "S", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL S
	case '\U0001D64F':
		return "T", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL T
	case '\U0001D650':
		return "U", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL U
	case '\U0001D651':
		return "V", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL V
	case '\U0001D652':
		return "W", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL W
	case '\U0001D653':
		return "X", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL X
	case '\U0001D654':
		return "Y", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL Y
	case '\U0001D655':
		return "Z", true // MATHEMATICAL SANS-SERIF BOLD ITALIC CAPITAL Z
	case '\U0001D656':
		return "a", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL A
	case '\U0001D657':
		return "b", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL B
	case '\U0001D658':
		return "c", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL C
	case '\U0001D659':
		return "d", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL D
	case '\U0001D65A':
		return "e", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL E
	case '\U0001D65B':
		return "f", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL F
	case '\U0001D65C':
		return "g", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL G
	case '\U0001D65D':
		return "h", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL H
	case '\U0001D65E':
		return "i", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL I
	case '\U0001D65F':
		return "j", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL J
	case '\U0001D660':
		return "k", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL K
	case '\U0001D661':
		return "l", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL L
	case '\U0001D662':
		return "m", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL M
	case '\U0001D663':
		return "n", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL N
	case '\U0001D664':
		return "o", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL O
	case '\U0001D665':
		return "p", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL P
	case '\U0001D666':
		return "q", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL Q
	case '\U0001D667':
		return "r", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL R
	case '\U0001D668':
		return "s", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL S
	case '\U0001D669':
		return "t", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL T
	case '\U0001D66A':
		return "u", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL U
	case '\U0001D66B':
		return "v", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL V
	case '\U0001D66C':
		return "w", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL W
	case '\U0001D66D':
		return "x", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL X
	case '\U0001D66E':
		return "y", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL Y
	case '\U0001D66F':
		return "z", true // MATHEMATICAL SANS-SERIF BOLD ITALIC SMALL Z
	case '\U0001D670':
		return "A", true // MATHEMATICAL MONOSPACE CAPITAL A
	case '\U0001D671':
		return "B", true // MATHEMATICAL MONOSPACE CAPITAL B
	case '\U0001D672':
		return "C", true // MATHEMATICAL MONOSPACE CAPITAL C
	case '\U0001D673':
		return "D", true // MATHEMATICAL MONOSPACE CAPITAL D
	case '\U0001D674':
		return "E", true // MATHEMATICAL MONOSPACE CAPITAL E
	case '\U0001D675':
		return "F", true // MATHEMATICAL MONOSPACE CAPITAL F
	case '\U0001D676':
		return "G", true // MATHEMATICAL MONOSPACE CAPITAL G
	case '\U0001D677':
		return "H", true // MATHEMATICAL MONOSPACE CAPITAL H
	case '\U0001D678':
		return "I", true // MATHEMATICAL MONOSPACE CAPITAL I
	case '\U0001D679':
		return "J", true // MATHEMATICAL MONOSPACE CAPITAL J
	case '\U0001D67A':
		return "K", true // MATHEMATICAL MONOSPACE CAPITAL K
	case '\U0001D67B':
		return "L", true // MATHEMATICAL MONOSPACE CAPITAL L
	case '\U0001D67C':
		return "M", true // MATHEMATICAL MONOSPACE CAPITAL M
	case '\U0001D67D':
		return "N", true // MATHEMATICAL MONOSPACE CAPITAL N
	case '\U0001D67E':
		return "O", true // MATHEMATICAL MONOSPACE CAPITAL O
	case '\U0001D67F':
		return "P", true // MATHEMATICAL MONOSPACE CAPITAL P
	case '\U0001D680':
		return "Q", true // MATHEMATICAL MONOSPACE CAPITAL Q
	case '\U0001D681':
		return "R", true // MATHEMATICAL MONOSPACE CAPITAL R
	case '\U0001D682':
		return "S", true // MATHEMATICAL MONOSPACE CAPITAL S
	case '\U0001D683':
		return "T", true // MATHEMATICAL MONOSPACE CAPITAL T
	case '\U0001D684':
		return "U", true // MATHEMATICAL MONOSPACE CAPITAL U
	case '\U0001D685':
		return "V", true // MATHEMATICAL MONOSPACE CAPITAL V
	case '\U0001D686':
		return "W", true // MATHEMATICAL MONOSPACE CAPITAL W
	case '\U0001D687':
		return "X", true // MATHEMATICAL MONOSPACE CAPITAL X
	case '\U0001D688':
		return "Y", true // MATHEMATICAL MONOSPACE CAPITAL Y
	case '\U0001D689':
		return "Z", true // MATHEMATICAL MONOSPACE CAPITAL Z
	case '\U0001D68A':
		return "a", true // MATHEMATICAL MONOSPACE SMALL A
	case '\U0001D68B':
		return "b", true // MATHEMATICAL MONOSPACE SMALL B
	case '\U0001D68C':
		return "c", true // MATHEMATICAL MONOSPACE SMALL C
	case '\U0001D68D':
		return "d", true // MATHEMATICAL MONOSPACE SMALL D
	case '\U0001D68E':
		return "e", true // MATHEMATICAL MONOSPACE SMALL E
	case '\U0001D68F':
		return "f", true // MATHEMATICAL MONOSPACE SMALL F
	case '\U0001D690':
		return "g", true // MATHEMATICAL MONOSPACE SMALL G
	case '\U0001D691':
		return "h", true // MATHEMATICAL MONOSPACE SMALL H
	case '\U0001D692':
		return "i", true // MATHEMATICAL MONOSPACE SMALL I
	case '\U0001D693':
		return "j", true // MATHEMATICAL MONOSPACE SMALL J
	case '\U0001D694':
		return "k", true // MATHEMATICAL MONOSPACE SMALL K
	case '\U0001D695':
		return "l", true // MATHEMATICAL MONOSPACE SMALL L
	case '\U0001D696':
		return "m", true // MATHEMATICAL MONOSPACE SMALL M
	case '\U0001D697':
		return "n", true // MATHEMATICAL MONOSPACE SMALL N
	case '\U0001D698':
		return "o", true // MATHEMATICAL MONOSPACE SMALL O
	case '\U0001D699':
		return "p", true // MATHEMATICAL MONOSPACE SMALL P
	case '\U0001D69A':
		return "q", true // MATHEMATICAL MONOSPACE SMALL Q
	case '\U0001D69B':
		return "r", true // MATHEMATICAL MONOSPACE SMALL R
	case '\U0001D69C':
		return "s", true // MATHEMATICAL MONOSPACE SMALL S
	case '\U0001D69D':
		return "t", true // MATHEMATICAL MONOSPACE SMALL T
	case '\U0001D69E':
		return "u", true // MATHEMATICAL MONOSPACE SMALL U
	case '\U0001D69F':
		return "v", true // MATHEMATICAL MONOSPACE SMALL V
	case '\U0001D6A0':
		return "w", true // MATHEMATICAL MONOSPACE SMALL W
	case '\U0001D6A1':
		return "x", true // MATHEMATICAL MONOSPACE SMALL X
	case '\U0001D6A2':
		return "y", true // MATHEMATICAL MONOSPACE SMALL Y
	case '\U0001D6A3':
		return "z", true // MATHEMATICAL MONOSPACE SMALL Z
	case '\U0001D7CE':
		return "0", true // MATHEMATICAL BOLD DIGIT ZERO
	case '\U0001D7CF':
		return "1", true // MATHEMATICAL BOLD DIGIT ONE
	case '\U0001D7D0':
		return "2", true // MATHEMATICAL BOLD DIGIT TWO
	case '\U0001D7D1':
		return "3", true // MATHEMATICAL BOLD DIGIT THREE
	case '\U0001D7D2':
		return "4", true // MATHEMATICAL BOLD DIGIT FOUR
	case '\U0001D7D3':
		return "5", true // MATHEMATICAL BOLD DIGIT FIVE
	case '\U0001D7D4':
		return "6", true // MATHEMATICAL BOLD DIGIT SIX
	case '\U0001D7D5':
		return "7", true // MATHEMATICAL BOLD DIGIT SEVEN
	case '\U0001D7D6':
		return "8", true // MATHEMATICAL BOLD DIGIT EIGHT
	case '\U0001D7D7':
		return "9", true // MATHEMATICAL BOLD DIGIT NINE
	case '\U0001D7D8':
		return "0", true // MATHEMATICAL DOUBLE-STRUCK DIGIT ZERO
	case '\U0001D7D9':
		return "1", true // MATHEMATICAL DOUBLE-STRUCK DIGIT ONE
	case '\U0001D7DA':
		return "2", true // MATHEMATICAL DOUBLE-STRUCK DIGIT TWO
	case '\U0001D7DB':
		return "3", true // MATHEMATICAL DOUBLE-STRUCK DIGIT THREE
	case '\U0001D7DC':
		return "4", true // MATHEMATICAL DOUBLE-STRUCK DIGIT FOUR
	case '\U0001D7DD':
		return "5", true // MATHEMATICAL DOUBLE-STRUCK DIGIT FIVE
	case '\U0001D7DE':
		return "6", true // MATHEMATICAL DOUBLE-STRUCK DIGIT SIX
	case '\U0001D7DF':
		return "7", true // MATHEMATICAL DOUBLE-STRUCK DIGIT SEVEN
	case '\U0001D7E0':
		return "8", true // MATHEMATICAL DOUBLE-STRUCK DIGIT EIGHT
	case '\U0001D7E1':
		return "9", true // MATHEMATICAL DOUBLE-STRUCK DIGIT NINE
	case '\U0001D7E2':
		return "0", true // MATHEMATICAL SANS-SERIF DIGIT ZERO
	case '\U0001D7E3':
		return "1", true // MATHEMATICAL SANS-SERIF DIGIT ONE
	case '\U0001D7E4':
		return "2", true // MATHEMATICAL SANS-SERIF DIGIT TWO
	case '\U0001D7E5':
		return "3", true // MATHEMATICAL SANS-SERIF DIGIT THREE
	case '\U0001D7E6':
		return "4", true // MATHEMATICAL SANS-SERIF DIGIT FOUR
	case '\U0001D7E7':
		return "5", true // MATHEMATICAL SANS-SERIF DIGIT FIVE
	case '\U0001D7E8':
		return "6", true // MATHEMATICAL SANS-SERIF DIGIT SIX
	case '\U0001D7E9':
		return "7", true // MATHEMATICAL SANS-SERIF DIGIT SEVEN
	case '\U0001D7EA':
		return "8", true // MATHEMATICAL SANS-SERIF DIGIT EIGHT
	case '\U0001D7EB':
		return "9", true // MATHEMATICAL SANS-SERIF DIGIT NINE
	case '\U0001D7EC':
		return "0", true // MATHEMATICAL SANS-SERIF BOLD DIGIT ZERO
	case '\U0001D7ED':
		return "1", true // MATHEMATICAL SANS-SERIF BOLD DIGIT ONE
	case '\U0001D7EE':
		return "2", true // MATHEMATICAL SANS-SERIF BOLD DIGIT TWO
	case '\U0001D7EF':
		return "3", true // MATHEMATICAL SANS-SERIF BOLD DIGIT THREE
	case '\U0001D7F0':
		return "4", true // MATHEMATICAL SANS-SERIF BOLD DIGIT FOUR
	case '\U0001D7F1':
		return "5", true // MATHEMATICAL SANS-SERIF BOLD DIGIT FIVE
	case '\U0001D7F2':
		return "6", true // MATHEMATICAL SANS-SERIF BOLD DIGIT SIX
	case '\U0001D7F3':
		return "7", true // MATHEMATICAL SANS-SERIF BOLD DIGIT SEVEN
	case '\U0001D7F4':
		return "8", true // MATHEMATICAL SANS-SERIF BOLD DIGIT EIGHT
	case '\U0001D7F5':
		return "9", true // MATHEMATICAL SANS-SERIF BOLD DIGIT NINE
	case '\U0001D7F6':
		return "0", true // MATHEMATICAL MONOSPACE DIGIT ZERO
	case '\U0001D7F7':
		return "1", true // MATHEMATICAL MONOSPACE DIGIT ONE
	case '\U0001D7F8':
		return "2", true // MATHEMATICAL MONOSPACE DIGIT TWO
	case '\U0001D7F9':
		return "3", true // MATHEMATICAL MONOSPACE DIGIT THREE
	case '\U0001D7FA':
		return "4", true // MATHEMATICAL MONOSPACE DIGIT FOUR
	case '\U0001D7FB':
		return "5", true // MATHEMATICAL MONOSPACE DIGIT FIVE
	case '\U0001D7FC':
		return "6", true // MATHEMATICAL MONOSPACE DIGIT SIX
	case '\U0001D7FD':
		return "7", true // MATHEMATICAL MONOSPACE DIGIT SEVEN
	case '\U0001D7FE':
		return "8", true // MATHEMATICAL MONOSPACE DIGIT EIGHT
	case '\U0001D7FF':
		return "9", true // MATHEMATICAL MONOSPACE DIGIT NINE
	case '\U0001F100':
		return "0.", true // DIGIT ZERO FULL STOP
	case '\U0001F101':
		return "0,", true // DIGIT ZERO COMMA
	case '\U0001F102':
		return "1,", true // DIGIT ONE COMMA
	case '\U0001F103':
		return "2,", true // DIGIT TWO COMMA
	case '\U0001F104':
		return "3,", true // DIGIT THREE COMMA
	case '\U0001F105':
		return "4,", true // DIGIT FOUR COMMA
	case '\U0001F106':
		return "5,", true // DIGIT FIVE COMMA
	case '\U0001F107':
		return "6,", true // DIGIT SIX COMMA
	case '\U0001F108':
		return "7,", true // DIGIT SEVEN COMMA
	case '\U0001F109':
		return "8,", true // DIGIT EIGHT COMMA
	case '\U0001F10A':
		return "9,", true // DIGIT NINE COMMA
	case '\U0001F110':
		return "(A)", true // PARENTHESIZED LATIN CAPITAL LETTER A
	case '\U0001F111':
		return "(B)", true // PARENTHESIZED LATIN CAPITAL LETTER B
	case '\U0001F112':
		return "(C)", true // PARENTHESIZED LATIN CAPITAL LETTER C
	case '\U0001F113':
		return "(D)", true // PARENTHESIZED LATIN CAPITAL LETTER D
	case '\U0001F114':
		return "(E)", true // PARENTHESIZED LATIN CAPITAL LETTER E
	case '\U0001F115':
		return "(F)", true // PARENTHESIZED LATIN CAPITAL LETTER F
	case '\U0001F116':
		return "(G)", true // PARENTHESIZED LATIN CAPITAL LETTER G
	case '\U0001F117':
		return "(H)", true // PARENTHESIZED LATIN CAPITAL LETTER H
	case '\U0001F118':
		return "(I)", true // PARENTHESIZED LATIN CAPITAL LETTER I
	case '\U0001F119':
		return "(J)", true // PARENTHESIZED LATIN CAPITAL LETTER J
	case '\U0001F11A':
		return "(K)", true // PARENTHESIZED LATIN CAPITAL LETTER K
	case '\U0001F11B':
		return "(L)", true // PARENTHESIZED LATIN CAPITAL LETTER L
	case '\U0001F11C':
		return "(M)", true // PARENTHESIZED LATIN CAPITAL LETTER M
	case '\U0001F11D':
		return "(N)", true // PARENTHESIZED LATIN CAPITAL LETTER N
	case '\U0001F11E':
		return "(O)", true // PARENTHESIZED LATIN CAPITAL LETTER O
	case '\U0001F11F':
		return "(P)", true // PARENTHESIZED LATIN CAPITAL LETTER P
	case '\U0001F120':
		return "(Q)", true // PARENTHESIZED LATIN CAPITAL LETTER Q
	case '\U0001F121':
		return "(R)", true // PARENTHESIZED LATIN CAPITAL LETTER R
	case '\U0001F122':
		return "(S)", true // PARENTHESIZED LATIN CAPITAL LETTER S
	case '\U0001F123':
		return "(T)", true // PARENTHESIZED LATIN CAPITAL LETTER T
	case '\U0001F124':
		return "(U)", true // PARENTHESIZED LATIN CAPITAL LETTER U
	case '\U0001F125':
		return "(V)", true // PARENTHESIZED LATIN CAPITAL LETTER V
	case '\U0001F126':
		return "(W)", true // PARENTHESIZED LATIN CAPITAL LETTER W
	case '\U0001F127':
		return "(X)", true // PARENTHESIZED LATIN CAPITAL LETTER X
	case '\U0001F128':
		return "(Y)", true // PARENTHESIZED LATIN CAPITAL LETTER Y
	case '\U0001F129':
		return "(Z)", true // PARENTHESIZED LATIN CAPITAL LETTER Z
	case '\U0001F12B':
		return "C", true // CIRCLED ITALIC LATIN CAPITAL LETTER C
	case '\U0001F12C':
		return "R", true // CIRCLED ITALIC LATIN CAPITAL LETTER R
	case '\U0001F12D':
		return "CD", true // CIRCLED CD
	case '\U0001F12E':
		return "WZ", true // CIRCLED WZ
	case '\U0001F130':
		return "A", true // SQUARED LATIN CAPITAL LETTER A
	case '\U0001F131':
		return "B", true // SQUARED LATIN CAPITAL LETTER B
	case '\U0001F132':
		return "C", true // SQUARED LATIN CAPITAL LETTER C
	case '\U0001F133':
		return "D", true // SQUARED LATIN CAPITAL LETTER D
	case '\U0001F134':
		return "E", true // SQUARED LATIN CAPITAL LETTER E
	case '\U0001F135':
		return "F", true // SQUARED LATIN CAPITAL LETTER F
	case '\U0001F136':
		return "G", true // SQUARED LATIN CAPITAL LETTER G
	case '\U0001F137':
		return "H", true // SQUARED LATIN CAPITAL LETTER H
	case '\U0001F138':
		return "I", true // SQUARED LATIN CAPITAL LETTER I
	case '\U0001F139':
		return "J", true // SQUARED LATIN CAPITAL LETTER J
	case '\U0001F13A':
		return "K", true // SQUARED LATIN CAPITAL LETTER K
	case '\U0001F13B':
		return "L", true // SQUARED LATIN CAPITAL LETTER L
	case '\U0001F13C':
		return "M", true // SQUARED LATIN CAPITAL LETTER M
	case '\U0001F13D':
		return "N", true // SQUARED LATIN CAPITAL LETTER N
	case '\U0001F13E':
		return "O", true // SQUARED LATIN CAPITAL LETTER O
	case '\U0001F13F':
		return "P", true // SQUARED LATIN CAPITAL LETTER P
	case '\U0001F140':
		return "Q", true // SQUARED LATIN CAPITAL LETTER Q
	case '\U0001F141':
		return "R", true // SQUARED LATIN CAPITAL LETTER R
	case '\U0001F142':
		return "S", true // SQUARED LATIN CAPITAL LETTER S
	case '\U0001F143':
		return "T", true // SQUARED LATIN CAPITAL LETTER T
	case '\U0001F144':
		return "U", true // SQUARED LATIN CAPITAL LETTER U
	case '\U0001F145':
		return "V", true // SQUARED LATIN CAPITAL LETTER V
	case '\U0001F146':
		return "W", true // SQUARED LATIN CAPITAL LETTER W
	case '\U0001F147':
		return "X", true // SQUARED LATIN CAPITAL LETTER X
	case '\U0001F148':
		return "Y", true // SQUARED LATIN CAPITAL LETTER Y
	case '\U0001F149':
		return "Z", true // SQUARED LATIN CAPITAL LETTER Z
	case '\U0001F14A':
		return "HV", true // SQUARED HV
	case '\U0001F14B':
		return "MV", true // SQUARED MV
	case '\U0001F14C':
		return "SD", true // SQUARED SD
	case '\U0001F14D':
		return "SS", true // SQUARED SS
	case '\U0001F14E':
		return "PPV", true // SQUARED PPV
	case '\U0001F14F':
		return "WC", true // SQUARED WC
	case '\U0001F16A':
		return "MC", true // RAISED MC SIGN
	case '\U0001F16B':
		return "MD", true // RAISED MD SIGN
	case '\U0001F16C':
		return "MR", true // RAISED MR SIGN
	case '\U0001F190':
		return "DJ", true // SQUARE DJ
	case '\U0001FBF0':
		return "0", true // SEGMENTED DIGIT ZERO
	case '\U0001FBF1':
		return "1", true // SEGMENTED DIGIT ONE
	case '\U0001FBF2':
		return "2", true // SEGMENTED DIGIT TWO
	case '\U0001FBF3':
		return "3", true // SEGMENTED DIGIT THREE
	case '\U0001FBF4':
		return "4", true // SEGMENTED DIGIT FOUR
	case '\U0001FBF5':
		return "5", true // SEGMENTED DIGIT FIVE
	case '\U0001FBF6':
		return "6", true // SEGMENTED DIGIT SIX
	case '\U0001FBF7':
		return "7", true // SEGMENTED DIGIT SEVEN
	case '\U0001FBF8':
		return "8", true // SEGMENTED DIGIT EIGHT
	case '\U0001FBF9':
		return "9", true // SEGMENTED DIGIT NINE
	default:
		return "", false
	}
}

package localization

// Language map for Accept-Language header
var LangMap = map[string]string{
	// Core global languages
	"en": "en-US,en;q=0.9",
	"fr": "fr-FR,fr;q=0.9,en;q=0.8",
	"de": "de-DE,de;q=0.9,en;q=0.8",
	"es": "es-ES,es;q=0.9,en;q=0.8",
	"it": "it-IT,it;q=0.9,en;q=0.8",
	"pt": "pt-BR,pt;q=0.9,en;q=0.8",
	"ru": "ru-RU,ru;q=0.9,en;q=0.8",
	"zh": "zh-CN,zh;q=0.9,en;q=0.8",     // Simplified Chinese
	"zh-TW": "zh-TW,zh;q=0.9,en;q=0.8", // Traditional Chinese
	"ja": "ja-JP,ja;q=0.9,en;q=0.8",
	"ko": "ko-KR,ko;q=0.9,en;q=0.8",
	"ar": "ar-SA,ar;q=0.9,en;q=0.8",
	"hi": "hi-IN,hi;q=0.9,en;q=0.8",

	// African languages
	"sw": "sw-KE,sw;q=0.9,en;q=0.8",   // Swahili
	"am": "am-ET,am;q=0.9,en;q=0.8",   // Amharic
	"yo": "yo-NG,yo;q=0.9,en;q=0.8",   // Yoruba
	"ha": "ha-NE,ha;q=0.9,en;q=0.8",   // Hausa
	"zu": "zu-ZA,zu;q=0.9,en;q=0.8",   // Zulu
	"xh": "xh-ZA,xh;q=0.9,en;q=0.8",   // Xhosa
	"st": "st-LS,st;q=0.9,en;q=0.8",   // Sesotho
	"tn": "tn-BW,tn;q=0.9,en;q=0.8",   // Setswana
	"lg": "lg-UG,lg;q=0.9,en;q=0.8",   // Luganda
	"rw": "rw-RW,rw;q=0.9,en;q=0.8",   // Kinyarwanda
	"so": "so-SO,so;q=0.9,en;q=0.8",   // Somali
	"ts": "ts-ZA,ts;q=0.9,en;q=0.8",   // Tsonga
	"ve": "ve-ZA,ve;q=0.9,en;q=0.8",   // Venda
	"sn": "sn-ZW,sn;q=0.9,en;q=0.8",   // Shona

	// Middle Eastern languages
	"fa": "fa-IR,fa;q=0.9,en;q=0.8",   // Persian (Farsi)
	"he": "he-IL,he;q=0.9,en;q=0.8",   // Hebrew
	"tr": "tr-TR,tr;q=0.9,en;q=0.8",   // Turkish
	"ku": "ku-IQ,ku;q=0.9,en;q=0.8",   // Kurdish

	// South & Southeast Asia
	"bn": "bn-BD,bn;q=0.9,en;q=0.8",   // Bengali
	"ta": "ta-IN,ta;q=0.9,en;q=0.8",   // Tamil
	"te": "te-IN,te;q=0.9,en;q=0.8",   // Telugu
	"ml": "ml-IN,ml;q=0.9,en;q=0.8",   // Malayalam
	"kn": "kn-IN,kn;q=0.9,en;q=0.8",   // Kannada
	"mr": "mr-IN,mr;q=0.9,en;q=0.8",   // Marathi
	"pa": "pa-IN,pa;q=0.9,en;q=0.8",   // Punjabi
	"gu": "gu-IN,gu;q=0.9,en;q=0.8",   // Gujarati
	"si": "si-LK,si;q=0.9,en;q=0.8",   // Sinhala
	"ne": "ne-NP,ne;q=0.9,en;q=0.8",   // Nepali
	"th": "th-TH,th;q=0.9,en;q=0.8",   // Thai
	"km": "km-KH,km;q=0.9,en;q=0.8",   // Khmer
	"my": "my-MM,my;q=0.9,en;q=0.8",   // Burmese
	"lo": "lo-LA,lo;q=0.9,en;q=0.8",   // Lao
	"vi": "vi-VN,vi;q=0.9,en;q=0.8",   // Vietnamese
	"id": "id-ID,id;q=0.9,en;q=0.8",   // Indonesian
	"ms": "ms-MY,ms;q=0.9,en;q=0.8",   // Malay
	"fil": "fil-PH,fil;q=0.9,en;q=0.8",// Filipino

	// European languages
	"pl": "pl-PL,pl;q=0.9,en;q=0.8",   // Polish
	"nl": "nl-NL,nl;q=0.9,en;q=0.8",   // Dutch
	"sv": "sv-SE,sv;q=0.9,en;q=0.8",   // Swedish
	"no": "no-NO,no;q=0.9,en;q=0.8",   // Norwegian
	"da": "da-DK,da;q=0.9,en;q=0.8",   // Danish
	"fi": "fi-FI,fi;q=0.9,en;q=0.8",   // Finnish
	"cs": "cs-CZ,cs;q=0.9,en;q=0.8",   // Czech
	"sk": "sk-SK,sk;q=0.9,en;q=0.8",   // Slovak
	"hu": "hu-HU,hu;q=0.9,en;q=0.8",   // Hungarian
	"el": "el-GR,el;q=0.9,en;q=0.8",   // Greek
	"uk": "uk-UA,uk;q=0.9,en;q=0.8",   // Ukrainian
	"ro": "ro-RO,ro;q=0.9,en;q=0.8",   // Romanian
	"bg": "bg-BG,bg;q=0.9,en;q=0.8",   // Bulgarian
	"sr": "sr-RS,sr;q=0.9,en;q=0.8",   // Serbian
	"hr": "hr-HR,hr;q=0.9,en;q=0.8",   // Croatian
	"sl": "sl-SI,sl;q=0.9,en;q=0.8",   // Slovenian
	"lt": "lt-LT,lt;q=0.9,en;q=0.8",   // Lithuanian
	"lv": "lv-LV,lv;q=0.9,en;q=0.8",   // Latvian
	"et": "et-EE,et;q=0.9,en;q=0.8",   // Estonian
	"is": "is-IS,is;q=0.9,en;q=0.8",   // Icelandic
}
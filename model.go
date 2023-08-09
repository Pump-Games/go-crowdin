package crowdin

// AddFileOptions used for AddFile() API call
type AddFileOptions struct {
	// Note: Used only when uploading CSV (or XLS/XLSX) file to define data columns mapping.
	// Acceptable value is the combination of the following constants:
	// "identifier" — Column contains string identifier.
	// "source_phrase" — Column contains only source string (in result string will contain same string).
	// "source_or_translation" — Column contains source string but when exporting same column should contain translation (also when uploading existing translations, the value from this column will be used as a translated string).
	// "translation" — Column contains translated string (when imported file already contains translations).
	// "context" — Column contains some comments on source string. Context information.
	// "max_length" — Column contains max. length of translation for this string.
	// "none" — Do not import column.
	Scheme string

	// Used when uploading CSV (or XLS/XLSX) files via API. Defines whether first line should be imported or it contains columns headers. May not contain value.
	FirstLineContainsHeader bool

	// Files array that should be added to Crowdin project. Array keys should contain file names with path in Crowdin project.
	Files map[string]string

	// Acceptable values are:
	// empty value or "auto" — Try to detect file type by extension or MIME type
	// "gettext" — GNU GetText (*.po, *.pot)
	// "qtts" — Nokia Qt (*.ts)
	// "dklang" — Delphi DKLang (*.dklang)
	// "android" — Android (*.xml)
	// "resx" — .NET (*.resx)
	// "properties" — Java (*.properties)
	// "macosx" — Mac OS X / iOS (*.strings)
	// "blackberry" — BlackBerry (*.rrc)
	// "Symbian" — Symbian (*.lXX)
	// "flex" — Adobe Flex (*.properties)
	// "bada" — Samsung Bada (*.xml)
	// "txt" — Plain Text (*.txt)
	// "srt" — SubRip .srt (*.srt)
	// "sbv" — Youtube .sbv (*.sbv)
	// "xliff" — XLIFF (*.xliff)
	// "html" — HTML (*.html, *.htm, *.xhtml, *.xhtm)
	// "dtd" — Mozilla DTD (*.dtd)
	// "chrome" — Google Chrome Extension (*.json)
	// "yaml" — Ruby On Rails (*.yaml)
	// "csv" — Comma Separated Values (*.csv)
	// "rc" — Windows Resources (*.rc)
	// "wxl" — WiX Installer Resources (*.wxl)
	// "nsh" — NSIS Installer Resources (*.nsh)
	// "joomla" — Joomla localizable resources (*.ini)
	// "ini" — Generic INI (*.ini)
	// "isl" — ISL (*.isl)
	// "resw" — Windows 8 Metro (*.resw)
	// "resjson" — Windows 8 Metro (*.resjson)
	// "docx" — Microsoft Office and OpenOffice.org Documents (*.docx, *.dotx, *.odt, *.ott, *.xslx, *.xltx, *.pptx, *.potx, *.ods, *.ots, *.odg, *.otg, *.odp, *.otp, *.idml)
	// "md" — Markdown (*.md, *.text, *.markdown...)
	// "mediawiki" — MediaWiki (*.wiki, *.wikitext, *.mediawiki)
	// "play" — Play Framework
	// "haml" — Haml (*.haml)
	// "arb" — Application Resource Bundle (*.arb)
	// "vtt" — Video Subtitling and WebVTT (*.vtt)
	Type string
}

// UpdateFileOptions used for UpdateFile() API call
type UpdateFileOptions struct {
	// Note: Used only when uploading CSV (or XLS/XLSX) file to define data columns mapping.
	// Acceptable value is the combination of the following constants:
	// "identifier" — Column contains string identifier.
	// "source_phrase" — Column contains only source string (in result string will contain same string).
	// "source_or_translation" — Column contains source string but when exporting same column should contain translation (also when uploading existing translations, the value from this column will be used as a translated string).
	// "translation" — Column contains translated string (when imported file already contains translations).
	// "context" — Column contains some comments on source string. Context information.
	// "max_length" — Column contains max. length of translation for this string.
	// "none" — Do not import column.
	Scheme string

	// Used when uploading CSV (or XLS/XLSX) files via API. Defines whether first line should be imported or it contains columns headers. May not contain value.
	FirstLineContainsHeader bool

	// Files array that should be added to Crowdin project. Array keys should contain file names with path in Crowdin project.
	Files map[string]string
}

// UploadTranslationsOptions are options for UploadTranslations api call
type UploadTranslationsOptions struct {
	// Target language. With a single call it's possible to upload translations for several files but only into one of the languages.
	Language string

	// Translated files array. Array keys should contain file names in Crowdin.
	Files map[string]string

	// Defines whether to add translation if there is the same translation previously added. Acceptable values are: 0 or 1. Default is 0.
	ImportDuplicates string
}

// ChangeDirectoryOptions are options for ChangeDirectory api call
type ChangeDirectoryOptions struct {
	// Full directory path that should be modified (e.g. /MainPage/AboutUs).
	Name string

	// New directory name.
	NewName string

	// New directory title to be displayed in Crowdin UI.
	Title string
}

// CreateProjectOptions are options for CreateProject api call
type CreateProjectOptions struct {
	// Project name.
	Name string

	// Project identifier. Should be unique among other Crowdin projects.
	Identifier string

	// Source files language. Should be a two letters language code.
	SourceLanguage string

	// An array of language codes project should be translate to.
	Languages []string

	// Project join policy. Acceptable values are: open, private
	JoinPolicy string
}

// EditProjectOptions are options for EditProject api call
type EditProjectOptions struct {
	// Project name.
	Name string

	// An array of language codes project should be translate to.
	Languages []string

	// Project join policy. Acceptable values are: open, private
	JoinPolicy string
}

// ExportFileOptions are options for ExportFile api call
type ExportFileOptions struct {
	// This parameter specifies a path to the file that should be exported from the project.
	CrowdinFile string

	// Crowdin language code.
	Language string

	// Path to the file name that file will be exported to.
	LocalPath string
}

// DownloadOptions are options for DownloadTranslations api call
type DownloadOptions struct {
	// Language code or "all" to download a bundle with translations to all languages.
	Package string

	// Path to the file name that file will be exported to.
	LocalPath string
}

type responseLanguageStatus struct {
	Files []struct {
		ID              string `json:"id"`
		Name            string `json:"name"`
		NodeType        string `json:"node_type"`
		Phrases         string `json:"phrases"`
		Translated      string `json:"translated"`
		Approved        string `json:"approved"`
		Words           string `json:"words"`
		WordsTranslated string `json:"words_translated"`
		WordsApproved   string `json:"words_approved"`
	} `json:"files"`
}

type responseAddFile struct {
	Success bool `json:"success"`
	Stats   struct {
		Files []struct {
			FileID  int    `json:"file_id"`
			Name    string `json:"name"`
			Strings int    `json:"strings"`
			Words   int    `json:"words"`
		} `json:"files"`
	} `json:"stats"`
}

type responseUploadTranslation struct {
	Success bool `json:"success"`
	Stats   struct {
		Files []struct {
			Name   string `json:"name"`
			Status string `json:"status"`
		} `json:"files"`
	} `json:"stats"`
}

type responseManageProject struct {
	Project struct {
		Success    bool   `json:"success"`
		Invitation string `json:"invitation"`
		URL        string `json:"url"`
		Key        string `json:"key"`
	} `json:"project"`
}

type responseDeleteProject struct {
	Project struct {
		Success bool `json:"success"`
	} `json:"project"`
}

type responseExportTranslations struct {
	Success struct {
		Status string `json:"status"`
	} `json:"success"`
}

// TranslationStatus is a response struct
type TranslationStatus struct {
	Name               string `json:"name"`
	Code               string `json:"code"`
	Phrases            string `json:"phrases"`
	Translated         string `json:"translated"`
	Approved           string `json:"approved"`
	Words              string `json:"words"`
	WordsTranslated    string `json:"words_translated"`
	WordsApproved      string `json:"words_approved"`
	TranslatedProgress int    `json:"translated_progress"`
	ApprovedProgress   int    `json:"approved_progress"`
}

// ExportStatus is a response struct
type ExportStatus struct {
	Status          string `json:"status"`
	Progress        int    `json:"progress"`
	LastBuild       string `json:"last_build"`
	Code            string `json:"code"`
	Message         string `json:"message"`
	CurrentFile     string `json:"current_file"`
	CurrentLanguage string `json:"current_language"`
}

// ProjectInfo is a response struct
type ProjectInfo struct {
	Files []struct {
		Name         string `json:"name"`
		NodeType     string `json:"node_type"`
		Created      string `json:"created"`
		LastUpdated  string `json:"last_updated"`
		LastAccessed string `json:"last_accessed"`
		LastRevision string `json:"last_revision"`
	} `json:"files"`
	Language struct {
		Name         string `json:"name"`
		Code         string `json:"code"`
		CanTranslate int    `json:"can_translate"`
		CanApprove   int    `json:"can_approve"`
	}
	Details struct {
		SourceLanguage struct {
			Name string `json:"name"`
			Code string `json:"code"`
		} `json:"source_language"`
		Name                  string `json:"name"`
		Identifier            string `json:"identifier"`
		Created               string `json:"created"`
		Description           string `json:"description"`
		JoinPolicy            string `json:"private"`
		LastBuild             string `json:"last_build"`
		LastActivity          string `json:"last_activity"`
		ParticipantsCount     string `json:"participants_count"`
		TotalStringsCount     string `json:"total_strings_count"`
		TotalWordsCount       string `json:"total_words_count"`
		DuplicateStringsCount int    `json:"duplicate_strings_count"`
		DuplicateWordsCount   int    `json:"duplicate_words_count"`
		InviteURL             struct {
			Translator  string `json:"translator"`
			Proofreader string `json:"proofreader"`
		} `json:"invite_url"`
	} `json:"details"`
}

// AccountDetails is a response struct
type AccountDetails struct {
	Success  bool `json:"success"`
	Projects []struct {
		Role         string `json:"role"`
		Name         string `json:"name"`
		Identifier   string `json:"identifier"`
		Downloadable int    `json:"downloadable"`
		Key          string `json:"key"`
	} `json:"projects"`
}

type responseGeneral struct {
	Success bool `json:"success"`
}

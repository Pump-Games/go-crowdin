package crowdin

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/mreiferson/go-httpclient"
)

var (
	apiBaseURL        = "https://api.crowdin.com/api/project/"
	apiAccountBaseURL = "https://api.crowdin.com/api/account/"
)

// Crowdin API wrapper
type Crowdin struct {
	config struct {
		apiBaseURL        string
		apiAccountBaseURL string
		token             string
		project           string
		client            *http.Client
	}
	debug     bool
	logWriter io.Writer
}

// New - create new instance of Crowdin API.
func New(token, project string) *Crowdin {

	transport := &httpclient.Transport{
		ConnectTimeout:   5 * time.Second,
		ReadWriteTimeout: 40 * time.Second,
	}
	defer transport.Close()

	s := &Crowdin{}
	s.config.apiBaseURL = apiBaseURL
	s.config.apiAccountBaseURL = apiAccountBaseURL
	s.config.token = token
	s.config.project = project
	s.config.client = &http.Client{
		Transport: transport,
	}
	return s
}

// SetProject set project details
func (crowdin *Crowdin) SetProject(token, project string) *Crowdin {
	crowdin.config.token = token
	crowdin.config.project = project
	return crowdin
}

// SetDebug - traces errors if it's set to true.
func (crowdin *Crowdin) SetDebug(debug bool, logWriter io.Writer) {
	crowdin.debug = debug
	crowdin.logWriter = logWriter
}

// SetClient sets a custom http client. Can be useful in App Engine case.
func (crowdin *Crowdin) SetClient(client *http.Client) {
	crowdin.config.client = client
}

// AddFile - Add new file to Crowdin project.
func (crowdin *Crowdin) AddFile(options *AddFileOptions) (*responseAddFile, error) {

	params := make(map[string]string)
	params["json"] = ""

	if options != nil {

		if options.Type != "" {
			params["type"] = options.Type
		}

		if options.Scheme != "" {
			params["scheme"] = options.Scheme
		}

		if options.FirstLineContainsHeader {
			params["first_line_contains_header"] = "true"
		} else {
			params["first_line_contains_header"] = "false"
		}

	}

	files := make(map[string]string)
	if options != nil && options.Files != nil {
		for k, path := range options.Files {
			files[fmt.Sprintf("files[%v]", k)] = path
		}
	}

	response, err := crowdin.post(&postOptions{
		urlStr: fmt.Sprintf(crowdin.config.apiBaseURL+"%v/add-file?key=%v", crowdin.config.project, crowdin.config.token),
		params: params,
		files:  files,
	})

	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	crowdin.log(string(response))

	var responseAPI responseAddFile
	err = json.Unmarshal(response, &responseAPI)
	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	return &responseAPI, nil

}

// UpdateFile - Upload latest version of your localization file to Crowdin
func (crowdin *Crowdin) UpdateFile(options *UpdateFileOptions) (*responseGeneral, error) {

	params := make(map[string]string)
	params["json"] = ""

	if options != nil {

		if options.Scheme != "" {
			params["scheme"] = options.Scheme
		}

		if options.FirstLineContainsHeader {
			params["first_line_contains_header"] = "true"
		} else {
			params["first_line_contains_header"] = "false"
		}

	}

	files := make(map[string]string)
	if options != nil && options.Files != nil {
		for k, path := range options.Files {
			files[fmt.Sprintf("files[%v]", k)] = path
		}
	}

	response, err := crowdin.post(&postOptions{
		urlStr: fmt.Sprintf(crowdin.config.apiBaseURL+"%v/update-file?key=%v", crowdin.config.project, crowdin.config.token),
		params: params,
		files:  files,
	})

	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	crowdin.log(string(response))

	var responseAPI responseGeneral
	err = json.Unmarshal(response, &responseAPI)
	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	return &responseAPI, nil

}

// DeleteFile - Delete file from Crowdin project. All the translations will be lost without ability to restore them
func (crowdin *Crowdin) DeleteFile(fileName string) (*responseGeneral, error) {

	params := make(map[string]string)
	params["json"] = ""
	params["file"] = fileName

	response, err := crowdin.post(&postOptions{
		urlStr: fmt.Sprintf(crowdin.config.apiBaseURL+"%v/delete-file?key=%v", crowdin.config.project, crowdin.config.token),
		params: params,
	})

	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	crowdin.log(string(response))

	var responseAPI responseGeneral
	err = json.Unmarshal(response, &responseAPI)
	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	return &responseAPI, nil

}

// UploadTranslations - Upload latest version of your localization file to Crowdin
func (crowdin *Crowdin) UploadTranslations(options *UploadTranslationsOptions) (*responseUploadTranslation, error) {

	params := make(map[string]string)
	params["json"] = ""

	if options != nil {

		if options.Language != "" {
			params["language"] = options.Language
		}

		params["import_duplicates"] = options.ImportDuplicates

	}

	files := make(map[string]string)
	if options != nil && options.Files != nil {
		for k, path := range options.Files {
			files[fmt.Sprintf("files[%v]", k)] = path
		}
	}

	response, err := crowdin.post(&postOptions{
		urlStr: fmt.Sprintf(crowdin.config.apiBaseURL+"%v/upload-translation?key=%v", crowdin.config.project, crowdin.config.token),
		params: params,
		files:  files,
	})

	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	crowdin.log(string(response))

	var responseAPI responseUploadTranslation
	err = json.Unmarshal(response, &responseAPI)
	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	return &responseAPI, nil

}

// GetTranslationsStatus - Track overall translation and proofreading progresses of each target language
func (crowdin *Crowdin) GetTranslationsStatus() ([]TranslationStatus, error) {

	response, err := crowdin.post(&postOptions{
		urlStr: fmt.Sprintf(crowdin.config.apiBaseURL+"%v/status?key=%v", crowdin.config.project, crowdin.config.token),
		params: map[string]string{
			"json": "",
		},
	})

	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	crowdin.log(string(response))

	var responseAPI []TranslationStatus
	err = json.Unmarshal(response, &responseAPI)
	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	return responseAPI, nil
} // GetTranslationsStatus - Track overall translation and proofreading progresses of each target language

// GetExportStatus - Get the status of translations export
func (crowdin *Crowdin) GetExportStatus() (*ExportStatus, error) {

	response, err := crowdin.post(&postOptions{
		urlStr: fmt.Sprintf(crowdin.config.apiBaseURL+"%v/export-status?key=%v", crowdin.config.project, crowdin.config.token),
		params: map[string]string{
			"json": "",
		},
	})

	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	crowdin.log(string(response))

	var responseAPI ExportStatus
	err = json.Unmarshal(response, &responseAPI)
	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	return &responseAPI, nil
}

// GetLanguageStatus - Get the detailed translation progress for specified language.
// Language codes - https://crowdin.com/page/api/language-codes
func (crowdin *Crowdin) GetLanguageStatus(languageCode string) (*responseLanguageStatus, error) {

	response, err := crowdin.post(&postOptions{
		urlStr: fmt.Sprintf(crowdin.config.apiBaseURL+"%v/language-status?key=%v", crowdin.config.project, crowdin.config.token),
		params: map[string]string{
			"language": languageCode,
			"json":     "",
		},
	})

	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	crowdin.log(string(response))

	var responseAPI responseLanguageStatus
	err = json.Unmarshal(response, &responseAPI)
	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	return &responseAPI, nil
}

// GetProjectDetails - Get Crowdin Project details
func (crowdin *Crowdin) GetProjectDetails() (*ProjectInfo, error) {

	response, err := crowdin.post(&postOptions{
		urlStr: fmt.Sprintf(crowdin.config.apiBaseURL+"%v/info?key=%v", crowdin.config.project, crowdin.config.token),
		params: map[string]string{
			"json": "",
		},
	})

	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	crowdin.log(string(response))

	var responseAPI ProjectInfo
	err = json.Unmarshal(response, &responseAPI)
	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	return &responseAPI, nil
}

// DownloadTranslations - Download ZIP file with translations. You can choose the language of translation you need or download all of them at once.
func (crowdin *Crowdin) DownloadTranslations(options *DownloadOptions) error {

	if options == nil || options.Package == "" {
		return errors.New("Package can't be empty")
	}

	response, err := crowdin.getResponse(&getOptions{
		urlStr: fmt.Sprintf(crowdin.config.apiBaseURL+"%v/download/%v.zip?key=%v", crowdin.config.project, options.Package, crowdin.config.token),
	})

	defer response.Body.Close()

	if err != nil {
		crowdin.log(err)
		return err
	}

	// create the file
	out, err := os.Create(options.LocalPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// writer the body to file
	_, err = io.Copy(out, response.Body)
	if err != nil {
		return err
	}

	return nil
}

// ExportFile - This method exports single translated files from Crowdin. Additionally, it can be applied to export XLIFF files for offline localization.
func (crowdin *Crowdin) ExportFile(options *ExportFileOptions) error {

	params := make(map[string]string)

	if options != nil {

		if options.Language != "" {
			params["language"] = options.Language
		}

		if options.CrowdinFile != "" {
			params["file"] = options.CrowdinFile
		}
	}

	response, err := crowdin.getResponse(&getOptions{
		urlStr: fmt.Sprintf(crowdin.config.apiBaseURL+"%v/export-file?key=%v", crowdin.config.project, crowdin.config.token),
		params: params,
	})

	defer response.Body.Close()

	if err != nil {
		crowdin.log(err)
		return err
	}

	// create the file
	out, err := os.Create(options.LocalPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// writer the body to file
	_, err = io.Copy(out, response.Body)
	if err != nil {
		return err
	}

	return nil
}

// ExportTranslations - Build ZIP archive with the latest translations. Please note that this method can be invoked only once per 30 minutes (there is no such restriction for organization plans). Also API call will be ignored if there were no changes in the project since previous export. You can see whether ZIP archive with latest translations was actually build by status attribute ("built" or "skipped") returned in response.
func (crowdin *Crowdin) ExportTranslations() (*responseExportTranslations, error) {

	response, err := crowdin.get(&getOptions{
		urlStr: fmt.Sprintf(crowdin.config.apiBaseURL+"%v/export?key=%v", crowdin.config.project, crowdin.config.token),
		params: map[string]string{
			"json": "",
		},
	})

	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	crowdin.log(string(response))

	var responseAPI responseExportTranslations
	err = json.Unmarshal(response, &responseAPI)
	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	return &responseAPI, nil
}

// GetAccountProjects - Get Crowdin Project details.
func (crowdin *Crowdin) GetAccountProjects(accountKey, loginUsername string) (*AccountDetails, error) {

	response, err := crowdin.post(&postOptions{
		urlStr: fmt.Sprintf(crowdin.config.apiAccountBaseURL+"get-projects?account-key=%v", accountKey),
		params: map[string]string{
			"login": loginUsername,
			"json":  "",
		},
	})

	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	crowdin.log(string(response))

	var responseAPI AccountDetails
	err = json.Unmarshal(response, &responseAPI)
	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	return &responseAPI, nil
}

// CreateProject - Create Crowdin project.
func (crowdin *Crowdin) CreateProject(accountKey, loginUsername string, options *CreateProjectOptions) (*responseManageProject, error) {

	params := make(map[string]string)
	params["json"] = ""
	params["login"] = loginUsername

	paramsArray := make(map[string][]string)

	if options != nil {

		if options.Name != "" {
			params["name"] = options.Name
		}

		if options.Identifier != "" {
			params["identifier"] = options.Identifier
		}

		if options.SourceLanguage != "" {
			params["source_language"] = options.SourceLanguage
		}

		if options.JoinPolicy != "" {
			params["join_policy"] = options.JoinPolicy
		}

		if options.Languages != nil {
			paramsArray["languages[]"] = options.Languages
		}
	}

	response, err := crowdin.post(&postOptions{
		urlStr:      fmt.Sprintf(crowdin.config.apiAccountBaseURL+"create-project?account-key=%v", accountKey),
		params:      params,
		paramsArray: paramsArray,
	})

	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	crowdin.log(string(response))

	var responseAPI responseManageProject
	err = json.Unmarshal(response, &responseAPI)
	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	return &responseAPI, nil
}

// EditProject - Edit Crowdin project.
func (crowdin *Crowdin) EditProject(options *EditProjectOptions) (*responseManageProject, error) {

	params := make(map[string]string)
	params["json"] = ""

	paramsArray := make(map[string][]string)

	if options != nil {

		if options.Name != "" {
			params["name"] = options.Name
		}

		if options.JoinPolicy != "" {
			params["join_policy"] = options.JoinPolicy
		}

		if options.Languages != nil {
			paramsArray["languages[]"] = options.Languages
		}
	}

	response, err := crowdin.post(&postOptions{
		urlStr:      fmt.Sprintf(crowdin.config.apiBaseURL+"%v/edit-project?key=%v", crowdin.config.project, crowdin.config.token),
		params:      params,
		paramsArray: paramsArray,
	})

	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	crowdin.log(string(response))

	var responseAPI responseManageProject
	err = json.Unmarshal(response, &responseAPI)
	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	return &responseAPI, nil
}

// DeleteProject - Delete Crowdin project with all translations.
func (crowdin *Crowdin) DeleteProject() (*responseDeleteProject, error) {

	params := make(map[string]string)
	params["json"] = ""

	response, err := crowdin.post(&postOptions{
		urlStr: fmt.Sprintf(crowdin.config.apiBaseURL+"%v/delete-project?key=%v", crowdin.config.project, crowdin.config.token),
		params: params,
	})

	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	crowdin.log(string(response))

	var responseAPI responseDeleteProject
	err = json.Unmarshal(response, &responseAPI)
	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	return &responseAPI, nil
}

// AddDirectory - Add directory to Crowdin project.
// name - Directory name (with path if nested directory should be created).
func (crowdin *Crowdin) AddDirectory(directoryName string) (*responseGeneral, error) {

	response, err := crowdin.post(&postOptions{
		urlStr: fmt.Sprintf(crowdin.config.apiBaseURL+"%v/add-directory?key=%v", crowdin.config.project, crowdin.config.token),
		params: map[string]string{
			"name": directoryName,
			"json": "",
		},
	})

	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	crowdin.log(string(response))

	var responseAPI responseGeneral
	err = json.Unmarshal(response, &responseAPI)
	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	return &responseAPI, nil
}

// ChangeDirectory - Rename directory or modify its attributes. When renaming directory the path can not be changed (it means new_name parameter can not contain path, name only).
func (crowdin *Crowdin) ChangeDirectory(options *ChangeDirectoryOptions) (*responseGeneral, error) {

	params := make(map[string]string)
	params["json"] = ""

	if options != nil {

		if options.Name != "" {
			params["name"] = options.Name
		}

		if options.NewName != "" {
			params["new_name"] = options.NewName
		}

		if options.Title != "" {
			params["title"] = options.Title
		}
	}

	response, err := crowdin.post(&postOptions{
		urlStr: fmt.Sprintf(crowdin.config.apiBaseURL+"%v/change-directory?key=%v", crowdin.config.project, crowdin.config.token),
		params: params,
	})

	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	crowdin.log(string(response))

	var responseAPI responseGeneral
	err = json.Unmarshal(response, &responseAPI)
	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	return &responseAPI, nil
}

// DeleteDirectory - Delete Crowdin project directory. All nested files and directories will be deleted too.
// name - Directory name (with path if nested directory should be created).
func (crowdin *Crowdin) DeleteDirectory(directoryName string) (*responseGeneral, error) {

	response, err := crowdin.post(&postOptions{
		urlStr: fmt.Sprintf(crowdin.config.apiBaseURL+"%v/delete-directory?key=%v", crowdin.config.project, crowdin.config.token),
		params: map[string]string{
			"name": directoryName,
			"json": "",
		},
	})

	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	crowdin.log(string(response))

	var responseAPI responseGeneral
	err = json.Unmarshal(response, &responseAPI)
	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	return &responseAPI, nil
}

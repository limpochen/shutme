package manif

import "encoding/xml"

/*
type VersionInfo struct {
	FixedFileInfo  FixedFileInfo  `json:"FixedFileInfo"`
	StringFileInfo StringFileInfo `json:"StringFileInfo"`
	VarFileInfo    VarFileInfo    `json:"VarFileInfo"`
	IconPath       string         `json:"IconPath"`
	ManifestPath   string         `json:"ManifestPath"`
}
type FileVersion struct {
	Major int `json:"Major"`
	Minor int `json:"Minor"`
	Patch int `json:"Patch"`
	Build int `json:"Build"`
}
type ProductVersion struct {
	Major int `json:"Major"`
	Minor int `json:"Minor"`
	Patch int `json:"Patch"`
	Build int `json:"Build"`
}
type FixedFileInfo struct {
	FileVersion    FileVersion    `json:"FileVersion"`
	ProductVersion ProductVersion `json:"ProductVersion"`
	FileFlagsMask  string         `json:"FileFlagsMask"`
	FileFlags      string         `json:"FileFlags "`
	FileOS         string         `json:"FileOS"`
	FileType       string         `json:"FileType"`
	FileSubType    string         `json:"FileSubType"`
}
type StringFileInfo struct {
	Comments         string `json:"Comments"`
	CompanyName      string `json:"CompanyName"`
	FileDescription  string `json:"FileDescription"`
	FileVersion      string `json:"FileVersion"`
	InternalName     string `json:"InternalName"`
	LegalCopyright   string `json:"LegalCopyright"`
	LegalTrademarks  string `json:"LegalTrademarks"`
	OriginalFilename string `json:"OriginalFilename"`
	PrivateBuild     string `json:"PrivateBuild"`
	ProductName      string `json:"ProductName"`
	ProductVersion   string `json:"ProductVersion"`
	SpecialBuild     string `json:"SpecialBuild"`
}
type Translation struct {
	LangID    string `json:"LangID"`
	CharsetID string `json:"CharsetID"`
}
type VarFileInfo struct {
	Translation Translation `json:"Translation"`
}
*/
type Assembly struct {
	XMLName          xml.Name `xml:"assembly"`
	Text             string   `xml:",chardata"`
	Xmlns            string   `xml:"xmlns,attr"`
	ManifestVersion  string   `xml:"manifestVersion,attr"`
	AssemblyIdentity struct {
		Text                  string `xml:",chardata"`
		Type                  string `xml:"type,attr"`
		Name                  string `xml:"name,attr"`
		Version               string `xml:"version,attr"`
		ProcessorArchitecture string `xml:"processorArchitecture,attr"`
	} `xml:"assemblyIdentity"`
	TrustInfo struct {
		Text     string `xml:",chardata"`
		Xmlns    string `xml:"xmlns,attr"`
		Security struct {
			Text                string `xml:",chardata"`
			RequestedPrivileges struct {
				Text                    string `xml:",chardata"`
				RequestedExecutionLevel struct {
					Text     string `xml:",chardata"`
					Level    string `xml:"level,attr"`
					UiAccess string `xml:"uiAccess,attr"`
				} `xml:"requestedExecutionLevel"`
			} `xml:"requestedPrivileges"`
		} `xml:"security"`
	} `xml:"trustInfo"`
}

type MyStructName struct {
	FixedFileInfo struct {
		FileFlags     string `json:"FileFlags "`
		FileFlagsMask string `json:"FileFlagsMask"`
		FileOS        string `json:"FileOS"`
		FileSubType   string `json:"FileSubType"`
		FileType      string `json:"FileType"`
		FileVersion   struct {
			Build int64 `json:"Build"`
			Major int64 `json:"Major"`
			Minor int64 `json:"Minor"`
			Patch int64 `json:"Patch"`
		} `json:"FileVersion"`
		ProductVersion struct {
			Build int64 `json:"Build"`
			Major int64 `json:"Major"`
			Minor int64 `json:"Minor"`
			Patch int64 `json:"Patch"`
		} `json:"ProductVersion"`
	} `json:"FixedFileInfo"`
	IconPath       string `json:"IconPath"`
	ManifestPath   string `json:"ManifestPath"`
	StringFileInfo struct {
		Comments         string `json:"Comments"`
		CompanyName      string `json:"CompanyName"`
		FileDescription  string `json:"FileDescription"`
		FileVersion      string `json:"FileVersion"`
		InternalName     string `json:"InternalName"`
		LegalCopyright   string `json:"LegalCopyright"`
		LegalTrademarks  string `json:"LegalTrademarks"`
		OriginalFilename string `json:"OriginalFilename"`
		PrivateBuild     string `json:"PrivateBuild"`
		ProductName      string `json:"ProductName"`
		ProductVersion   string `json:"ProductVersion"`
		SpecialBuild     string `json:"SpecialBuild"`
	} `json:"StringFileInfo"`
	VarFileInfo struct {
		Translation struct {
			CharsetID string `json:"CharsetID"`
			LangID    string `json:"LangID"`
		} `json:"Translation"`
	} `json:"VarFileInfo"`
}

var assembly Assembly
var fileinfo MyStructName

func main() {

}

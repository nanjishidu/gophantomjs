package gophantomjs

import (
	"bytes"
	"fmt"
	. "github.com/gophper/gomini"
	"net/url"
	"os"
	"sync"
	"text/template"
	"time"
)

var (
	settingMutex                 sync.Mutex
	tempDir                      = "gophantomjs/temp/"
	defaultPhantomJsPageSettings = &PhantomJsPageSettings{
		JavascriptEnabled:  false,
		LoadImages:         true,
		UserAgent:          "GoPhantomJsServer",
		UserName:           "",
		Password:           "",
		XSSAuditingEnabled: false,
		ResourceTimeout:    3000,
		CustomHeaders:      map[string]string{},
	}
)

func init() {
	if IsExist(tempDir) {
		RemoveAll(tempDir)
	}
	Mkdir(tempDir)
}

type PhantomJsRequest struct {
	binpath  string //phatomjs 二进制文件地址
	jspath   string //js 文件地址
	url      string
	method   string
	params   map[string][]string
	data     map[string]interface{}
	cookies  string
	settings *PhantomJsPageSettings
}

type PhantomJsPageCookies struct {
	Domain   string `json:"domain"`
	Expires  string `json:"expires"`
	Expiry   int64  `json:"expiry"`
	HttpOnly bool   `json:"httponly"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Secure   bool   `json:"secure"`
	Value    string `json:"value"`
}
type PhantomJsPageSettings struct {
	JavascriptEnabled  bool
	LoadImages         bool
	UserAgent          string
	UserName           string
	Password           string
	XSSAuditingEnabled bool
	ResourceTimeout    int
	CustomHeaders      map[string]string
}

func SetDefaultPageSettings(settings *PhantomJsPageSettings) {
	settingMutex.Lock()
	defer settingMutex.Unlock()
	defaultPhantomJsPageSettings = settings
}
func NewPhantomJsRequest(url, method string) *PhantomJsRequest {
	return &PhantomJsRequest{
		binpath:  "/usr/local/bin/phantomjs",
		url:      url,
		method:   method,
		params:   map[string][]string{},
		settings: defaultPhantomJsPageSettings,
	}

}
func Head(url string) *PhantomJsRequest {
	return NewPhantomJsRequest(url, "HEAD")
}
func Get(url string) *PhantomJsRequest {
	return NewPhantomJsRequest(url, "GET")

}
func Post(url string) *PhantomJsRequest {
	return NewPhantomJsRequest(url, "POST")
}

//设置请求参数
func (b *PhantomJsRequest) Param(key, value string) *PhantomJsRequest {
	if param, ok := b.params[key]; ok {
		b.params[key] = append(param, value)
	} else {
		b.params[key] = []string{value}
	}
	return b
}

//设置请求header
func (b *PhantomJsRequest) Header(key, value string) *PhantomJsRequest {
	b.settings.CustomHeaders[key] = value
	return b
}
func (b *PhantomJsRequest) Setting(setting *PhantomJsPageSettings) *PhantomJsRequest {
	b.settings = setting
	return b
}
func (b *PhantomJsRequest) SetUserAgent(useragent string) *PhantomJsRequest {
	b.settings.UserAgent = useragent
	return b
}
func (b *PhantomJsRequest) SetBasicAuth(username, password string) *PhantomJsRequest {
	b.settings.UserName = username
	b.settings.Password = password
	return b
}

func (b *PhantomJsRequest) SetJavascriptEnabled(isJavascriptEnabled bool) *PhantomJsRequest {
	b.settings.JavascriptEnabled = isJavascriptEnabled
	return b
}
func (b *PhantomJsRequest) SetLoadImages(isLoadImages bool) *PhantomJsRequest {
	b.settings.LoadImages = isLoadImages
	return b
}
func (b *PhantomJsRequest) SetXSSAuditingEnabled(isXSSAuditingEnabled bool) *PhantomJsRequest {
	b.settings.XSSAuditingEnabled = isXSSAuditingEnabled
	return b
}
func (b *PhantomJsRequest) SetResourceTimeout(resourceTimeout int) *PhantomJsRequest {
	b.settings.ResourceTimeout = resourceTimeout
	return b
}
func (b *PhantomJsRequest) SetCookies(cookies string) *PhantomJsRequest {
	b.cookies = cookies
	return b
}
func (b *PhantomJsRequest) SetBinPath(binpath string) *PhantomJsRequest {
	b.binpath = binpath
	return b
}
func (b *PhantomJsRequest) SetJsPath(jspath string) *PhantomJsRequest {
	b.jspath = jspath
	return b
}
func (b *PhantomJsRequest) GetParamBody() string {
	var paramBody = ""
	if len(b.params) > 0 {
		var buf bytes.Buffer
		for k, v := range b.params {
			for _, vv := range v {
				buf.WriteString(url.QueryEscape(k))
				buf.WriteByte('=')
				buf.WriteString(url.QueryEscape(vv))
				buf.WriteByte('&')
			}
		}
		paramBody = buf.String()
		paramBody = paramBody[0 : len(paramBody)-1]
	}
	return paramBody
}

func (b *PhantomJsRequest) CreateJs(p string) {
	var data = make(map[string]interface{})
	data["url"] = b.url
	data["method"] = b.method
	data["paramBody"] = b.GetParamBody()
	data["setting"] = b.settings
	if b.cookies != "" {
		data["cookies"] = b.cookies
	}
	b.jspath = tempDir + Md5(GetInt64Str(time.Now().Unix())) + ".js"
	f, err := os.Create(b.jspath)
	if err != nil {
		fmt.Println("create file error", err)
		return
	}
	defer f.Close()
	t, _ := template.New("pageContent").Parse(p)
	t.Execute(f, data)
}
func (b *PhantomJsRequest) PageContent() (string, error) {
	b.CreateJs(commonTemplate + pageContent)
	return Exec(b.binpath, b.jspath)
}
func (b *PhantomJsRequest) GetCookies() (string, error) {
	b.CreateJs(commonTemplate + getCookies)
	return Exec(b.binpath, b.jspath)
}

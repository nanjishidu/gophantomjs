// template.go
package gophantomjs

const commonTemplate = `
var page = require('webpage').create();
{{$setting := .setting}}
{{if $setting}}
page.settings.JavascriptEnabled = {{$setting.JavascriptEnabled}};
page.settings.LoadImages = {{$setting.LoadImages}};
page.settings.UserAgent = "{{$setting.UserAgent}}";
{{if $setting.UserName}}
page.settings.UserName = {{$setting.UserName}};
page.settings.Password = {{$setting.Password}};
{{end}}
page.settings.XSSAuditingEnabled = {{$setting.XSSAuditingEnabled}};
page.settings.ResourceTimeout = {{.setting.ResourceTimeout}};
{{end}}
{{$cookies := .cookies}}
{{if $cookies}}
page.addCookie({{$cookies}})
{{end}}
`
const pageContent = `
page.open("{{.url}}","{{.method}}","{{.paramBody}}",function(status) {
    console.log(page.content);
    phantom.exit();
});
`
const getCookies = `
page.open("{{.url}}","{{.method}}","{{.paramBody}}",function(status) {
    // console.log(page.cookies);
 var cookies = page.cookies;
  
  console.log('Listing cookies:');
  for(var i in cookies) {
    console.log(cookies[i].name + '=' + cookies[i].value);
  }
    phantom.exit();
});
`

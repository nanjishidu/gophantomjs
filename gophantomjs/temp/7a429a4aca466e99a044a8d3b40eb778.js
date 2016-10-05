
var page = require('webpage').create();


page.settings.JavascriptEnabled = false;
page.settings.LoadImages = false;
page.settings.UserAgent = "PhantomJsServer";

page.settings.XSSAuditingEnabled = false;
page.settings.ResourceTimeout = 3000;




page.open("http://www.oschina.net","GET","",function(status) {
    console.log(page.content);
    phantom.exit();
});

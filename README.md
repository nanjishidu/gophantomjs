# gophantomjs

PhantomJS可以把它看作一个“虚拟浏览器”，除了不能浏览，其他与正常浏览器一样。
它的内核是WebKit引擎，不提供图形界面，只能在命令行下使用。

参数
URL    //请求URL
Method //请求Method GET POST
Params //请求参数
UserAgent 
Cookies
LibraryPath

phantom

	injectJs
	libraryPath
	cookies

web page

	cookies
	settings 
		javascriptEnabled
		loadImages
		userAgent
		userName
		password
		XSSAuditingEnabled
		resourceTimeout
		customHeaders
	libraryPath
	navigationLocked
	zoomFactor
	clipRect
	
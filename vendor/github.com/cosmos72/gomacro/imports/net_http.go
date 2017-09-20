// this file was generated by gomacro command: import _b "net/http"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"bufio"
	"net"
	"net/http"
	"net/url"
	"os"
)

// reflection: allow interpreted code to import "net/http"
func init() {
	Packages["net/http"] = Package{
	Binds: map[string]Value{
		"CanonicalHeaderKey":	ValueOf(http.CanonicalHeaderKey),
		"DefaultClient":	ValueOf(&http.DefaultClient).Elem(),
		"DefaultMaxHeaderBytes":	ValueOf(http.DefaultMaxHeaderBytes),
		"DefaultMaxIdleConnsPerHost":	ValueOf(http.DefaultMaxIdleConnsPerHost),
		"DefaultServeMux":	ValueOf(&http.DefaultServeMux).Elem(),
		"DefaultTransport":	ValueOf(&http.DefaultTransport).Elem(),
		"DetectContentType":	ValueOf(http.DetectContentType),
		"ErrAbortHandler":	ValueOf(&http.ErrAbortHandler).Elem(),
		"ErrBodyNotAllowed":	ValueOf(&http.ErrBodyNotAllowed).Elem(),
		"ErrBodyReadAfterClose":	ValueOf(&http.ErrBodyReadAfterClose).Elem(),
		"ErrContentLength":	ValueOf(&http.ErrContentLength).Elem(),
		"ErrHandlerTimeout":	ValueOf(&http.ErrHandlerTimeout).Elem(),
		"ErrHeaderTooLong":	ValueOf(&http.ErrHeaderTooLong).Elem(),
		"ErrHijacked":	ValueOf(&http.ErrHijacked).Elem(),
		"ErrLineTooLong":	ValueOf(&http.ErrLineTooLong).Elem(),
		"ErrMissingBoundary":	ValueOf(&http.ErrMissingBoundary).Elem(),
		"ErrMissingContentLength":	ValueOf(&http.ErrMissingContentLength).Elem(),
		"ErrMissingFile":	ValueOf(&http.ErrMissingFile).Elem(),
		"ErrNoCookie":	ValueOf(&http.ErrNoCookie).Elem(),
		"ErrNoLocation":	ValueOf(&http.ErrNoLocation).Elem(),
		"ErrNotMultipart":	ValueOf(&http.ErrNotMultipart).Elem(),
		"ErrNotSupported":	ValueOf(&http.ErrNotSupported).Elem(),
		"ErrServerClosed":	ValueOf(&http.ErrServerClosed).Elem(),
		"ErrShortBody":	ValueOf(&http.ErrShortBody).Elem(),
		"ErrSkipAltProtocol":	ValueOf(&http.ErrSkipAltProtocol).Elem(),
		"ErrUnexpectedTrailer":	ValueOf(&http.ErrUnexpectedTrailer).Elem(),
		"ErrUseLastResponse":	ValueOf(&http.ErrUseLastResponse).Elem(),
		"ErrWriteAfterFlush":	ValueOf(&http.ErrWriteAfterFlush).Elem(),
		"Error":	ValueOf(http.Error),
		"FileServer":	ValueOf(http.FileServer),
		"Get":	ValueOf(http.Get),
		"Handle":	ValueOf(http.Handle),
		"HandleFunc":	ValueOf(http.HandleFunc),
		"Head":	ValueOf(http.Head),
		"ListenAndServe":	ValueOf(http.ListenAndServe),
		"ListenAndServeTLS":	ValueOf(http.ListenAndServeTLS),
		"LocalAddrContextKey":	ValueOf(&http.LocalAddrContextKey).Elem(),
		"MaxBytesReader":	ValueOf(http.MaxBytesReader),
		"MethodConnect":	ValueOf(http.MethodConnect),
		"MethodDelete":	ValueOf(http.MethodDelete),
		"MethodGet":	ValueOf(http.MethodGet),
		"MethodHead":	ValueOf(http.MethodHead),
		"MethodOptions":	ValueOf(http.MethodOptions),
		"MethodPatch":	ValueOf(http.MethodPatch),
		"MethodPost":	ValueOf(http.MethodPost),
		"MethodPut":	ValueOf(http.MethodPut),
		"MethodTrace":	ValueOf(http.MethodTrace),
		"NewFileTransport":	ValueOf(http.NewFileTransport),
		"NewRequest":	ValueOf(http.NewRequest),
		"NewServeMux":	ValueOf(http.NewServeMux),
		"NoBody":	ValueOf(&http.NoBody).Elem(),
		"NotFound":	ValueOf(http.NotFound),
		"NotFoundHandler":	ValueOf(http.NotFoundHandler),
		"ParseHTTPVersion":	ValueOf(http.ParseHTTPVersion),
		"ParseTime":	ValueOf(http.ParseTime),
		"Post":	ValueOf(http.Post),
		"PostForm":	ValueOf(http.PostForm),
		"ProxyFromEnvironment":	ValueOf(http.ProxyFromEnvironment),
		"ProxyURL":	ValueOf(http.ProxyURL),
		"ReadRequest":	ValueOf(http.ReadRequest),
		"ReadResponse":	ValueOf(http.ReadResponse),
		"Redirect":	ValueOf(http.Redirect),
		"RedirectHandler":	ValueOf(http.RedirectHandler),
		"Serve":	ValueOf(http.Serve),
		"ServeContent":	ValueOf(http.ServeContent),
		"ServeFile":	ValueOf(http.ServeFile),
		"ServerContextKey":	ValueOf(&http.ServerContextKey).Elem(),
		"SetCookie":	ValueOf(http.SetCookie),
		"StateActive":	ValueOf(http.StateActive),
		"StateClosed":	ValueOf(http.StateClosed),
		"StateHijacked":	ValueOf(http.StateHijacked),
		"StateIdle":	ValueOf(http.StateIdle),
		"StateNew":	ValueOf(http.StateNew),
		"StatusAccepted":	ValueOf(http.StatusAccepted),
		"StatusAlreadyReported":	ValueOf(http.StatusAlreadyReported),
		"StatusBadGateway":	ValueOf(http.StatusBadGateway),
		"StatusBadRequest":	ValueOf(http.StatusBadRequest),
		"StatusConflict":	ValueOf(http.StatusConflict),
		"StatusContinue":	ValueOf(http.StatusContinue),
		"StatusCreated":	ValueOf(http.StatusCreated),
		"StatusExpectationFailed":	ValueOf(http.StatusExpectationFailed),
		"StatusFailedDependency":	ValueOf(http.StatusFailedDependency),
		"StatusForbidden":	ValueOf(http.StatusForbidden),
		"StatusFound":	ValueOf(http.StatusFound),
		"StatusGatewayTimeout":	ValueOf(http.StatusGatewayTimeout),
		"StatusGone":	ValueOf(http.StatusGone),
		"StatusHTTPVersionNotSupported":	ValueOf(http.StatusHTTPVersionNotSupported),
		"StatusIMUsed":	ValueOf(http.StatusIMUsed),
		"StatusInsufficientStorage":	ValueOf(http.StatusInsufficientStorage),
		"StatusInternalServerError":	ValueOf(http.StatusInternalServerError),
		"StatusLengthRequired":	ValueOf(http.StatusLengthRequired),
		"StatusLocked":	ValueOf(http.StatusLocked),
		"StatusLoopDetected":	ValueOf(http.StatusLoopDetected),
		"StatusMethodNotAllowed":	ValueOf(http.StatusMethodNotAllowed),
		"StatusMovedPermanently":	ValueOf(http.StatusMovedPermanently),
		"StatusMultiStatus":	ValueOf(http.StatusMultiStatus),
		"StatusMultipleChoices":	ValueOf(http.StatusMultipleChoices),
		"StatusNetworkAuthenticationRequired":	ValueOf(http.StatusNetworkAuthenticationRequired),
		"StatusNoContent":	ValueOf(http.StatusNoContent),
		"StatusNonAuthoritativeInfo":	ValueOf(http.StatusNonAuthoritativeInfo),
		"StatusNotAcceptable":	ValueOf(http.StatusNotAcceptable),
		"StatusNotExtended":	ValueOf(http.StatusNotExtended),
		"StatusNotFound":	ValueOf(http.StatusNotFound),
		"StatusNotImplemented":	ValueOf(http.StatusNotImplemented),
		"StatusNotModified":	ValueOf(http.StatusNotModified),
		"StatusOK":	ValueOf(http.StatusOK),
		"StatusPartialContent":	ValueOf(http.StatusPartialContent),
		"StatusPaymentRequired":	ValueOf(http.StatusPaymentRequired),
		"StatusPermanentRedirect":	ValueOf(http.StatusPermanentRedirect),
		"StatusPreconditionFailed":	ValueOf(http.StatusPreconditionFailed),
		"StatusPreconditionRequired":	ValueOf(http.StatusPreconditionRequired),
		"StatusProcessing":	ValueOf(http.StatusProcessing),
		"StatusProxyAuthRequired":	ValueOf(http.StatusProxyAuthRequired),
		"StatusRequestEntityTooLarge":	ValueOf(http.StatusRequestEntityTooLarge),
		"StatusRequestHeaderFieldsTooLarge":	ValueOf(http.StatusRequestHeaderFieldsTooLarge),
		"StatusRequestTimeout":	ValueOf(http.StatusRequestTimeout),
		"StatusRequestURITooLong":	ValueOf(http.StatusRequestURITooLong),
		"StatusRequestedRangeNotSatisfiable":	ValueOf(http.StatusRequestedRangeNotSatisfiable),
		"StatusResetContent":	ValueOf(http.StatusResetContent),
		"StatusSeeOther":	ValueOf(http.StatusSeeOther),
		"StatusServiceUnavailable":	ValueOf(http.StatusServiceUnavailable),
		"StatusSwitchingProtocols":	ValueOf(http.StatusSwitchingProtocols),
		"StatusTeapot":	ValueOf(http.StatusTeapot),
		"StatusTemporaryRedirect":	ValueOf(http.StatusTemporaryRedirect),
		"StatusText":	ValueOf(http.StatusText),
		"StatusTooManyRequests":	ValueOf(http.StatusTooManyRequests),
		"StatusUnauthorized":	ValueOf(http.StatusUnauthorized),
		"StatusUnavailableForLegalReasons":	ValueOf(http.StatusUnavailableForLegalReasons),
		"StatusUnprocessableEntity":	ValueOf(http.StatusUnprocessableEntity),
		"StatusUnsupportedMediaType":	ValueOf(http.StatusUnsupportedMediaType),
		"StatusUpgradeRequired":	ValueOf(http.StatusUpgradeRequired),
		"StatusUseProxy":	ValueOf(http.StatusUseProxy),
		"StatusVariantAlsoNegotiates":	ValueOf(http.StatusVariantAlsoNegotiates),
		"StripPrefix":	ValueOf(http.StripPrefix),
		"TimeFormat":	ValueOf(http.TimeFormat),
		"TimeoutHandler":	ValueOf(http.TimeoutHandler),
		"TrailerPrefix":	ValueOf(http.TrailerPrefix),
	},Types: map[string]Type{
		"Client":	TypeOf((*http.Client)(nil)).Elem(),
		"CloseNotifier":	TypeOf((*http.CloseNotifier)(nil)).Elem(),
		"ConnState":	TypeOf((*http.ConnState)(nil)).Elem(),
		"Cookie":	TypeOf((*http.Cookie)(nil)).Elem(),
		"CookieJar":	TypeOf((*http.CookieJar)(nil)).Elem(),
		"Dir":	TypeOf((*http.Dir)(nil)).Elem(),
		"File":	TypeOf((*http.File)(nil)).Elem(),
		"FileSystem":	TypeOf((*http.FileSystem)(nil)).Elem(),
		"Flusher":	TypeOf((*http.Flusher)(nil)).Elem(),
		"Handler":	TypeOf((*http.Handler)(nil)).Elem(),
		"HandlerFunc":	TypeOf((*http.HandlerFunc)(nil)).Elem(),
		"Header":	TypeOf((*http.Header)(nil)).Elem(),
		"Hijacker":	TypeOf((*http.Hijacker)(nil)).Elem(),
		"ProtocolError":	TypeOf((*http.ProtocolError)(nil)).Elem(),
		"PushOptions":	TypeOf((*http.PushOptions)(nil)).Elem(),
		"Pusher":	TypeOf((*http.Pusher)(nil)).Elem(),
		"Request":	TypeOf((*http.Request)(nil)).Elem(),
		"Response":	TypeOf((*http.Response)(nil)).Elem(),
		"ResponseWriter":	TypeOf((*http.ResponseWriter)(nil)).Elem(),
		"RoundTripper":	TypeOf((*http.RoundTripper)(nil)).Elem(),
		"ServeMux":	TypeOf((*http.ServeMux)(nil)).Elem(),
		"Server":	TypeOf((*http.Server)(nil)).Elem(),
		"Transport":	TypeOf((*http.Transport)(nil)).Elem(),
	},Proxies: map[string]Type{
		"CloseNotifier":	TypeOf((*CloseNotifier_net_http)(nil)).Elem(),
		"CookieJar":	TypeOf((*CookieJar_net_http)(nil)).Elem(),
		"File":	TypeOf((*File_net_http)(nil)).Elem(),
		"FileSystem":	TypeOf((*FileSystem_net_http)(nil)).Elem(),
		"Flusher":	TypeOf((*Flusher_net_http)(nil)).Elem(),
		"Handler":	TypeOf((*Handler_net_http)(nil)).Elem(),
		"Hijacker":	TypeOf((*Hijacker_net_http)(nil)).Elem(),
		"Pusher":	TypeOf((*Pusher_net_http)(nil)).Elem(),
		"ResponseWriter":	TypeOf((*ResponseWriter_net_http)(nil)).Elem(),
		"RoundTripper":	TypeOf((*RoundTripper_net_http)(nil)).Elem(),
	},Untypeds: map[string]string{
		"DefaultMaxHeaderBytes":	"int:1048576",
		"DefaultMaxIdleConnsPerHost":	"int:2",
		"MethodConnect":	"string:CONNECT",
		"MethodDelete":	"string:DELETE",
		"MethodGet":	"string:GET",
		"MethodHead":	"string:HEAD",
		"MethodOptions":	"string:OPTIONS",
		"MethodPatch":	"string:PATCH",
		"MethodPost":	"string:POST",
		"MethodPut":	"string:PUT",
		"MethodTrace":	"string:TRACE",
		"StatusAccepted":	"int:202",
		"StatusAlreadyReported":	"int:208",
		"StatusBadGateway":	"int:502",
		"StatusBadRequest":	"int:400",
		"StatusConflict":	"int:409",
		"StatusContinue":	"int:100",
		"StatusCreated":	"int:201",
		"StatusExpectationFailed":	"int:417",
		"StatusFailedDependency":	"int:424",
		"StatusForbidden":	"int:403",
		"StatusFound":	"int:302",
		"StatusGatewayTimeout":	"int:504",
		"StatusGone":	"int:410",
		"StatusHTTPVersionNotSupported":	"int:505",
		"StatusIMUsed":	"int:226",
		"StatusInsufficientStorage":	"int:507",
		"StatusInternalServerError":	"int:500",
		"StatusLengthRequired":	"int:411",
		"StatusLocked":	"int:423",
		"StatusLoopDetected":	"int:508",
		"StatusMethodNotAllowed":	"int:405",
		"StatusMovedPermanently":	"int:301",
		"StatusMultiStatus":	"int:207",
		"StatusMultipleChoices":	"int:300",
		"StatusNetworkAuthenticationRequired":	"int:511",
		"StatusNoContent":	"int:204",
		"StatusNonAuthoritativeInfo":	"int:203",
		"StatusNotAcceptable":	"int:406",
		"StatusNotExtended":	"int:510",
		"StatusNotFound":	"int:404",
		"StatusNotImplemented":	"int:501",
		"StatusNotModified":	"int:304",
		"StatusOK":	"int:200",
		"StatusPartialContent":	"int:206",
		"StatusPaymentRequired":	"int:402",
		"StatusPermanentRedirect":	"int:308",
		"StatusPreconditionFailed":	"int:412",
		"StatusPreconditionRequired":	"int:428",
		"StatusProcessing":	"int:102",
		"StatusProxyAuthRequired":	"int:407",
		"StatusRequestEntityTooLarge":	"int:413",
		"StatusRequestHeaderFieldsTooLarge":	"int:431",
		"StatusRequestTimeout":	"int:408",
		"StatusRequestURITooLong":	"int:414",
		"StatusRequestedRangeNotSatisfiable":	"int:416",
		"StatusResetContent":	"int:205",
		"StatusSeeOther":	"int:303",
		"StatusServiceUnavailable":	"int:503",
		"StatusSwitchingProtocols":	"int:101",
		"StatusTeapot":	"int:418",
		"StatusTemporaryRedirect":	"int:307",
		"StatusTooManyRequests":	"int:429",
		"StatusUnauthorized":	"int:401",
		"StatusUnavailableForLegalReasons":	"int:451",
		"StatusUnprocessableEntity":	"int:422",
		"StatusUnsupportedMediaType":	"int:415",
		"StatusUpgradeRequired":	"int:426",
		"StatusUseProxy":	"int:305",
		"StatusVariantAlsoNegotiates":	"int:506",
		"TimeFormat":	"string:Mon, 02 Jan 2006 15:04:05 GMT",
		"TrailerPrefix":	"string:Trailer:",
	},
	}
}

// --------------- proxy for net/http.CloseNotifier ---------------
type CloseNotifier_net_http struct {
	Object	interface{}
	CloseNotify_	func(interface{}) <-chan bool
}
func (Proxy *CloseNotifier_net_http) CloseNotify() <-chan bool {
	return Proxy.CloseNotify_(Proxy.Object)
}

// --------------- proxy for net/http.CookieJar ---------------
type CookieJar_net_http struct {
	Object	interface{}
	Cookies_	func(_proxy_obj_ interface{}, u *url.URL) []*http.Cookie
	SetCookies_	func(_proxy_obj_ interface{}, u *url.URL, cookies []*http.Cookie) 
}
func (Proxy *CookieJar_net_http) Cookies(u *url.URL) []*http.Cookie {
	return Proxy.Cookies_(Proxy.Object, u)
}
func (Proxy *CookieJar_net_http) SetCookies(u *url.URL, cookies []*http.Cookie)  {
	Proxy.SetCookies_(Proxy.Object, u, cookies)
}

// --------------- proxy for net/http.File ---------------
type File_net_http struct {
	Object	interface{}
	Close_	func(interface{}) error
	Read_	func(_proxy_obj_ interface{}, p []byte) (n int, err error)
	Readdir_	func(_proxy_obj_ interface{}, count int) ([]os.FileInfo, error)
	Seek_	func(_proxy_obj_ interface{}, offset int64, whence int) (int64, error)
	Stat_	func(interface{}) (os.FileInfo, error)
}
func (Proxy *File_net_http) Close() error {
	return Proxy.Close_(Proxy.Object)
}
func (Proxy *File_net_http) Read(p []byte) (n int, err error) {
	return Proxy.Read_(Proxy.Object, p)
}
func (Proxy *File_net_http) Readdir(count int) ([]os.FileInfo, error) {
	return Proxy.Readdir_(Proxy.Object, count)
}
func (Proxy *File_net_http) Seek(offset int64, whence int) (int64, error) {
	return Proxy.Seek_(Proxy.Object, offset, whence)
}
func (Proxy *File_net_http) Stat() (os.FileInfo, error) {
	return Proxy.Stat_(Proxy.Object)
}

// --------------- proxy for net/http.FileSystem ---------------
type FileSystem_net_http struct {
	Object	interface{}
	Open_	func(_proxy_obj_ interface{}, name string) (http.File, error)
}
func (Proxy *FileSystem_net_http) Open(name string) (http.File, error) {
	return Proxy.Open_(Proxy.Object, name)
}

// --------------- proxy for net/http.Flusher ---------------
type Flusher_net_http struct {
	Object	interface{}
	Flush_	func(interface{}) 
}
func (Proxy *Flusher_net_http) Flush()  {
	Proxy.Flush_(Proxy.Object)
}

// --------------- proxy for net/http.Handler ---------------
type Handler_net_http struct {
	Object	interface{}
	ServeHTTP_	func(interface{}, http.ResponseWriter, *http.Request) 
}
func (Proxy *Handler_net_http) ServeHTTP(unnamed0 http.ResponseWriter, unnamed1 *http.Request)  {
	Proxy.ServeHTTP_(Proxy.Object, unnamed0, unnamed1)
}

// --------------- proxy for net/http.Hijacker ---------------
type Hijacker_net_http struct {
	Object	interface{}
	Hijack_	func(interface{}) (net.Conn, *bufio.ReadWriter, error)
}
func (Proxy *Hijacker_net_http) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return Proxy.Hijack_(Proxy.Object)
}

// --------------- proxy for net/http.Pusher ---------------
type Pusher_net_http struct {
	Object	interface{}
	Push_	func(_proxy_obj_ interface{}, target string, opts *http.PushOptions) error
}
func (Proxy *Pusher_net_http) Push(target string, opts *http.PushOptions) error {
	return Proxy.Push_(Proxy.Object, target, opts)
}

// --------------- proxy for net/http.ResponseWriter ---------------
type ResponseWriter_net_http struct {
	Object	interface{}
	Header_	func(interface{}) http.Header
	Write_	func(interface{}, []byte) (int, error)
	WriteHeader_	func(interface{}, int) 
}
func (Proxy *ResponseWriter_net_http) Header() http.Header {
	return Proxy.Header_(Proxy.Object)
}
func (Proxy *ResponseWriter_net_http) Write(unnamed0 []byte) (int, error) {
	return Proxy.Write_(Proxy.Object, unnamed0)
}
func (Proxy *ResponseWriter_net_http) WriteHeader(unnamed0 int)  {
	Proxy.WriteHeader_(Proxy.Object, unnamed0)
}

// --------------- proxy for net/http.RoundTripper ---------------
type RoundTripper_net_http struct {
	Object	interface{}
	RoundTrip_	func(interface{}, *http.Request) (*http.Response, error)
}
func (Proxy *RoundTripper_net_http) RoundTrip(unnamed0 *http.Request) (*http.Response, error) {
	return Proxy.RoundTrip_(Proxy.Object, unnamed0)
}

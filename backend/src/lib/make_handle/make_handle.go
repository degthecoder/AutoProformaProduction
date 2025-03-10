package make_handle

import (
	"fmt"
	"log"
	"net/http"
)

type FuncType func(http.ResponseWriter, *http.Request)

func MakeHandle(hf FuncType) FuncType {
	return func(rw http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				switch err := err.(type) {
				case string:
					log.Println(err)
					fatalError(rw, req, err)
				case error:
					log.Println(err.Error())
					fatalError(rw, req, err.Error())
				//case *app.BadRequestError:
				//	badRequest(rw, req)
				//case *app.NotFoundError:
				//	notFound(rw, req)
				default:
					fmt.Printf("%s %v", "unknown error", err)
					log.Printf("%s %v\n", "unknown error", err)
				}
			}
		}()

		hf(rw, req)
	}
}

func fatalError(rw http.ResponseWriter, req *http.Request, msg string) {
	if req.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		//list := strings.Split(string(debug.Stack()[:]), "\n")
		//errorMsg := fmt.Sprintf("%s %s", app.T("Fatal Error:"), app.T(msg))

		//if app.Settings.Debug {
		//	errorMsg += "\n" + strings.Join(list[6:], "\n")
		//}

		//http.Error(rw, errorMsg, http.StatusInternalServerError)
		return
	}

	//ctx := context.NewContext(rw, req)
	//ctx.AddStr("vlErrorCode", "")
	//ctx.AddStr("rsTitle", app.T("Fatal Error"))
	//ctx.AddStr("rsHomePage", app.T("Anasayfa"))
	//ctx.AddStr("rsLoginPage", app.T("Giriş"))
	//
	//if app.Settings.Debug {
	//	buf := new(util.Buf)
	//	list := strings.Split(string(debug.Stack()[:]), "\n")
	//
	//	msg = msg + "<br>"
	//	msg = msg + strings.Join(list[6:], "<br>")
	//	buf.Add("<p style='max-height:300px;overflow-y:auto'>%s</p>", msg)
	//	ctx.AddHtml("vlExp", buf.String())
	//}
	//content.Include(ctx)
	//content.Default(ctx)
	//ctx.Render("error.html")
}

func badRequest(rw http.ResponseWriter, req *http.Request) {
	if req.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	//ctx := context.NewContext(rw, req)
	//ctx.Rw.WriteHeader(http.StatusBadRequest)
	//
	//ctx.AddStr("vlErrorCode", strconv.Itoa(http.StatusBadRequest))
	//ctx.AddStr("rsTitle", app.T("Yetki Yok"))
	//ctx.AddStr("rsHomePage", app.T("Anasayfa"))
	//ctx.AddStr("rsLoginPage", app.T("Giriş"))
	//ctx.AddStr("vlExp", "<p>"+app.T("Bu sayfaya erişim hakkınız yok.")+"</p>")
	//
	//content.Include(ctx)
	//content.Default(ctx)
	//ctx.Render("error.html")
}

func notFound(rw http.ResponseWriter, req *http.Request) {
	if req.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		http.Error(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	//ctx := context.NewContext(rw, req)
	//ctx.Rw.WriteHeader(http.StatusNotFound)
	//ctx.AddStr("vlErrorCode", strconv.Itoa(http.StatusNotFound))
	//ctx.AddStr("rsTitle", app.T("404 hatası oluştu.."))
	//ctx.AddStr("rsHomePage", app.T("Anasayfa"))
	//ctx.AddStr("rsLoginPage", app.T("Giriş"))
	//ctx.AddStr("vlExp", "<p>"+app.T("Hata! Sayfa bulunamadı.")+"</p>")
	//
	//content.Include(ctx)
	//content.Default(ctx)
	//ctx.Render("error.html")
}

package api

import (
	"bytes"
	"j2ee/constant"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

type CaptchaType struct {
	CaptchaID string `json:"captchaID"` //图形验证码ID
	CaptchaURL  string `json:"captchaURL"`  //图形验证码图片url
}

// GetCaptcha 获取图形验证码（验证码ID和图片url）
// [get] /api/v0/get_captcha
func GetCaptcha(c *gin.Context) {
	length := 5 // captcha.DefaultLen
	captchaID := captcha.NewLen(length)
	var captchaServer CaptchaType
	captchaServer.CaptchaID = captchaID
	captchaServer.CaptchaURL = "/api/v0" + "/captcha/" + captchaID + ".png"
	if captchaID == "" {
		c.JSON(constant.ERROR_FAIL_GETTING_CAPTCHA, captchaServer)
	} else {
		c.JSON(constant.SUCCESS, captchaServer)
	}
}

// GetCaptchaImage 获取验证码图片
// [get] /api/v0/captcha/:captchaID.png 即 /api/v0/CaptchaURL
func GetCaptchaImage(c *gin.Context) {
	ServeHTTP(c.Writer, c.Request)
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dir, file := path.Split(r.URL.Path)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	if ext == "" || id == "" {
		http.NotFound(w, r)
		return
	}
	if r.FormValue("reload") != "" {
		captcha.Reload(id)
	}
	lang := strings.ToLower(r.FormValue("lang"))
	download := path.Base(dir) == "download"
	if Serve(w, r, id, ext, lang, download, captcha.StdWidth, captcha.StdHeight) == captcha.ErrNotFound {
		http.NotFound(w, r)
	}
}

func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		err := captcha.WriteImage(&content, id, width, height)
		if err != nil {
			return err
		}
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		err := captcha.WriteAudio(&content, id, lang)
		if err != nil {
			return err
		}
	default:
		return captcha.ErrNotFound
	}
	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}

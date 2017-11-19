package admin

import (
	"github.com/genshen/blog/components/context/admin"
	"github.com/genshen/blog/components/utils"
	"github.com/astaxie/beego"
	"html/template"
)

const (
	UserId   = "AdminUserID"
	Username = "AdminUsername"
)

type AuthController struct {
	BaseController
}

func (this *AuthController) SignIn() {
	if this.isUserLogin() {
		this.Redirect(beego.URLFor("PanelController.Get"), 302)
		return;
	}
	if (this.Ctx.Request.Method == "POST") {
		sign_in_form := admin.SignInForm{Email: this.GetString("email"), Password: this.GetString("password")}
		if errs := sign_in_form.Valid(); errs != nil {
			s := utils.NewInstant(errs, map[string]string{"email": sign_in_form.Email, "password": ""})
			this.Data["json"] = &utils.SimpleJsonResponse{Status: 0, Error: &s}
		} else {
			this.LoginUser(sign_in_form.ID, sign_in_form.Username)
			next := this.GetString("next")
			if len(next) > 0 && next[0] != '/' {
				next = "/" + next
			} else if next == "" {
				next = AdminPrefix
			}
			this.Data["json"] = &utils.SimpleJsonResponse{Status: 1, Addition: next}
		}
		this.ServeJSON()
	} else {
		this.Data["dev"] = false;
		this.TplName = "admin/auth/sign_in.html"
	}
}

/*only in Dev Mode*/
type SignUpForm struct {
	Email    string
	Username string
	Password string
}

func (this *AuthController) SignUp() {
	if (this.Ctx.Request.Method == "POST") { //todo
		email := this.GetString("Email")
		username := this.GetString("Username")
		password := this.GetString("Password")
		admin.CreateUser(username, email, password)
		this.Data["json"] = &utils.SimpleJsonResponse{Status: 1, Addition: ""};
		this.ServeJSON()
	} else {
		this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())
		this.Data["Form"] = &SignUpForm{}
		this.TplName = "admin/auth/sign_up.html"
	}
}

func (this *AuthController) SignOut() {
	this.DelSession(UserId)
	this.DelSession(Username)
	this.Redirect(AdminAuthUri, 302)
}

func (this *AuthController) LoginUser(id string, username string) {
	this.SetSession(UserId, id)
	this.SetSession(Username, username)
}

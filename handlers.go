//PostShowLogin handles logging the user in
func (m *Repository) PostShowLogin(w http.ResponseWriter, r *http.Request){
	_ = m.App.Session.RenewToken(r.Context())

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")
	
	form := forms.New(r.PostForm)
	form.Required("email", "password")
	form.IsEmail("email")

	if !form.Valid(){
		render.Template(w, r, "login.page.tmpl", &models.TemplateData{
			Form: form,
		})
		return
	}
}

id, _, err := m.DB.Authenticate(email, password)
if err != nil {
	log.Println(err)

	m.App.Session.Put(r.Context(), "error", "Invalid login credentials")
	http.Redirect(W, r, "/user/login", http.StatusSeeOther)
	return
}

m.App.Session.Put(r.Context(), "user_id", id)
m.App.Session.Put(r.Context(), "flash", "Logged in successfully")
http.Redirect(w, r, "/", http.SatusSeeOther)
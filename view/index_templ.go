// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package view

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/de4et/command-constructor/types"

func Index(commandTemplates []types.CommandTemplate, user *types.User) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			if len(commandTemplates) == 0 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"no-command-templates\"><div class=\"no-command-templates-common\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"nothing-img\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" version=\"1.1\" x=\"0px\" y=\"0px\" viewBox=\"0 0 32 40\" style=\"enable-background:new 0 0 32 32;\" xml:space=\"preserve\"><g><path d=\"M27.28,25.86C29.59,23.22,31,19.78,31,16c0-8.27-6.73-15-15-15S1,7.73,1,16s6.73,15,15,15c3.78,0,7.22-1.41,9.86-3.72   l3.43,3.43C29.49,30.9,29.74,31,30,31s0.51-0.1,0.71-0.29c0.39-0.39,0.39-1.02,0-1.41L27.28,25.86z M12,11c1.1,0,2,0.9,2,2   s-0.9,2-2,2s-2-0.9-2-2S10.9,11,12,11z M21,20.91c-0.13,0.06-0.27,0.09-0.4,0.09c-0.38,0-0.75-0.22-0.92-0.6   C19.03,18.94,17.58,18,16,18s-3.03,0.94-3.68,2.41c-0.22,0.5-0.81,0.74-1.32,0.51c-0.5-0.22-0.73-0.81-0.51-1.32   C11.46,17.41,13.63,16,16,16s4.54,1.41,5.5,3.59C21.73,20.1,21.5,20.69,21,20.91z M20,15c-1.1,0-2-0.9-2-2s0.9-2,2-2s2,0.9,2,2   S21.1,15,20,15z\"></path></g><text x=\"0\" y=\"47\" fill=\"#000000\" font-size=\"5px\" font-weight=\"bold\" font-family=\"&#39;Helvetica Neue&#39;, Helvetica, Arial-Unicode, Arial, Sans-serif\">Created by ZAK</text><text x=\"0\" y=\"52\" fill=\"#000000\" font-size=\"5px\" font-weight=\"bold\" font-family=\"&#39;Helvetica Neue&#39;, Helvetica, Arial-Unicode, Arial, Sans-serif\">from the Noun Project</text></svg> <span class=\"nothing-text-common\">Здесь пока ничего нет. ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if user == nil {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a id=\"advise-button-login\" class=\"advise-button\">Войдите</a> или <a id=\"advise-button-reg\" class=\"advise-button\">зарегистрируйтесь</a><script>\r\n\t\t\t\t\t\t\t\taddOpenableBlockListeners(\".login_block\", \"#advise-button-login\", \"#login_content\");\r\n\t\t\t\t\t\t\t\taddOpenableBlockListeners(\".reg_block\", \"#advise-button-reg\", \"#reg_content\");\r\n\t\t\t\t\t\t\t\taddFocusListeners(\"#advise-button-login\", \"#login_button\");\r\n\t\t\t\t\t\t\t\taddFocusListeners(\"#advise-button-reg\", \"#reg_button\");\r\n\t\t\t\t\t\t\t</script>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a class=\"advise-button\">Создайте</a> свой первый шаблон")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"command-templates\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				for _, v := range commandTemplates {
					templ_7745c5c3_Err = CommandTemplate(v).Render(ctx, templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = page("Command-constructor", user).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate

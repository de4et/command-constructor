// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package view

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/de4et/command-constructor/types"

func CreateTemplate(user *types.User) templ.Component {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"create-menu\"><h2 class=\"body-title\">Создание шаблона</h2><div class=\"preview-command\"><span class=\"preview-command-text preview-command-commandname\">pscp</span><div data-paramtype=\"constant\" class=\"preview-command-commandparam\"><span class=\"preview-command-text preview-commandparam-empty\">-i</span></div><div data-paramtype=\"template\" class=\"preview-command-commandparam\"><div class=\"preview-command-popup\"><select name=\"preview-command-popup\" class=\"preview-command-popup-select\"><option value=\"Youtube\">Youtube</option> <option value=\"Twitter\" selected=\"true\">Twitter</option> <option value=\"Google\">Google</option></select><script>\r\n\t\t\t\t\t\tdocument.querySelector(\".preview-command-popup-select\").selectedIndex=1;\r\n\t\t\t\t\t</script></div></div><div data-paramtype=\"template\" class=\"preview-command-commandparam\"><input class=\"preview-command-nameless\" value=\"root@127.0.0.1\"></div><div data-paramtype=\"template\" class=\"preview-command-commandparam\"><span class=\"preview-command-string-name\">-r</span> <input class=\"preview-command-string-input\" value=\"~/\"></div></div><form class=\"create-form\"><input placeholder=\"Название\" class=\"create-form-input form-input-name\"> <textarea placeholder=\"Описание\" class=\"create-form-input form-input-description\"></textarea> <input placeholder=\"Имя команды\" class=\"create-form-input form-input-commandName\"><div class=\"create-form-params\"><div class=\"create-form-param\"></div></div><button type=\"submit\" class=\"create-form-button\">Создать</button></form></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = page("Создание шаблона", user).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate

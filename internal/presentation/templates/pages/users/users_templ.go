// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package users

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"github.com/iota-agency/iota-erp/internal/domain/user"
	"github.com/iota-agency/iota-erp/internal/presentation/templates/components/base"
	"github.com/iota-agency/iota-erp/internal/presentation/templates/components/base/button"
	"github.com/iota-agency/iota-erp/internal/presentation/templates/icons"
	"github.com/iota-agency/iota-erp/internal/presentation/templates/layouts"
	"github.com/iota-agency/iota-erp/internal/presentation/types"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func Index(pageContext *types.PageContext, users []*user.User) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"m-6\"><h1 class=\"text-2xl font-medium\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(pageContext.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "NavigationLinks.Users"}))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/presentation/templates/pages/users/users.templ`, Line: 18, Col: 98}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1><div class=\"mt-5 table-bg\"><section></section>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var4 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
				for _, user := range users {
					templ_7745c5c3_Var5 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
						templ_7745c5c3_Var6 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
							var templ_7745c5c3_Var7 string
							templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(user.FullName())
							if templ_7745c5c3_Err != nil {
								return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/presentation/templates/pages/users/users.templ`, Line: 31, Col: 24}
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							return templ_7745c5c3_Err
						})
						templ_7745c5c3_Err = base.TableCell().Render(templ.WithChildren(ctx, templ_7745c5c3_Var6), templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						templ_7745c5c3_Var8 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
							if v := user.LastAction; v != nil {
								_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-data=\"relativeformat\"><span x-text=\"")
								if templ_7745c5c3_Err != nil {
									return templ_7745c5c3_Err
								}
								var templ_7745c5c3_Var9 string
								templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("format('%s')", v.String()))
								if templ_7745c5c3_Err != nil {
									return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/presentation/templates/pages/users/users.templ`, Line: 36, Col: 63}
								}
								_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
								if templ_7745c5c3_Err != nil {
									return templ_7745c5c3_Err
								}
								_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></span></div>")
								if templ_7745c5c3_Err != nil {
									return templ_7745c5c3_Err
								}
							} else {
								var templ_7745c5c3_Var10 string
								templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(pageContext.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Unknown"}))
								if templ_7745c5c3_Err != nil {
									return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/presentation/templates/pages/users/users.templ`, Line: 39, Col: 88}
								}
								_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
								if templ_7745c5c3_Err != nil {
									return templ_7745c5c3_Err
								}
							}
							return templ_7745c5c3_Err
						})
						templ_7745c5c3_Err = base.TableCell().Render(templ.WithChildren(ctx, templ_7745c5c3_Var8), templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						templ_7745c5c3_Var11 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-data=\"relativeformat\"><span class=\"capitalize\" x-text=\"")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							var templ_7745c5c3_Var12 string
							templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("format('%s')", user.UpdatedAt.String()))
							if templ_7745c5c3_Err != nil {
								return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/presentation/templates/pages/users/users.templ`, Line: 44, Col: 94}
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></span></div>")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							return templ_7745c5c3_Err
						})
						templ_7745c5c3_Err = base.TableCell().Render(templ.WithChildren(ctx, templ_7745c5c3_Var11), templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						templ_7745c5c3_Var13 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
							templ_7745c5c3_Var14 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
								templ_7745c5c3_Err = icons.Pencil(icons.Props{Size: "20"}).Render(ctx, templ_7745c5c3_Buffer)
								if templ_7745c5c3_Err != nil {
									return templ_7745c5c3_Err
								}
								return templ_7745c5c3_Err
							})
							templ_7745c5c3_Err = button.Secondary(button.Props{Fixed: true, Size: button.SizeSM, Class: "btn-fixed", Href: fmt.Sprintf("/users/%d", user.Id)}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var14), templ_7745c5c3_Buffer)
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							return templ_7745c5c3_Err
						})
						templ_7745c5c3_Err = base.TableCell().Render(templ.WithChildren(ctx, templ_7745c5c3_Var13), templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						return templ_7745c5c3_Err
					})
					templ_7745c5c3_Err = base.TableRow().Render(templ.WithChildren(ctx, templ_7745c5c3_Var5), templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				return templ_7745c5c3_Err
			})
			templ_7745c5c3_Err = base.Table([]*base.TableColumn{
				{Label: pageContext.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "FullName"}), Key: "fullName"},
				{Label: pageContext.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "LastAction"}), Key: "lastAction"},
				{Label: pageContext.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "UpdatedAt"}), Key: "updatedAt"},
				{Label: pageContext.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Actions"}), Class: "w-16"},
			}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var4), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layouts.Authenticated(pageContext).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

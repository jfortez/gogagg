// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.696
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/jfortez/gogagg/web/templates/layouts"
import "github.com/jfortez/gogagg/model"
import "fmt"
import "strconv"

func Index(requestedMessages []model.RequestedMessages) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script type=\"text/javascript\">\n   window.onload = function () {\n    var conn;\n    var msg = document.getElementById(\"msg\");\n    var log = document.getElementById(\"log\");\n\n    function appendLog(item) {\n        var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;\n        log.appendChild(item);\n        if (doScroll) {\n            log.scrollTop = log.scrollHeight - log.clientHeight;\n        }\n    }\n\n    document.getElementById(\"form\").onsubmit = function () {\n        if (!conn) {\n            return false;\n        }\n        if (!msg.value) {\n            return false;\n        }\n        conn.send(msg.value);\n        msg.value = \"\";\n        return false;\n    };\n\n    if (window[\"WebSocket\"]) {\n        conn = new WebSocket(\"ws://\" + document.location.host + \"/ws\");\n        conn.onclose = function (evt) {\n            var item = document.createElement(\"div\");\n            item.innerHTML = \"<b>Connection closed.</b>\";\n            appendLog(item);\n        };\n        conn.onmessage = function (evt) {\n            var messages = evt.data.split('\\n');\n            for (var i = 0; i < messages.length; i++) {\n                var item = document.createElement(\"div\");\n                item.innerText = messages[i];\n                appendLog(item);\n            }\n        };\n    } else {\n        var item = document.createElement(\"div\");\n        item.innerHTML = \"<b>Your browser does not support WebSockets.</b>\";\n        appendLog(item);\n    }\n};\n  </script> <div class=\"h-screen grid grid-cols-12 grid-rows-1\"><div class=\"col-span-3 row-span-12  border-r border-gray-100 dark:border-gray-700\"><div class=\"p-4\"><div class=\"flex items-center justify-between mb-4\"><h1 class=\"text-2xl font-bold text-white\">Messages <span class=\"text-xl text-gray-300\"><span id=\"messages-count\">(")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(len(requestedMessages)))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/templates/index.templ`, Line: 62, Col: 159}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(")</span></span></h1><div class=\"flex items-center gap-2\"><button type=\"button\" class=\"font-medium rounded-lg text-xs p-2 text-center inline-flex items-center justify-center text-gray-900 bg-white hover:bg-gray-100 dark:text-white dark:hover:bg-gray-700 dark:bg-gray-900 dark:hover:text-white outline-0\"><svg width=\"15\" height=\"15\" viewBox=\"0 0 15 15\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M3.625 7.5C3.625 8.12132 3.12132 8.625 2.5 8.625C1.87868 8.625 1.375 8.12132 1.375 7.5C1.375 6.87868 1.87868 6.375 2.5 6.375C3.12132 6.375 3.625 6.87868 3.625 7.5ZM8.625 7.5C8.625 8.12132 8.12132 8.625 7.5 8.625C6.87868 8.625 6.375 8.12132 6.375 7.5C6.375 6.87868 6.87868 6.375 7.5 6.375C8.12132 6.375 8.625 6.87868 8.625 7.5ZM12.5 8.625C13.1213 8.625 13.625 8.12132 13.625 7.5C13.625 6.87868 13.1213 6.375 12.5 6.375C11.8787 6.375 11.375 6.87868 11.375 7.5C11.375 8.12132 11.8787 8.625 12.5 8.625Z\" fill=\"currentColor\" fill-rule=\"evenodd\" clip-rule=\"evenodd\"></path></svg> <span class=\"sr-only\">Icon description</span></button></div></div><ul id=\"messages-list\" class=\"divide-y divide-gray-100 dark:divide-gray-700 space-y-1 -mx-4\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, requestedMessage := range requestedMessages {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li><a href=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 templ.SafeURL = templ.URL(fmt.Sprintf("/message/%v", requestedMessage.MessageId))
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var4)))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"flex items-center gap-2 p-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white hover:cursor-pointer\"><img src=\"https://cdn-icons-png.freepik.com/512/6596/6596121.png\" class=\"w-10 h-10 min-w-10 min-h-10 rounded-full\" alt=\"avatar\"><div class=\"flex-1 min-w-0\"><div class=\"text-sm font-medium text-gray-900 dark:text-white truncate\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(requestedMessage.RequestedUserName)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/templates/index.templ`, Line: 76, Col: 118}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"text-xs text-gray-500 dark:text-gray-300 truncate w-full\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(requestedMessage.Content)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/templates/index.templ`, Line: 78, Col: 37}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul></div></div><div class=\"col-span-9 row-span-11\"><div class=\"flex items-center justify-between py-2 px-4 border-b border-gray-100 dark:border-gray-700\"><div class=\"flex items-center gap-2\"><img src=\"https://cdn-icons-png.freepik.com/512/6596/6596121.png\" class=\"w-10 h-10 min-w-10 min-h-10 rounded-full\" alt=\"avatar\"><div class=\"flex flex-col\"><div class=\"text-sm font-medium text-gray-900 dark:text-white\">Bonnie Green</div><div class=\"text-xs text-gray-500 dark:text-gray-300\">Active since 12:00 PM</div></div></div><div class=\"flex items-center gap-2\"><svg class=\"w-6 h-6 text-gray-800 dark:text-white\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" fill=\"none\" viewBox=\"0 0 24 24\"><path stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M10 11h2v5m-2 0h4m-2.592-8.5h.01M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z\"></path></svg></div></div><ul id=\"messages-list\" class=\"p-4 flex flex-col gap-4\"><li class=\"flex items-start gap-2.5\"><img src=\"https://cdn-icons-png.freepik.com/512/6596/6596121.png\" class=\"w-10 h-10 min-w-10 min-h-10 rounded-full\" alt=\"avatar\"><div class=\"flex flex-col w-full max-w-[320px] leading-1.5 p-4 border-gray-200 bg-gray-100 rounded-e-xl rounded-es-xl dark:bg-gray-700\"><div class=\"flex items-center space-x-2 rtl:space-x-reverse\"><span class=\"text-sm font-semibold text-gray-900 dark:text-white\">Bonnie Green</span> <span class=\"text-sm font-normal text-gray-500 dark:text-gray-400\">11:46</span></div><p class=\"text-sm font-normal py-2.5 text-gray-900 dark:text-white\">That's awesome. I think our users will really appreciate the improvements.</p><span class=\"text-sm font-normal text-gray-500 dark:text-gray-400\">Delivered</span></div><button id=\"dropdownMenuIconButton\" data-dropdown-toggle=\"dropdownDots\" data-dropdown-placement=\"bottom-start\" class=\"inline-flex self-center items-center p-2 text-sm font-medium text-center text-gray-900 bg-white rounded-lg hover:bg-gray-100 focus:ring-4 focus:outline-none dark:text-white focus:ring-gray-50 dark:bg-gray-900 dark:hover:bg-gray-800 dark:focus:ring-gray-600\" type=\"button\"><svg class=\"w-4 h-4 text-gray-500 dark:text-gray-400\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 4 15\"><path d=\"M3.5 1.5a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0Zm0 6.041a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0Zm0 5.959a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0Z\"></path></svg></button><div id=\"dropdownDots\" class=\"z-10 hidden bg-white divide-y divide-gray-100 rounded-lg shadow w-40 dark:bg-gray-700 dark:divide-gray-600\"><ul class=\"py-2 text-sm text-gray-700 dark:text-gray-200\" aria-labelledby=\"dropdownMenuIconButton\"><li><a href=\"#\" class=\"block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white\">Reply</a></li><li><a href=\"#\" class=\"block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white\">Forward</a></li><li><a href=\"#\" class=\"block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white\">Copy</a></li><li><a href=\"#\" class=\"block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white\">Report</a></li><li><a href=\"#\" class=\"block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white\">Delete</a></li></ul></div></li><li class=\"flex items-start gap-2.5\"><img src=\"https://cdn-icons-png.freepik.com/512/6596/6596121.png\" class=\"w-10 h-10 min-w-10 min-h-10 rounded-full\" alt=\"avatar\"><div class=\"flex flex-col w-full max-w-[320px] leading-1.5 p-4 border-gray-200 bg-gray-100 rounded-e-xl rounded-es-xl dark:bg-gray-700\"><div class=\"flex items-center space-x-2 rtl:space-x-reverse\"><span class=\"text-sm font-semibold text-gray-900 dark:text-white\">Bonnie Green</span> <span class=\"text-sm font-normal text-gray-500 dark:text-gray-400\">11:46</span></div><p class=\"text-sm font-normal py-2.5 text-gray-900 dark:text-white\">That's awesome. I think our users will really appreciate the improvements.</p><span class=\"text-sm font-normal text-gray-500 dark:text-gray-400\">Delivered</span></div><button id=\"dropdownMenuIconButton\" data-dropdown-toggle=\"dropdownDots\" data-dropdown-placement=\"bottom-start\" class=\"inline-flex self-center items-center p-2 text-sm font-medium text-center text-gray-900 bg-white rounded-lg hover:bg-gray-100 focus:ring-4 focus:outline-none dark:text-white focus:ring-gray-50 dark:bg-gray-900 dark:hover:bg-gray-800 dark:focus:ring-gray-600\" type=\"button\"><svg class=\"w-4 h-4 text-gray-500 dark:text-gray-400\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 4 15\"><path d=\"M3.5 1.5a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0Zm0 6.041a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0Zm0 5.959a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0Z\"></path></svg></button><div id=\"dropdownDots\" class=\"z-10 hidden bg-white divide-y divide-gray-100 rounded-lg shadow w-40 dark:bg-gray-700 dark:divide-gray-600\"><ul class=\"py-2 text-sm text-gray-700 dark:text-gray-200\" aria-labelledby=\"dropdownMenuIconButton\"><li><a href=\"#\" class=\"block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white\">Reply</a></li><li><a href=\"#\" class=\"block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white\">Forward</a></li><li><a href=\"#\" class=\"block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white\">Copy</a></li><li><a href=\"#\" class=\"block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white\">Report</a></li><li><a href=\"#\" class=\"block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white\">Delete</a></li></ul></div></li><li class=\"flex items-start gap-2.5 justify-start flex-row-reverse\"><img src=\"https://cdn-icons-png.freepik.com/512/6596/6596121.png\" class=\"w-10 h-10 min-w-10 min-h-10 rounded-full\" alt=\"avatar\"><div class=\"flex flex-col w-full max-w-[320px] leading-1.5 p-4 border-gray-200 bg-gray-100 rounded-s-xl rounded-b-xl dark:bg-gray-700\"><div class=\"flex items-center space-x-2 rtl:space-x-reverse\"><span class=\"text-sm font-semibold text-gray-900 dark:text-white\">Bonnie Green</span> <span class=\"text-sm font-normal text-gray-500 dark:text-gray-400\">11:46</span></div><p class=\"text-sm font-normal py-2.5 text-gray-900 dark:text-white\">That's awesome. I think our users will really appreciate the improvements.</p><span class=\"text-sm font-normal text-gray-500 dark:text-gray-400\">Delivered</span></div><button id=\"dropdownMenuIconButton\" data-dropdown-toggle=\"dropdownDots\" data-dropdown-placement=\"bottom-start\" class=\"inline-flex self-center items-center p-2 text-sm font-medium text-center text-gray-900 bg-white rounded-lg hover:bg-gray-100 focus:ring-4 focus:outline-none dark:text-white focus:ring-gray-50 dark:bg-gray-900 dark:hover:bg-gray-800 dark:focus:ring-gray-600\" type=\"button\"><svg class=\"w-4 h-4 text-gray-500 dark:text-gray-400\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 4 15\"><path d=\"M3.5 1.5a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0Zm0 6.041a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0Zm0 5.959a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0Z\"></path></svg></button><div id=\"dropdownDots\" class=\"z-10 hidden bg-white divide-y divide-gray-100 rounded-lg shadow w-40 dark:bg-gray-700 dark:divide-gray-600\"><ul class=\"py-2 text-sm text-gray-700 dark:text-gray-200\" aria-labelledby=\"dropdownMenuIconButton\"><li><a href=\"#\" class=\"block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white\">Reply</a></li><li><a href=\"#\" class=\"block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white\">Forward</a></li><li><a href=\"#\" class=\"block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white\">Copy</a></li><li><a href=\"#\" class=\"block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white\">Report</a></li><li><a href=\"#\" class=\"block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white\">Delete</a></li></ul></div></li></ul></div><div class=\"col-span-9 row-span-1\"><form class=\"px-2 pb-2\"><label for=\"chat\" class=\"sr-only\">Your message</label><div class=\"flex items-center px-3 py-2 rounded-lg bg-gray-50 dark:bg-gray-700\"><button type=\"button\" class=\"inline-flex justify-center p-2 text-gray-500 rounded-lg cursor-pointer hover:text-gray-900 hover:bg-gray-100 dark:text-gray-400 dark:hover:text-white dark:hover:bg-gray-600\"><svg class=\"w-5 h-5\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 20 18\"><path fill=\"currentColor\" d=\"M13 5.5a.5.5 0 1 1-1 0 .5.5 0 0 1 1 0ZM7.565 7.423 4.5 14h11.518l-2.516-3.71L11 13 7.565 7.423Z\"></path> <path stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M18 1H2a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1Z\"></path> <path stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M13 5.5a.5.5 0 1 1-1 0 .5.5 0 0 1 1 0ZM7.565 7.423 4.5 14h11.518l-2.516-3.71L11 13 7.565 7.423Z\"></path></svg> <span class=\"sr-only\">Upload image</span></button> <button type=\"button\" class=\"p-2 text-gray-500 rounded-lg cursor-pointer hover:text-gray-900 hover:bg-gray-100 dark:text-gray-400 dark:hover:text-white dark:hover:bg-gray-600\"><svg class=\"w-5 h-5\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 20 20\"><path stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M13.408 7.5h.01m-6.876 0h.01M19 10a9 9 0 1 1-18 0 9 9 0 0 1 18 0ZM4.6 11a5.5 5.5 0 0 0 10.81 0H4.6Z\"></path></svg> <span class=\"sr-only\">Add emoji</span></button> <textarea id=\"chat\" rows=\"1\" class=\"block mx-4 p-2.5 w-full text-sm text-gray-900 bg-white rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-800 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500\" placeholder=\"Your message...\"></textarea> <button type=\"submit\" class=\"inline-flex justify-center p-2 text-blue-600 rounded-full cursor-pointer hover:bg-blue-100 dark:text-blue-500 dark:hover:bg-gray-600\"><svg class=\"w-5 h-5 rotate-90 rtl:-rotate-90\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 18 20\"><path d=\"m17.914 18.594-8-18a1 1 0 0 0-1.828 0l-8 18a1 1 0 0 0 1.157 1.376L8 18.281V9a1 1 0 0 1 2 0v9.281l6.758 1.689a1 1 0 0 0 1.156-1.376Z\"></path></svg> <span class=\"sr-only\">Send message</span></button></div></form></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layouts.Main().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

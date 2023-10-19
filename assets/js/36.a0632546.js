(window.webpackJsonp=window.webpackJsonp||[]).push([[36],{321:function(a,t,s){"use strict";s.r(t);var e=s(10),n=Object(e.a)({},(function(){var a=this,t=a._self._c;return t("ContentSlotsDistributor",{attrs:{"slot-key":a.$parent.slotKey}},[t("h1",{attrs:{id:"快速上手"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#快速上手"}},[a._v("#")]),a._v(" 快速上手")]),a._v(" "),t("h2",{attrs:{id:"安装"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#安装"}},[a._v("#")]),a._v(" 安装")]),a._v(" "),t("div",{staticClass:"language-shell extra-class"},[t("pre",{pre:!0,attrs:{class:"language-shell"}},[t("code",[a._v("go "),t("span",{pre:!0,attrs:{class:"token function"}},[a._v("install")]),a._v(" github.com/go-zing/gozz@latest\n")])])]),t("h2",{attrs:{id:"使用"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#使用"}},[a._v("#")]),a._v(" 使用")]),a._v(" "),t("p",[t("code",[a._v("Gozz")]),a._v(" CLI 基于 "),t("a",{attrs:{href:"https://github.com/spf13/cobra",target:"_blank",rel:"noopener noreferrer"}},[a._v("cobra"),t("OutboundLink")],1),a._v(" 构建，命令行交互语法遵循格式：")]),a._v(" "),t("div",{staticClass:"language-shell extra-class"},[t("pre",{pre:!0,attrs:{class:"language-shell"}},[t("code",[a._v("gozz "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("--GLOBAL-FLAGS"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("COMMAND"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("--COMMAND-FLAGS"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("ARGS"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v("\n")])])]),t("h2",{attrs:{id:"环境"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#环境"}},[a._v("#")]),a._v(" 环境")]),a._v(" "),t("p",[t("code",[a._v("Gozz")]),a._v(" 在启动时会自动加载用户目录 "),t("code",[a._v("~/.gozz/plugins/")]),a._v(" 下的 "),t("code",[a._v(".so")]),a._v(" 插件，期间发生的异常会被忽略。")]),a._v(" "),t("p",[a._v("使用者可以通过指定环境变量 "),t("code",[a._v("GOZZ_PLUGINS_DIR")]),a._v(" 来变更默认的插件安装目录。")]),a._v(" "),t("h2",{attrs:{id:"指令"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#指令"}},[a._v("#")]),a._v(" 指令")]),a._v(" "),t("p",[t("code",[a._v("Gozz")]),a._v(" 支持以下指令：")]),a._v(" "),t("h3",{attrs:{id:"gozz-list"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#gozz-list"}},[a._v("#")]),a._v(" "),t("code",[a._v("gozz list")])]),a._v(" "),t("p",[a._v("该指令会列出已经被正确注册到内核和可使用的插件，并且输出插件和参数相关的介绍到控制台，使用者也可以通过该指令来检查插件是否被正确加载。")]),a._v(" "),t("h3",{attrs:{id:"gozz-run"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#gozz-run"}},[a._v("#")]),a._v(" "),t("code",[a._v("gozz run")])]),a._v(" "),t("p",[a._v("用法：")]),a._v(" "),t("div",{staticClass:"language-shell extra-class"},[t("pre",{pre:!0,attrs:{class:"language-shell"}},[t("code",[a._v("gozz run "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("--plugin/-p"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("filename"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" \n")])])]),t("p",[a._v("该指令会启动对 "),t("code",[a._v("filename")]),a._v(" 文件或目录注解的分析，\n并将分析的结构化注解及上下文，提交至指定的若干插件进行下一步工作。这将会是使用者最常用的指令。")]),a._v(" "),t("h4",{attrs:{id:"参数-filename"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#参数-filename"}},[a._v("#")]),a._v(" 参数 "),t("code",[a._v("filename")])]),a._v(" "),t("p",[a._v("当 "),t("code",[a._v("filename")]),a._v(" 为文件时，解析器只会解析当前单个文件内容")]),a._v(" "),t("p",[a._v("当 "),t("code",[a._v("filename")]),a._v(" 为目录时，解析器会遍历该目录以及嵌套子目录")]),a._v(" "),t("h4",{attrs:{id:"可选参数-plugin-p-name-options"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#可选参数-plugin-p-name-options"}},[a._v("#")]),a._v(" 可选参数 "),t("code",[a._v('--plugin / -p "name:options..."')])]),a._v(" "),t("p",[a._v("使用 "),t("code",[a._v("gozz run")]),a._v(" 必须指定 1 ~ N 个插件运行")]),a._v(" "),t("div",{staticClass:"language-shell extra-class"},[t("pre",{pre:!0,attrs:{class:"language-shell"}},[t("code",[t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("# 指定单个插件")]),a._v("\ngozz run "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("--plugin")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"foo"')]),a._v(" ./\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("# 简写参数")]),a._v("\ngozz run "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("-p")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"foo"')]),a._v(" ./\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("# 指定多个插件")]),a._v("\ngozz run "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("-p")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"foo"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("-p")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"bar"')]),a._v(" ./\n")])])]),t("p",[a._v("单次运行指定多个插件时，插件们会按参数顺序串行执行，尽管每个插件运行前都会重新进行文件解析，\n但基于文件元信息的版本解析缓存将会在进程内复用，因此同时指定多个插件会有更好体验。")]),a._v(" "),t("h4",{attrs:{id:"追加插件默认参数"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#追加插件默认参数"}},[a._v("#")]),a._v(" 追加插件默认参数")]),a._v(" "),t("p",[a._v("在指定插件名后方使用 "),t("code",[a._v(":")]),a._v(" 间隔，可以通过该参数追加插件默认参数：")]),a._v(" "),t("div",{staticClass:"language-shell extra-class"},[t("pre",{pre:!0,attrs:{class:"language-shell"}},[t("code",[t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("# 使用foo插件 并为所有匹配的注解 添加两个默认选项 [ key:value key2:value2 ]")]),a._v("\ngozz run "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("-p")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"foo:key=value:key2=value2"')]),a._v(" ./ \n")])])]),t("p",[a._v("注解 "),t("code",[a._v("+zz:foo:key=value3")]),a._v(" 添加默认选项 "),t("code",[a._v("key=value:key2=value2")]),a._v(" 后 等价于：")]),a._v(" "),t("p",[t("code",[a._v("+zz:foo:key=value3(已有值未被覆盖):key2=value2(缺省值使用默认)")])]),a._v(" "),t("h3",{attrs:{id:"wip-gozz-install"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#wip-gozz-install"}},[a._v("#")]),a._v(" [WIP] "),t("code",[a._v("gozz install")])]),a._v(" "),t("p",[a._v("用法")]),a._v(" "),t("div",{staticClass:"language-shell extra-class"},[t("pre",{pre:!0,attrs:{class:"language-shell"}},[t("code",[a._v("gozz "),t("span",{pre:!0,attrs:{class:"token function"}},[a._v("install")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("--output/-o"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("--filename/-f"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("repository"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" \n")])])]),t("p",[a._v("运行该指令会尝试将提供的 "),t("code",[a._v("repository")]),a._v(" 代码编译为 "),t("code",[a._v(".so")]),a._v(" 插件文件，并安装至用户目录 "),t("code",[a._v("~/.gozz/plugins")]),a._v("\n( 或 "),t("code",[a._v("GOZZ_PLUGINS_DIR")]),a._v(" 指定目录 ) 下。")]),a._v(" "),t("h4",{attrs:{id:"参数-repository"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#参数-repository"}},[a._v("#")]),a._v(" 参数 "),t("code",[a._v("repository")])]),a._v(" "),t("p",[a._v("当 "),t("code",[a._v("repository")]),a._v(" 带有网络协议前缀时，如 "),t("code",[a._v("ssh://、git://、http://、https://")]),a._v(" ，会使用 "),t("code",[a._v("git")]),a._v(" 进行仓库远程下载，\n否则会视为本地文件路径。")]),a._v(" "),t("h4",{attrs:{id:"可选参数-output-o-filename"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#可选参数-output-o-filename"}},[a._v("#")]),a._v(" 可选参数 "),t("code",[a._v('--output / -o "filename"')])]),a._v(" "),t("p",[a._v("指定该参数时，编译结果会输出为指定文件名，不再自动安装到用户目录。")]),a._v(" "),t("h4",{attrs:{id:"可选参数-filename-f-filename"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#可选参数-filename-f-filename"}},[a._v("#")]),a._v(" 可选参数 "),t("code",[a._v('--filename / -f "filename"')])]),a._v(" "),t("p",[a._v("指定该参数时，编译会使用该参数作为相对路径进行编译。")]),a._v(" "),t("p",[a._v("例：")]),a._v(" "),t("div",{staticClass:"language-shell extra-class"},[t("pre",{pre:!0,attrs:{class:"language-shell"}},[t("code",[a._v("gozz "),t("span",{pre:!0,attrs:{class:"token function"}},[a._v("install")]),a._v(" https://github.com/go-zing/gozz-plugins "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("-f")]),a._v(" ./contrib/sqlite "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("-o")]),a._v(" sqlite.so\n")])])]),t("p",[a._v("则会下载远程项目，在项目内进行插件编译")]),a._v(" "),t("div",{staticClass:"language- extra-class"},[t("pre",{pre:!0,attrs:{class:"language-text"}},[t("code",[a._v("go build --buildmode=plugin -o sqlite.so ./contrib/sqlite\n")])])]),t("h4",{attrs:{id:"使用该指令成功安装外部插件需要满足以下前提"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#使用该指令成功安装外部插件需要满足以下前提"}},[a._v("#")]),a._v(" 使用该指令成功安装外部插件需要满足以下前提：")]),a._v(" "),t("ul",[t("li",[a._v("编译当前 "),t("code",[a._v("gozz")]),a._v(" 使用的 Golang 版本和当前执行环境 Golang 版本一致。")]),a._v(" "),t("li",[a._v("指定 "),t("code",[a._v("repository")]),a._v(" 依赖兼容当前 "),t("code",[a._v("gozz")]),a._v(" 版本使用的 "),t("code",[a._v("gozz-core")]),a._v("。")]),a._v(" "),t("li",[a._v("其他影响 Golang 插件机制的环境因素：\n"),t("a",{attrs:{href:"https://tonybai.com/2021/07/19/understand-go-plugin/",target:"_blank",rel:"noopener noreferrer"}},[a._v("一文搞懂Go语言的plugin"),t("OutboundLink")],1),a._v("。")])]),a._v(" "),t("h2",{attrs:{id:"全局参数"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#全局参数"}},[a._v("#")]),a._v(" 全局参数")]),a._v(" "),t("h3",{attrs:{id:"x-extension-filename"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#x-extension-filename"}},[a._v("#")]),a._v(" "),t("code",[a._v("-x / --extension [filename]")])]),a._v(" "),t("p",[a._v("例：")]),a._v(" "),t("div",{staticClass:"language-shell extra-class"},[t("pre",{pre:!0,attrs:{class:"language-shell"}},[t("code",[t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("# 运行 list")]),a._v("\ngozz "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("-x")]),a._v(" plugin.so list\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("# 运行 run")]),a._v("\ngozz "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("-x")]),a._v(" plugin.so run "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[a._v("-p")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"plugin"')]),a._v(" ./\n")])])]),t("p",[a._v("使用该参数可以在工具启动时 额外加载指定文件路径的 "),t("code",[a._v(".so")]),a._v(" 插件，通过此方式加载的插件如果发生异常，会输出异常信息并立即退出进程。")]),a._v(" "),t("h2",{attrs:{id:"注解语法"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#注解语法"}},[a._v("#")]),a._v(" 注解语法")]),a._v(" "),t("p",[a._v("注解遵循以下语法：")]),a._v(" "),t("p",[a._v("以"),t("strong",[a._v("注释")]),a._v("形式紧贴"),t("strong",[a._v("代码声明对象")])]),a._v(" "),t("div",{staticClass:"language-go extra-class"},[t("pre",{pre:!0,attrs:{class:"language-go"}},[t("code",[t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:[PLUGIN]:[ARGS]:[OPTIONS]")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" T "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n")])])]),t("h3",{attrs:{id:"plugin-插件名"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#plugin-插件名"}},[a._v("#")]),a._v(" PLUGIN - 插件名")]),a._v(" "),t("p",[a._v("不同插件在注册到内核时会有唯一的插件名标识，这个标识也会用来为各个插件匹配注解，忽略大小写。")]),a._v(" "),t("p",[a._v("如：插件 "),t("code",[a._v("Foo")]),a._v(" 会使用 "),t("code",[a._v("+zz:foo")]),a._v(" 匹配注解")]),a._v(" "),t("h3",{attrs:{id:"args-必填参数"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#args-必填参数"}},[a._v("#")]),a._v(" ARGS - 必填参数")]),a._v(" "),t("p",[a._v("不同插件会指定固定数量的必填参数，参数会以 "),t("code",[a._v(":")]),a._v(" 间隔按插件指定顺序追加在注解后，\n如果该注解的固定参数数量不足，将会被忽略。")]),a._v(" "),t("p",[a._v("如：")]),a._v(" "),t("p",[a._v("插件 "),t("code",[a._v("foo")]),a._v(" 指定的参数数量为 2，符合要求注解为 "),t("code",[a._v("+zz:foo:arg1:arg2")])]),a._v(" "),t("p",[a._v("而 "),t("code",[a._v("+zz:foo:arg1")]),a._v(" 或 "),t("code",[a._v("+zz:foo")]),a._v(" 会被忽略")]),a._v(" "),t("h4",{attrs:{id:"为什么要忽略"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#为什么要忽略"}},[a._v("#")]),a._v(" 为什么要忽略")]),a._v(" "),t("p",[a._v("如果一个缺省参数是可以被推断默认值的，按"),t("RouterLink",{attrs:{to:"/zh/guide/#设计理念"}},[a._v("设计理念")]),a._v("，此参数不应被列为必填参数。")],1),a._v(" "),t("h3",{attrs:{id:"options-可填参数"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#options-可填参数"}},[a._v("#")]),a._v(" OPTIONS - 可填参数")]),a._v(" "),t("p",[a._v("超出插件指定必填参数后的内容，会被按 "),t("code",[a._v("Key = Value")]),a._v(" 组解析为可选参数提供给插件")]),a._v(" "),t("p",[a._v("如：")]),a._v(" "),t("p",[a._v("插件 "),t("code",[a._v("foo")]),a._v(" 指定的参数数量为 2")]),a._v(" "),t("p",[t("code",[a._v("+zz:foo:arg1:arg2:arg3=value:arg4")])]),a._v(" "),t("p",[a._v("将会被解析出 可选参数 "),t("code",[a._v('{"arg3":"value","arg4":""}')])]),a._v(" "),t("h4",{attrs:{id:"重复参数"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#重复参数"}},[a._v("#")]),a._v(" 重复参数")]),a._v(" "),t("p",[a._v("在可选参数 Value 中，如果存在数组类值，一般都会使用 "),t("code",[a._v(",")]),a._v(" 进行分隔，如：")]),a._v(" "),t("div",{staticClass:"language-go extra-class"},[t("pre",{pre:!0,attrs:{class:"language-go"}},[t("code",[t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:foo:set=a,b,c,d")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" T "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n")])])]),t("p",[a._v("解析器会将，相同 Key 的 Value 以这方式进行聚合 (包括空值)")]),a._v(" "),t("div",{staticClass:"language-go extra-class"},[t("pre",{pre:!0,attrs:{class:"language-go"}},[t("code",[t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:foo:set=a:set=:set=b:set=c,d")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" T "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n")])])]),t("p",[a._v("等价于")]),a._v(" "),t("div",{staticClass:"language-go extra-class"},[t("pre",{pre:!0,attrs:{class:"language-go"}},[t("code",[t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:foo:set=a,,b,c,d")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" T "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n")])])]),t("h4",{attrs:{id:"boolean-option"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#boolean-option"}},[a._v("#")]),a._v(" Boolean Option")]),a._v(" "),t("p",[a._v("部分可选参数，插件判断以某个Key是否存在，来判断特性是否开启。")]),a._v(" "),t("p",[a._v("此类参数如果 Value 不为空，会额外判断 Value 是否为 "),t("code",[a._v("0 / false / null")]),a._v(" 等否定含义。")]),a._v(" "),t("p",[a._v("即:")]),a._v(" "),t("ul",[t("li",[t("code",[a._v("+zz:bar:arg3=value:arg4")]),a._v(" -> "),t("code",[a._v('{"arg3":"value","arg4":""}')])]),a._v(" "),t("li",[t("code",[a._v("+zz:bar:arg3=value:arg4=true")]),a._v(" -> "),t("code",[a._v('{"arg3":"value","arg4":"true"}')])]),a._v(" "),t("li",[t("code",[a._v("+zz:bar:arg3=value:arg4=false")]),a._v(" -> "),t("code",[a._v('{"arg3":"value"}')])]),a._v(" "),t("li",[t("code",[a._v("+zz:bar:arg3=value:arg4=0")]),a._v(" -> "),t("code",[a._v('{"arg3":"value"}')])])]),a._v(" "),t("h4",{attrs:{id:"参数转义"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#参数转义"}},[a._v("#")]),a._v(" 参数转义")]),a._v(" "),t("p",[a._v("部分插件参数值可能会出现 "),t("code",[a._v(":")]),a._v(" 字符串的情况，可以使用 "),t("code",[a._v("\\")]),a._v(" 对 "),t("code",[a._v(":")]),a._v(" 进行转义，对命令行参数及注解内参数均生效。")]),a._v(" "),t("p",[a._v("即：")]),a._v(" "),t("p",[t("code",[a._v("+zz:plugin:addr=localhost:8080")]),a._v(" -> "),t("code",[a._v("+zz:plugin:addr=localhost\\:8080")])]),a._v(" "),t("p",[a._v("转义后，解析器在分隔注解参数时会先将 "),t("code",[a._v("\\:")]),a._v(" 替换为 "),t("code",[a._v("\\u003A")]),a._v("，\n在分隔后再将子串转义替换成 "),t("code",[a._v(":")]),a._v(" 。")]),a._v(" "),t("h3",{attrs:{id:"声明对象"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#声明对象"}},[a._v("#")]),a._v(" 声明对象")]),a._v(" "),t("p",[a._v("注解可以添加在 "),t("code",[a._v("Decl")]),a._v(" 代码块，也可以添加给指定的 "),t("code",[a._v("Spec")]),a._v(" 对象。")]),a._v(" "),t("p",[a._v("不理解 "),t("code",[a._v("Decl")]),a._v(" 和 "),t("code",[a._v("Spec")]),a._v(" 定义的可以先看下 "),t("a",{attrs:{href:"./how-it-works"}},[a._v("原理")]),a._v(" 部分。")]),a._v(" "),t("div",{staticClass:"language-go extra-class"},[t("pre",{pre:!0,attrs:{class:"language-go"}},[t("code",[t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("package")]),a._v(" t\n\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:foo")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("(")]),a._v("\n\tT0 "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n\tT1 "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n\n\t"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:bar")]),a._v("\n\tT2 "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v(")")]),a._v("\n\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:foo")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" T3 "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n")])])]),t("p",[a._v("在 "),t("code",[a._v("Decl")]),a._v(" 范围声明的注解，会被复制给 "),t("code",[a._v("Decl")]),a._v(" 内定义的所有 "),t("code",[a._v("Spec")]),a._v(" 对象内，即上例等价于：")]),a._v(" "),t("div",{staticClass:"language-go extra-class"},[t("pre",{pre:!0,attrs:{class:"language-go"}},[t("code",[t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("package")]),a._v(" t\n\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("(")]),a._v("\n\t"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:foo")]),a._v("\n\tT0 "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n\t"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:foo")]),a._v("\n\tT1 "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n\n\t"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:foo")]),a._v("\n\t"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:bar")]),a._v("\n\tT2 "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v(")")]),a._v("\n\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:foo")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" T3 "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n")])])]),t("h3",{attrs:{id:"多行注解"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#多行注解"}},[a._v("#")]),a._v(" 多行注解")]),a._v(" "),t("h4",{attrs:{id:"注解判定"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#注解判定"}},[a._v("#")]),a._v(" 注解判定")]),a._v(" "),t("p",[a._v("只要是和声明对象关联的注释，在每行行首 (忽略空白字符串) 存在注解前缀 "),t("code",[a._v("+zz:")]),a._v(" 都会被认为是注解行。")]),a._v(" "),t("div",{staticClass:"language-go extra-class"},[t("pre",{pre:!0,attrs:{class:"language-go"}},[t("code",[t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("/*\n+zz:foo  <- 注解行\n  +zz:foo  <- 注解行\n x +zoo:foo <- 非注解行\n//   +zoo:foo  <- 非注解行\n*/")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" T1 "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:foo  <- 注解行")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("//   +zoo:  <- 注解行")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("//   +zoo-  <- 非注解行")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// x +zoo:foo <- 非注解行")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" T2 "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n")])])]),t("h4",{attrs:{id:"重复注解"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#重复注解"}},[a._v("#")]),a._v(" 重复注解")]),a._v(" "),t("p",[a._v("一个代码声明对象上可能会有多个注解，且可能会是相同插件的重复注解，这些注解不会被认为是异常，而是会被重复提供给插件。")]),a._v(" "),t("p",[a._v("具体对于同对象重复注解的处理方式和结果，会基于不同插件实现逻辑而不同。")]),a._v(" "),t("div",{staticClass:"language-go extra-class"},[t("pre",{pre:!0,attrs:{class:"language-go"}},[t("code",[t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:foo")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:bar")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:bar:filename=./bar")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" T "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n")])])]),t("h2",{attrs:{id:"其他约定"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#其他约定"}},[a._v("#")]),a._v(" 其他约定")]),a._v(" "),t("h3",{attrs:{id:"生成文件路径"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#生成文件路径"}},[a._v("#")]),a._v(" 生成文件路径")]),a._v(" "),t("p",[a._v("一部分的内置插件会有生成目标文件的必填或可选参数，如：")]),a._v(" "),t("div",{staticClass:"language-go extra-class"},[t("pre",{pre:!0,attrs:{class:"language-go"}},[t("code",[t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:api:./")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:impl:./impl.go")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:wire:inject=/")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" API "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n")])])]),t("p",[a._v("这些目录文件目标路径会遵循以下规则：")]),a._v(" "),t("p",[a._v("假设项目根路径在 "),t("code",[a._v("/go/src/project/")])]),a._v(" "),t("div",{staticClass:"language- extra-class"},[t("pre",{pre:!0,attrs:{class:"language-text"}},[t("code",[a._v("/go/src/project/\n├── go.mod\n└── types\n    └── api.go\n")])])]),t("h4",{attrs:{id:"_1-如果包含-go-后缀-则以该文件名为生成目标文件名"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#_1-如果包含-go-后缀-则以该文件名为生成目标文件名"}},[a._v("#")]),a._v(" 1. 如果包含 "),t("code",[a._v(".go")]),a._v(" 后缀，则以该文件名为生成目标文件名")]),a._v(" "),t("div",{staticClass:"language-go extra-class"},[t("pre",{pre:!0,attrs:{class:"language-go"}},[t("code",[t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// /go/src/project/types/api.go")]),a._v("\n\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:api:./api.go")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" API "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n")])])]),t("p",[a._v("以上例子生成的文件名为 "),t("code",[a._v("api.go")]),a._v(" 具体目录路径会在后续规则说明")]),a._v(" "),t("h4",{attrs:{id:"_2-不包含-go-后缀-则会使用插件提供的默认文件名"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#_2-不包含-go-后缀-则会使用插件提供的默认文件名"}},[a._v("#")]),a._v(" 2. 不包含 "),t("code",[a._v(".go")]),a._v(" 后缀，则会使用插件提供的默认文件名")]),a._v(" "),t("p",[a._v("插件使用的默认文件名大部分是  "),t("code",[a._v("zzgen.${plugin}.go")])]),a._v(" "),t("div",{staticClass:"language-go extra-class"},[t("pre",{pre:!0,attrs:{class:"language-go"}},[t("code",[t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// /go/src/project/types/api.go")]),a._v("\n\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:api:./")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" API "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n")])])]),t("p",[a._v("以上例子生成的文件名为 "),t("code",[a._v("zzgen.api.go")])]),a._v(" "),t("h4",{attrs:{id:"_3-路径参数支持-golang-模版语法"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#_3-路径参数支持-golang-模版语法"}},[a._v("#")]),a._v(" 3. 路径参数支持 Golang 模版语法")]),a._v(" "),t("p",[a._v("支持的模版对象字段有 "),t("code",[a._v("Name")]),a._v(" "),t("code",[a._v("Package")])]),a._v(" "),t("p",[a._v("同时支持一系列字符串函数 "),t("code",[a._v("snake / camel ...")]),a._v(" "),t("a",{attrs:{href:"https://github.com/go-zing/gozz-core/blob/main/generate.go#L32",target:"_blank",rel:"noopener noreferrer"}},[a._v("详情见此"),t("OutboundLink")],1)]),a._v(" "),t("div",{staticClass:"language-go extra-class"},[t("pre",{pre:!0,attrs:{class:"language-go"}},[t("code",[t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// /go/src/project/types/api.go")]),a._v("\n\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:api:./{{ lower .Name }}.go")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" Foo "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n")])])]),t("p",[a._v("以上例子生成的文件名为 "),t("code",[a._v("foo.go")])]),a._v(" "),t("h4",{attrs:{id:"_4-如果文件名是相对路径-那么起点是-注解当前的文件所在目录"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#_4-如果文件名是相对路径-那么起点是-注解当前的文件所在目录"}},[a._v("#")]),a._v(" 4. 如果文件名是相对路径，那么起点是 注解当前的文件所在目录")]),a._v(" "),t("div",{staticClass:"language-go extra-class"},[t("pre",{pre:!0,attrs:{class:"language-go"}},[t("code",[t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// /go/src/project/types/api.go")]),a._v("\n\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:api:./apix.go")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" API "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n")])])]),t("p",[a._v("以上指定 "),t("code",[a._v(".go")]),a._v(" 文件名例子生成的文件完整路径为  "),t("code",[a._v("/go/src/project/types/apix.go")])]),a._v(" "),t("br"),a._v(" "),t("div",{staticClass:"language-go extra-class"},[t("pre",{pre:!0,attrs:{class:"language-go"}},[t("code",[t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// /go/src/project/types/api.go")]),a._v("\n\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:api:./apix")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" API "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n")])])]),t("p",[a._v("以上不指定 "),t("code",[a._v(".go")]),a._v(" 文件名例子生成的文件完整路径为  "),t("code",[a._v("/go/src/project/types/apix/zzgen.api.go")])]),a._v(" "),t("h4",{attrs:{id:"_5-如果文件名是绝对路径-那么起点是-注解文件所在-module-项目根目录-go-mod-所在目录"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#_5-如果文件名是绝对路径-那么起点是-注解文件所在-module-项目根目录-go-mod-所在目录"}},[a._v("#")]),a._v(" 5. 如果文件名是绝对路径，那么起点是 注解文件所在 "),t("code",[a._v("module")]),a._v(" 项目根目录 ( "),t("code",[a._v("go.mod")]),a._v(" 所在目录 )")]),a._v(" "),t("p",[a._v("该目录通过在注解文件所在目录执行 "),t("code",[a._v("go env GOMOD")]),a._v(" 获取")]),a._v(" "),t("div",{staticClass:"language-go extra-class"},[t("pre",{pre:!0,attrs:{class:"language-go"}},[t("code",[t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// /go/src/project/types/api.go")]),a._v("\n\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:api:/apix.go")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" API "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n")])])]),t("p",[a._v("以上指定 "),t("code",[a._v(".go")]),a._v(" 文件名例子生成的文件完整路径为  "),t("code",[a._v("/go/src/project/apix.go")])]),a._v(" "),t("br"),a._v(" "),t("div",{staticClass:"language-go extra-class"},[t("pre",{pre:!0,attrs:{class:"language-go"}},[t("code",[t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// /go/src/project/types/api.go")]),a._v("\n\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:api:/apix")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" API "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n")])])]),t("p",[a._v("以上不指定 "),t("code",[a._v(".go")]),a._v(" 文件名例子生成的文件完整路径为  "),t("code",[a._v("/go/src/project/apix/zzgen.api.go")])]),a._v(" "),t("br"),a._v(" "),t("p",[a._v("假设在 "),t("code",[a._v("types")]),a._v(" 目录内存在子 "),t("code",[a._v("module")])]),a._v(" "),t("div",{staticClass:"language- extra-class"},[t("pre",{pre:!0,attrs:{class:"language-text"}},[t("code",[a._v("/go/src/project/\n├── go.mod\n└── types\n    ├── api.go\n    └── go.mod\n")])])]),t("div",{staticClass:"language-go extra-class"},[t("pre",{pre:!0,attrs:{class:"language-go"}},[t("code",[t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// /go/src/project/types/api.go")]),a._v("\n\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("// +zz:api:/")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("type")]),a._v(" API "),t("span",{pre:!0,attrs:{class:"token keyword"}},[a._v("interface")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("{")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("}")]),a._v("\n")])])]),t("p",[a._v("以上例子生成的文件完整路径会变为  "),t("code",[a._v("/go/src/project/types/zzgen.api.go")])]),a._v(" "),t("h3",{attrs:{id:"生成模版"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#生成模版"}},[a._v("#")]),a._v(" 生成模版")]),a._v(" "),t("p",[a._v("对于覆盖式模版生成代码，生成代码文件前会检查该目录是否存在已有的 "),t("code",[a._v("${filename}.impl")]),a._v(" 文件。")]),a._v(" "),t("p",[a._v("若有，则会直接读取该模版文件作为生成模版。")]),a._v(" "),t("p",[a._v("否则，将会使用插件内建的默认模版文本，并输出为 "),t("code",[a._v("${filename}.impl")]),a._v(" 模版文件在生成代码同目录。")])])}),[],!1,null,null,null);t.default=n.exports}}]);
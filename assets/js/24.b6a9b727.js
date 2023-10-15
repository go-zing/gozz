(window.webpackJsonp=window.webpackJsonp||[]).push([[24],{307:function(_,v,e){"use strict";e.r(v);var o=e(14),t=Object(o.a)({},(function(){var _=this,v=_._self._c;return v("ContentSlotsDistributor",{attrs:{"slot-key":_.$parent.slotKey}},[v("h1",{attrs:{id:"介绍"}},[v("a",{staticClass:"header-anchor",attrs:{href:"#介绍"}},[_._v("#")]),_._v(" 介绍")]),_._v(" "),v("p",[v("code",[_._v("Gozz")]),_._v(" 由两部分组成：")]),_._v(" "),v("ul",[v("li",[v("p",[v("a",{attrs:{href:"https://github.com/go-zing/gozz-core",target:"_blank",rel:"noopener noreferrer"}},[_._v("gozz-core"),v("OutboundLink")],1),_._v(" 是核心依赖库，包含了代码文件解析，注解解析，结构化注解对象，\n运行时对象索引，缓存等核心功能依赖。同时为各个插件提供包含模版、生成、编辑、追加文件相关的工具依赖。")])]),_._v(" "),v("li",[v("p",[v("a",{attrs:{href:"https://github.com/go-zing/gozz",target:"_blank",rel:"noopener noreferrer"}},[_._v("gozz"),v("OutboundLink")],1),_._v(" 是用户使用该项目的 cli 工具运行入口，\n包含了对用户命令的解析，插件的加载和实例化，以及对外部扩展的加载。")])])]),_._v(" "),v("p",[_._v("笔者还提供了一系列强大的内置插件，囊括了在以往工作中开发这些插件雏形时，\n希望能够规范化解决，以及将这些优秀设计实践在团队快速落地的一些技术性需求，包括：")]),_._v(" "),v("ul",[v("li",[_._v("优化系统依赖架构设计提升项目质量和可维护性，提升多人大型项目人效协作效率")]),_._v(" "),v("li",[_._v("快速构建API和输出API文档、业务领域文档，降低API及文档持续迭代和多分发的成本")]),_._v(" "),v("li",[_._v("通过灵活切面化注入设计提升可观测性和降低变更复杂度，降低开发者工作量及扰动感知")])]),_._v(" "),v("h2",{attrs:{id:"理念"}},[v("a",{staticClass:"header-anchor",attrs:{href:"#理念"}},[_._v("#")]),_._v(" 理念")]),_._v(" "),v("p",[_._v("编程 就是将 "),v("strong",[_._v("重复、繁琐的事情")]),_._v(" ，将成熟的方案 通过程序化实现高效、低成本、规范化地解决。")]),_._v(" "),v("p",[_._v("而我们的设计和开发过程中本身有很多工作，本身也是 "),v("strong",[_._v("重复、繁琐的事情")]),_._v("。")]),_._v(" "),v("p",[_._v("在我们一遍又一遍地去浪费时间做这些 低效开发时 本身就违背了编程的初衷和思想。")]),_._v(" "),v("p",[_._v("因此 "),v("code",[_._v("Gozz")]),_._v(" 希望 将成熟的技术方法论 和 优秀的系统架构方案 高效、低成本、规范化地传播和落地。")]),_._v(" "),v("p",[_._v("并希望大家由此能：")]),_._v(" "),v("blockquote",[v("p",[_._v("Less, But Better")])]),_._v(" "),v("p",[_._v("而 "),v("code",[_._v("Gozz")]),_._v(" 实现这点的理念就是：")]),_._v(" "),v("blockquote",[v("p",[_._v("最好的代码，是不需要每个人都写的代码")])]),_._v(" "),v("p",[_._v("通过代码生成和编辑，提供最佳实践，或提供最佳实践的模板化入口。")]),_._v(" "),v("h3",{attrs:{id:"交互设计理念"}},[v("a",{staticClass:"header-anchor",attrs:{href:"#交互设计理念"}},[_._v("#")]),_._v(" 交互设计理念")]),_._v(" "),v("p",[v("code",[_._v("Gozz")]),_._v(" 遵循 "),v("a",{attrs:{href:"https://zh.wikipedia.org/wiki/%E7%BA%A6%E5%AE%9A%E4%BC%98%E4%BA%8E%E9%85%8D%E7%BD%AE",target:"_blank",rel:"noopener noreferrer"}},[_._v("约定优于配置"),v("OutboundLink")],1)]),_._v(" "),v("p",[_._v("会尽可能使用 "),v("strong",[_._v("简洁")]),_._v(" 和 "),v("strong",[_._v("符合人类直觉")]),_._v(" 的 命令 / 注解 / 参数")]),_._v(" "),v("h2",{attrs:{id:"适用场景"}},[v("a",{staticClass:"header-anchor",attrs:{href:"#适用场景"}},[_._v("#")]),_._v(" 适用场景")]),_._v(" "),v("h3",{attrs:{id:"个人"}},[v("a",{staticClass:"header-anchor",attrs:{href:"#个人"}},[_._v("#")]),_._v(" 个人")]),_._v(" "),v("p",[v("code",[_._v("Gozz")]),_._v(" 不会在 Feature 中强调对初阶 Gopher 的友好性，因为 "),v("strong",[v("code",[_._v("注释注解")]),_._v("是非官方的特性")]),_._v(" ，我们不希望引入额外认知成本 从而影响\nGopher 对 Golang 本身的学习曲线。")]),_._v(" "),v("p",[_._v("一方面 "),v("code",[_._v("Gozz")]),_._v(" 部分插件的生成结果的使用 需要开发者对 Golang 类型和 "),v("code",[_._v("interface")]),_._v(" 系统 甚至 "),v("code",[_._v("reflect")]),_._v(" 有一定深入认知。")]),_._v(" "),v("p",[_._v("另一方面 对于 "),v("code",[_._v("Gozz 解决的到底是什么问题")]),_._v(" 也需要对 "),v("code",[_._v("团队协作 / 系统架构 / 设计模式")]),_._v(" 有一定的前置认知。")]),_._v(" "),v("p",[_._v("如果你有使用 "),v("code",[_._v("JAVA Spring")]),_._v(" 的经验，那么恭喜，你可以有一定的认知优势。")]),_._v(" "),v("p",[v("code",[_._v("Gozz")]),_._v(" 解决需求的思路学习了 "),v("code",[_._v("JAVA Spring")]),_._v(" 一些重要设计思想，"),v("strong",[_._v("但绝不是拙劣地模仿")]),_._v("。")]),_._v(" "),v("h3",{attrs:{id:"团队"}},[v("a",{staticClass:"header-anchor",attrs:{href:"#团队"}},[_._v("#")]),_._v(" 团队")]),_._v(" "),v("p",[v("code",[_._v("人月神话")]),_._v(" 中有 "),v("code",[_._v("保有概念整体性")]),_._v(" 的说法 这个系统设计理念也是笔者比较认同 即：")]),_._v(" "),v("p",[v("strong",[_._v("系统的核心架构设计 需要由少数人专制控制")])]),_._v(" "),v("p",[_._v("因此，使用 "),v("code",[_._v("Gozz")]),_._v(" 的时候，团队需要由一个核心的资深角色给出适合团队的 "),v("code",[_._v("Gozz")]),_._v(" 配置 以及 维护生成模板和适配层。")]),_._v(" "),v("p",[_._v("在团队内，即使是不同的业务项目，在 "),v("code",[_._v("Gozz")]),_._v(" 的使用都不应该产生过多分歧。")]),_._v(" "),v("p",[_._v("而其他成员只需要遵循团队的规范，将运行 "),v("code",[_._v("Gozz")]),_._v(" 的指令写到 项目 "),v("code",[_._v("Makefile")]),_._v(" 或构建工具，在 代码变更 及 提交之前 去执行。")]),_._v(" "),v("hr"),_._v(" "),v("p",[_._v("在确保企业生产稳定性后，"),v("code",[_._v("Gozz")]),_._v(" 鼓励所有团队角色去主动探索和学习，如何去优化业务和项目的架构设，用更低的成本办更多更好的事。")]),_._v(" "),v("p",[v("code",[_._v("Gozz")]),_._v(" 也将会在项目示例中提供一些最佳实践供开发者们参考。")]),_._v(" "),v("h3",{attrs:{id:"微服务-or-单体"}},[v("a",{staticClass:"header-anchor",attrs:{href:"#微服务-or-单体"}},[_._v("#")]),_._v(" 微服务 or 单体")]),_._v(" "),v("p",[v("code",[_._v("Gozz")]),_._v(" 的应用并不局限于微服务场景，相反，越大型的项目和协作团队，相信会越容易从 "),v("code",[_._v("Gozz")]),_._v(" 中受益更多。")]),_._v(" "),v("p",[v("a",{attrs:{href:"../past-and-present"}},[_._v("前世今生")]),_._v(" 中有提到：一些内置插件的前身，就是为了优化数十万行级系统代码重构需求而生。")]),_._v(" "),v("p",[_._v("微服务，只代表个别服务在业务依赖领域设计中的 "),v("code",[_._v("Micro / Pluggable / Extensible")]),_._v(" ，但从不意味着该项目代码架构层级和设计简单，\n或该服务维系人员较少。")]),_._v(" "),v("p",[_._v("如果你希望团队的微服务项目代码质量和协作效率能够得到一定的提升，也欢迎使用 "),v("code",[_._v("Gozz")]),_._v(" 和我们提供的各个内置插件功能。")])])}),[],!1,null,null,null);v.default=t.exports}}]);
(window.webpackJsonp=window.webpackJsonp||[]).push([[25],{309:function(e,t,o){"use strict";o.r(t);var n=o(10),a=Object(n.a)({},(function(){var e=this,t=e._self._c;return t("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[t("h1",{attrs:{id:"introduction"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#introduction"}},[e._v("#")]),e._v(" Introduction")]),e._v(" "),t("p",[t("code",[e._v("Gozz")]),e._v(" was made from two parts：")]),e._v(" "),t("ul",[t("li",[t("a",{attrs:{href:"https://github.com/go-zing/gozz",target:"_blank",rel:"noopener noreferrer"}},[e._v("gozz"),t("OutboundLink")],1),e._v(" is CLI (Command-Line Interface) for the project.\nIt manages the runtime of process,\nparses commands and invokes plugins or extension plugins.")]),e._v(" "),t("li",[t("a",{attrs:{href:"https://github.com/go-zing/gozz-core",target:"_blank",rel:"noopener noreferrer"}},[e._v("gozz-core"),t("OutboundLink")],1),e._v(" provides core library,\ncontains code and annotations parsing, object types,\ncaches and so one core functional dependencies.\nIt also provides tools on handing codes and files for plugins.")])]),e._v(" "),t("p",[t("code",[e._v("Gozz")]),e._v(" provides a series of "),t("a",{attrs:{href:"plugins"}},[e._v("built-in plugins")]),e._v(",\ndedicated to providing tool-based solutions for some of the technical needs,\nthat have been summarized and accumulated in the past.")]),e._v(" "),t("p",[t("a",{attrs:{href:"https://github.com/go-zing/gozz-doc-examples",target:"_blank",rel:"noopener noreferrer"}},[e._v("gozz-doc-examples"),t("OutboundLink")],1),e._v(" contains examples for all builtin plugins,\n"),t("a",{attrs:{href:"plugins"}},[e._v("Plugins")]),e._v(" would provide comparative reading.")]),e._v(" "),t("p",[e._v("Also, you can explore more external plugins in "),t("a",{attrs:{href:"https://github.com/go-zing/gozz-plugins",target:"_blank",rel:"noopener noreferrer"}},[e._v("gozz-plugins"),t("OutboundLink")],1),e._v(".\nThese externals were based on "),t("a",{attrs:{href:"https://pkg.go.dev/plugin",target:"_blank",rel:"noopener noreferrer"}},[e._v("Golang plugin"),t("OutboundLink")],1),e._v(",\nand we provide command tool to "),t("RouterLink",{attrs:{to:"/guide/getting-started.html#gozz-install"}},[e._v("install plugin")]),e._v(" automatically.\nSo developers could do expand development easily.")],1),e._v(" "),t("div",{staticClass:"custom-block tip"},[t("p",{staticClass:"custom-block-title"},[e._v("Style")]),e._v(" "),t("p",[t("code",[e._v("Gozz")]),e._v(" emphasizes progressive idempotent equation generation,\nand provide stable automated code iteration integration through consistent commands and configurations to fulfill\nrequirements continuous changing.")])]),e._v(" "),t("h2",{attrs:{id:"vision"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#vision"}},[e._v("#")]),e._v(" Vision")]),e._v(" "),t("p",[e._v("Programming is to take "),t("strong",[e._v("repetitive and tedious things")]),e._v(" and program mature solutions to achieve efficient, low-cost,\nand standardized solutions.")]),e._v(" "),t("p",[e._v("There is a lot of work in our design and development process, which is also "),t("strong",[e._v("repetitive and tedious")]),e._v(".")]),e._v(" "),t("p",[e._v("When we waste time doing these inefficient development over and over again, it itself goes against the original\nintention and idea of programming.")]),e._v(" "),t("p",[e._v("Through tooling, "),t("code",[e._v("Gozz")]),e._v(" hopes to help spread and implement "),t("strong",[e._v("mature technology methodologies")]),e._v(" and "),t("strong",[e._v("excellent practice\nplans")]),e._v(":")]),e._v(" "),t("p",[e._v("Provide best practices through code generation and editing, or provide access to templated best practices.")]),e._v(" "),t("blockquote",[t("p",[e._v("The best code is code that nobody repeat")])]),e._v(" "),t("p",[e._v("And hope that we can:")]),e._v(" "),t("blockquote",[t("p",[e._v("Less, But Better")])]),e._v(" "),t("h2",{attrs:{id:"design-concept"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#design-concept"}},[e._v("#")]),e._v(" Design Concept")]),e._v(" "),t("p",[t("code",[e._v("Gozz")]),e._v(" follows "),t("a",{attrs:{href:"https://en.wikipedia.org/wiki/Convention_over_configuration",target:"_blank",rel:"noopener noreferrer"}},[e._v("Convention Over Configuration"),t("OutboundLink")],1)]),e._v(" "),t("p",[e._v("And prefer commands/annotations/parameters that are "),t("strong",[e._v("concise")]),e._v(" and\n"),t("strong",[e._v("in line with human intuition")]),e._v(" as much as possible")]),e._v(" "),t("h2",{attrs:{id:"applicable-scene"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#applicable-scene"}},[e._v("#")]),e._v(" Applicable scene")]),e._v(" "),t("h3",{attrs:{id:"personal"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#personal"}},[e._v("#")]),e._v(" personal")]),e._v(" "),t("p",[e._v("First of all, before using "),t("code",[e._v("Gozz")]),e._v(", I hope you can clearly understand: "),t("strong",[e._v("Annotations are not an official Golang feature")])]),e._v(" "),t("p",[e._v("For developers who don't know this, I would recommend learning Golang through other tutorials first.")]),e._v(" "),t("br"),e._v(" "),t("p",[e._v("Then, using "),t("code",[e._v("Gozz")]),e._v(" well requires a certain amount of pre-knowledge:")]),e._v(" "),t("ul",[t("li",[e._v("To understand the use of the code generated by the "),t("code",[e._v("Gozz")]),e._v(" plugins, you need to have a certain understanding of Golang\n"),t("code",[e._v("type / interface / reflect")]),e._v(".")]),e._v(" "),t("li",[e._v("Understanding "),t("code",[e._v("what problem Gozz solves")]),e._v(" requires a certain understanding\nof "),t("code",[e._v("team collaboration / system architecture / design patterns")]),e._v(".")])]),e._v(" "),t("br"),e._v(" "),t("div",{staticClass:"custom-block tip"},[t("p",{staticClass:"custom-block-title"},[e._v("TIP")]),e._v(" "),t("p",[e._v("If you have experience using "),t("code",[e._v("JAVA Spring")]),e._v(", you can have certain cognitive advantages.")]),e._v(" "),t("p",[t("code",[e._v("Gozz")]),e._v(" learned some important design ideas of "),t("code",[e._v("JAVA Spring")]),e._v(", but "),t("strong",[e._v("it is by no means a poor imitation.")])])]),e._v(" "),t("h3",{attrs:{id:"team"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#team"}},[e._v("#")]),e._v(" Team")]),e._v(" "),t("p",[e._v("There is a point in "),t("a",{attrs:{href:"https://en.wikipedia.org/wiki/The_Mythical_Man-Month",target:"_blank",rel:"noopener noreferrer"}},[e._v("The mythical man-month"),t("OutboundLink")],1),e._v("\nabout "),t("code",[e._v("conceptual integrity")]),e._v(".\nThis system design concept is also one that the author agrees with:")]),e._v(" "),t("p",[t("strong",[e._v("The core architecture design of the system needs to be autocratically controlled by a few people.")])]),e._v(" "),t("p",[e._v("Therefore, the team needs a core senior role to provide the "),t("code",[e._v("Gozz")]),e._v(" configuration suitable for the team,\nand maintain the generated templates and adaptation layers.")]),e._v(" "),t("p",[e._v("Within the team, how to configure "),t("code",[e._v("Gozz")]),e._v(" and its templates should not cause too much disagreement.")]),e._v(" "),t("p",[e._v("Other members should follow the team's specifications to run "),t("code",[e._v("Gozz")]),e._v(" within "),t("code",[e._v("Makefile")]),e._v(" or DevOps pipeline,\nand execute them before commit or build.")]),e._v(" "),t("br"),e._v(" "),t("p",[e._v("Ensuring production stability,\n"),t("code",[e._v("Gozz")]),e._v(" still encourages every team roles to actively explore and learn how to optimize the\narchitectural design of business and projects, and do more and better things at lower costs.")]),e._v(" "),t("p",[e._v("We would also provide some best practices in project examples for developers to refer to.")]),e._v(" "),t("h3",{attrs:{id:"micro-or-monolith"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#micro-or-monolith"}},[e._v("#")]),e._v(" Micro or Monolith")]),e._v(" "),t("p",[t("code",[e._v("Gozz")]),e._v(" is never limited to micro. On the contrary, the larger the projects and collaborative teams,\nthe better it will be to benefit more from "),t("code",[e._v("Gozz")]),e._v(".")]),e._v(" "),t("p",[t("RouterLink",{attrs:{to:"/story.html"}},[e._v("Story Parts")]),e._v(" would introduce:\nThe predecessors of some built-in plugins were created to optimize the\nreconstruction needs of some code system projects with lines and lines of codes.")],1),e._v(" "),t("p",[e._v("Microservices only represent the "),t("code",[e._v("Micro / Pluggable / Extensible")]),e._v("\nof individual services in the design of business dependency areas,\nbut it never means that the project code architecture and design may be simple.\nOr the service has fewer maintenance staff.")]),e._v(" "),t("p",[e._v("If you want to improve the code quality and collaboration efficiency of your team's projects,\nyou are also welcome to try "),t("code",[e._v("Gozz")]),e._v(" and the awesome built-in plugins it provides.")])])}),[],!1,null,null,null);t.default=a.exports}}]);
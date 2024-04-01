"use strict";(self.webpackChunkrenovate_issue_action=self.webpackChunkrenovate_issue_action||[]).push([[358],{5680:(e,n,a)=>{a.d(n,{xA:()=>c,yg:()=>f});var t=a(6540);function i(e,n,a){return n in e?Object.defineProperty(e,n,{value:a,enumerable:!0,configurable:!0,writable:!0}):e[n]=a,e}function o(e,n){var a=Object.keys(e);if(Object.getOwnPropertySymbols){var t=Object.getOwnPropertySymbols(e);n&&(t=t.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),a.push.apply(a,t)}return a}function r(e){for(var n=1;n<arguments.length;n++){var a=null!=arguments[n]?arguments[n]:{};n%2?o(Object(a),!0).forEach((function(n){i(e,n,a[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(a)):o(Object(a)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(a,n))}))}return e}function l(e,n){if(null==e)return{};var a,t,i=function(e,n){if(null==e)return{};var a,t,i={},o=Object.keys(e);for(t=0;t<o.length;t++)a=o[t],n.indexOf(a)>=0||(i[a]=e[a]);return i}(e,n);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(t=0;t<o.length;t++)a=o[t],n.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(e,a)&&(i[a]=e[a])}return i}var s=t.createContext({}),u=function(e){var n=t.useContext(s),a=n;return e&&(a="function"==typeof e?e(n):r(r({},n),e)),a},c=function(e){var n=u(e.components);return t.createElement(s.Provider,{value:n},e.children)},p="mdxType",g={inlineCode:"code",wrapper:function(e){var n=e.children;return t.createElement(t.Fragment,{},n)}},d=t.forwardRef((function(e,n){var a=e.components,i=e.mdxType,o=e.originalType,s=e.parentName,c=l(e,["components","mdxType","originalType","parentName"]),p=u(a),d=i,f=p["".concat(s,".").concat(d)]||p[d]||g[d]||o;return a?t.createElement(f,r(r({ref:n},c),{},{components:a})):t.createElement(f,r({ref:n},c))}));function f(e,n){var a=arguments,i=n&&n.mdxType;if("string"==typeof e||i){var o=a.length,r=new Array(o);r[0]=d;var l={};for(var s in n)hasOwnProperty.call(n,s)&&(l[s]=n[s]);l.originalType=e,l[p]="string"==typeof e?e:i,r[1]=l;for(var u=2;u<o;u++)r[u]=a[u];return t.createElement.apply(null,r)}return t.createElement.apply(null,a)}d.displayName="MDXCreateElement"},6750:(e,n,a)=>{a.r(n),a.d(n,{assets:()=>s,contentTitle:()=>r,default:()=>g,frontMatter:()=>o,metadata:()=>l,toc:()=>u});var t=a(8168),i=(a(6540),a(5680));const o={sidebar_position:100},r="Configuration file",l={unversionedId:"config/config-file",id:"config/config-file",title:"Configuration file",description:"Configuration file path",source:"@site/docs/config/config-file.md",sourceDirName:"config",slug:"/config/config-file",permalink:"/renovate-issue-action/config/config-file",draft:!1,editUrl:"https://github.com/suzuki-shunsuke/renovate-issue-action-docs/edit/main/docs/config/config-file.md",tags:[],version:"current",sidebarPosition:100,frontMatter:{sidebar_position:100},sidebar:"tutorialSidebar",previous:{title:"Configuration",permalink:"/renovate-issue-action/config/"},next:{title:"Template",permalink:"/renovate-issue-action/config/template"}},s={},u=[{value:"Configuration file path",id:"configuration-file-path",level:2},{value:"Overview",id:"overview",level:2},{value:"Configuration Overlays",id:"configuration-overlays",level:2},{value:"Example",id:"example",level:2}],c={toc:u},p="wrapper";function g(e){let{components:n,...a}=e;return(0,i.yg)(p,(0,t.A)({},c,a,{components:n,mdxType:"MDXLayout"}),(0,i.yg)("h1",{id:"configuration-file"},"Configuration file"),(0,i.yg)("h2",{id:"configuration-file-path"},"Configuration file path"),(0,i.yg)("p",null,"Configuration file is optional.\nBy default, the file ",(0,i.yg)("inlineCode",{parentName:"p"},"^\\.renovate-issue-action\\.ya?ml$")," in the current directory is read if it exists.\nYou can specify configuration file path with the command line option ",(0,i.yg)("inlineCode",{parentName:"p"},"--config, -c"),"."),(0,i.yg)("pre",null,(0,i.yg)("code",{parentName:"pre",className:"language-console"},"$ renovate-issue-action --config config.yaml\n")),(0,i.yg)("h2",{id:"overview"},"Overview"),(0,i.yg)("p",null,"You can configure the following fields of issues created by renovate-issue-action."),(0,i.yg)("ul",null,(0,i.yg)("li",{parentName:"ul"},"Repository where issue is created"),(0,i.yg)("li",{parentName:"ul"},"Issue title, description, labels, assignees, and so on")),(0,i.yg)("p",null,"You can change the above setting conditionally per Pull Request according to the following Pull Request metadata."),(0,i.yg)("ul",null,(0,i.yg)("li",{parentName:"ul"},"Pull Request Labels"),(0,i.yg)("li",{parentName:"ul"},"Renovate metadata such as ",(0,i.yg)("inlineCode",{parentName:"li"},"packageName"),", ",(0,i.yg)("inlineCode",{parentName:"li"},"depName"),", ",(0,i.yg)("inlineCode",{parentName:"li"},"groupName"),", ",(0,i.yg)("inlineCode",{parentName:"li"},"packageFileDir"))),(0,i.yg)("p",null,"You can also ignore Pull Requests conditionally.\nFor example, you can ignore Pull Requests regarding to specific packages."),(0,i.yg)("h2",{id:"configuration-overlays"},"Configuration Overlays"),(0,i.yg)("p",null,"There are some configuration layers."),(0,i.yg)("ol",null,(0,i.yg)("li",{parentName:"ol"},"conditional configuration in ",(0,i.yg)("inlineCode",{parentName:"li"},"entries"))),(0,i.yg)("p",null,"e.g."),(0,i.yg)("pre",null,(0,i.yg)("code",{parentName:"pre",className:"language-yaml"},'entries:\n- if: Metadata.PackageName == "actions/checkout"\n  issue:\n    labels: ["checkout"]\n')),(0,i.yg)("ol",{start:2},(0,i.yg)("li",{parentName:"ol"},"global configuration")),(0,i.yg)("p",null,"e.g."),(0,i.yg)("pre",null,(0,i.yg)("code",{parentName:"pre",className:"language-yaml"},'issue:\n  assignees: ["suzuki-shunsuke"]\n')),(0,i.yg)("ol",{start:3},(0,i.yg)("li",{parentName:"ol"},"Built-in default Configuration (Configuration file is optional)")),(0,i.yg)("p",null,"Each configuration are merged and duplicated configuration fields are overridden.\nElements of ",(0,i.yg)("inlineCode",{parentName:"p"},"entries")," are evaluated one by one from the first element.\nEach element of ",(0,i.yg)("inlineCode",{parentName:"p"},"entries")," has ",(0,i.yg)("inlineCode",{parentName:"p"},"if")," field, and the evaluation result of ",(0,i.yg)("inlineCode",{parentName:"p"},"if")," field must be ",(0,i.yg)("inlineCode",{parentName:"p"},"true")," or ",(0,i.yg)("inlineCode",{parentName:"p"},"false"),".\nIf the evaluation result of ",(0,i.yg)("inlineCode",{parentName:"p"},"if")," field is ",(0,i.yg)("inlineCode",{parentName:"p"},"false"),", the entry is ignored and the next entry is evaluated.\nIf the evaluation result of ",(0,i.yg)("inlineCode",{parentName:"p"},"if")," field is ",(0,i.yg)("inlineCode",{parentName:"p"},"true"),", the entry is selected and subsequent entries are ignored.\nIf no element is selected, ",(0,i.yg)("inlineCode",{parentName:"p"},"entries")," is ignored."),(0,i.yg)("h2",{id:"example"},"Example"),(0,i.yg)("p",null,"e.g. renovate-issue-action.yaml"),(0,i.yg)("pre",null,(0,i.yg)("code",{parentName:"pre",className:"language-yaml"},'projects: # By default, empty\n- alias: sre\n  column_id: PC_*** # GitHub Project Column ID\n- alias: service-a\n  next_id: PN_*** # Project Next Item ID\nrenovate_login: \'renovate[bot]\' # By default, \'renovate[bot]\'. If you use Self-hosted Renovate, you have to set this field.\nissue:\n  repo_owner: suzuki-shunsuke # By default, the repository which GitHub Actions is run\n  repo_name: renovate-issue-action # By default, the repository which GitHub Actions is run\n  title: \'Renovate Automerge Failure({{.RepoOwner}}/{{.RepoName}}): {{if .Metadata.GroupName}}{{.Metadata.GroupName}}{{else}}{{.Metadata.PackageName}}{{.Metadata.DepName}}{{end}} {{if .Metadata.PackageFileDir}}({{.Metadata.PackageFileDir}}){{end}}\'\n  body: |\n    _This pull request was created by [renovate-issue-action](https://github.com/suzuki-shunsuke/renovate-issue-action)._\n\n    :warning: Please don\'t edit the Issue title, because renovate-issue-action searches issue with Issue title.\n\n    {{if .Metadata.PackageName}}packageName: {{.Metadata.PackageName}}{{end}}\n    {{if .Metadata.GroupName}}groupName: {{.Metadata.GroupName}}{{end}}\n    {{if .Metadata.DepName}}depName: {{.Metadata.DepName}}{{end}}\n  assignees: ["suzuki-shunsuke"] # By default, null\n  labels: ["renovate-issue-action"] # By default, null\nentries:\n- if: Metadata.PackageName == "actions/checkout"\n  issue:\n    labels: []\n    assignees: []\n    additional_body: |\n      ## What to do\n\n      * Check closed Pull Requests\n      * Fix the problem and update the package\n- if: Metadata.PackageName == "actions/cache"\n  ignore: true # If `ignore` is true, do nothing.\n- if: |\n    any(Labels, {# in Vars.labels})\n  vars:\n  - name: labels\n    value: ["sre"]\n  issue:\n    repo_name: sre-issues\n    addtional_labels: ["sre"]\n    addtional_assignees: ["octocat"]\n    additional_body: "@suzuki-shunsuke"\n    projects:\n    - sre\n    - service-a\n')))}g.isMDXComponent=!0}}]);
"use strict";(self.webpackChunkrenovate_issue_action=self.webpackChunkrenovate_issue_action||[]).push([[285],{3905:(e,t,r)=>{r.d(t,{Zo:()=>s,kt:()=>d});var n=r(7294);function a(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function i(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function o(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?i(Object(r),!0).forEach((function(t){a(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function l(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},i=Object.keys(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var p=n.createContext({}),c=function(e){var t=n.useContext(p),r=t;return e&&(r="function"==typeof e?e(t):o(o({},t),e)),r},s=function(e){var t=c(e.components);return n.createElement(p.Provider,{value:t},e.children)},u="mdxType",m={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},f=n.forwardRef((function(e,t){var r=e.components,a=e.mdxType,i=e.originalType,p=e.parentName,s=l(e,["components","mdxType","originalType","parentName"]),u=c(r),f=a,d=u["".concat(p,".").concat(f)]||u[f]||m[f]||i;return r?n.createElement(d,o(o({ref:t},s),{},{components:r})):n.createElement(d,o({ref:t},s))}));function d(e,t){var r=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=r.length,o=new Array(i);o[0]=f;var l={};for(var p in t)hasOwnProperty.call(t,p)&&(l[p]=t[p]);l.originalType=e,l[u]="string"==typeof e?e:a,o[1]=l;for(var c=2;c<i;c++)o[c]=r[c];return n.createElement.apply(null,o)}return n.createElement.apply(null,r)}f.displayName="MDXCreateElement"},3903:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>p,contentTitle:()=>o,default:()=>m,frontMatter:()=>i,metadata:()=>l,toc:()=>c});var n=r(7462),a=(r(7294),r(3905));const i={sidebar_position:200},o="Template",l={unversionedId:"config/template",id:"config/template",title:"Template",description:"Some configuration fields are parsed with Go's text/template.",source:"@site/docs/config/template.md",sourceDirName:"config",slug:"/config/template",permalink:"/renovate-issue-action/config/template",draft:!1,editUrl:"https://github.com/suzuki-shunsuke/renovate-issue-action-docs/edit/main/docs/config/template.md",tags:[],version:"current",sidebarPosition:200,frontMatter:{sidebar_position:200},sidebar:"tutorialSidebar",previous:{title:"Configuration file",permalink:"/renovate-issue-action/config/config-file"},next:{title:"Expression",permalink:"/renovate-issue-action/config/expression"}},p={},c=[{value:"Template Variables",id:"template-variables",level:2}],s={toc:c},u="wrapper";function m(e){let{components:t,...r}=e;return(0,a.kt)(u,(0,n.Z)({},s,r,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"template"},"Template"),(0,a.kt)("p",null,"Some configuration fields are parsed with Go's ",(0,a.kt)("a",{parentName:"p",href:"https://pkg.go.dev/text/template"},"text/template"),".\nIn the template, ",(0,a.kt)("a",{parentName:"p",href:"http://masterminds.github.io/sprig/"},"sprig Function")," can be used."),(0,a.kt)("h2",{id:"template-variables"},"Template Variables"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"RepoOwner"),(0,a.kt)("li",{parentName:"ul"},"RepoName"),(0,a.kt)("li",{parentName:"ul"},"Metadata",(0,a.kt)("ul",{parentName:"li"},(0,a.kt)("li",{parentName:"ul"},"GroupName"),(0,a.kt)("li",{parentName:"ul"},"PackageName"),(0,a.kt)("li",{parentName:"ul"},"DepName"),(0,a.kt)("li",{parentName:"ul"},"PackageFileDir"),(0,a.kt)("li",{parentName:"ul"},"UpdateType")))))}m.isMDXComponent=!0}}]);
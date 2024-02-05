"use strict";(self.webpackChunkrenovate_issue_action=self.webpackChunkrenovate_issue_action||[]).push([[488],{5788:(e,n,t)=>{t.d(n,{Iu:()=>c,yg:()=>f});var r=t(1504);function a(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function i(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,r)}return t}function o(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?i(Object(t),!0).forEach((function(n){a(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):i(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function s(e,n){if(null==e)return{};var t,r,a=function(e,n){if(null==e)return{};var t,r,a={},i=Object.keys(e);for(r=0;r<i.length;r++)t=i[r],n.indexOf(t)>=0||(a[t]=e[t]);return a}(e,n);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)t=i[r],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(a[t]=e[t])}return a}var p=r.createContext({}),l=function(e){var n=r.useContext(p),t=n;return e&&(t="function"==typeof e?e(n):o(o({},n),e)),t},c=function(e){var n=l(e.components);return r.createElement(p.Provider,{value:n},e.children)},u="mdxType",m={inlineCode:"code",wrapper:function(e){var n=e.children;return r.createElement(r.Fragment,{},n)}},g=r.forwardRef((function(e,n){var t=e.components,a=e.mdxType,i=e.originalType,p=e.parentName,c=s(e,["components","mdxType","originalType","parentName"]),u=l(t),g=a,f=u["".concat(p,".").concat(g)]||u[g]||m[g]||i;return t?r.createElement(f,o(o({ref:n},c),{},{components:t})):r.createElement(f,o({ref:n},c))}));function f(e,n){var t=arguments,a=n&&n.mdxType;if("string"==typeof e||a){var i=t.length,o=new Array(i);o[0]=g;var s={};for(var p in n)hasOwnProperty.call(n,p)&&(s[p]=n[p]);s.originalType=e,s[u]="string"==typeof e?e:a,o[1]=s;for(var l=2;l<i;l++)o[l]=t[l];return r.createElement.apply(null,o)}return r.createElement.apply(null,t)}g.displayName="MDXCreateElement"},2928:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>p,contentTitle:()=>o,default:()=>m,frontMatter:()=>i,metadata:()=>s,toc:()=>l});var r=t(5072),a=(t(1504),t(5788));const i={sidebar_position:300},o="Expression",s={unversionedId:"config/expression",id:"config/expression",title:"Expression",description:"antonmedv/expr is used.",source:"@site/docs/config/expression.md",sourceDirName:"config",slug:"/config/expression",permalink:"/renovate-issue-action/config/expression",draft:!1,editUrl:"https://github.com/suzuki-shunsuke/renovate-issue-action-docs/edit/main/docs/config/expression.md",tags:[],version:"current",sidebarPosition:300,frontMatter:{sidebar_position:300},sidebar:"tutorialSidebar",previous:{title:"Template",permalink:"/renovate-issue-action/config/template"},next:{title:"Environment variable",permalink:"/renovate-issue-action/config/environment-variable"}},p={},l=[{value:"Expression Variables",id:"expression-variables",level:3}],c={toc:l},u="wrapper";function m(e){let{components:n,...t}=e;return(0,a.yg)(u,(0,r.c)({},c,t,{components:n,mdxType:"MDXLayout"}),(0,a.yg)("h1",{id:"expression"},"Expression"),(0,a.yg)("p",null,(0,a.yg)("a",{parentName:"p",href:"https://github.com/antonmedv/expr"},"antonmedv/expr")," is used.\nAbout expr, please see ",(0,a.yg)("a",{parentName:"p",href:"https://github.com/antonmedv/expr/blob/master/docs/Language-Definition.md"},"Language Definition")," too."),(0,a.yg)("p",null,"e.g."),(0,a.yg)("pre",null,(0,a.yg)("code",{parentName:"pre",className:"language-yaml"},'entries:\n- if: Metadata.PackageName == "actions/checkout"\n')),(0,a.yg)("h3",{id:"expression-variables"},"Expression Variables"),(0,a.yg)("ul",null,(0,a.yg)("li",{parentName:"ul"},"Labels []string: Pull Request Label Names"),(0,a.yg)("li",{parentName:"ul"},"Metadata",(0,a.yg)("ul",{parentName:"li"},(0,a.yg)("li",{parentName:"ul"},"GroupName"),(0,a.yg)("li",{parentName:"ul"},"PackageName"),(0,a.yg)("li",{parentName:"ul"},"DepName"),(0,a.yg)("li",{parentName:"ul"},"PackageFileDir"),(0,a.yg)("li",{parentName:"ul"},"UpdateType")))))}m.isMDXComponent=!0}}]);
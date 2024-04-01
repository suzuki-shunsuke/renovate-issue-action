"use strict";(self.webpackChunkrenovate_issue_action=self.webpackChunkrenovate_issue_action||[]).push([[533],{5680:(e,t,n)=>{n.d(t,{xA:()=>u,yg:()=>g});var r=n(6540);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function a(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var c=r.createContext({}),l=function(e){var t=r.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):a(a({},t),e)),n},u=function(e){var t=l(e.components);return r.createElement(c.Provider,{value:t},e.children)},p="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},m=r.forwardRef((function(e,t){var n=e.components,o=e.mdxType,i=e.originalType,c=e.parentName,u=s(e,["components","mdxType","originalType","parentName"]),p=l(n),m=o,g=p["".concat(c,".").concat(m)]||p[m]||d[m]||i;return n?r.createElement(g,a(a({ref:t},u),{},{components:n})):r.createElement(g,a({ref:t},u))}));function g(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var i=n.length,a=new Array(i);a[0]=m;var s={};for(var c in t)hasOwnProperty.call(t,c)&&(s[c]=t[c]);s.originalType=e,s[p]="string"==typeof e?e:o,a[1]=s;for(var l=2;l<i;l++)a[l]=n[l];return r.createElement.apply(null,a)}return r.createElement.apply(null,n)}m.displayName="MDXCreateElement"},1119:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>c,contentTitle:()=>a,default:()=>d,frontMatter:()=>i,metadata:()=>s,toc:()=>l});var r=n(8168),o=(n(6540),n(5680));const i={sidebar_position:500},a="Add Issue to GitHub Projects",s={unversionedId:"config/add-issue-to-projects",id:"config/add-issue-to-projects",title:"Add Issue to GitHub Projects",description:"e.g.",source:"@site/docs/config/add-issue-to-projects.md",sourceDirName:"config",slug:"/config/add-issue-to-projects",permalink:"/renovate-issue-action/config/add-issue-to-projects",draft:!1,editUrl:"https://github.com/suzuki-shunsuke/renovate-issue-action-docs/edit/main/docs/config/add-issue-to-projects.md",tags:[],version:"current",sidebarPosition:500,frontMatter:{sidebar_position:500},sidebar:"tutorialSidebar",previous:{title:"Environment variable",permalink:"/renovate-issue-action/config/environment-variable"},next:{title:"How to Get GitHub Project ID",permalink:"/renovate-issue-action/config/get-project-id"}},c={},l=[{value:"How to get <code>column_id</code> and <code>next_id</code>",id:"how-to-get-column_id-and-next_id",level:2}],u={toc:l},p="wrapper";function d(e){let{components:t,...n}=e;return(0,o.yg)(p,(0,r.A)({},u,n,{components:t,mdxType:"MDXLayout"}),(0,o.yg)("h1",{id:"add-issue-to-github-projects"},"Add Issue to GitHub Projects"),(0,o.yg)("p",null,"e.g."),(0,o.yg)("pre",null,(0,o.yg)("code",{parentName:"pre",className:"language-yaml"},"projects:\n- alias: sre\n  column_id: PC_*** # GitHub Project Column ID\n- alias: service-a\n  next_id: PN_*** # Project Next Item ID\nissue:\n  projects:\n  - sre # alias\n  - service-a # alias\n")),(0,o.yg)("p",null,"You can add a crated issue to GitHub Projects.\nID of GitHub Project isn't human friendly, so in renovate-issue-action you specify projects with aliases.\nPairs of alias and Project are defined in the field ",(0,o.yg)("inlineCode",{parentName:"p"},"projects"),", and in ",(0,o.yg)("inlineCode",{parentName:"p"},"issue.projects")," projects are specified with aliases.\n",(0,o.yg)("inlineCode",{parentName:"p"},"alias")," in ",(0,o.yg)("inlineCode",{parentName:"p"},"projects")," must be unique."),(0,o.yg)("p",null,"Note that ",(0,o.yg)("inlineCode",{parentName:"p"},"alias")," in ",(0,o.yg)("inlineCode",{parentName:"p"},"projects")," is ",(0,o.yg)("strong",{parentName:"p"},"not")," GitHub Project Name but renovate-issue-action specific alias of GitHub Project."),(0,o.yg)("p",null,"renovate-issue-action supports both GitHub Project and GitHub Project Beta."),(0,o.yg)("p",null,"In case of GitHub Project, please specify Project Column ID as ",(0,o.yg)("inlineCode",{parentName:"p"},"column_id"),"."),(0,o.yg)("pre",null,(0,o.yg)("code",{parentName:"pre",className:"language-yaml"},"projects:\n- alias: sre\n  column_id: PC_*** # GitHub Project Column ID\n")),(0,o.yg)("p",null,"In case of GitHub Project Beta, please specify Project ID as ",(0,o.yg)("inlineCode",{parentName:"p"},"next_id"),"."),(0,o.yg)("pre",null,(0,o.yg)("code",{parentName:"pre",className:"language-yaml"},"projects:\n- alias: service-a\n  next_id: PN_*** # Project Next Item ID\n")),(0,o.yg)("h2",{id:"how-to-get-column_id-and-next_id"},"How to get ",(0,o.yg)("inlineCode",{parentName:"h2"},"column_id")," and ",(0,o.yg)("inlineCode",{parentName:"h2"},"next_id")),(0,o.yg)("p",null,"Please see ",(0,o.yg)("a",{parentName:"p",href:"get-project-id"},"How to Get GitHub Project ID")))}d.isMDXComponent=!0}}]);
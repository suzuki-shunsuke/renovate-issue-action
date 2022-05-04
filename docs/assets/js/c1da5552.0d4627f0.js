"use strict";(self.webpackChunkrenovate_issue_action=self.webpackChunkrenovate_issue_action||[]).push([[194],{3905:function(e,t,n){n.d(t,{Zo:function(){return l},kt:function(){return f}});var r=n(7294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function a(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function c(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var s=r.createContext({}),u=function(e){var t=r.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):a(a({},t),e)),n},l=function(e){var t=u(e.components);return r.createElement(s.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},d=r.forwardRef((function(e,t){var n=e.components,o=e.mdxType,i=e.originalType,s=e.parentName,l=c(e,["components","mdxType","originalType","parentName"]),d=u(n),f=o,m=d["".concat(s,".").concat(f)]||d[f]||p[f]||i;return n?r.createElement(m,a(a({ref:t},l),{},{components:n})):r.createElement(m,a({ref:t},l))}));function f(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var i=n.length,a=new Array(i);a[0]=d;var c={};for(var s in t)hasOwnProperty.call(t,s)&&(c[s]=t[s]);c.originalType=e,c.mdxType="string"==typeof e?e:o,a[1]=c;for(var u=2;u<i;u++)a[u]=n[u];return r.createElement.apply(null,a)}return r.createElement.apply(null,n)}d.displayName="MDXCreateElement"},3073:function(e,t,n){n.r(t),n.d(t,{assets:function(){return l},contentTitle:function(){return s},default:function(){return f},frontMatter:function(){return c},metadata:function(){return u},toc:function(){return p}});var r=n(7462),o=n(3366),i=(n(7294),n(3905)),a=["components"],c={sidebar_position:500},s="Add Issue to GitHub Projects",u={unversionedId:"config/add-issue-to-projects",id:"config/add-issue-to-projects",title:"Add Issue to GitHub Projects",description:"e.g.",source:"@site/docs/config/add-issue-to-projects.md",sourceDirName:"config",slug:"/config/add-issue-to-projects",permalink:"/renovate-issue-action/config/add-issue-to-projects",draft:!1,editUrl:"https://github.com/suzuki-shunsuke/renovate-issue-action-docs/edit/main/docs/config/add-issue-to-projects.md",tags:[],version:"current",sidebarPosition:500,frontMatter:{sidebar_position:500},sidebar:"tutorialSidebar",previous:{title:"Environment variable",permalink:"/renovate-issue-action/config/environment-variable"},next:{title:"How to Get GitHub Project ID",permalink:"/renovate-issue-action/config/get-project-id"}},l={},p=[{value:"How to get <code>column_id</code> and <code>next_id</code>",id:"how-to-get-column_id-and-next_id",level:2}],d={toc:p};function f(e){var t=e.components,n=(0,o.Z)(e,a);return(0,i.kt)("wrapper",(0,r.Z)({},d,n,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"add-issue-to-github-projects"},"Add Issue to GitHub Projects"),(0,i.kt)("p",null,"e.g."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-yaml"},"projects:\n- alias: sre\n  column_id: PC_*** # GitHub Project Column ID\n- alias: service-a\n  next_id: PN_*** # Project Next Item ID\nissue:\n  projects:\n  - sre # alias\n  - service-a # alias\n")),(0,i.kt)("p",null,"You can add a crated issue to GitHub Projects.\nID of GitHub Project isn't human friendly, so in renovate-issue-action you specify projects with aliases.\nPairs of alias and Project are defined in the field ",(0,i.kt)("inlineCode",{parentName:"p"},"projects"),", and in ",(0,i.kt)("inlineCode",{parentName:"p"},"issue.projects")," projects are specified with aliases.\n",(0,i.kt)("inlineCode",{parentName:"p"},"alias")," in ",(0,i.kt)("inlineCode",{parentName:"p"},"projects")," must be unique."),(0,i.kt)("p",null,"Note that ",(0,i.kt)("inlineCode",{parentName:"p"},"alias")," in ",(0,i.kt)("inlineCode",{parentName:"p"},"projects")," is ",(0,i.kt)("strong",{parentName:"p"},"not")," GitHub Project Name but renovate-issue-action specific alias of GitHub Project."),(0,i.kt)("p",null,"renovate-issue-action supports both GitHub Project and GitHub Project Beta."),(0,i.kt)("p",null,"In case of GitHub Project, please specify Project Column ID as ",(0,i.kt)("inlineCode",{parentName:"p"},"column_id"),"."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-yaml"},"projects:\n- alias: sre\n  column_id: PC_*** # GitHub Project Column ID\n")),(0,i.kt)("p",null,"In case of GitHub Project Beta, please specify Project ID as ",(0,i.kt)("inlineCode",{parentName:"p"},"next_id"),"."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-yaml"},"projects:\n- alias: service-a\n  next_id: PN_*** # Project Next Item ID\n")),(0,i.kt)("h2",{id:"how-to-get-column_id-and-next_id"},"How to get ",(0,i.kt)("inlineCode",{parentName:"h2"},"column_id")," and ",(0,i.kt)("inlineCode",{parentName:"h2"},"next_id")),(0,i.kt)("p",null,"Please see ",(0,i.kt)("a",{parentName:"p",href:"get-project-id"},"How to Get GitHub Project ID")))}f.isMDXComponent=!0}}]);
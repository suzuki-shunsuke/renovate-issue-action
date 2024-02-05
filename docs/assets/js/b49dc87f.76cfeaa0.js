"use strict";(self.webpackChunkrenovate_issue_action=self.webpackChunkrenovate_issue_action||[]).push([[650],{5788:(e,n,r)=>{r.d(n,{Iu:()=>s,yg:()=>d});var t=r(1504);function o(e,n,r){return n in e?Object.defineProperty(e,n,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[n]=r,e}function i(e,n){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var t=Object.getOwnPropertySymbols(e);n&&(t=t.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),r.push.apply(r,t)}return r}function a(e){for(var n=1;n<arguments.length;n++){var r=null!=arguments[n]?arguments[n]:{};n%2?i(Object(r),!0).forEach((function(n){o(e,n,r[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(r,n))}))}return e}function u(e,n){if(null==e)return{};var r,t,o=function(e,n){if(null==e)return{};var r,t,o={},i=Object.keys(e);for(t=0;t<i.length;t++)r=i[t],n.indexOf(r)>=0||(o[r]=e[r]);return o}(e,n);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(t=0;t<i.length;t++)r=i[t],n.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}var l=t.createContext({}),c=function(e){var n=t.useContext(l),r=n;return e&&(r="function"==typeof e?e(n):a(a({},n),e)),r},s=function(e){var n=c(e.components);return t.createElement(l.Provider,{value:n},e.children)},g="mdxType",p={inlineCode:"code",wrapper:function(e){var n=e.children;return t.createElement(t.Fragment,{},n)}},m=t.forwardRef((function(e,n){var r=e.components,o=e.mdxType,i=e.originalType,l=e.parentName,s=u(e,["components","mdxType","originalType","parentName"]),g=c(r),m=o,d=g["".concat(l,".").concat(m)]||g[m]||p[m]||i;return r?t.createElement(d,a(a({ref:n},s),{},{components:r})):t.createElement(d,a({ref:n},s))}));function d(e,n){var r=arguments,o=n&&n.mdxType;if("string"==typeof e||o){var i=r.length,a=new Array(i);a[0]=m;var u={};for(var l in n)hasOwnProperty.call(n,l)&&(u[l]=n[l]);u.originalType=e,u[g]="string"==typeof e?e:o,a[1]=u;for(var c=2;c<i;c++)a[c]=r[c];return t.createElement.apply(null,a)}return t.createElement.apply(null,r)}m.displayName="MDXCreateElement"},9868:(e,n,r)=>{r.r(n),r.d(n,{assets:()=>l,contentTitle:()=>a,default:()=>p,frontMatter:()=>i,metadata:()=>u,toc:()=>c});var t=r(5072),o=(r(1504),r(5788));const i={sidebar_position:600},a="How to Get GitHub Project ID",u={unversionedId:"config/get-project-id",id:"config/get-project-id",title:"How to Get GitHub Project ID",description:"In our understanding, to get GitHub Project Column ID and GitHub Project Next Item ID, you have to use GitHub API.",source:"@site/docs/config/get-project-id.md",sourceDirName:"config",slug:"/config/get-project-id",permalink:"/renovate-issue-action/config/get-project-id",draft:!1,editUrl:"https://github.com/suzuki-shunsuke/renovate-issue-action-docs/edit/main/docs/config/get-project-id.md",tags:[],version:"current",sidebarPosition:600,frontMatter:{sidebar_position:600},sidebar:"tutorialSidebar",previous:{title:"Add Issue to GitHub Projects",permalink:"/renovate-issue-action/config/add-issue-to-projects"}},l={},c=[{value:"Project Column ID",id:"project-column-id",level:2},{value:"Repository Project",id:"repository-project",level:3},{value:"User Project",id:"user-project",level:3},{value:"Organization Project",id:"organization-project",level:3},{value:"Project Next Item ID",id:"project-next-item-id",level:2},{value:"User Project",id:"user-project-1",level:3},{value:"Organization Project",id:"organization-project-1",level:3}],s={toc:c},g="wrapper";function p(e){let{components:n,...r}=e;return(0,o.yg)(g,(0,t.c)({},s,r,{components:n,mdxType:"MDXLayout"}),(0,o.yg)("h1",{id:"how-to-get-github-project-id"},"How to Get GitHub Project ID"),(0,o.yg)("p",null,"In our understanding, to get GitHub Project Column ID and GitHub Project Next Item ID, you have to use GitHub API."),(0,o.yg)("p",null,"In the following scripts, we use ",(0,o.yg)("a",{parentName:"p",href:"https://cli.github.com/"},"GitHub CLI"),"."),(0,o.yg)("h2",{id:"project-column-id"},"Project Column ID"),(0,o.yg)("h3",{id:"repository-project"},"Repository Project"),(0,o.yg)("pre",null,(0,o.yg)("code",{parentName:"pre",className:"language-shell"},'#!/usr/bin/env bash\n\nset -eu\n\nowner=$1\nrepo=$2\nnumber=$3\n\n# shellcheck disable=SC2016\nQUERY=\'\nquery($owner: String!, $name: String!, $number: Int!) {\n  repository(owner: $owner, name: $name){\n    project(number: $number) {\n      columns(first: 100) {\n        nodes {\n          databaseId\n          name\n        }\n      }\n    }\n  }\n}\'\n\ngh api graphql -F "owner=$owner" -F "name=$repo" -F "number=$number" -f query="$QUERY"\n')),(0,o.yg)("h3",{id:"user-project"},"User Project"),(0,o.yg)("pre",null,(0,o.yg)("code",{parentName:"pre",className:"language-shell"},'#!/usr/bin/env bash\n\nset -eu\n\nlogin=$1\nnumber=$2\n\n# shellcheck disable=SC2016\nQUERY=\'\nquery($login: String!, $number: Int!) {\n  user(login: $login){\n    project(number: $number) {\n      columns(first: 100) {\n        nodes {\n          databaseId\n          name\n        }\n      }\n    }\n  }\n}\'\n\ngh api graphql -F "login=$login" -F "number=$number" -f query="$QUERY"\n')),(0,o.yg)("h3",{id:"organization-project"},"Organization Project"),(0,o.yg)("pre",null,(0,o.yg)("code",{parentName:"pre",className:"language-shell"},'#!/usr/bin/env bash\n\nset -eu\n\nlogin=$1\nnumber=$2\n\n# shellcheck disable=SC2016\nQUERY=\'\nquery($login: String!, $number: Int!) {\n  organization(login: $login){\n    project(number: $number) {\n      columns(first: 100) {\n        nodes {\n          databaseId\n          name\n        }\n      }\n    }\n  }\n}\'\n\ngh api graphql -F "login=$login" -F "number=$number" -f query="$QUERY"\n')),(0,o.yg)("h2",{id:"project-next-item-id"},"Project Next Item ID"),(0,o.yg)("h3",{id:"user-project-1"},"User Project"),(0,o.yg)("pre",null,(0,o.yg)("code",{parentName:"pre",className:"language-shell"},'#!/usr/bin/env bash\n\nset -eu\n\nlogin=$1\nnumber=$2\n\n# shellcheck disable=SC2016\nQUERY=\'\nquery($login: String!, $number: Int!) {\n  user(login: $login){\n    projectNext(number: $number) {\n      id\n      title\n    }\n  }\n}\'\n\ngh api graphql -F "login=$login" -F "number=$number" -f query="$QUERY"\n')),(0,o.yg)("h3",{id:"organization-project-1"},"Organization Project"),(0,o.yg)("pre",null,(0,o.yg)("code",{parentName:"pre",className:"language-shell"},'#!/usr/bin/env bash\n\nset -eu\n\nlogin=$1\nnumber=$2\n\n# shellcheck disable=SC2016\nQUERY=\'\nquery($login: String!, $number: Int!) {\n  organization(login: $login){\n    projectNext(number: $number) {\n      id\n      title\n    }\n  }\n}\'\n\ngh api graphql -F "login=$login" -F "number=$number" -f query="$QUERY"\n')))}p.isMDXComponent=!0}}]);
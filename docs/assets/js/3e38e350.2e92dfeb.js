"use strict";(self.webpackChunkrenovate_issue_action=self.webpackChunkrenovate_issue_action||[]).push([[896],{3905:function(e,n,t){t.d(n,{Zo:function(){return s},kt:function(){return m}});var r=t(7294);function i(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function o(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,r)}return t}function a(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?o(Object(t),!0).forEach((function(n){i(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):o(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function u(e,n){if(null==e)return{};var t,r,i=function(e,n){if(null==e)return{};var t,r,i={},o=Object.keys(e);for(r=0;r<o.length;r++)t=o[r],n.indexOf(t)>=0||(i[t]=e[t]);return i}(e,n);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)t=o[r],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(i[t]=e[t])}return i}var l=r.createContext({}),c=function(e){var n=r.useContext(l),t=n;return e&&(t="function"==typeof e?e(n):a(a({},n),e)),t},s=function(e){var n=c(e.components);return r.createElement(l.Provider,{value:n},e.children)},p={inlineCode:"code",wrapper:function(e){var n=e.children;return r.createElement(r.Fragment,{},n)}},f=r.forwardRef((function(e,n){var t=e.components,i=e.mdxType,o=e.originalType,l=e.parentName,s=u(e,["components","mdxType","originalType","parentName"]),f=c(t),m=i,v=f["".concat(l,".").concat(m)]||f[m]||p[m]||o;return t?r.createElement(v,a(a({ref:n},s),{},{components:t})):r.createElement(v,a({ref:n},s))}));function m(e,n){var t=arguments,i=n&&n.mdxType;if("string"==typeof e||i){var o=t.length,a=new Array(o);a[0]=f;var u={};for(var l in n)hasOwnProperty.call(n,l)&&(u[l]=n[l]);u.originalType=e,u.mdxType="string"==typeof e?e:i,a[1]=u;for(var c=2;c<o;c++)a[c]=t[c];return r.createElement.apply(null,a)}return r.createElement.apply(null,t)}f.displayName="MDXCreateElement"},3614:function(e,n,t){t.r(n),t.d(n,{assets:function(){return s},contentTitle:function(){return l},default:function(){return m},frontMatter:function(){return u},metadata:function(){return c},toc:function(){return p}});var r=t(7462),i=t(3366),o=(t(7294),t(3905)),a=["components"],u={sidebar_position:400},l="Environment variable",c={unversionedId:"config/environment-variable",id:"config/environment-variable",title:"Environment variable",description:"The following environment variables are required.",source:"@site/docs/config/environment-variable.md",sourceDirName:"config",slug:"/config/environment-variable",permalink:"/renovate-issue-action/config/environment-variable",editUrl:"https://github.com/suzuki-shunsuke/renovate-issue-action-docs/edit/main/docs/config/environment-variable.md",tags:[],version:"current",sidebarPosition:400,frontMatter:{sidebar_position:400},sidebar:"tutorialSidebar",previous:{title:"Expression",permalink:"/renovate-issue-action/config/expression"}},s={},p=[{value:"GitHub Token",id:"github-token",level:2}],f={toc:p};function m(e){var n=e.components,t=(0,i.Z)(e,a);return(0,o.kt)("wrapper",(0,r.Z)({},f,t,{components:n,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"environment-variable"},"Environment variable"),(0,o.kt)("p",null,"The following environment variables are required."),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},"GITHUB_TOKEN")),(0,o.kt)("p",null,"Furthermore, the following ",(0,o.kt)("a",{parentName:"p",href:"https://docs.github.com/en/actions/learn-github-actions/environment-variables#default-environment-variables"},"GitHub Actions Default environment variables")," are used."),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},"GITHUB_EVENT_PATH"),(0,o.kt)("li",{parentName:"ul"},"GITHUB_RUN_ID")),(0,o.kt)("h2",{id:"github-token"},"GitHub Token"),(0,o.kt)("p",null,"Basically, you can use GitHub Actions Access Token ",(0,o.kt)("inlineCode",{parentName:"p"},"github.token"),".\nBut if you want to create issues in other repositories or avoid API Rate Limiting,\nyou have to use Personal Access Token or GitHub App's token."),(0,o.kt)("p",null,"The following permissions are required."),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("inlineCode",{parentName:"li"},"issues:write"))))}m.isMDXComponent=!0}}]);
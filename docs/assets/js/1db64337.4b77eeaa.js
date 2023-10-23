"use strict";(self.webpackChunkrenovate_issue_action=self.webpackChunkrenovate_issue_action||[]).push([[372],{3905:(e,t,n)=>{n.d(t,{Zo:()=>c,kt:()=>v});var r=n(7294);function s(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach((function(t){s(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,r,s=function(e,t){if(null==e)return{};var n,r,s={},a=Object.keys(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||(s[n]=e[n]);return s}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(s[n]=e[n])}return s}var l=r.createContext({}),u=function(e){var t=r.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},c=function(e){var t=u(e.components);return r.createElement(l.Provider,{value:t},e.children)},p="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},m=r.forwardRef((function(e,t){var n=e.components,s=e.mdxType,a=e.originalType,l=e.parentName,c=i(e,["components","mdxType","originalType","parentName"]),p=u(n),m=s,v=p["".concat(l,".").concat(m)]||p[m]||d[m]||a;return n?r.createElement(v,o(o({ref:t},c),{},{components:n})):r.createElement(v,o({ref:t},c))}));function v(e,t){var n=arguments,s=t&&t.mdxType;if("string"==typeof e||s){var a=n.length,o=new Array(a);o[0]=m;var i={};for(var l in t)hasOwnProperty.call(t,l)&&(i[l]=t[l]);i.originalType=e,i[p]="string"==typeof e?e:s,o[1]=i;for(var u=2;u<a;u++)o[u]=n[u];return r.createElement.apply(null,o)}return r.createElement.apply(null,n)}m.displayName="MDXCreateElement"},6777:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>o,default:()=>d,frontMatter:()=>a,metadata:()=>i,toc:()=>u});var r=n(7462),s=(n(7294),n(3905));const a={sidebar_position:100},o="Overview",i={unversionedId:"overview",id:"overview",title:"Overview",description:"Background",source:"@site/docs/overview.md",sourceDirName:".",slug:"/overview",permalink:"/renovate-issue-action/overview",draft:!1,editUrl:"https://github.com/suzuki-shunsuke/renovate-issue-action-docs/edit/main/docs/overview.md",tags:[],version:"current",sidebarPosition:100,frontMatter:{sidebar_position:100},sidebar:"tutorialSidebar",previous:{title:"renovate-issue-action",permalink:"/renovate-issue-action/"},next:{title:"How to use",permalink:"/renovate-issue-action/how-to-use"}},l={},u=[{value:"Background",id:"background",level:2},{value:"Overview",id:"overview-1",level:2},{value:"Why don&#39;t you use closed Pull Requests instead of Issues?",id:"why-dont-you-use-closed-pull-requests-instead-of-issues",level:2}],c={toc:u},p="wrapper";function d(e){let{components:t,...n}=e;return(0,s.kt)(p,(0,r.Z)({},c,n,{components:t,mdxType:"MDXLayout"}),(0,s.kt)("h1",{id:"overview"},"Overview"),(0,s.kt)("h2",{id:"background"},"Background"),(0,s.kt)("p",null,"We wrote blog posts about how to handle a number of Pull Requests by Renovate automatically."),(0,s.kt)("ul",null,(0,s.kt)("li",{parentName:"ul"},(0,s.kt)("a",{parentName:"li",href:"https://blog.studysapuri.jp/entry/2022/02/18/080000"},"2022-02-18 Renovate \u306e\u5927\u91cf\u306e Pull Request \u3092\u51e6\u7406\u3059\u308b\u6280\u8853")),(0,s.kt)("li",{parentName:"ul"},(0,s.kt)("a",{parentName:"li",href:"https://devs.quipper.com/2022/03/29/automate-handling-a-number-of-pull-requests-by-renovate-in-terraform-monorepo.html"},"2022-03-29 Automate handling a number of Pull Requests by Renovate in Terraform Monorepo"))),(0,s.kt)("p",null,"As described in blog posts, leaving Renovate pull requests open will limit the number of new pull requests that can be created.\nTherefore, you could close pull requests that could not be automerged and delete feature branches automatically."),(0,s.kt)("p",null,(0,s.kt)("inlineCode",{parentName:"p"},"renovate-issue-action")," helps you managing tasks as GitHub Issues to handle closed pull requests."),(0,s.kt)("h2",{id:"overview-1"},"Overview"),(0,s.kt)("p",null,"Please run ",(0,s.kt)("inlineCode",{parentName:"p"},"renovate-issue-action")," with GitHub Actions triggered by ",(0,s.kt)("inlineCode",{parentName:"p"},"pull_request"),"'s ",(0,s.kt)("inlineCode",{parentName:"p"},"closed")," events.\nAn Issue is created when Renovate's Pull Request was closed by other than Renovate."),(0,s.kt)("p",null,(0,s.kt)("img",{parentName:"p",src:"https://user-images.githubusercontent.com/13323303/164878251-96020cd9-052c-4e33-a17d-c6201ebcaa94.png",alt:"image"})),(0,s.kt)("p",null,"--"),(0,s.kt)("p",null,(0,s.kt)("img",{parentName:"p",src:"https://user-images.githubusercontent.com/13323303/164878264-641b05ab-3a4d-42d8-82b8-76d806751ebe.png",alt:"image"})),(0,s.kt)("p",null,"--"),(0,s.kt)("p",null,(0,s.kt)("img",{parentName:"p",src:"https://user-images.githubusercontent.com/13323303/164878275-ba0264a6-043b-473a-b1a4-9c8c58054662.png",alt:"image"})),(0,s.kt)("p",null,"The issue description has a list of closed Pull Requests.\nIf the issue already exists, closed pull request is added to the list of closed Pull Request instead of creating a new Issue."),(0,s.kt)("p",null,(0,s.kt)("img",{parentName:"p",src:"https://user-images.githubusercontent.com/13323303/164878350-74ae61b2-0a22-4dbd-a06c-cc3d9345b54b.png",alt:"image"})),(0,s.kt)("p",null,"If a Pull Request is merged or Renovate closes it, the related issue would be closed."),(0,s.kt)("p",null,(0,s.kt)("img",{parentName:"p",src:"https://user-images.githubusercontent.com/13323303/164878427-eb5a9d48-634e-4099-a38f-bbb0a7b894f6.png",alt:"image"})),(0,s.kt)("p",null,"--"),(0,s.kt)("p",null,(0,s.kt)("img",{parentName:"p",src:"https://user-images.githubusercontent.com/13323303/164878444-0df765eb-1e2b-4a84-9b21-6ff911511bd1.png",alt:"image"})),(0,s.kt)("h2",{id:"why-dont-you-use-closed-pull-requests-instead-of-issues"},"Why don't you use closed Pull Requests instead of Issues?"),(0,s.kt)("p",null,"You could use closed Pull Requests instead of Issues without ",(0,s.kt)("inlineCode",{parentName:"p"},"renovate-issue-action"),".\nThen why do we develop ",(0,s.kt)("inlineCode",{parentName:"p"},"renovate-issue-action"),"?"),(0,s.kt)("p",null,"We think it is inconvenient to treat closed Pull Requests as tasks.\nRenovate creates a new Pull Request when a new version is released,\nso there would be multiple Pull Requests against the same task.\nIt is difficult to understand the list of unresolved Pull Requests."))}d.isMDXComponent=!0}}]);
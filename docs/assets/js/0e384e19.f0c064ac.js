"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[671],{7876:(e,t,s)=>{s.r(t),s.d(t,{assets:()=>l,contentTitle:()=>o,default:()=>h,frontMatter:()=>r,metadata:()=>a,toc:()=>c});var n=s(5893),i=s(1151);const r={title:"Introduction",slug:"/"},o=void 0,a={id:"intro",title:"Introduction",description:"AIKit is a quick, easy, and local or cloud-agnostic way to get started to host and deploy large language models (LLMs) for inference. No GPU, internet access or additional tools are needed to get started except for Docker!",source:"@site/docs/intro.md",sourceDirName:".",slug:"/",permalink:"/aikit/",draft:!1,unlisted:!1,editUrl:"https://github.com/sozercan/aikit/blob/main/website/docs/docs/intro.md",tags:[],version:"current",frontMatter:{title:"Introduction",slug:"/"},sidebar:"sidebar",next:{title:"Quick Start",permalink:"/aikit/quick-start"}},l={},c=[{value:"Features",id:"features",level:2}];function d(e){const t={a:"a",code:"code",h2:"h2",li:"li",p:"p",ul:"ul",...(0,i.a)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsxs)(t.p,{children:["AIKit is a quick, easy, and local or cloud-agnostic way to get started to host and deploy large language models (LLMs) for inference. No GPU, internet access or additional tools are needed to get started except for ",(0,n.jsx)(t.a,{href:"https://docs.docker.com/desktop/install/linux-install/",children:"Docker"}),"!"]}),"\n",(0,n.jsxs)(t.p,{children:["AIKit uses ",(0,n.jsx)(t.a,{href:"https://localai.io/",children:"LocalAI"})," under-the-hood to run inference. LocalAI provides a drop-in replacement REST API that is OpenAI API compatible, so you can use any OpenAI API compatible client, such as ",(0,n.jsx)(t.a,{href:"https://github.com/sozercan/kubectl-ai",children:"Kubectl AI"}),", ",(0,n.jsx)(t.a,{href:"https://github.com/sozercan/chatbot-ui",children:"Chatbot-UI"})," and many more, to send requests to open-source LLMs powered by AIKit!"]}),"\n",(0,n.jsx)(t.h2,{id:"features",children:"Features"}),"\n",(0,n.jsxs)(t.ul,{children:["\n",(0,n.jsxs)(t.li,{children:["\ud83d\udc33 No GPU, Internet access or additional tools needed except for ",(0,n.jsx)(t.a,{href:"https://docs.docker.com/desktop/install/linux-install/",children:"Docker"}),"!"]}),"\n",(0,n.jsxs)(t.li,{children:["\ud83e\udd0f Minimal image size, resulting in less vulnerabilities and smaller attack surface with a custom ",(0,n.jsx)(t.a,{href:"https://github.com/GoogleContainerTools/distroless",children:"distroless"}),"-based image"]}),"\n",(0,n.jsxs)(t.li,{children:["\ud83d\ude80 ",(0,n.jsx)(t.a,{href:"/aikit/specs",children:"Easy to use declarative configuration"})]}),"\n",(0,n.jsx)(t.li,{children:"\u2728 OpenAI API compatible to use with any OpenAI API compatible client"}),"\n",(0,n.jsxs)(t.li,{children:["\ud83d\udcf8 ",(0,n.jsx)(t.a,{href:"/aikit/vision",children:"Multi-modal model support"})]}),"\n",(0,n.jsx)(t.li,{children:"\ud83d\uddbc\ufe0f Image generation support with Stable Diffusion"}),"\n",(0,n.jsxs)(t.li,{children:["\ud83e\udd99 Support for GGUF (",(0,n.jsx)(t.a,{href:"https://github.com/ggerganov/llama.cpp",children:(0,n.jsx)(t.code,{children:"llama"})}),"), GPTQ (",(0,n.jsx)(t.a,{href:"https://github.com/turboderp/exllama",children:(0,n.jsx)(t.code,{children:"exllama"})})," or ",(0,n.jsx)(t.a,{href:"https://github.com/turboderp/exllamav2",children:(0,n.jsx)(t.code,{children:"exllama2"})}),"), EXL2 (",(0,n.jsx)(t.a,{href:"https://github.com/turboderp/exllamav2",children:(0,n.jsx)(t.code,{children:"exllama2"})}),"), and GGML (",(0,n.jsx)(t.a,{href:"https://github.com/ggerganov/llama.cpp",children:(0,n.jsx)(t.code,{children:"llama-ggml"})}),") and ",(0,n.jsx)(t.a,{href:"https://github.com/state-spaces/mamba",children:"Mamba"})," models"]}),"\n",(0,n.jsxs)(t.li,{children:["\ud83d\udea2 ",(0,n.jsx)(t.a,{href:"#kubernetes-deployment",children:"Kubernetes deployment ready"})]}),"\n",(0,n.jsx)(t.li,{children:"\ud83d\udce6 Supports multiple models with a single image"}),"\n",(0,n.jsxs)(t.li,{children:["\ud83d\udda5\ufe0f ",(0,n.jsx)(t.a,{href:"/aikit/gpu",children:"Supports GPU-accelerated inferencing with NVIDIA GPUs"})]}),"\n",(0,n.jsxs)(t.li,{children:["\ud83d\udd10 ",(0,n.jsxs)(t.a,{href:"/aikit/cosign",children:["Signed images for ",(0,n.jsx)(t.code,{children:"aikit"})," and pre-made models"]})]}),"\n",(0,n.jsx)(t.li,{children:"\ud83c\udf08 Support for non-proprietary and self-hosted container registries to store model images"}),"\n"]}),"\n",(0,n.jsxs)(t.p,{children:["To get started, please see ",(0,n.jsx)(t.a,{href:"/aikit/quick-start",children:"Quick Start"}),"!"]})]})}function h(e={}){const{wrapper:t}={...(0,i.a)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(d,{...e})}):d(e)}},1151:(e,t,s)=>{s.d(t,{Z:()=>a,a:()=>o});var n=s(7294);const i={},r=n.createContext(i);function o(e){const t=n.useContext(r);return n.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}function a(e){let t;return t=e.disableParentContext?"function"==typeof e.components?e.components(i):e.components||i:o(e.components),n.createElement(r.Provider,{value:t},e.children)}}}]);
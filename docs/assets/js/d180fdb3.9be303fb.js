"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[136],{5639:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>c,contentTitle:()=>a,default:()=>m,frontMatter:()=>r,metadata:()=>i,toc:()=>l});var s=n(5893),o=n(1151);const r={title:"Kubernetes Deployment"},a=void 0,i={id:"kubernetes",title:"Kubernetes Deployment",description:"It is easy to get started to deploy your models to Kubernetes!",source:"@site/docs/kubernetes.md",sourceDirName:".",slug:"/kubernetes",permalink:"/aikit/kubernetes",draft:!1,unlisted:!1,editUrl:"https://github.com/sozercan/aikit/blob/main/website/docs/docs/kubernetes.md",tags:[],version:"current",frontMatter:{title:"Kubernetes Deployment"},sidebar:"sidebar",previous:{title:"GPU Acceleration",permalink:"/aikit/gpu"},next:{title:"Image Verification",permalink:"/aikit/cosign"}},c={},l=[];function d(e){const t={a:"a",admonition:"admonition",code:"code",p:"p",pre:"pre",...(0,o.a)(),...e.components};return(0,s.jsxs)(s.Fragment,{children:[(0,s.jsx)(t.p,{children:"It is easy to get started to deploy your models to Kubernetes!"}),"\n",(0,s.jsxs)(t.p,{children:["Make sure you have a Kubernetes cluster running and ",(0,s.jsx)(t.code,{children:"kubectl"})," is configured to talk to it, and your model images are accessible from the cluster."]}),"\n",(0,s.jsx)(t.admonition,{type:"tip",children:(0,s.jsxs)(t.p,{children:["You can use ",(0,s.jsx)(t.a,{href:"https://kind.sigs.k8s.io/",children:"kind"})," to create a local Kubernetes cluster for testing purposes."]})}),"\n",(0,s.jsx)(t.pre,{children:(0,s.jsx)(t.code,{className:"language-bash",children:'# create a deployment\n# for pre-made models, replace "my-model" with the image name\nkubectl create deployment my-llm-deployment --image=my-model\n\n# expose it as a service\nkubectl expose deployment my-llm-deployment --port=8080 --target-port=8080 --name=my-llm-service\n\n# easy to scale up and down as needed\nkubectl scale deployment my-llm-deployment --replicas=3\n\n# port-forward for testing locally\nkubectl port-forward service/my-llm-service 8080:8080\n\n# send requests to your model\ncurl http://localhost:8080/v1/chat/completions -H "Content-Type: application/json" -d \'{\n     "model": "llama-2-7b-chat",\n     "messages": [{"role": "user", "content": "explain kubernetes in a sentence"}]\n   }\'\n{"created":1701236489,"object":"chat.completion","id":"dd1ff40b-31a7-4418-9e32-42151ab6875a","model":"llama-2-7b-chat","choices":[{"index":0,"finish_reason":"stop","message":{"role":"assistant","content":"\\nKubernetes is a container orchestration system that automates the deployment, scaling, and management of containerized applications in a microservices architecture."}}],"usage":{"prompt_tokens":0,"completion_tokens":0,"total_tokens":0}}\n'})}),"\n",(0,s.jsx)(t.admonition,{type:"tip",children:(0,s.jsxs)(t.p,{children:["For an example Kubernetes deployment and service YAML, see ",(0,s.jsx)(t.a,{href:"./kubernetes/",children:"kubernetes folder"}),". Please note that these are examples, you may need to customize them (such as properly configured resource requests and limits) based on your needs."]})})]})}function m(e={}){const{wrapper:t}={...(0,o.a)(),...e.components};return t?(0,s.jsx)(t,{...e,children:(0,s.jsx)(d,{...e})}):d(e)}},1151:(e,t,n)=>{n.d(t,{Z:()=>i,a:()=>a});var s=n(7294);const o={},r=s.createContext(o);function a(e){const t=s.useContext(r);return s.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}function i(e){let t;return t=e.disableParentContext?"function"==typeof e.components?e.components(o):e.components||o:a(e.components),s.createElement(r.Provider,{value:t},e.children)}}}]);
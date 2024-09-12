# Personal Website

Here is my personal website, it is not fancy and not even complex. In future I will definately put some interacticity into some project, but for now LGTM.

## How to run

If for some reason you want to run this website on your computer or use it as a service to deploy, Here is instructions:

1. Clone the repo
2. Build docker image ``` docker build --tag web-app . ```
3. Run docker image ``` docker run -p 8080:8080 web-app ```
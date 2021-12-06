![GitHub Workflow Status](https://img.shields.io/github/workflow/status/dirien/gitlab-wrap/build?style=for-the-badge)
![GitHub](https://img.shields.io/github/license/dirien/gitlab-wrap?style=for-the-badge)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/dirien/gitlab-wrap?style=for-the-badge)

# GitLab Wrap

GitLab Wrap 2021 is a smol web application that creates a card with your stats from your public GitLab account for 2021.

## Installation TL;DR

Easy, just call `docker build -t wrap .` in the root of the project.

To start the container, just call ` docker run -p 8080:8080 -e GITLAB_TOKEN=<gitlab-api-token> wrap`.

## Assets

In the internal directory (`internal/card/assets`) of the project you can find the following assets:

- 404.png
- card.png
- SourceSansPro-Regular.ttf

assets to render the card.

## UI

The Ui is based on the [Vue.js](https://vuejs.org/) framework and is located in the `ui` directory.

You can start the development server with `npm run dev`

I did not configure the calls to the backend server, so you need to change the string `https://gitlabwrap.fly.dev/card/`
by hand in the UI.

## Backend

Run `npm run build` to build the dist folder of the UI. This path needs to be set as flag `--static` in the `go run`
command.

You can start the go backend with `go run . --static ui/dist`. It listens on port 8080 per default, but you can set the
environment variable `PORT` to change the port.

# Caveats

The calls from the UI to the backend server are not configurable, so you need to change the
string `https://gitlabwrap.fly.dev/card/` by hand in the UI.

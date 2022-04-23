# renovate-issue-action

Create, update, and close GitHub Issues with GitHub Actions according to [Renovate](https://github.com/renovatebot/renovate) Pull Requests.
Especially, this tool is useful for the usecase which you have to handle a number of Pull Requests automatically and some pull requests can't be merged automatically due to CI failure.

<img width="1068" alt="image" src="https://user-images.githubusercontent.com/13323303/164878956-45d9ba65-436b-48a8-ae7d-d76712822007.png">

## :warning: Don't use this yet

This is still under development.
The development status is alpha.

## Background

I wrote blog posts about how to handle a number of Pull Requests by Renovate automatically.

* [2022-02-18 Renovate の大量の Pull Request を処理する技術](https://blog.studysapuri.jp/entry/2022/02/18/080000)
* [2022-03-29 Automate handling a number of Pull Requests by Renovate in Terraform Monorepo](https://devs.quipper.com/2022/03/29/automate-handling-a-number-of-pull-requests-by-renovate-in-terraform-monorepo.html)

As described in blog posts, leaving Renovate pull requests open will limit the number of new pull requests that can be created.
Therefore, you could close pull requests that could not be automerged and delete feature branches automatically.

`renovate-issue-action` helps you managing tasks as GitHub Issues to handle closed pull requests.

## Overview

Please run `renovate-issue-action` with GitHub Actions triggered by `pull_request`'s `closed` events.
An Issue is created when Renovate's Pull Request was closed by other than Renovate.

<img width="955" alt="image" src="https://user-images.githubusercontent.com/13323303/164878251-96020cd9-052c-4e33-a17d-c6201ebcaa94.png">

--

<img width="952" alt="image" src="https://user-images.githubusercontent.com/13323303/164878264-641b05ab-3a4d-42d8-82b8-76d806751ebe.png">

--

<img width="1071" alt="image" src="https://user-images.githubusercontent.com/13323303/164878275-ba0264a6-043b-473a-b1a4-9c8c58054662.png">

The issue description has a list of closed Pull Requests.
If the issue already exists, closed pull request is added to the list of closed Pull Request instead of creating a new Issue.

<img width="1046" alt="image" src="https://user-images.githubusercontent.com/13323303/164878350-74ae61b2-0a22-4dbd-a06c-cc3d9345b54b.png">

If a Pull Request is merged or Renovate closes it, the related issue would be closed.

<img width="906" alt="image" src="https://user-images.githubusercontent.com/13323303/164878427-eb5a9d48-634e-4099-a38f-bbb0a7b894f6.png">

--

<img width="935" alt="image" src="https://user-images.githubusercontent.com/13323303/164878444-0df765eb-1e2b-4a84-9b21-6ff911511bd1.png">

## Why don't you use closed Pull Requests instead of Issues?

You could use closed Pull Requests instead of Issues without `renovate-issue-action`.
Then why do we develop `renovate-issue-action`?

We think it is inconvenient to treat closed Pull Requests as tasks.
Renovate creates a new Pull Request when a new version is released,
so there would be multiple Pull Requests against the same task.
It is difficult to understand the list of unresolved Pull Requests.

## How to use

Please add GitHub Actions Workflow and run renovate-issue-action.

[Example](https://github.com/suzuki-shunsuke/example-renovate-issue-action/blob/main/.github/workflows/close-renovate-pr.yaml)

Please update Renovate Configuration using [prBodyNotes](https://docs.renovatebot.com/configuration-options/#prbodynotes).

```json
  "prBodyNotes": [
    "<!-- renovate-issue-action: {\"packageFileDir\": \"{{packageFileDir}}\", \"packageName\": \"{{packageName}}\", \"groupName\": \"{{groupName}}\", \"depName\": \"{{depName}}\", \"manager\": \"{{manager}}\", \"updateType\": \"{{updateType}}\"} -->"
  ]
```

renovate-issue-action gets metadata from this HTML comment.

## Install

Note that `renovate-issue-action` isn't a Action but a CLI tool.
Please download the asset from [GitHub Releases](https://github.com/suzuki-shunsuke/renovate-issue-action/releases).
We recommend [aqua](https://aquaproj.github.io/) to install `renovate-issue-action`.

## Configuration

Configuration file is optional.
By default, the file `^\.renovate-issue-action\.ya?ml$` in the current directory is read if it exists.
You can specify configuration file path with the command line option `--config, -c`.

```console
$ renovate-issue-action --config config.yaml
```

renovate-issue-action.yaml

```yaml
renovate_login: 'renovate[bot]' # By default, 'renovate[bot]'
issue:
  repo_owner: suzuki-shunsuke # By default, the repository which GitHub Actions is run
  repo_name: renovate-issue-action # By default, the repository which GitHub Actions is run
  title: 'Renovate Automerge Failure({{.RepoOwner}}/{{.RepoName}}): {{if .Metadata.GroupName}}{{.Metadata.GroupName}}{{else}}{{.Metadata.PackageName}}{{.Metadata.DepName}}{{end}} {{if .Metadata.PackageFileDir}}({{.Metadata.PackageFileDir}}){{end}}'
  description_header: |
    _This pull request was created by [renovate-issue-action](https://github.com/suzuki-shunsuke/renovate-issue-action)._

    :warning: Please don't edit the Issue title, because renovate-issue-action searches issue with Issue title.

    {{if .Metadata.PackageName}}packageName: {{.Metadata.PackageName}}{{end}}
    {{if .Metadata.GroupName}}groupName: {{.Metadata.GroupName}}{{end}}
    {{if .Metadata.DepName}}depName: {{.Metadata.DepName}}{{end}}
  description_body: "" # By default, empty
  assignees: ["suzuki-shunsuke"] # By default, null
```

### Template

Some configuration fields are parsed with Go's [text/template](https://pkg.go.dev/text/template).
In the template, [sprig Function](http://masterminds.github.io/sprig/) can be used.

#### Template Variables

* RepoOwner
* RepoName
* Metadata
  * GroupName
  * PackageName
  * DepName
  * PackageFileDir

### Environment variable

The following environment variables are required.

* GITHUB_TOKEN

Furthermore, the following [GitHub Actions Default environment variables](https://docs.github.com/en/actions/learn-github-actions/environment-variables#default-environment-variables) are used.

* GITHUB_EVENT_PATH
* GITHUB_RUN_ID

### GitHub Token

Basically, you can use GitHub Actions Access Token `github.token`.
But if you want to create issues in other repositories or avoid API Rate Limiting,
you have to use Personal Access Token or GitHub App's token.

The following permissions are required.

* `issues:write`

## Question (GitHub Discussions)

If you have any question, please create a Discussion. https://github.com/suzuki-shunsuke/renovate-issue-action/discussions

## Versioning Policy

[suzuki-shunsuke/versioning-policy v0.1.0](https://github.com/suzuki-shunsuke/versioning-policy/blob/v0.1.0/POLICY.md), which is compatible with [Semantic Versioning 2.0.0](https://semver.org/).

## LICENSE

[MIT](LICENSE)

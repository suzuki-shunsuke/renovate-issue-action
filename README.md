# renovate-issue-action

Manage update failures by [Renovate](https://github.com/renovatebot/renovate) as GitHub Issues

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
The issue description has a list of closed Pull Requests.
If the issue already exists, closed pull request is added to the list of closed Pull Request instead of creating a new Issue.
If a Pull Request is merged or Renovate closes it, the related issue would be closed.

## Why don't you use closed Pull Requests instead of Issues?

You could use closed Pull Requests instead of Issues without `renovate-issue-action`.
Then why do we develop `renovate-issue-action`?

We think it is inconvenient to treat closed Pull Requests as tasks.
Renovate creates a new Pull Request when a new version is released,
so there would be multiple Pull Requests against the same task.
It is difficult to understand the list of unresolved Pull Requests.

## How to use

Please add GitHub Actions Workflow and run renovate-issue-action.

e.g. [.github/workflows/close-renovate-pr.yaml](.github/workflows/close-renovate-pr.yaml)

Please update Renovate Configuration using [prBodyNotes](https://docs.renovatebot.com/configuration-options/#prbodynotes).

```json
  "prBodyNotes": [
    "<!-- renovate-issue-action: {\"packageFileDir\": \"{{packageFileDir}}\", \"packageName\": \"{{packageName}}\", \"groupName\": \"{{groupName}}\", \"depName\": \"{{depName}}\", \"manager\": \"{{manager}}\", \"updateType\": \"{{updateType}}\"} -->"
  ]
```

renovate-issue-action gets metadata from this HTML comment.

## Configuration

Configuration file is optional.
By default, the file `^\.renovate-issue-action\.ya?ml$` is read if it exists.
You can specify configuration file path with the command line option `--config`.

```console
$ renovate-issue-action --config config.yaml
```

renovate-issue-action.yaml

```yaml
renovate_login: 'renovate[bot]'
```

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

## LICENSE

[MIT](LICENSE)

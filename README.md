# renovate-issue-action

Manage update failures by [Renovate](https://github.com/renovatebot/renovate) as GitHub Issues

## :warning: Don't use this yet

This is still under development.
The development status is alpha.

## Motivation

## Install

## How to use

Add GitHub Actions Workflow and run renovate-issue-action.

e.g. [.github/workflows/close-renovate-pr.yaml](.github/workflows/close-renovate-pr.yaml)

Please update Renovate Configuration using [prBodyNotes](https://docs.renovatebot.com/configuration-options/#prbodynotes).

```json
  "prBodyNotes": [
    "<!-- {\"packageFileDir\": \"{{packageFileDir}}\", \"packageName\": \"{{packageName}}\", \"groupName\": \"{{groupName}}\", \"depName\": \"{{depName}}\", \"manager\": \"{{manager}}\", \"updateType\": \"{{updateType}}\"} -->"
  ]
```

renovate-issue-action gets metadata from this HTML comment.

## Configuration

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

### GitHub Token's permission

* `issues:write`

## LICENSE

[MIT](LICENSE)

# renovate-issue-action

Manage update failures by Renovate as GitHub Issues

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

### Environment variable

* GITHUB_TOKEN
* EVENT_PAYLOAD

### GitHub Token's permission

* issues:write

## LICENSE

[MIT](LICENSE)

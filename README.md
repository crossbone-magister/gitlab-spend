# gitlab-spend

A [Timewarrior](https://timewarrior.net/) extension that automatically reports tracked time to GitLab issues via the GitLab API.

## How it works

For each Timewarrior interval that contains a `gitlab:<project>#<iid>` tag, `gitlab-spend` calls the GitLab "add spent time" API endpoint. Already-reported intervals are tracked in a local state file to prevent duplicate submissions.

## Installation

Build from source:

```sh
go build -o gitlab-spend .
```

Place the binary in your Timewarrior extensions directory:

```sh
cp gitlab-spend ~/.timewarrior/extensions/
```

## Configuration

Add to your Timewarrior config (`~/.timewarrior/timewarrior.cfg`):

```
reports.gitlabspend.host  = gitlab.example.com
reports.gitlabspend.token = <your-personal-access-token>

# Optional: custom state file path (default: $XDG_DATA_HOME/gitlab-spend/state.json)
reports.gitlabspend.state_file = /path/to/state.json
```

The token needs the `api` scope.

## Usage

Tag your Timewarrior intervals with a GitLab issue reference:

```sh
timew track gitlab:my-group/my-project#42
```

Then run the extension:

```sh
timew report gitlab-spend
```

Intervals without a `gitlab:` tag are silently skipped.

## License

See [LICENSE](LICENSE).

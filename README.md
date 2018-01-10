# backlog-issues
backlog-issues

## Setup

- Download here https://github.com/mitakeck/backlog-issues/releases/tag/v0.0.1

```
$ wget https://github.com/mitakeck/backlog-issues/releases/download/v0.0.1/backlog-issues_darwin_amd64
$ mv backlog-issues* backlog-issues

# generate api token from  https://xxxxx.backlog.jp/EditApiSettings.action
$ export BACKLOG_TOKEN=xxxxx
$ export BACKLOG_SPACE=xxxxx
```

## Usage

```
$ backlog-issues ls
                URL                |       SUMMARY       |           DESCRIPTION          
+----------------------------------+---------------------+----------------------------------
https://xxxxx.backlog.jp/view/zzzz | issue summary 1     | hogehogehoge..........
https://xxxxx.backlog.jp/view/zzzz | issue summary 2     | hogehogehoge..........
https://xxxxx.backlog.jp/view/zzzz | issue summary 3     | hogehogehoge..........
...
```

### Usage with `peco`

```
alias row='() {awk "{print \$${1:-1}}"}'
alias bls='(){ echo $(backlog-issues ls | peco) | row 1 | xargs open}'
```

```
$ bls
-> select to open in web browser
```

# 1. 统计数据

共 {{ .ContributorsNumber }} 位作者提交了 {{ CommitNumber }} 个 Commit 。感谢以下作者的贡献：

{{ range $i, $contributor := .Contributors }}{{ $contributor }}{{ if $i  }}, {{ end }}{{ end }}

最近 2 周，共修改新增代码行 {{ .NewLine }} ，删除代码行 {{ .DelLine }} 。

# 2. 主要进展

## 2.1 新增功能

{{ range $url, $content := .Improves }}
- {{ $url }}

  {{ $content }}
{{ end }}

## 2.3 WIP

{{ range $url, $content := .WIP }}
- {{ $url }}

  {{ $content }}
{{ end }}

## 2.2 Bug 修复

{{ range $url, $content := .BugFixes }}
- {{ $url }}

  {{ $content }}
{{ end }}

## 2.3 功能改进

{{ range $url, $content := .Improvements }}
- {{ $url }}

  {{ $content }}
{{ end }}

## 2.3 其他

{{ range $url, $content := .Others }}
- {{ $url }}

  {{ $content }}
{{ end }}


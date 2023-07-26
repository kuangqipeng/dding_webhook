{{ define "alert-notify" -}}
## 告警通知
> **告警名称**: {{ .AlertNotifyName }}
>
> **告警时间**: {{ .AlertTime }}
>
> **告警信息**
>
> instance: {{ .Instance }}
>
> status: {{ .Status }}
>
> **Labels**
{{- range $k,$v := .CommonLabels }}
> - {{ $k }}: {{ $v }}
{{- end }}
>
> **Annotations**
{{- range $k,$v := .CommonAnnotations }}
> - {{ $k }}: {{ $v }}
{{- end }}
{{ end -}}
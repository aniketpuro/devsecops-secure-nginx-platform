{{- define "devsecops-secure-nginx-platform.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "devsecops-secure-nginx-platform.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- include "devsecops-secure-nginx-platform.name" . -}}
{{- end -}}
{{- end -}}

{{- define "devsecops-secure-nginx-platform.labels" -}}
app.kubernetes.io/name: {{ include "devsecops-secure-nginx-platform.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
helm.sh/chart: {{ .Chart.Name }}-{{ .Chart.Version | replace "+" "_" }}
{{- end -}}

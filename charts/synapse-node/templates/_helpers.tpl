{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "synapse-node.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "synapse-node.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "synapse-node.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Common labels
*/}}
{{- define "synapse-node.labels" -}}
app.kubernetes.io/name: {{ include "synapse-node.name" . }}
helm.sh/chart: {{ include "synapse-node.chart" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{/*
Create the name of the service account to use
*/}}
{{- define "synapse-node.serviceAccountName" -}}
{{- if .Values.serviceAccount.create -}}
    {{ default (include "synapse-node.fullname" .) .Values.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.serviceAccount.name }}
{{- end -}}
{{- end -}}

{{/*
Service constants
*/}}
{{- define "synapse-node.svcSpec" -}}
ports:
{{- range .Values.service.ports }}
    - name: {{ default "porter" .name }}
      {{- if .protocol }}
      protocol: {{ .protocol }}
      {{- end }}
      port: {{ .port }}
      targetPort: {{ .targetPort }}
{{- end }}
selector:
    app.kubernetes.io/name: {{ include "synapse-node.name" . }}
    app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{- define "synapse-node.svcLabels" -}}
{{- with .Values.service.extraLabels }}
{{- range $key, $value := . }}
{{ $key }}: {{ $value }}
{{- end -}}
{{- end }}
{{ include "synapse-node.labels" . }}
{{- end -}}

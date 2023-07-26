package main

import "time"

type Notify struct {
	Version           string                 `json:"version,omitempty"`
	GroupKey          string                 `json:"groupKey,omitempty"`
	TruncatedAlerts   int                    `json:"truncatedAlerts,omitempty"`
	Status            string                 `json:"status,omitempty"`
	Receiver          string                 `json:"receiver,omitempty"`
	GroupLabels       map[string]string      `json:"groupLabels,omitempty"`
	CommonLabels      map[string]interface{} `json:"commonLabels,omitempty"`
	CommonAnnotations map[string]interface{} `json:"commonAnnotations,omitempty"`
	ExternalURL       string                 `json:"externalURL,omitempty"`
	Alerts            []Alert                `json:"alerts,omitempty"`
}

type Alert struct {
	Status       string                 `json:"status,omitempty"`
	Labels       map[string]interface{} `json:"labels,omitempty"`
	Annotations  map[string]interface{} `json:"annotations,omitempty"`
	StartsAt     time.Time              `json:"startsAt,omitempty"`
	EndsAt       time.Time              `json:"endsAt,omitempty"`
	GeneratorURL string                 `json:"generatorURL,omitempty"`
	Fingerprint  string                 `json:"fingerprint,omitempty"`
}

type NotifyTpl struct {
	AlertNotifyName   string
	AlertTime         string
	Instance          string
	Status            string
	CommonLabels      map[string]interface{}
	CommonAnnotations map[string]interface{}
}

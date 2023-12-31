package schemata 

{{- $operationGroup := pascalize .Name -}}
{{- $hasID := false -}}

{{/* Some property fields can be characterized as more "complex" objects whose data must be retrieved with specialized helper functions */}}
{{- $isReadOnlyModel := true -}}
{{- $needsUtils := false -}}
{{- range .Properties -}}
	{{- if not .ReadOnly -}} {{- $isReadOnlyModel = false -}} {{- end -}}
	{{- if eq .Name "customProperties" -}} {{- $needsUtils = true -}} {{- end -}}
	{{- if eq .Name "id" -}} {{- $hasID = true -}} {{- end -}}
{{- end }}

import (
	"vestahealthcare/models"
	{{- if $needsUtils }}
	"vestahealthcare/salesforce/utils"
	{{- end }}
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the {{ $operationGroup }} resource defined in the Terraform configuration
func {{ $operationGroup }}Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		{{- range .Properties }}
		"{{ snakize .Name }}": {
				{{- if (eq (camelize .Name) "id") }}
			Type: schema.TypeString,
			Computed: true,
				{{- else }}
					{{- if eq .GoType "string" }}
			Type: schema.TypeString,
						{{- if .Default }}
			Default: "{{ .Default }}",
						{{- end }}
					{{- else if eq .GoType "bool" }}
			Type: schema.TypeBool,
					{{- if .Default }}
			Default: {{ .Default }},
						{{- end }}
					{{- else if eq .GoType "int64" }}
			Type: schema.TypeInt,
					{{- else if eq .GoType "int32" }}
			Type: schema.TypeInt,
					{{- else if eq .GoType "interface{}" }}
			Type: schema.TypeMap, //GoType: {{ .GoType }}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
					{{- else }}
			Type: schema.TypeList, //GoType: {{ .GoType }}
						{{- if and .Items .Items.HasAdditionalProperties }}
			Elem: &schema.Schema{
				Type: schema.TypeMap,
				Description: "{{ .Description }}",
				Elem: &schema.Schema{
					Type: schema.Type{{ with .Items.AdditionalProperties }}{{ pascalize .GoType }}{{ end }},
				},
			},
						{{- else if stringContains .GoType "[]*" }} 
			Elem: &schema.Resource{
				Schema: {{ .Items.GoType }}Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
						{{- else }}
			Elem: &schema.Resource{
				Schema: {{ .GoType }}Schema(),
			},
						{{- end }}
					{{- end }}
					{{- if and .Required (not .ReadOnly) }}
			Required: true,
					{{- else if .ReadOnly }}
			Computed: true,
					{{- else }}
			Optional: true,
					{{- end }}
				{{- end }}
		},
		{{ end }}
	}
}

{{/* Only resources have respective data sources that need to be mapped to an appropriate schema */}}
{{- if eq .Example "isResource" }}
// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and {{ $operationGroup }}Schema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSource{{ $operationGroup }}Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		{{- range .Properties }}
		"{{ snakize .Name }}": {
				{{- if (eq .Name "id") }}
			Type: schema.TypeString,
			Computed: true,
			Optional: true,
				{{- else }}
					{{- if eq .GoType "string" }}
			Type: schema.TypeString,
						{{- if .Default }}
			Default: "{{ .Default }}",
						{{- end }}
					{{- else if eq .GoType "bool" }}
			Type: schema.TypeBool,
					{{- if .Default }}
			Default: {{ .Default }},
						{{- end }}
					{{- else if eq .GoType "int64" }}
			Type: schema.TypeInt,
					{{- else if eq .GoType "int32" }}
			Type: schema.TypeInt,
					{{- else if eq .GoType "interface{}" }}
			Type: schema.TypeMap, //GoType: {{ .GoType }}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
					{{- else }}
			Type: schema.TypeList, //GoType: {{ .GoType }}
						{{- if and .Items .Items.HasAdditionalProperties }}
			Elem: &schema.Schema{
				Type: schema.TypeMap,
				Description: "{{ .Description }}",
				Elem: &schema.Schema{
					Type: schema.Type{{ with .Items.AdditionalProperties }}{{ pascalize .GoType }}{{ end }},
				},
			},
						{{- else if stringContains .GoType "[]*" }} 
			Elem: &schema.Resource{
				Schema: {{ .Items.GoType }}Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
						{{- else }}
			Elem: &schema.Resource{
				Schema: {{ .GoType }}Schema(),
			},
						{{- end }}
					{{- end }}
			Optional: true,
			Computed: true,
				{{- end }}
		},
		{{ end }}
		"filter": { {{/* all of LM's GetList endpoints support a filter parameter for, well, filtering! Check our endpoint filtering documentation below */}}
			Type:     schema.TypeString, {{/* https://www.salesforce.com/support/rest-api-developers-guide/v1/devices/get-devices */}}
            Optional: true,
		},
	}
}
{{- end }}

// Update the underlying {{ $operationGroup }} resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func Set{{ $operationGroup }}ResourceData(d *schema.ResourceData, m *models.{{ $operationGroup }}, isDataResource bool) {
	{{- $hasId := 0 -}}
	{{- range .Properties }}{{ if eq (snakize .Name) "id" }}{{ $hasId = 1 }}{{ end }}{{ end -}}
	{{- if eq $hasId 0 }}
	if isDataResource {
		d.SetId("-")
	}
	{{- end -}}
	{{- range .Properties }}
		{{- if eq (snakize .Name) "id" }}
			{{- if not (eq .SwaggerType "string") }}
	d.SetId(strconv.Itoa(int(m.ID)))
			{{- else }}
	if m.ID == "" && isDataResource {
		d.SetId("-")
	} else {
		d.SetId(m.ID)
	}
			{{- end }}
		{{- else if or .IsComplexObject }}
	d.Set("{{ snakize .Name }}", Set{{ pascalize .GoType }}SubResourceData([]*models.{{ pascalize .GoType }}{m.{{ pascalize .Name }}}))
		{{- else if stringContains .GoType "[]*" }}
	d.Set("{{ snakize .Name }}", Set{{ pascalize .Items.GoType }}SubResourceData(m.{{ pascalize .Name }}))
		{{- else }}
	d.Set("{{ snakize .Name }}", m.{{ pascalize .Name }})
		{{- end }}
	{{- end }}
}

// Iterate throught and update the {{ $operationGroup }} resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func Set{{ $operationGroup }}SubResourceData(m []*models.{{ $operationGroup }}) (d []*map[string]interface{}) {
	{{- $model := camelize $operationGroup }}
	for _, {{ $model }} := range m {
		if {{ $model }} != nil {
			properties := make(map[string]interface{})
			{{- range .Properties }}
				{{- if or .IsComplexObject }}
			properties["{{snakize .Name}}"] = Set{{ pascalize .GoType }}SubResourceData([]*models.{{ pascalize .GoType }}{ {{- $model }}.{{ pascalize .Name -}} })
				{{- else if stringContains .GoType "[]*" }}
			properties["{{snakize .Name}}"] = Set{{ pascalize .Items.GoType }}SubResourceData({{ $model }}.{{ pascalize .Name }})
				{{- else }}
			properties["{{ snakize .Name }}"] = {{ $model }}.{{ pascalize .Name }}
				{{- end }}
			{{- end }}
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate {{ $operationGroup }} resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func {{ $operationGroup }}Model(d *schema.ResourceData) *models.{{ $operationGroup }} {
	{{- range .Properties }}
		{{- if or (not .ReadOnly) (and (eq .Name "id") (not $isReadOnlyModel)) }}
			{{- if and (eq .Name "id") (not (eq .SwaggerType "string")) }}
	id, _ := strconv.Atoi(d.Get("id").(string))
			{{- else if .IsComplexObject }}
	var {{ varname .Name }} *models.{{ pascalize .GoType }} = nil//hit complex
	{{ .Name }}Interface, {{ .Name }}IsSet := d.GetOk("{{ snakize .Name }}")
	if {{ .Name }}IsSet {
		{{ .Name }}Map := {{ .Name }}Interface.([]interface{})[0].(map[string]interface{})
		{{ varname .Name }} = {{ .GoType }}ModelFromMap({{ .Name }}Map)
	}
			{{- else if stringContains .Name "Properties" }}
	{{ varname .Name }} := utils.GetPropertiesFromResource(d, "{{ snakize .Name }}")
			{{- else if eq .GoType "interface{}" }}
	{{ varname .Name }} := d.Get("{{ snakize .Name }}")
			{{- else if or (eq .GoType "int32") ( eq .GoType "int64") }}
	{{ varname .Name }} := {{ .GoType }}(d.Get("{{ snakize .Name }}").(int))
			{{- else if and .Items .Items.HasAdditionalProperties }}
	{{ varname .Name }} := d.Get("{{ snakize .Name }}").([]{{ .Items.GoType }})
			{{- else if and (not (eq .GoType "string")) (not (eq .GoType "[]string")) (not (eq .GoType "bool")) (not (eq .GoType "int")) (not (eq .GoType "float32")) (not (eq .GoType "float64")) (not (eq .GoType "uint32")) (not (eq .GoType "uint64")) }}
	{{ varname .Name }} := d.Get("{{ snakize .Name }}").({{ if hasPrefix .GoType "[]" }}[]{{ end }}*models.{{ pascalize .GoType }})
			{{- else }}
	{{ varname .Name }} := d.Get("{{ snakize .Name }}").({{ .GoType }})
			{{- end }}
		{{- end }}
	{{- end }}
	
	return &models.{{ $operationGroup }} {
		{{- if not $isReadOnlyModel }}
		{{- range .Properties }}
			{{- if or (not .ReadOnly) (eq .Name "id") }}
				{{- if and (eq .Name "id") (not (eq .SwaggerType "string")) }}
		{{ pascalize .Name }}: int32({{ varname .Name }}),
				{{- else if and (and (not .IsMap) .IsNullable (not .IsSuperAlias)) (not .IsComplexObject) }}
		{{ pascalize .Name }}: &{{ varname .Name }},
				{{- else }}
		{{ pascalize .Name }}: {{ varname .Name }},
				{{- end }}
			{{- end }}
		{{- end }}
		{{- end }}
	}
}

// Function to perform the following actions:
func {{ $operationGroup }}ModelFromMap(m map[string]interface{}) *models.{{ $operationGroup }} {
	{{- range .Properties }}
		{{- if or (not .ReadOnly) (and (eq .Name "id") (not $isReadOnlyModel)) }}
			{{- if and (eq .Name "id") (not (eq .SwaggerType "string")) }}
	id, _ := strconv.Atoi(m["id"].(string))
			{{- else if .IsComplexObject }}
	var {{ varname .Name }} *models.{{ pascalize .GoType }} = nil//hit complex
	if m["{{ snakize .Name }}"] != nil {
		{{ varname .Name }} = {{ .GoType }}ModelFromMap(m["{{ snakize .Name }}"].(map[string]interface{}))
	}
			{{- else if stringContains .Name "Properties" }}
	{{ varname .Name }} := utils.GetPropertiesFromResource(d, "{{ snakize .Name }}")
			{{- else if eq .GoType "interface{}" }}
	{{ varname .Name }} := m["{{ snakize .Name }}"]
			{{- else if or (eq .GoType "int32") ( eq .GoType "int64") }}
	{{ varname .Name }} := {{ .GoType }}(m["{{ snakize .Name }}"].(int))
			{{- else if and .Items .Items.HasAdditionalProperties }}
	{{ varname .Name }} := m["{{ snakize .Name }}"].([]{{ .Items.GoType }})
			{{- else if and (not (eq .GoType "string")) (not (eq .GoType "[]string")) (not (eq .GoType "bool")) (not (eq .GoType "int")) (not (eq .GoType "float32")) (not (eq .GoType "float64")) (not (eq .GoType "uint32")) (not (eq .GoType "uint64")) }}
	var {{ varname .Name }} {{ if hasPrefix .GoType "[]" }}[]{{ end }}*models.{{ pascalize .GoType }} = nil//hit complex
	if m["{{ snakize .Name }}"] != nil {
		{{ varname .Name }} = {{ pascalize .GoType }}ModelFrom{{ if hasPrefix .GoType "[]" }}ArrayOf{{ end }}Map(m["{{ snakize .Name }}"].({{ if hasPrefix .GoType "[]" }}[]{{ else }}map[string]{{ end }}interface{}))
	}
			{{- else }}
	{{ varname .Name }} := m["{{ snakize .Name }}"].({{ .GoType }})
			{{- end }}
		{{- end }}
	{{- end }}
	
	return &models.{{ $operationGroup }} {
		{{- if not $isReadOnlyModel }}
		{{- range .Properties }}
			{{- if or (not .ReadOnly) (eq .Name "id") }}
				{{- if and (eq .Name "id") (not (eq .SwaggerType "string")) }}
		{{ pascalize .Name }}: int32({{ varname .Name }}),
				{{- else if and (and (not .IsMap) .IsNullable (not .IsSuperAlias)) (not .IsComplexObject) }}
		{{ pascalize .Name }}: &{{ varname .Name }},
				{{- else }}
		{{ pascalize .Name }}: {{ varname .Name }},
				{{- end }}
			{{- end }}
		{{- end }}
		{{- end }}
	}
}

func {{ $operationGroup }}ModelFromArrayOfMap(m []interface{}) []*models.{{ $operationGroup }} {
	mapped := make([]*models.{{ $operationGroup }}, len(m))
  for i, v := range m {
    mapped[i] = {{ $operationGroup }}ModelFromMap(v.(map[string]interface{}))
  }
  return mapped
}

// Retrieve property field names for updating the {{ $operationGroup }} resource 
func Get{{ $operationGroup }}PropertyFields() (t []string) {
	return []string{
		{{- range .Properties }}
			{{- if or (not .ReadOnly) (eq .Name "id") }}
		"{{ snakize .Name }}",
			{{- end }}
		{{- end }}
	}
}
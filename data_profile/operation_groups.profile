/*
&generator.GenOperationGroup{GenCommon:generator.GenCommon{Copyright:"", TargetImportPath:"vestahealthcare"}, Name:"query_result", Operations:generator.GenOperations{generator.GenOperation{GenCommon:generator.GenCommon{Copyright:"", TargetImportPath:"vestahealthcare"}, Package:"query_result", ReceiverName:"o", Name:"getQuery", Summary:"Get QueryResult", Description:"", Method:"GET", Path:"/query", BasePath:"/services/data/v57.0", Tags:[]string{"QueryResult"}, UseTags:true, RootPackage:"restapi", Imports:map[string]string{"client":"vestahealthcare/client"}, DefaultImports:map[string]string{"event_relay_config":"vestahealthcare/client/event_relay_config", "event_relay_feedback":"vestahealthcare/client/event_relay_feedback", "models":"vestahealthcare/models", "named_credential":"vestahealthcare/client/named_credential", "platform_event_channel":"vestahealthcare/client/platform_event_channel", "platform_event_channel_member":"vestahealthcare/client/platform_event_channel_member", "query_result":"vestahealthcare/client/query_result"}, ExtraSchemas:generator.GenSchemaList{}, PackageAlias:"query_result", Authorized:true, Security:[]generator.GenSecurityRequirements{generator.GenSecurityRequirements{generator.GenSecurityRequirement{Name:"oauth2_password", Scopes:[]string{}}}}, SecurityDefinitions:generator.GenSecuritySchemes{generator.GenSecurityScheme{AppName:"getQuery", ID:"oauth2_password", Name:"", ReceiverName:"o", IsBasicAuth:false, IsAPIKeyAuth:false, IsOAuth2:true, Scopes:[]string{}, Source:"", Principal:"interface{}", PrincipalIsNullable:false, Description:"", Type:"oauth2", In:"", Flow:"password", AuthorizationURL:"", TokenURL:"https://vestahealthcare--uat.sandbox.my.salesforce.com/services/oauth2/token", Extensions:map[string]interface {}(nil), ScopesDesc:[]generator.GenSecurityScope{}}}, SecurityRequirements:[]analysis.SecurityRequirement{analysis.SecurityRequirement{Name:"oauth2_password", Scopes:[]string{}}}, Principal:"interface{}", PrincipalIsNullable:false, SuccessResponse:(*generator.GenResponse)(0x14000e273b0), SuccessResponses:[]generator.GenResponse{generator.GenResponse{Package:"query_result", ModelsPackage:"models", ReceiverName:"o", Name:"getQueryOK", Description:"successful operation", IsSuccess:true, Code:200, Method:"GET", Path:"/query", Headers:generator.GenHeaders(nil), Schema:(*generator.GenSchema)(0x140030daa80), AllowsForStreaming:false, Imports:map[string]string{"client":"vestahealthcare/client"}, DefaultImports:map[string]string{"event_relay_config":"vestahealthcare/client/event_relay_config", "event_relay_feedback":"vestahealthcare/client/event_relay_feedback", "models":"vestahealthcare/models", "named_credential":"vestahealthcare/client/named_credential", "platform_event_channel":"vestahealthcare/client/platform_event_channel", "platform_event_channel_member":"vestahealthcare/client/platform_event_channel_member", "query_result":"vestahealthcare/client/query_result"}, Extensions:map[string]interface {}(nil), StrictResponders:false, OperationName:"getQuery", Examples:generator.GenResponseExamples{}}}, Responses:generator.GenStatusCodeResponses{generator.GenResponse{Package:"query_result", ModelsPackage:"models", ReceiverName:"o", Name:"getQueryOK", Description:"successful operation", IsSuccess:true, Code:200, Method:"GET", Path:"/query", Headers:generator.GenHeaders(nil), Schema:(*generator.GenSchema)(0x140030daa80), AllowsForStreaming:false, Imports:map[string]string{"client":"vestahealthcare/client"}, DefaultImports:map[string]string{"event_relay_config":"vestahealthcare/client/event_relay_config", "event_relay_feedback":"vestahealthcare/client/event_relay_feedback", "models":"vestahealthcare/models", "named_credential":"vestahealthcare/client/named_credential", "platform_event_channel":"vestahealthcare/client/platform_event_channel", "platform_event_channel_member":"vestahealthcare/client/platform_event_channel_member", "query_result":"vestahealthcare/client/query_result"}, Extensions:map[string]interface {}(nil), StrictResponders:false, OperationName:"getQuery", Examples:generator.GenResponseExamples{}}}, DefaultResponse:(*generator.GenResponse)(0x14000e27770), Params:generator.GenParameters{generator.GenParameter{resolvedType:generator.resolvedType{IsAnonymous:false, IsArray:false, IsMap:false, IsInterface:false, IsPrimitive:true, IsCustomFormatter:false, IsAliased:false, IsNullable:false, IsStream:false, IsEmptyOmitted:false, IsJSONString:false, IsEnumCI:false, IsBase64:false, IsExternal:false, IsTuple:false, HasAdditionalItems:false, IsComplexObject:false, IsBaseType:false, HasDiscriminator:false, GoType:"string", Pkg:"", PkgAlias:"", AliasedType:"", SwaggerType:"string", SwaggerFormat:"", Extensions:spec.Extensions(nil), ElemType:(*generator.resolvedType)(nil), IsMapNullOverride:false, IsSuperAlias:false, IsEmbedded:false, SkipExternalValidation:false}, sharedValidations:generator.sharedValidations{SchemaValidations:spec.SchemaValidations{CommonValidations:spec.CommonValidations{Maximum:(*float64)(nil), ExclusiveMaximum:false, Minimum:(*float64)(nil), ExclusiveMinimum:false, MaxLength:(*int64)(nil), MinLength:(*int64)(nil), Pattern:"", MaxItems:(*int64)(nil), MinItems:(*int64)(nil), UniqueItems:false, MultipleOf:(*float64)(nil), Enum:[]interface {}(nil)}, PatternProperties:spec.SchemaProperties(nil), MaxProperties:(*int64)(nil), MinProperties:(*int64)(nil)}, HasValidations:false, HasContextValidations:false, Required:true, HasSliceValidations:false, ItemsEnum:[]interface {}(nil)}, ID:"Q", Name:"q", ModelsPackage:"models", Path:"\"q\"", ValueExpression:"o.Q", IndexVar:"i", KeyVar:"", ReceiverName:"o", Location:"query", Title:"", Description:"A SOQL query that returns data", Converter:"", Formatter:"", Schema:(*generator.GenSchema)(nil), CollectionFormat:"", Child:(*generator.GenItems)(nil), Parent:(*generator.GenItems)(nil), Default:interface {}(nil), HasDefault:false, ZeroValue:"\"\"", AllowEmptyValue:false, HasSimpleBodyParams:false, HasModelBodyParams:false, HasSimpleBodyItems:false, HasModelBodyItems:false, HasSimpleBodyMap:false, HasModelBodyMap:false, Extensions:map[string]interface {}(nil)}}, QueryParams:generator.GenParameters{generator.GenParameter{resolvedType:generator.resolvedType{IsAnonymous:false, IsArray:false, IsMap:false, IsInterface:false, IsPrimitive:true, IsCustomFormatter:false, IsAliased:false, IsNullable:false, IsStream:false, IsEmptyOmitted:false, IsJSONString:false, IsEnumCI:false, IsBase64:false, IsExternal:false, IsTuple:false, HasAdditionalItems:false, IsComplexObject:false, IsBaseType:false, HasDiscriminator:false, GoType:"string", Pkg:"", PkgAlias:"", AliasedType:"", SwaggerType:"string", SwaggerFormat:"", Extensions:spec.Extensions(nil), ElemType:(*generator.resolvedType)(nil), IsMapNullOverride:false, IsSuperAlias:false, IsEmbedded:false, SkipExternalValidation:false}, sharedValidations:generator.sharedValidations{SchemaValidations:spec.SchemaValidations{CommonValidations:spec.CommonValidations{Maximum:(*float64)(nil), ExclusiveMaximum:false, Minimum:(*float64)(nil), ExclusiveMinimum:false, MaxLength:(*int64)(nil), MinLength:(*int64)(nil), Pattern:"", MaxItems:(*int64)(nil), MinItems:(*int64)(nil), UniqueItems:false, MultipleOf:(*float64)(nil), Enum:[]interface {}(nil)}, PatternProperties:spec.SchemaProperties(nil), MaxProperties:(*int64)(nil), MinProperties:(*int64)(nil)}, HasValidations:false, HasContextValidations:false, Required:true, HasSliceValidations:false, ItemsEnum:[]interface {}(nil)}, ID:"Q", Name:"q", ModelsPackage:"models", Path:"\"q\"", ValueExpression:"o.Q", IndexVar:"i", KeyVar:"", ReceiverName:"o", Location:"query", Title:"", Description:"A SOQL query that returns data", Converter:"", Formatter:"", Schema:(*generator.GenSchema)(nil), CollectionFormat:"", Child:(*generator.GenItems)(nil), Parent:(*generator.GenItems)(nil), Default:interface {}(nil), HasDefault:false, ZeroValue:"\"\"", AllowEmptyValue:false, HasSimpleBodyParams:false, HasModelBodyParams:false, HasSimpleBodyItems:false, HasModelBodyItems:false, HasSimpleBodyMap:false, HasModelBodyMap:false, Extensions:map[string]interface {}(nil)}}, PathParams:generator.GenParameters(nil), HeaderParams:generator.GenParameters(nil), FormParams:generator.GenParameters(nil), HasQueryParams:true, HasPathParams:false, HasHeaderParams:false, HasFormParams:false, HasFormValueParams:false, HasFileParams:false, HasBodyParams:false, HasStreamingResponse:false, Schemes:[]string{"https"}, ExtraSchemes:[]string{}, SchemeOverrides:[]string(nil), ExtraSchemeOverrides:[]string(nil), ProducesMediaTypes:[]string{"application/json"}, ConsumesMediaTypes:[]string{"application/json"}, TimeoutName:"timeout", Extensions:map[string]interface {}(nil), StrictResponders:false, ExternalDocs:(*spec.ExternalDocumentation)(nil), Produces:[]string{"application/json"}, Consumes:[]string(nil)}}, Summary:"", Description:"", Imports:map[string]string{"client":"vestahealthcare/client"}, DefaultImports:map[string]string{"event_relay_config":"vestahealthcare/client/event_relay_config", "event_relay_feedback":"vestahealthcare/client/event_relay_feedback", "models":"vestahealthcare/models", "named_credential":"vestahealthcare/client/named_credential", "platform_event_channel":"vestahealthcare/client/platform_event_channel", "platform_event_channel_member":"vestahealthcare/client/platform_event_channel_member", "query_result":"vestahealthcare/client/query_result"}, RootPackage:"operations", GenOpts:(*generator.GenOpts)(0x14000164580), PackageAlias:"query_result"}
*/

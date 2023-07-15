/*
&generator.GenOperation{GenCommon:generator.GenCommon{Copyright:"", TargetImportPath:"vestahealthcare"}, Package:"platform_event_channel_member", ReceiverName:"o", Name:"updatePlatformEventChannelMember", Summary:"Update PlatformEventChannelMember", Description:"", Method:"PATCH", Path:"/PlatformEventChannelMember/{Id}", BasePath:"/services/data/v57.0/tooling/sobjects", Tags:[]string{"PlatformEventChannelMember"}, UseTags:true, RootPackage:"restapi", Imports:map[string]string{"client":"vestahealthcare/client"}, DefaultImports:map[string]string{"event_relay_config":"vestahealthcare/client/event_relay_config", "models":"vestahealthcare/models", "named_credential":"vestahealthcare/client/named_credential", "platform_event_channel":"vestahealthcare/client/platform_event_channel", "platform_event_channel_member":"vestahealthcare/client/platform_event_channel_member"}, ExtraSchemas:generator.GenSchemaList{}, PackageAlias:"platform_event_channel_member", Authorized:true, Security:[]generator.GenSecurityRequirements{generator.GenSecurityRequirements{generator.GenSecurityRequirement{Name:"oauth2_password", Scopes:[]string{}}}}, SecurityDefinitions:generator.GenSecuritySchemes{generator.GenSecurityScheme{AppName:"updatePlatformEventChannelMember", ID:"oauth2_password", Name:"", ReceiverName:"o", IsBasicAuth:false, IsAPIKeyAuth:false, IsOAuth2:true, Scopes:[]string{}, Source:"", Principal:"interface{}", PrincipalIsNullable:false, Description:"", Type:"oauth2", In:"", Flow:"password", AuthorizationURL:"", TokenURL:"https://vestahealthcare--uat.sandbox.my.salesforce.com/services/oauth2/token", Extensions:map[string]interface {}(nil), ScopesDesc:[]generator.GenSecurityScope{}}}, SecurityRequirements:[]analysis.SecurityRequirement{analysis.SecurityRequirement{Name:"oauth2_password", Scopes:[]string{}}}, Principal:"interface{}", PrincipalIsNullable:false, SuccessResponse:(*generator.GenResponse)(0x140029d2a50), SuccessResponses:[]generator.GenResponse{generator.GenResponse{Package:"platform_event_channel_member", ModelsPackage:"models", ReceiverName:"o", Name:"updatePlatformEventChannelMemberNoContent", Description:"successful operation", IsSuccess:true, Code:204, Method:"PATCH", Path:"/PlatformEventChannelMember/{Id}", Headers:generator.GenHeaders(nil), Schema:(*generator.GenSchema)(nil), AllowsForStreaming:false, Imports:map[string]string{"client":"vestahealthcare/client"}, DefaultImports:map[string]string{"event_relay_config":"vestahealthcare/client/event_relay_config", "models":"vestahealthcare/models", "named_credential":"vestahealthcare/client/named_credential", "platform_event_channel":"vestahealthcare/client/platform_event_channel", "platform_event_channel_member":"vestahealthcare/client/platform_event_channel_member"}, Extensions:map[string]interface {}(nil), StrictResponders:false, OperationName:"updatePlatformEventChannelMember", Examples:generator.GenResponseExamples{}}}, Responses:generator.GenStatusCodeResponses{generator.GenResponse{Package:"platform_event_channel_member", ModelsPackage:"models", ReceiverName:"o", Name:"updatePlatformEventChannelMemberNoContent", Description:"successful operation", IsSuccess:true, Code:204, Method:"PATCH", Path:"/PlatformEventChannelMember/{Id}", Headers:generator.GenHeaders(nil), Schema:(*generator.GenSchema)(nil), AllowsForStreaming:false, Imports:map[string]string{"client":"vestahealthcare/client"}, DefaultImports:map[string]string{"event_relay_config":"vestahealthcare/client/event_relay_config", "models":"vestahealthcare/models", "named_credential":"vestahealthcare/client/named_credential", "platform_event_channel":"vestahealthcare/client/platform_event_channel", "platform_event_channel_member":"vestahealthcare/client/platform_event_channel_member"}, Extensions:map[string]interface {}(nil), StrictResponders:false, OperationName:"updatePlatformEventChannelMember", Examples:generator.GenResponseExamples{}}}, DefaultResponse:(*generator.GenResponse)(0x140029d3950), Params:generator.GenParameters{generator.GenParameter{resolvedType:generator.resolvedType{IsAnonymous:false, IsArray:false, IsMap:false, IsInterface:false, IsPrimitive:true, IsCustomFormatter:false, IsAliased:false, IsNullable:false, IsStream:false, IsEmptyOmitted:false, IsJSONString:false, IsEnumCI:false, IsBase64:false, IsExternal:false, IsTuple:false, HasAdditionalItems:false, IsComplexObject:false, IsBaseType:false, HasDiscriminator:false, GoType:"string", Pkg:"", PkgAlias:"", AliasedType:"", SwaggerType:"string", SwaggerFormat:"", Extensions:spec.Extensions(nil), ElemType:(*generator.resolvedType)(nil), IsMapNullOverride:false, IsSuperAlias:false, IsEmbedded:false, SkipExternalValidation:false}, sharedValidations:generator.sharedValidations{SchemaValidations:spec.SchemaValidations{CommonValidations:spec.CommonValidations{Maximum:(*float64)(nil), ExclusiveMaximum:false, Minimum:(*float64)(nil), ExclusiveMinimum:false, MaxLength:(*int64)(nil), MinLength:(*int64)(nil), Pattern:"", MaxItems:(*int64)(nil), MinItems:(*int64)(nil), UniqueItems:false, MultipleOf:(*float64)(nil), Enum:[]interface {}(nil)}, PatternProperties:spec.SchemaProperties(nil), MaxProperties:(*int64)(nil), MinProperties:(*int64)(nil)}, HasValidations:false, HasContextValidations:false, Required:true, HasSliceValidations:false, ItemsEnum:[]interface {}(nil)}, ID:"ID", Name:"Id", ModelsPackage:"models", Path:"\"Id\"", ValueExpression:"o.ID", IndexVar:"i", KeyVar:"", ReceiverName:"o", Location:"path", Title:"", Description:"", Converter:"", Formatter:"", Schema:(*generator.GenSchema)(nil), CollectionFormat:"", Child:(*generator.GenItems)(nil), Parent:(*generator.GenItems)(nil), Default:interface {}(nil), HasDefault:false, ZeroValue:"\"\"", AllowEmptyValue:false, HasSimpleBodyParams:false, HasModelBodyParams:false, HasSimpleBodyItems:false, HasModelBodyItems:false, HasSimpleBodyMap:false, HasModelBodyMap:false, Extensions:map[string]interface {}(nil)}, generator.GenParameter{resolvedType:generator.resolvedType{IsAnonymous:false, IsArray:false, IsMap:false, IsInterface:false, IsPrimitive:false, IsCustomFormatter:false, IsAliased:false, IsNullable:true, IsStream:false, IsEmptyOmitted:true, IsJSONString:false, IsEnumCI:false, IsBase64:false, IsExternal:false, IsTuple:false, HasAdditionalItems:false, IsComplexObject:true, IsBaseType:false, HasDiscriminator:false, GoType:"models.PlatformEventChannelMemberUpdateRequest", Pkg:"platform_event_channel_member", PkgAlias:"", AliasedType:"", SwaggerType:"object", SwaggerFormat:"", Extensions:spec.Extensions(nil), ElemType:(*generator.resolvedType)(nil), IsMapNullOverride:false, IsSuperAlias:false, IsEmbedded:false, SkipExternalValidation:false}, sharedValidations:generator.sharedValidations{SchemaValidations:spec.SchemaValidations{CommonValidations:spec.CommonValidations{Maximum:(*float64)(nil), ExclusiveMaximum:false, Minimum:(*float64)(nil), ExclusiveMinimum:false, MaxLength:(*int64)(nil), MinLength:(*int64)(nil), Pattern:"", MaxItems:(*int64)(nil), MinItems:(*int64)(nil), UniqueItems:false, MultipleOf:(*float64)(nil), Enum:[]interface {}(nil)}, PatternProperties:spec.SchemaProperties(nil), MaxProperties:(*int64)(nil), MinProperties:(*int64)(nil)}, HasValidations:false, HasContextValidations:true, Required:true, HasSliceValidations:false, ItemsEnum:[]interface {}(nil)}, ID:"Body", Name:"body", ModelsPackage:"models", Path:"\"body\"", ValueExpression:"o.Body", IndexVar:"i", KeyVar:"k", ReceiverName:"o", Location:"body", Title:"", Description:"", Converter:"", Formatter:"", Schema:(*generator.GenSchema)(0x14002323500), CollectionFormat:"", Child:(*generator.GenItems)(0x140009b8a00), Parent:(*generator.GenItems)(nil), Default:interface {}(nil), HasDefault:false, ZeroValue:"new(models.PlatformEventChannelMemberUpdateRequest)", AllowEmptyValue:false, HasSimpleBodyParams:false, HasModelBodyParams:true, HasSimpleBodyItems:false, HasModelBodyItems:false, HasSimpleBodyMap:false, HasModelBodyMap:false, Extensions:map[string]interface {}(nil)}}, QueryParams:generator.GenParameters(nil), PathParams:generator.GenParameters{generator.GenParameter{resolvedType:generator.resolvedType{IsAnonymous:false, IsArray:false, IsMap:false, IsInterface:false, IsPrimitive:true, IsCustomFormatter:false, IsAliased:false, IsNullable:false, IsStream:false, IsEmptyOmitted:false, IsJSONString:false, IsEnumCI:false, IsBase64:false, IsExternal:false, IsTuple:false, HasAdditionalItems:false, IsComplexObject:false, IsBaseType:false, HasDiscriminator:false, GoType:"string", Pkg:"", PkgAlias:"", AliasedType:"", SwaggerType:"string", SwaggerFormat:"", Extensions:spec.Extensions(nil), ElemType:(*generator.resolvedType)(nil), IsMapNullOverride:false, IsSuperAlias:false, IsEmbedded:false, SkipExternalValidation:false}, sharedValidations:generator.sharedValidations{SchemaValidations:spec.SchemaValidations{CommonValidations:spec.CommonValidations{Maximum:(*float64)(nil), ExclusiveMaximum:false, Minimum:(*float64)(nil), ExclusiveMinimum:false, MaxLength:(*int64)(nil), MinLength:(*int64)(nil), Pattern:"", MaxItems:(*int64)(nil), MinItems:(*int64)(nil), UniqueItems:false, MultipleOf:(*float64)(nil), Enum:[]interface {}(nil)}, PatternProperties:spec.SchemaProperties(nil), MaxProperties:(*int64)(nil), MinProperties:(*int64)(nil)}, HasValidations:false, HasContextValidations:false, Required:true, HasSliceValidations:false, ItemsEnum:[]interface {}(nil)}, ID:"ID", Name:"Id", ModelsPackage:"models", Path:"\"Id\"", ValueExpression:"o.ID", IndexVar:"i", KeyVar:"", ReceiverName:"o", Location:"path", Title:"", Description:"", Converter:"", Formatter:"", Schema:(*generator.GenSchema)(nil), CollectionFormat:"", Child:(*generator.GenItems)(nil), Parent:(*generator.GenItems)(nil), Default:interface {}(nil), HasDefault:false, ZeroValue:"\"\"", AllowEmptyValue:false, HasSimpleBodyParams:false, HasModelBodyParams:false, HasSimpleBodyItems:false, HasModelBodyItems:false, HasSimpleBodyMap:false, HasModelBodyMap:false, Extensions:map[string]interface {}(nil)}}, HeaderParams:generator.GenParameters(nil), FormParams:generator.GenParameters(nil), HasQueryParams:false, HasPathParams:true, HasHeaderParams:false, HasFormParams:false, HasFormValueParams:false, HasFileParams:false, HasBodyParams:true, HasStreamingResponse:false, Schemes:[]string{"https"}, ExtraSchemes:[]string{}, SchemeOverrides:[]string(nil), ExtraSchemeOverrides:[]string(nil), ProducesMediaTypes:[]string{"application/json"}, ConsumesMediaTypes:[]string{"application/json"}, TimeoutName:"timeout", Extensions:map[string]interface {}(nil), StrictResponders:false, ExternalDocs:(*spec.ExternalDocumentation)(nil), Produces:[]string{"application/json"}, Consumes:[]string(nil)}
*/

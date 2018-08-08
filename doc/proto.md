# Protocol Documentation
<a name="top"/>

## Table of Contents

- [analyze.proto](#analyze.proto)
    - [AnalyzeApiRequest](#types.AnalyzeApiRequest)
    - [AnalyzeRequest](#types.AnalyzeRequest)
    - [AnalyzeResponse](#types.AnalyzeResponse)
  
  
  
    - [AnalyzeService](#types.AnalyzeService)
  

- [anonymize.proto](#anonymize.proto)
    - [AnonymizeApiRequest](#types.AnonymizeApiRequest)
    - [AnonymizeRequest](#types.AnonymizeRequest)
    - [AnonymizeResponse](#types.AnonymizeResponse)
  
  
  
    - [AnonymizeService](#types.AnonymizeService)
  

- [common.proto](#common.proto)
    - [AnalyzeResult](#types.AnalyzeResult)
    - [FieldTypes](#types.FieldTypes)
    - [Location](#types.Location)
  
    - [FieldTypesEnum](#types.FieldTypesEnum)
  
  
  

- [cronjob.proto](#cronjob.proto)
    - [CronJobApiRequest](#types.CronJobApiRequest)
    - [CronJobRequest](#types.CronJobRequest)
    - [CronJobResponse](#types.CronJobResponse)
  
  
  
    - [CronJobService](#types.CronJobService)
  

- [databinder.proto](#databinder.proto)
    - [CompletionMessage](#types.CompletionMessage)
    - [DatabinderRequest](#types.DatabinderRequest)
    - [DatabinderResponse](#types.DatabinderResponse)
  
    - [DataBinderTypesEnum](#types.DataBinderTypesEnum)
  
  
    - [DatabinderService](#types.DatabinderService)
  

- [scan.proto](#scan.proto)
    - [ScanRequest](#types.ScanRequest)
  
  
  
  

- [stream.proto](#stream.proto)
    - [StreamRequest](#types.StreamRequest)
  
  
  
  

- [template.proto](#template.proto)
    - [AnalyzeTemplate](#types.AnalyzeTemplate)
    - [AnonymizeTemplate](#types.AnonymizeTemplate)
    - [BlobStorageConfig](#types.BlobStorageConfig)
    - [CloudStorageConfig](#types.CloudStorageConfig)
    - [CronJobTemplate](#types.CronJobTemplate)
    - [DBConfig](#types.DBConfig)
    - [Databinder](#types.Databinder)
    - [DatabinderTemplate](#types.DatabinderTemplate)
    - [EHConfig](#types.EHConfig)
    - [FieldTypeTransformation](#types.FieldTypeTransformation)
    - [HashValue](#types.HashValue)
    - [ImageLocations](#types.ImageLocations)
    - [ImageWord](#types.ImageWord)
    - [JobTemplate](#types.JobTemplate)
    - [KafkaConfig](#types.KafkaConfig)
    - [KinesisConfig](#types.KinesisConfig)
    - [MaskImage](#types.MaskImage)
    - [MaskValue](#types.MaskValue)
    - [RedactValue](#types.RedactValue)
    - [ReplaceValue](#types.ReplaceValue)
    - [S3Config](#types.S3Config)
    - [ScanTemplate](#types.ScanTemplate)
    - [Schedule](#types.Schedule)
    - [StreamConfig](#types.StreamConfig)
    - [StreamTemplate](#types.StreamTemplate)
    - [Transformation](#types.Transformation)
    - [Trigger](#types.Trigger)
  
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="analyze.proto"/>
<p align="right"><a href="#top">Top</a></p>

## analyze.proto



<a name="types.AnalyzeApiRequest"/>

### AnalyzeApiRequest
AnalyzeApiRequest represents the request to the API HTTP service


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| text | [string](#string) |  |  |
| analyzeTemplateId | [string](#string) |  |  |
| analyzeTemplate | [AnalyzeTemplate](#types.AnalyzeTemplate) |  | Optional parameter for running the analyzer without creating a template |






<a name="types.AnalyzeRequest"/>

### AnalyzeRequest
AnalyzeRequest represents the request to the analyze service via GRPC


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| text | [string](#string) |  |  |
| analyzeTemplate | [AnalyzeTemplate](#types.AnalyzeTemplate) |  |  |
| minProbability | [string](#string) |  |  |






<a name="types.AnalyzeResponse"/>

### AnalyzeResponse
AnalyzeResponse represents the analyze service response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| analyzeResults | [AnalyzeResult](#types.AnalyzeResult) | repeated |  |





 

 

 


<a name="types.AnalyzeService"/>

### AnalyzeService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Apply | [AnalyzeRequest](#types.AnalyzeRequest) | [AnalyzeResponse](#types.AnalyzeRequest) |  |

 



<a name="anonymize.proto"/>
<p align="right"><a href="#top">Top</a></p>

## anonymize.proto



<a name="types.AnonymizeApiRequest"/>

### AnonymizeApiRequest
AnonymizeApiRequest represents the request to the API HTTP service


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| text | [string](#string) |  |  |
| analyzeTemplateId | [string](#string) |  |  |
| anonymizeTemplateId | [string](#string) |  |  |
| analyzeTemplate | [AnalyzeTemplate](#types.AnalyzeTemplate) |  | Optional parameter for running the analyzer without creating a template |
| anonymizeTemplate | [AnonymizeTemplate](#types.AnonymizeTemplate) |  | Optional parameter for running the anonymizer without creating a template |






<a name="types.AnonymizeRequest"/>

### AnonymizeRequest
AnonymizeRequest represents the request to the anonymize service via GRPC


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| text | [string](#string) |  |  |
| template | [AnonymizeTemplate](#types.AnonymizeTemplate) |  |  |
| analyzeResults | [AnalyzeResult](#types.AnalyzeResult) | repeated |  |






<a name="types.AnonymizeResponse"/>

### AnonymizeResponse
AnonymizeResponse represents the anonymize service response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| text | [string](#string) |  |  |





 

 

 


<a name="types.AnonymizeService"/>

### AnonymizeService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Apply | [AnonymizeRequest](#types.AnonymizeRequest) | [AnonymizeResponse](#types.AnonymizeRequest) |  |

 



<a name="common.proto"/>
<p align="right"><a href="#top">Top</a></p>

## common.proto



<a name="types.AnalyzeResult"/>

### AnalyzeResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| text | [string](#string) |  |  |
| field | [FieldTypes](#types.FieldTypes) |  |  |
| probability | [float](#float) |  |  |
| location | [Location](#types.Location) |  |  |






<a name="types.FieldTypes"/>

### FieldTypes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| languageCode | [string](#string) |  |  |






<a name="types.Location"/>

### Location



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start | [sint32](#sint32) |  |  |
| end | [sint32](#sint32) |  |  |
| length | [sint32](#sint32) |  |  |





 


<a name="types.FieldTypesEnum"/>

### FieldTypesEnum


| Name | Number | Description |
| ---- | ------ | ----------- |
| CREDIT_CARD | 0 |  |
| CRYPTO | 1 |  |
| DATE_TIME | 2 |  |
| DOMAIN_NAME | 3 |  |
| EMAIL_ADDRESS | 4 |  |
| IBAN_CODE | 5 |  |
| IP_ADDRESS | 6 |  |
| NRP | 7 |  |
| LOCATION | 8 |  |
| PERSON | 9 |  |
| PHONE_NUMBER | 10 |  |
| US_BANK_NUMBER | 11 |  |
| US_DRIVER_LICENSE | 12 |  |
| US_ITIN | 13 |  |
| US_PASSPORT | 14 |  |
| US_SSN | 15 |  |


 

 

 



<a name="cronjob.proto"/>
<p align="right"><a href="#top">Top</a></p>

## cronjob.proto



<a name="types.CronJobApiRequest"/>

### CronJobApiRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CronJobTemplateId | [string](#string) |  |  |






<a name="types.CronJobRequest"/>

### CronJobRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| trigger | [Trigger](#types.Trigger) |  |  |
| scanRequest | [ScanRequest](#types.ScanRequest) |  |  |






<a name="types.CronJobResponse"/>

### CronJobResponse






 

 

 


<a name="types.CronJobService"/>

### CronJobService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Apply | [CronJobRequest](#types.CronJobRequest) | [CronJobResponse](#types.CronJobRequest) |  |

 



<a name="databinder.proto"/>
<p align="right"><a href="#top">Top</a></p>

## databinder.proto



<a name="types.CompletionMessage"/>

### CompletionMessage







<a name="types.DatabinderRequest"/>

### DatabinderRequest
DatabinderRequest represents the request to the data-binder service via GRPC


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| analyzeResults | [AnalyzeResult](#types.AnalyzeResult) | repeated |  |
| anonymizeResult | [AnonymizeResponse](#types.AnonymizeResponse) |  |  |
| path | [string](#string) |  |  |






<a name="types.DatabinderResponse"/>

### DatabinderResponse
DatabinderResponse represents the response from the data-binder service via GRPC





 


<a name="types.DataBinderTypesEnum"/>

### DataBinderTypesEnum


| Name | Number | Description |
| ---- | ------ | ----------- |
| mysql | 0 |  |
| mssql | 1 |  |
| postgres | 2 |  |
| sqlite3 | 3 |  |
| oracle | 4 |  |
| kafka | 5 |  |
| eventhub | 6 |  |
| kinesis | 7 |  |
| s3 | 8 |  |
| azureblob | 9 |  |
| googlestorage | 10 |  |


 

 


<a name="types.DatabinderService"/>

### DatabinderService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Apply | [DatabinderRequest](#types.DatabinderRequest) | [DatabinderResponse](#types.DatabinderRequest) |  |
| Init | [DatabinderTemplate](#types.DatabinderTemplate) | [DatabinderResponse](#types.DatabinderTemplate) |  |
| Completion | [CompletionMessage](#types.CompletionMessage) | [DatabinderResponse](#types.CompletionMessage) |  |

 



<a name="scan.proto"/>
<p align="right"><a href="#top">Top</a></p>

## scan.proto



<a name="types.ScanRequest"/>

### ScanRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kind | [string](#string) |  |  |
| cloudStorageConfig | [CloudStorageConfig](#types.CloudStorageConfig) |  |  |
| minProbability | [string](#string) |  |  |
| analyzeTemplate | [AnalyzeTemplate](#types.AnalyzeTemplate) |  |  |
| anonymizeTemplate | [AnonymizeTemplate](#types.AnonymizeTemplate) |  |  |
| databinderTemplate | [DatabinderTemplate](#types.DatabinderTemplate) |  |  |





 

 

 

 



<a name="stream.proto"/>
<p align="right"><a href="#top">Top</a></p>

## stream.proto



<a name="types.StreamRequest"/>

### StreamRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kind | [string](#string) |  |  |
| streamConfig | [StreamConfig](#types.StreamConfig) |  |  |
| minProbability | [string](#string) |  |  |
| analyzeTemplate | [AnalyzeTemplate](#types.AnalyzeTemplate) |  |  |
| anonymizeTemplate | [AnonymizeTemplate](#types.AnonymizeTemplate) |  |  |
| databinderTemplate | [DatabinderTemplate](#types.DatabinderTemplate) |  |  |





 

 

 

 



<a name="template.proto"/>
<p align="right"><a href="#top">Top</a></p>

## template.proto



<a name="types.AnalyzeTemplate"/>

### AnalyzeTemplate
AnalyzeTemplate represents the analyze service template definition


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fields | [FieldTypes](#types.FieldTypes) | repeated |  |






<a name="types.AnonymizeTemplate"/>

### AnonymizeTemplate
AnonymizeTemplate represents the anonymize service template definition


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| displayName | [string](#string) |  |  |
| description | [string](#string) |  |  |
| createTime | [string](#string) |  |  |
| modifiedTime | [string](#string) |  |  |
| fieldTypeTransformations | [FieldTypeTransformation](#types.FieldTypeTransformation) | repeated |  |






<a name="types.BlobStorageConfig"/>

### BlobStorageConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accountName | [string](#string) |  |  |
| accountKey | [string](#string) |  |  |
| containerName | [string](#string) |  |  |






<a name="types.CloudStorageConfig"/>

### CloudStorageConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| blobStorageConfig | [BlobStorageConfig](#types.BlobStorageConfig) |  |  |
| s3Config | [S3Config](#types.S3Config) |  |  |






<a name="types.CronJobTemplate"/>

### CronJobTemplate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| trigger | [Trigger](#types.Trigger) |  |  |
| scanTemplateId | [string](#string) |  |  |
| analyzeTemplateId | [string](#string) |  |  |
| anonymizeTemplateId | [string](#string) |  |  |
| databinderTemplateId | [string](#string) |  |  |






<a name="types.DBConfig"/>

### DBConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| connectionString | [string](#string) |  |  |
| tableName | [string](#string) |  |  |






<a name="types.Databinder"/>

### Databinder



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dbConfig | [DBConfig](#types.DBConfig) |  |  |
| cloudStorageConfig | [CloudStorageConfig](#types.CloudStorageConfig) |  |  |






<a name="types.DatabinderTemplate"/>

### DatabinderTemplate
DatabinderTemplate represents the analyzer service outputs definition


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| analyzerKind | [string](#string) |  |  |
| anonymizerKind | [string](#string) |  |  |
| databinder | [Databinder](#types.Databinder) |  |  |






<a name="types.EHConfig"/>

### EHConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ehNamespace | [string](#string) |  |  |
| ehName | [string](#string) |  |  |
| ehConnectionString | [string](#string) |  |  |
| ehKeyName | [string](#string) |  |  |
| ehKeyValue | [string](#string) |  |  |
| partitionCount | [int32](#int32) |  |  |






<a name="types.FieldTypeTransformation"/>

### FieldTypeTransformation
FieldTypeTransformation represents the transformation for array of fields types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fields | [FieldTypes](#types.FieldTypes) | repeated |  |
| transformation | [Transformation](#types.Transformation) |  |  |






<a name="types.HashValue"/>

### HashValue







<a name="types.ImageLocations"/>

### ImageLocations



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Words | [ImageWord](#types.ImageWord) | repeated |  |
| ImageText | [string](#string) |  |  |






<a name="types.ImageWord"/>

### ImageWord



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| XLocation | [string](#string) |  |  |
| Width | [string](#string) |  |  |
| Height | [string](#string) |  |  |
| YLocation | [string](#string) |  |  |
| Text | [string](#string) |  |  |
| startPosition | [int32](#int32) |  |  |
| endPosition | [int32](#int32) |  |  |






<a name="types.JobTemplate"/>

### JobTemplate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| trigger | [Trigger](#types.Trigger) |  |  |
| scanTemplateId | [string](#string) |  |  |
| streamTemplateId | [string](#string) |  |  |






<a name="types.KafkaConfig"/>

### KafkaConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| topic | [string](#string) |  |  |
| partitionCount | [int32](#int32) |  |  |
| saslUsername | [string](#string) |  |  |
| saslPassword | [string](#string) |  |  |






<a name="types.KinesisConfig"/>

### KinesisConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| aWsSecretAccessKey | [string](#string) |  |  |
| aWsRegion | [string](#string) |  |  |
| aWsSecretKey | [string](#string) |  |  |
| redisUrl | [string](#string) |  |  |
| streamName | [string](#string) |  |  |






<a name="types.MaskImage"/>

### MaskImage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| imageLocations | [ImageLocations](#types.ImageLocations) |  |  |
| imageToMaskBase64 | [string](#string) |  |  |






<a name="types.MaskValue"/>

### MaskValue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| maskingCharacter | [string](#string) |  |  |
| charsToMask | [int32](#int32) |  |  |
| fromEnd | [bool](#bool) |  |  |






<a name="types.RedactValue"/>

### RedactValue







<a name="types.ReplaceValue"/>

### ReplaceValue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| newValue | [string](#string) |  |  |






<a name="types.S3Config"/>

### S3Config



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accessId | [string](#string) |  |  |
| accessKey | [string](#string) |  |  |
| region | [string](#string) |  |  |
| bucketName | [string](#string) |  |  |
| endpoint | [string](#string) |  |  |






<a name="types.ScanTemplate"/>

### ScanTemplate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kind | [string](#string) |  |  |
| cloudStorageConfig | [CloudStorageConfig](#types.CloudStorageConfig) |  |  |
| minProbability | [string](#string) |  |  |






<a name="types.Schedule"/>

### Schedule



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| recurrencePeriodDuration | [string](#string) |  |  |






<a name="types.StreamConfig"/>

### StreamConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kafkaConfig | [KafkaConfig](#types.KafkaConfig) |  |  |
| ehConfig | [EHConfig](#types.EHConfig) |  |  |






<a name="types.StreamTemplate"/>

### StreamTemplate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kind | [string](#string) |  |  |
| streamConfig | [StreamConfig](#types.StreamConfig) |  |  |
| minProbability | [string](#string) |  |  |
| analyzeTemplateId | [string](#string) |  |  |
| anonymizeTemplateId | [string](#string) |  |  |
| databinderTemplateId | [string](#string) |  |  |






<a name="types.Transformation"/>

### Transformation
Transformation represents the transformation type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| replaceValue | [ReplaceValue](#types.ReplaceValue) |  |  |
| redactValue | [RedactValue](#types.RedactValue) |  |  |
| hashValue | [HashValue](#types.HashValue) |  |  |
| maskValue | [MaskValue](#types.MaskValue) |  |  |
| maskImage | [MaskImage](#types.MaskImage) |  |  |






<a name="types.Trigger"/>

### Trigger



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schedule | [Schedule](#types.Schedule) |  |  |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |


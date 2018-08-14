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
  
  
  

- [datasink.proto](#datasink.proto)
    - [CompletionMessage](#types.CompletionMessage)
    - [DatasinkRequest](#types.DatasinkRequest)
    - [DatasinkResponse](#types.DatasinkResponse)
  
    - [DatasinkTypesEnum](#types.DatasinkTypesEnum)
  
  
    - [DatasinkService](#types.DatasinkService)
  

- [scan.proto](#scan.proto)
    - [ScanRequest](#types.ScanRequest)
  
  
  
  

- [scheduler.proto](#scheduler.proto)
    - [ScannerCronJobApiRequest](#types.ScannerCronJobApiRequest)
    - [ScannerCronJobRequest](#types.ScannerCronJobRequest)
    - [ScannerCronJobResponse](#types.ScannerCronJobResponse)
    - [StreamsJobApiRequest](#types.StreamsJobApiRequest)
    - [StreamsJobRequest](#types.StreamsJobRequest)
    - [StreamsJobResponse](#types.StreamsJobResponse)
  
  
  
    - [SchedulerService](#types.SchedulerService)
  

- [stream.proto](#stream.proto)
    - [StreamRequest](#types.StreamRequest)
  
  
  
  

- [template.proto](#template.proto)
    - [AnalyzeTemplate](#types.AnalyzeTemplate)
    - [AnonymizeTemplate](#types.AnonymizeTemplate)
    - [BlobStorageConfig](#types.BlobStorageConfig)
    - [CloudStorageConfig](#types.CloudStorageConfig)
    - [DBConfig](#types.DBConfig)
    - [Datasink](#types.Datasink)
    - [DatasinkTemplate](#types.DatasinkTemplate)
    - [EHConfig](#types.EHConfig)
    - [FieldTypeTransformation](#types.FieldTypeTransformation)
    - [GoogleStorageConfig](#types.GoogleStorageConfig)
    - [HashValue](#types.HashValue)
    - [KafkaConfig](#types.KafkaConfig)
    - [KinesisConfig](#types.KinesisConfig)
    - [MaskValue](#types.MaskValue)
    - [RedactValue](#types.RedactValue)
    - [ReplaceValue](#types.ReplaceValue)
    - [S3Config](#types.S3Config)
    - [ScanTemplate](#types.ScanTemplate)
    - [ScannerCronJobTemplate](#types.ScannerCronJobTemplate)
    - [Schedule](#types.Schedule)
    - [StreamConfig](#types.StreamConfig)
    - [StreamTemplate](#types.StreamTemplate)
    - [StreamsJobTemplate](#types.StreamsJobTemplate)
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
| text | [string](#string) |  | The text to analyze |
| analyzeTemplateId | [string](#string) |  | The analyze template id - that hold the analyze configuration |
| analyzeTemplate | [AnalyzeTemplate](#types.AnalyzeTemplate) |  | Optional parameter for running the analyze service without creating a template |






<a name="types.AnalyzeRequest"/>

### AnalyzeRequest
AnalyzeRequest represents the request to the analyze service via GRPC


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| text | [string](#string) |  | The text to analyze |
| analyzeTemplate | [AnalyzeTemplate](#types.AnalyzeTemplate) |  | The analyze template, which configures which sensitive data should be analyzed |
| minProbability | [string](#string) |  | The analyze service defines a degree of certainty (0 to 1) for each result. The minProbability will filter results which has lower certainty than the provided value. |






<a name="types.AnalyzeResponse"/>

### AnalyzeResponse
AnalyzeResponse represents the analyze service response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| analyzeResults | [AnalyzeResult](#types.AnalyzeResult) | repeated | Array of the analyze results finding |





 

 

 


<a name="types.AnalyzeService"/>

### AnalyzeService
The Analyze Service is a service that analyze a given the text using predefined analyzers fields to identify patterns, 
formats, and checksums with relevant context.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Apply | [AnalyzeRequest](#types.AnalyzeRequest) | [AnalyzeResponse](#types.AnalyzeRequest) | Apply method will execute on the given request and return the analyze response with the sensitive text findings |

 



<a name="anonymize.proto"/>
<p align="right"><a href="#top">Top</a></p>

## anonymize.proto



<a name="types.AnonymizeApiRequest"/>

### AnonymizeApiRequest
AnonymizeApiRequest represents the request to the API HTTP service


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| text | [string](#string) |  | The text to anonymize |
| analyzeTemplateId | [string](#string) |  | The analyze template id - anonymization is done according to analyzing results. One of analyzeTemplateId or analyzeTemplate have to be configured. |
| anonymizeTemplateId | [string](#string) |  | The anonymize template id - represents the anonymize configuration, which fields to anonymize and how. |
| analyzeTemplate | [AnalyzeTemplate](#types.AnalyzeTemplate) |  | Optional parameter for running the analyzer without creating a template. |
| anonymizeTemplate | [AnonymizeTemplate](#types.AnonymizeTemplate) |  | Optional parameter for running the anonymizer without creating a template. |






<a name="types.AnonymizeRequest"/>

### AnonymizeRequest
AnonymizeRequest represents the request to the anonymize service via GRPC


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| text | [string](#string) |  | The text to anonymize |
| template | [AnonymizeTemplate](#types.AnonymizeTemplate) |  | The anonymize template represent the anonymize configuration, which fields to anonymize and how |
| analyzeResults | [AnalyzeResult](#types.AnalyzeResult) | repeated |  |






<a name="types.AnonymizeResponse"/>

### AnonymizeResponse
AnonymizeResponse represents the anonymize service response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| text | [string](#string) |  | The anonymized text |





 

 

 


<a name="types.AnonymizeService"/>

### AnonymizeService
The Anonymize Service is a service that anonymizes a given the text using predefined analyzers fields and anonymize configurations.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Apply | [AnonymizeRequest](#types.AnonymizeRequest) | [AnonymizeResponse](#types.AnonymizeRequest) | Apply method will execute on the given request and return the anonymize response with the sensitive text anonymized |

 



<a name="common.proto"/>
<p align="right"><a href="#top">Top</a></p>

## common.proto



<a name="types.AnalyzeResult"/>

### AnalyzeResult
AnalyzeResult represents the Analyze serivce findings


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| text | [string](#string) |  | The sensitive text result |
| field | [FieldTypes](#types.FieldTypes) |  | The sensitive text type (supported types: FieldTypesEnum) |
| probability | [float](#float) |  | The certainty of the result |
| location | [Location](#types.Location) |  | The loaction in the text of the finding |






<a name="types.FieldTypes"/>

### FieldTypes
FieldType strucy


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Field type name |
| languageCode | [string](#string) |  | Field type language code |






<a name="types.Location"/>

### Location
The location in the text of the finding


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start | [sint32](#sint32) |  | The location start |
| end | [sint32](#sint32) |  | The location end |
| length | [sint32](#sint32) |  | The location length |





 


<a name="types.FieldTypesEnum"/>

### FieldTypesEnum
FieldTypes for Analyzing and Anonymizing

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


 

 

 



<a name="datasink.proto"/>
<p align="right"><a href="#top">Top</a></p>

## datasink.proto



<a name="types.CompletionMessage"/>

### CompletionMessage
CompletionMessage represents an indication to the data sink service that the scanning job is done service via GRPC






<a name="types.DatasinkRequest"/>

### DatasinkRequest
DatasinkRequest represents the request to the data-sink service via GRPC


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| analyzeResults | [AnalyzeResult](#types.AnalyzeResult) | repeated | Array of the analyzed results |
| anonymizeResult | [AnonymizeResponse](#types.AnonymizeResponse) |  | The anonymized text |
| path | [string](#string) |  | The path where the anonymized text is located |






<a name="types.DatasinkResponse"/>

### DatasinkResponse
DatasinkResponse represents the response from the data-sink service via GRPC





 


<a name="types.DatasinkTypesEnum"/>

### DatasinkTypesEnum
The data sink supported destenation types

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


 

 


<a name="types.DatasinkService"/>

### DatasinkService
The data sink service represents the service for writing the results of the analyzing and anonymizng service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Apply | [DatasinkRequest](#types.DatasinkRequest) | [DatasinkResponse](#types.DatasinkRequest) | Apply method will execute on the given request and return whether the result where written successfully to the destination |
| Init | [DatasinkTemplate](#types.DatasinkTemplate) | [DatasinkResponse](#types.DatasinkTemplate) | Init the data sink service with the provided data sink template |
| Completion | [CompletionMessage](#types.CompletionMessage) | [DatasinkResponse](#types.CompletionMessage) | Completion method for indicating that the scanning job is done |

 



<a name="scan.proto"/>
<p align="right"><a href="#top">Top</a></p>

## scan.proto



<a name="types.ScanRequest"/>

### ScanRequest
ScanRequest represents the request to the scanner service via GRPC


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kind | [string](#string) |  | The scanner input kind |
| cloudStorageConfig | [CloudStorageConfig](#types.CloudStorageConfig) |  | The cloud storage configuration for scanning cloud storage such as AzureBlobStorage, AWS S3 And Google Storage |
| minProbability | [string](#string) |  | The minProbability will filter results which has lower certainty than the provided value. |
| analyzeTemplate | [AnalyzeTemplate](#types.AnalyzeTemplate) |  | The analyzer template configures the fields to analyze |
| anonymizeTemplate | [AnonymizeTemplate](#types.AnonymizeTemplate) |  | The anonymizer template configures how to anonymize the sensitive data [optional] |
| datasinkTemplate | [DatasinkTemplate](#types.DatasinkTemplate) |  | The datasinkTemplate configures the output destination of the analyzer/anonymizer results |





 

 

 

 



<a name="scheduler.proto"/>
<p align="right"><a href="#top">Top</a></p>

## scheduler.proto



<a name="types.ScannerCronJobApiRequest"/>

### ScannerCronJobApiRequest
ScannerCronJobApiRequest represents the request to the API HTTP service


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ScannerCronJobTemplateId | [string](#string) |  | The scanner cron job template id |






<a name="types.ScannerCronJobRequest"/>

### ScannerCronJobRequest
ScannerCronJobRequest represents the request to the scheduler service via GRPC


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| description | [string](#string) |  | The cron job description |
| trigger | [Trigger](#types.Trigger) |  | The trigger for a new job to start |
| scanRequest | [ScanRequest](#types.ScanRequest) |  | The scanner requeset that hold the scanning details |






<a name="types.ScannerCronJobResponse"/>

### ScannerCronJobResponse







<a name="types.StreamsJobApiRequest"/>

### StreamsJobApiRequest
StreamsJobApiRequest represents the request to the API HTTP service


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| StreamsJobTemplateId | [string](#string) |  | The streams job template id |






<a name="types.StreamsJobRequest"/>

### StreamsJobRequest
StreamsJobRequest represents the request to the scheduler service via GRPC


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| description | [string](#string) |  | The cron job description |
| streamsRequest | [StreamRequest](#types.StreamRequest) |  | The streams requeset that hold the streaming details |






<a name="types.StreamsJobResponse"/>

### StreamsJobResponse






 

 

 


<a name="types.SchedulerService"/>

### SchedulerService
The CronJob Service is a service that triggers a new cronjob for scanning a given storage periodcallly

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ApplyStream | [StreamsJobRequest](#types.StreamsJobRequest) | [StreamsJobResponse](#types.StreamsJobRequest) | ApplyStream method will trigger a new scanning cron job and will return if it was triggered successfully |
| ApplyScan | [ScannerCronJobRequest](#types.ScannerCronJobRequest) | [ScannerCronJobResponse](#types.ScannerCronJobRequest) | Apply method will trigger a new scanning cron job and will return if it was triggered successfully |

 



<a name="stream.proto"/>
<p align="right"><a href="#top">Top</a></p>

## stream.proto



<a name="types.StreamRequest"/>

### StreamRequest
StreamRequest represents the request to the stream service via GRPC


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kind | [string](#string) |  | The scanner input kind |
| streamConfig | [StreamConfig](#types.StreamConfig) |  | The stream configuration for scanning streams such as Azure EventHub, Kafka and Kinesis |
| minProbability | [string](#string) |  | The minProbability will filter results which has lower certainty than the provided value. |
| analyzeTemplate | [AnalyzeTemplate](#types.AnalyzeTemplate) |  | The analyzer template configures the fields to analyze |
| anonymizeTemplate | [AnonymizeTemplate](#types.AnonymizeTemplate) |  | The anonymizer template configures how to anonymize the sensitive data [optional] |
| datasinkTemplate | [DatasinkTemplate](#types.DatasinkTemplate) |  | The datasinkTemplate configures the output destination of the analyzer/anonymizer results |





 

 

 

 



<a name="template.proto"/>
<p align="right"><a href="#top">Top</a></p>

## template.proto



<a name="types.AnalyzeTemplate"/>

### AnalyzeTemplate
AnalyzeTemplate represents the template definition of the Analyze service- for analyzing sensitive text.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fields | [FieldTypes](#types.FieldTypes) | repeated | Array of the fields to analyze |
| description | [string](#string) |  | Template description |
| createTime | [string](#string) |  | Template Creation date |
| modifiedTime | [string](#string) |  | Template modification date |






<a name="types.AnonymizeTemplate"/>

### AnonymizeTemplate
AnonymizeTemplate represents the template definition of the Anonymize service for anonymizying the sensitive data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| description | [string](#string) |  | Template description |
| createTime | [string](#string) |  | Template Creation date |
| modifiedTime | [string](#string) |  | Template modification date |
| fieldTypeTransformations | [FieldTypeTransformation](#types.FieldTypeTransformation) | repeated | FieldTypeTransformation represents the transformation for an array of fields types |






<a name="types.BlobStorageConfig"/>

### BlobStorageConfig
Azure Blob Storage configuration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accountName | [string](#string) |  | Azure account name |
| accountKey | [string](#string) |  | Azure account key |
| containerName | [string](#string) |  | The blob storage container Name |






<a name="types.CloudStorageConfig"/>

### CloudStorageConfig
Represents the cloud storage config - supported storage: Azure blob storage, AWS S3, Google storage


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| blobStorageConfig | [BlobStorageConfig](#types.BlobStorageConfig) |  | The azure blob storage config |
| s3Config | [S3Config](#types.S3Config) |  | The s3 config |
| GoogleStorageConfig | [GoogleStorageConfig](#types.GoogleStorageConfig) |  | The google storage config |






<a name="types.DBConfig"/>

### DBConfig
DBConfig represents the Database configuration - the DB connection string and the table name
Supported database: mssql, mysql, sqlite3, postgreSQL, oracle


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| connectionString | [string](#string) |  | The database connection string |
| tableName | [string](#string) |  | The table name |






<a name="types.Datasink"/>

### Datasink
Datasink represents the configuration for storing the scanner output to the selected destination such as cloud storage, database, etc.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dbConfig | [DBConfig](#types.DBConfig) |  | The database configuration |
| cloudStorageConfig | [CloudStorageConfig](#types.CloudStorageConfig) |  | The cloud storage configuration |
| streamConfig | [StreamConfig](#types.StreamConfig) |  | The stream configuration |






<a name="types.DatasinkTemplate"/>

### DatasinkTemplate
DatasinkTemplate represents the scanner service outputs definition.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| analyzerKind | [string](#string) |  | analyzer output kind - supported types defined in DatasinkTypesEnum |
| anonymizerKind | [string](#string) |  | anonymizer output kind - supported types defined in DatasinkTypesEnum |
| datasink | [Datasink](#types.Datasink) |  | Datasink represents represents the configuration for storing the scanner output. Needs to be configured corresponding the selected analyzer/anonymizer types |






<a name="types.EHConfig"/>

### EHConfig
Azure EventHub configuration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ehNamespace | [string](#string) |  | EventHub namespace |
| ehName | [string](#string) |  | EventHub name |
| ehConnectionString | [string](#string) |  | Eventhub connection string |
| ehKeyName | [string](#string) |  | Eventhub key name (a key name and a key value can provided instead of the full connection string) |
| ehKeyValue | [string](#string) |  | Eventhub key value |






<a name="types.FieldTypeTransformation"/>

### FieldTypeTransformation
FieldTypeTransformation represents the transformation for an array of fields types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fields | [FieldTypes](#types.FieldTypes) | repeated | The array of field types |
| transformation | [Transformation](#types.Transformation) |  | The transformation for the array of fields |






<a name="types.GoogleStorageConfig"/>

### GoogleStorageConfig
Represents the Google Storage configuration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| json | [string](#string) |  | The json |
| projectId | [string](#string) |  | The project id |
| scopes | [string](#string) |  | The scopes authentication [there are different scopes, which you can find here https://cloud.google.com/storage/docs/authentication] |
| bucketName | [string](#string) |  | The bucket name |






<a name="types.HashValue"/>

### HashValue
Uses cryptographich hash on the given value with FNV-1 hash.






<a name="types.KafkaConfig"/>

### KafkaConfig
The Kafka configuration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | Kafka address |
| topic | [string](#string) |  | Kafka topic |
| saslUsername | [string](#string) |  | SASL authentication user name |
| saslPassword | [string](#string) |  | SASL authentication password |






<a name="types.KinesisConfig"/>

### KinesisConfig
The Kinesis cinfiguration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| awsAccessKeyId | [string](#string) |  | AWS secret access key |
| awsRegion | [string](#string) |  | AWS secret region |
| awsSecretAccessKey | [string](#string) |  | AWS secret secret key |
| redisUrl | [string](#string) |  | The Redis url |
| streamName | [string](#string) |  | The stream name |
| endpointAddress | [string](#string) |  | The endpoint address |






<a name="types.MaskValue"/>

### MaskValue
Mask the given value with the selected characters.
For example: my credit card number is 4961-2765-5327-5913
maskingCharacter: &#39;*&#39;, chars to mask: 9, fromEnd: true
Will output:  my credit card number is 4961-2765-*********


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| maskingCharacter | [string](#string) |  | the masking char |
| charsToMask | [int32](#int32) |  | number of chars to mask |
| fromEnd | [bool](#bool) |  | Should start masking from end |






<a name="types.RedactValue"/>

### RedactValue
Redacts the given value. For example We met in Seattle would change to &#39;We me in &#39;.






<a name="types.ReplaceValue"/>

### ReplaceValue
Replace the value with the given new value. 
For example: We met in Seattle.
new value: &lt;Location&gt;
replace to &#39;We met in &lt;Location&gt;&#39;


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| newValue | [string](#string) |  | The value to replace |






<a name="types.S3Config"/>

### S3Config
AWS S3 configuration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accessId | [string](#string) |  | The access id |
| accessKey | [string](#string) |  | The access key |
| region | [string](#string) |  | The region |
| bucketName | [string](#string) |  | The bucket name |
| endpoint | [string](#string) |  | The s3 endpoint |






<a name="types.ScanTemplate"/>

### ScanTemplate
ScanTemplate represents the scan configuration for scanning data and analyzing/anonymizing it.
And sending the output to the selected destination


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kind | [string](#string) |  | The input scan kind, supported kinds are: Azure blob storage, AWS S3, Google Storage |
| cloudStorageConfig | [CloudStorageConfig](#types.CloudStorageConfig) |  | The selected cloudstorage configuration |
| minProbability | [string](#string) |  | The minProbability will filter results which has lower certainty than the provided value. |






<a name="types.ScannerCronJobTemplate"/>

### ScannerCronJobTemplate
The kuberenetes Cronjob template. Creates a time based scanning job scheduler that runs periodcally
on the selected time


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| description | [string](#string) |  | The job description |
| trigger | [Trigger](#types.Trigger) |  | The trigger for a new job to start |
| scanTemplateId | [string](#string) |  | The scan template id configures job input source |
| analyzeTemplateId | [string](#string) |  | The analyzer template id configures the fields to analyze |
| anonymizeTemplateId | [string](#string) |  | The anonymizer template id configures how to anonymize the sensitive data [optional] |
| datasinkTemplateId | [string](#string) |  | The datasink template id configures the output destination of the analyzer/anonymizer results |






<a name="types.Schedule"/>

### Schedule
Defines the job schedule


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| recurrencePeriod | [string](#string) |  | The recurrence period of the job, set as a cron expression |






<a name="types.StreamConfig"/>

### StreamConfig
Represents the streams configuration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kafkaConfig | [KafkaConfig](#types.KafkaConfig) |  | The kafka configuration |
| ehConfig | [EHConfig](#types.EHConfig) |  | The Azure Event Hub configuration |
| kinesisConfig | [KinesisConfig](#types.KinesisConfig) |  | The Kinesis configuration |






<a name="types.StreamTemplate"/>

### StreamTemplate
StreamTemplate represents the stream configuration for streaming data and analyzing/anonymizing it.
And sending the output to the selected destination


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kind | [string](#string) |  | The stream kind. Supported Kinds: EventHub, Kafka, Kinesis |
| streamConfig | [StreamConfig](#types.StreamConfig) |  | The selected stream configuration |
| minProbability | [string](#string) |  | The minProbability will filter results which has lower certainty than the provided value. |
| analyzeTemplateId | [string](#string) |  | The analyzer template id configures the fields to analyze |
| anonymizeTemplateId | [string](#string) |  | The anonymizer template id configures how to anonymize the sensitive data [optional] |
| datasinkTemplateId | [string](#string) |  | The datasinkTemplateId configures the output destination of the analyzer/anonymizer results |
| partitionCount | [int32](#int32) |  | Number of partitions |






<a name="types.StreamsJobTemplate"/>

### StreamsJobTemplate
The kuberenetes job template. Creates a job that creates streams containers


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| description | [string](#string) |  | The job description |
| streamsTemplateId | [string](#string) |  | The scan template id configures job input source |
| analyzeTemplateId | [string](#string) |  | The analyzer template id configures the fields to analyze |
| anonymizeTemplateId | [string](#string) |  | The anonymizer template id configures how to anonymize the sensitive data [optional] |
| datasinkTemplateId | [string](#string) |  | The datasink template id configures the output destination of the analyzer/anonymizer results |






<a name="types.Transformation"/>

### Transformation
Transformation represents the transformation types - how the sensitive text will change


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| replaceValue | [ReplaceValue](#types.ReplaceValue) |  | Relace the text with the defined value |
| redactValue | [RedactValue](#types.RedactValue) |  | Redact the text |
| hashValue | [HashValue](#types.HashValue) |  | Hashes the text |
| maskValue | [MaskValue](#types.MaskValue) |  | Mask n characters of the text |






<a name="types.Trigger"/>

### Trigger
Trigger defines the trigger for a new job to start


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schedule | [Schedule](#types.Schedule) |  | Defines the job schedule |





 

 

 

 



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


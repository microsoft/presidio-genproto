# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: anonymize-image.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import common_pb2 as common__pb2
import template_pb2 as template__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='anonymize-image.proto',
  package='types',
  syntax='proto3',
  serialized_pb=_b('\n\x15\x61nonymize-image.proto\x12\x05types\x1a\x0c\x63ommon.proto\x1a\x0etemplate.proto\"\xe8\x01\n\x18\x41nonymizeImageApiRequest\x12\x0c\n\x04\x64\x61ta\x18\x01 \x01(\x0c\x12\x11\n\timageType\x18\x02 \x01(\t\x12\x19\n\x11\x61nalyzeTemplateId\x18\x03 \x01(\t\x12 \n\x18\x61nonymizeImageTemplateId\x18\x04 \x01(\t\x12/\n\x0f\x61nalyzeTemplate\x18\x05 \x01(\x0b\x32\x16.types.AnalyzeTemplate\x12=\n\x16\x61nonymizeImageTemplate\x18\x06 \x01(\x0b\x32\x1d.types.AnonymizeImageTemplate\"\xd2\x01\n\x15\x41nonymizeImageRequest\x12\x1b\n\x05image\x18\x01 \x01(\x0b\x32\x0c.types.Image\x12/\n\x08template\x18\x02 \x01(\x0b\x32\x1d.types.AnonymizeImageTemplate\x12=\n\x16\x61nonymizeImageTypeEnum\x18\x03 \x01(\x0e\x32\x1d.types.AnonymizeImageTypeEnum\x12,\n\x0e\x61nalyzeResults\x18\x04 \x03(\x0b\x32\x14.types.AnalyzeResult\"5\n\x16\x41nonymizeImageResponse\x12\x1b\n\x05image\x18\x01 \x01(\x0b\x32\x0c.types.Image2_\n\x15\x41nonymizeImageService\x12\x46\n\x05\x41pply\x12\x1c.types.AnonymizeImageRequest\x1a\x1d.types.AnonymizeImageResponse\"\x00\x62\x06proto3')
  ,
  dependencies=[common__pb2.DESCRIPTOR,template__pb2.DESCRIPTOR,])




_ANONYMIZEIMAGEAPIREQUEST = _descriptor.Descriptor(
  name='AnonymizeImageApiRequest',
  full_name='types.AnonymizeImageApiRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='data', full_name='types.AnonymizeImageApiRequest.data', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='imageType', full_name='types.AnonymizeImageApiRequest.imageType', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='analyzeTemplateId', full_name='types.AnonymizeImageApiRequest.analyzeTemplateId', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='anonymizeImageTemplateId', full_name='types.AnonymizeImageApiRequest.anonymizeImageTemplateId', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='analyzeTemplate', full_name='types.AnonymizeImageApiRequest.analyzeTemplate', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='anonymizeImageTemplate', full_name='types.AnonymizeImageApiRequest.anonymizeImageTemplate', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=63,
  serialized_end=295,
)


_ANONYMIZEIMAGEREQUEST = _descriptor.Descriptor(
  name='AnonymizeImageRequest',
  full_name='types.AnonymizeImageRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='image', full_name='types.AnonymizeImageRequest.image', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='template', full_name='types.AnonymizeImageRequest.template', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='anonymizeImageTypeEnum', full_name='types.AnonymizeImageRequest.anonymizeImageTypeEnum', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='analyzeResults', full_name='types.AnonymizeImageRequest.analyzeResults', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=298,
  serialized_end=508,
)


_ANONYMIZEIMAGERESPONSE = _descriptor.Descriptor(
  name='AnonymizeImageResponse',
  full_name='types.AnonymizeImageResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='image', full_name='types.AnonymizeImageResponse.image', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=510,
  serialized_end=563,
)

_ANONYMIZEIMAGEAPIREQUEST.fields_by_name['analyzeTemplate'].message_type = template__pb2._ANALYZETEMPLATE
_ANONYMIZEIMAGEAPIREQUEST.fields_by_name['anonymizeImageTemplate'].message_type = template__pb2._ANONYMIZEIMAGETEMPLATE
_ANONYMIZEIMAGEREQUEST.fields_by_name['image'].message_type = common__pb2._IMAGE
_ANONYMIZEIMAGEREQUEST.fields_by_name['template'].message_type = template__pb2._ANONYMIZEIMAGETEMPLATE
_ANONYMIZEIMAGEREQUEST.fields_by_name['anonymizeImageTypeEnum'].enum_type = common__pb2._ANONYMIZEIMAGETYPEENUM
_ANONYMIZEIMAGEREQUEST.fields_by_name['analyzeResults'].message_type = common__pb2._ANALYZERESULT
_ANONYMIZEIMAGERESPONSE.fields_by_name['image'].message_type = common__pb2._IMAGE
DESCRIPTOR.message_types_by_name['AnonymizeImageApiRequest'] = _ANONYMIZEIMAGEAPIREQUEST
DESCRIPTOR.message_types_by_name['AnonymizeImageRequest'] = _ANONYMIZEIMAGEREQUEST
DESCRIPTOR.message_types_by_name['AnonymizeImageResponse'] = _ANONYMIZEIMAGERESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

AnonymizeImageApiRequest = _reflection.GeneratedProtocolMessageType('AnonymizeImageApiRequest', (_message.Message,), dict(
  DESCRIPTOR = _ANONYMIZEIMAGEAPIREQUEST,
  __module__ = 'anonymize_image_pb2'
  # @@protoc_insertion_point(class_scope:types.AnonymizeImageApiRequest)
  ))
_sym_db.RegisterMessage(AnonymizeImageApiRequest)

AnonymizeImageRequest = _reflection.GeneratedProtocolMessageType('AnonymizeImageRequest', (_message.Message,), dict(
  DESCRIPTOR = _ANONYMIZEIMAGEREQUEST,
  __module__ = 'anonymize_image_pb2'
  # @@protoc_insertion_point(class_scope:types.AnonymizeImageRequest)
  ))
_sym_db.RegisterMessage(AnonymizeImageRequest)

AnonymizeImageResponse = _reflection.GeneratedProtocolMessageType('AnonymizeImageResponse', (_message.Message,), dict(
  DESCRIPTOR = _ANONYMIZEIMAGERESPONSE,
  __module__ = 'anonymize_image_pb2'
  # @@protoc_insertion_point(class_scope:types.AnonymizeImageResponse)
  ))
_sym_db.RegisterMessage(AnonymizeImageResponse)



_ANONYMIZEIMAGESERVICE = _descriptor.ServiceDescriptor(
  name='AnonymizeImageService',
  full_name='types.AnonymizeImageService',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=565,
  serialized_end=660,
  methods=[
  _descriptor.MethodDescriptor(
    name='Apply',
    full_name='types.AnonymizeImageService.Apply',
    index=0,
    containing_service=None,
    input_type=_ANONYMIZEIMAGEREQUEST,
    output_type=_ANONYMIZEIMAGERESPONSE,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_ANONYMIZEIMAGESERVICE)

DESCRIPTOR.services_by_name['AnonymizeImageService'] = _ANONYMIZEIMAGESERVICE

# @@protoc_insertion_point(module_scope)

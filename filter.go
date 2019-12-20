package main

import (
	"github.com/Workiva/go-datastructures/set"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc/builder"
	"github.com/wdullaer/proto-filter/filter"
)

const (
	fileFilter      = "file_filter"
	messageFilter   = "message_filter"
	serviceFilter   = "service_filter"
	methodFilter    = "method_filter"
	fieldFilter     = "field_filter"
	enumFilter      = "enum_filter"
	enumValueFilter = "enum_value_filter"
)

// filterFile recursively applies the ValueFilter to the proto file and all its
// contents. It will return `true` if the file is to be removed from the output
//
// filterFile mutates the FileBuilder (and child Builders) in place: this
// simplified the code quite a bit, since there is no convenience method to
// remove all children from a Builder.
func filterFile(fileBuilder *builder.FileBuilder, terms *set.Set) (bool, error) {
	// Use the regular protobuf stuff to extract the extension value and compare
	fDesc, err := fileBuilder.Build()
	if err != nil {
		return false, err
	}
	if fOptions := fDesc.GetFileOptions(); fOptions != nil {
		extVal, err := proto.GetExtension(fOptions, filter.E_File)
		if err != proto.ErrMissingExtension {
			if err != nil {
				return false, err
			}
			if isExcluded(extVal, terms) {
				return true, nil
			}
		}
	}

	for _, child := range fileBuilder.GetChildren() {
		if isExcluded, err := filterChild(child, terms); err != nil {
			return false, err
		} else if isExcluded {
			removeFileChild(fileBuilder, child)
		}
	}
	return false, nil
}

func removeFileChild(fileBuilder *builder.FileBuilder, child builder.Builder) {
	switch child.(type) {
	case *builder.MessageBuilder:
		fileBuilder.RemoveMessage(child.GetName())
	case *builder.EnumBuilder:
		fileBuilder.RemoveEnum(child.GetName())
	case *builder.ServiceBuilder:
		fileBuilder.RemoveService(child.GetName())
	case *builder.FieldBuilder:
		fileBuilder.RemoveExtension(child.GetName())
	}
}

func filterMessage(messageBuilder *builder.MessageBuilder, terms *set.Set) (bool, error) {
	mDesc, err := messageBuilder.Build()
	if err != nil {
		return false, err
	}
	if mOptions := mDesc.GetMessageOptions(); mOptions != nil {
		extVal, err := proto.GetExtension(mOptions, filter.E_Message)
		if err != proto.ErrMissingExtension {
			if err != nil {
				return false, err
			}
			if isExcluded(extVal, terms) {
				return true, nil
			}
		}
	}

	for _, child := range messageBuilder.GetChildren() {
		if isExcluded, err := filterChild(child, terms); err != nil {
			return false, err
		} else if isExcluded {
			removeMessageChild(messageBuilder, child)
		}
	}

	return false, nil
}

func removeMessageChild(messageBuilder *builder.MessageBuilder, child builder.Builder) {
	switch c := child.(type) {
	case *builder.FieldBuilder:
		if c.IsExtension() {
			messageBuilder.RemoveNestedExtension(c.GetName())
		} else {
			messageBuilder.RemoveField(c.GetName())
		}
	case *builder.OneOfBuilder:
		messageBuilder.RemoveOneOf(c.GetName())
	case *builder.MessageBuilder:
		messageBuilder.RemoveNestedMessage(c.GetName())
	case *builder.EnumBuilder:
		messageBuilder.RemoveNestedEnum(c.GetName())
	}
}

func filterEnum(enumBuilder *builder.EnumBuilder, terms *set.Set) (bool, error) {
	eDesc, err := enumBuilder.Build()
	if err != nil {
		return false, err
	}
	if eOptions := eDesc.GetEnumOptions(); eOptions != nil {
		extVal, err := proto.GetExtension(eOptions, filter.E_Enum)
		if err != proto.ErrMissingExtension {
			if err != nil {
				return false, err
			}
			if isExcluded(extVal, terms) {
				return true, nil
			}
		}
	}

	for _, child := range enumBuilder.GetChildren() {
		if isExcluded, err := filterChild(child, terms); err != nil {
			return false, err
		} else if isExcluded {
			removeEnumChild(enumBuilder, child)
		}
	}

	return false, nil
}

func removeEnumChild(enumBuilder *builder.EnumBuilder, child builder.Builder) {
	switch c := child.(type) {
	case *builder.EnumValueBuilder:
		enumBuilder.RemoveValue(c.GetName())
	}
}

func filterEnumValue(enumValueBuilder *builder.EnumValueBuilder, terms *set.Set) (bool, error) {
	evDesc, err := enumValueBuilder.Build()
	if err != nil {
		return false, err
	}
	if evOptions := evDesc.GetEnumValueOptions(); evOptions != nil {
		extVal, err := proto.GetExtension(evOptions, filter.E_EnumValue)
		if err != proto.ErrMissingExtension {
			if err != nil {
				return false, err
			}
			if isExcluded(extVal, terms) {
				return true, nil
			}
		}
	}

	for _, child := range enumValueBuilder.GetChildren() {
		if isExcluded, err := filterChild(child, terms); err != nil {
			return false, err
		} else if isExcluded {
			removeEnumValueChild(enumValueBuilder, child)
		}
	}

	return false, nil
}

func removeEnumValueChild(_ *builder.EnumValueBuilder, _ builder.Builder) {
	return
}

func filterService(serviceBuilder *builder.ServiceBuilder, terms *set.Set) (bool, error) {
	sDesc, err := serviceBuilder.Build()
	if err != nil {
		return false, err
	}
	if sOptions := sDesc.GetServiceOptions(); sOptions != nil {
		extVal, err := proto.GetExtension(sOptions, filter.E_Service)
		if err != proto.ErrMissingExtension {
			if err != nil {
				return false, err
			}
			if isExcluded(extVal, terms) {
				return true, nil
			}
		}
	}

	for _, child := range serviceBuilder.GetChildren() {
		if isExcluded, err := filterChild(child, terms); err != nil {
			return false, err
		} else if isExcluded {
			removeServiceChild(serviceBuilder, child)
		}
	}

	return false, nil
}

func removeServiceChild(serviceBuilder *builder.ServiceBuilder, child builder.Builder) {
	switch c := child.(type) {
	case *builder.MethodBuilder:
		serviceBuilder.RemoveMethod(c.GetName())
	}
}

func filterMethod(methodBuilder *builder.MethodBuilder, terms *set.Set) (bool, error) {
	mDesc, err := methodBuilder.Build()
	if err != nil {
		return false, err
	}
	if mOptions := mDesc.GetMethodOptions(); mOptions != nil {
		extVal, err := proto.GetExtension(mOptions, filter.E_Method)
		if err != proto.ErrMissingExtension {
			if err != nil {
				return false, err
			}
			if isExcluded(extVal, terms) {
				return true, nil
			}
		}
	}

	for _, child := range methodBuilder.GetChildren() {
		if isExcluded, err := filterChild(child, terms); err != nil {
			return false, err
		} else if isExcluded {
			removeMethodChild(methodBuilder, child)
		}
	}

	return false, nil
}

func removeMethodChild(_ *builder.MethodBuilder, _ builder.Builder) {
	return
}

func filterField(fieldBuilder *builder.FieldBuilder, terms *set.Set) (bool, error) {
	fDesc, err := fieldBuilder.Build()
	if err != nil {
		return false, err
	}
	if fOptions := fDesc.GetFieldOptions(); fOptions != nil {
		extVal, err := proto.GetExtension(fOptions, filter.E_Field)
		if err != proto.ErrMissingExtension {
			if err != nil {
				return false, err
			}
			if isExcluded(extVal, terms) {
				return true, nil
			}
		}
	}

	for _, child := range fieldBuilder.GetChildren() {
		if isExcluded, err := filterChild(child, terms); err != nil {
			return false, err
		} else if isExcluded {
			removeFieldChild(fieldBuilder, child)
		}
	}

	return false, nil
}

func removeFieldChild(_ *builder.FieldBuilder, _ builder.Builder) {
	return
}

func filterOneOf(oneOfBuilder *builder.OneOfBuilder, terms *set.Set) (bool, error) {
	oDesc, err := oneOfBuilder.Build()
	if err != nil {
		return false, err
	}
	if oOptions := oDesc.GetOneOfOptions(); oOptions != nil {
		extVal, err := proto.GetExtension(oOptions, filter.E_OneOf)
		if err != proto.ErrMissingExtension {
			if err != nil {
				return false, err
			}
			if isExcluded(extVal, terms) {
				return true, nil
			}
		}
	}

	for _, child := range oneOfBuilder.GetChildren() {
		if isExcluded, err := filterChild(child, terms); err != nil {
			return false, err
		} else if isExcluded {
			removeOneOfChild(oneOfBuilder, child)
		}
	}

	return false, nil
}

func removeOneOfChild(oneOfBuilder *builder.OneOfBuilder, child builder.Builder) {
	switch c := child.(type) {
	case *builder.FieldBuilder:
		oneOfBuilder.RemoveChoice(c.GetName())
	}
}

func filterChild(child builder.Builder, terms *set.Set) (bool, error) {
	switch c := child.(type) {
	case *builder.MessageBuilder:
		return filterMessage(c, terms)
	case *builder.EnumBuilder:
		return filterEnum(c, terms)
	case *builder.ServiceBuilder:
		return filterService(c, terms)
	case *builder.FieldBuilder:
		return filterField(c, terms)
	case *builder.MethodBuilder:
		return filterMethod(c, terms)
	case *builder.EnumValueBuilder:
		return filterEnumValue(c, terms)
	case *builder.OneOfBuilder:
		return filterOneOf(c, terms)
	default:
		return false, nil
	}
}

// isExcluded evaluates the filter rules based on the data in the ValueFilter
func isExcluded(extVal interface{}, terms *set.Set) bool {
	if terms == nil || terms.Len() == 0 {
		return false
	}
	filterVal := extVal.(*filter.ValueFilter)
	for _, item := range filterVal.GetExclude() {
		if terms.Exists(item) {
			return true
		}
	}
	for _, item := range filterVal.GetInclude() {
		if terms.Exists(item) {
			return false
		}
	}
	// If Include is empty we don't want to exclude the item by default.
	// If Include is not empty, we should only include it if is explicitly matching
	return len(filterVal.Include) != 0
}

package main

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
	dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc/builder"

	"github.com/Workiva/go-datastructures/set"
	"github.com/stretchr/testify/assert"
	"github.com/wdullaer/proto-filter/filter"
)

func TestIsExcluded(t *testing.T) {
	cases := []struct {
		name   string
		input  *filter.ValueFilter
		terms  *set.Set
		output bool
	}{
		{
			name:   "Should return `false` when ValueFilter and Terms are empty",
			input:  &filter.ValueFilter{},
			terms:  set.New(),
			output: false,
		},
		{
			name:   "Should return `false` when ValueFilter is empty and Terms is not",
			input:  &filter.ValueFilter{},
			terms:  set.New("foo"),
			output: false,
		},
		{
			name:   "Should return `false` when term is in ValueFilter.Include",
			input:  &filter.ValueFilter{Include: []string{"foo"}},
			terms:  set.New("foo"),
			output: false,
		},
		{
			name:   "Should return `true` when term is in ValueFilter.Exclude",
			input:  &filter.ValueFilter{Exclude: []string{"foo"}},
			terms:  set.New("foo"),
			output: true,
		},
		{
			name:   "Should return `true` when terms are in both ValueFilter.Exclude and ValueFilter.Include",
			input:  &filter.ValueFilter{Exclude: []string{"foo"}, Include: []string{"bar"}},
			terms:  set.New("foo", "bar"),
			output: true,
		},
		{
			name:   "Should return `true` when terms are not in ValueFilter.Include, but ValueFilter.Include is not empty",
			input:  &filter.ValueFilter{Exclude: []string{}, Include: []string{"foo"}},
			terms:  set.New("bar"),
			output: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.output, isExcluded(tc.input, tc.terms))
		})
	}
}

func getEnumValueFilter(exclude []string, include []string) *dpb.EnumValueOptions {
	filt := &filter.ValueFilter{
		Include: include,
		Exclude: exclude,
	}

	result := &dpb.EnumValueOptions{}
	if err := proto.SetExtension(result, filter.E_EnumValue, filt); err != nil {
		fmt.Println(err)
	}
	return result
}

// TestFilterEnumValue does not exhaustively test the filter logic (that happens in TestIsExcluded)
// The test cases are there to test the data transformation logic that happens
func TestFilterEnumValue(t *testing.T) {
	cases := []struct {
		name    string
		input   *builder.EnumValueBuilder
		terms   *set.Set
		output  bool
		isError bool
	}{
		{
			name:    "Should return `false` if input does not have the extension",
			input:   builder.NewEnumValue("enumvalue"),
			terms:   set.New(),
			output:  false,
			isError: false,
		},
		{
			name:    "Should return `false` if input has extension with term in Included",
			input:   builder.NewEnumValue("enumvalue").SetOptions(getEnumValueFilter([]string{}, []string{"foo"})),
			terms:   set.New("foo"),
			output:  false,
			isError: false,
		},
		{
			name:    "Should return `true` if input has extension with term in Excluded",
			input:   builder.NewEnumValue("enumvalue").SetOptions(getEnumValueFilter([]string{"foo"}, []string{})),
			terms:   set.New("foo"),
			output:  true,
			isError: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// EnumValue must be part of an enum for filterEnumValue to work
			builder.NewEnum("enum").AddValue(tc.input)
			if result, err := filterEnumValue(tc.input, tc.terms); assert.NoError(t, err) {
				assert.Equal(t, tc.output, result)
			}
		})
	}
}

func getFieldFilter(exclude []string, include []string) *dpb.FieldOptions {
	filt := &filter.ValueFilter{
		Include: include,
		Exclude: exclude,
	}

	result := &dpb.FieldOptions{}
	if err := proto.SetExtension(result, filter.E_Field, filt); err != nil {
		fmt.Println(err)
	}
	return result
}

// TestFilterEnumValue does not exhaustively test the filter logic (that happens in TestIsExcluded)
// The test cases are there to test the data transformation logic that happens
func TestFilterField(t *testing.T) {
	cases := []struct {
		name    string
		input   *builder.FieldBuilder
		terms   *set.Set
		output  bool
		isError bool
	}{
		{
			name:    "Should return `false` if input does not have the extension",
			input:   builder.NewField("field", builder.FieldTypeString()),
			terms:   set.New(),
			output:  false,
			isError: false,
		},
		{
			name:    "Should return `false` if input has extension with term in Included",
			input:   builder.NewField("field", builder.FieldTypeString()).SetOptions(getFieldFilter([]string{}, []string{"foo"})),
			terms:   set.New("foo"),
			output:  false,
			isError: false,
		},
		{
			name:    "Should return `true` if input has extension with term in Excluded",
			input:   builder.NewField("field", builder.FieldTypeString()).SetOptions(getFieldFilter([]string{"foo"}, []string{})),
			terms:   set.New("foo"),
			output:  true,
			isError: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// EnumValue must be part of an enum for filterEnumValue to work
			builder.NewMessage("message").AddField(tc.input)
			if result, err := filterField(tc.input, tc.terms); assert.NoError(t, err) {
				assert.Equal(t, tc.output, result)
			}
		})
	}
}

func getMethodFilter(exclude []string, include []string) *dpb.MethodOptions {
	filt := &filter.ValueFilter{
		Include: include,
		Exclude: exclude,
	}

	result := &dpb.MethodOptions{}
	if err := proto.SetExtension(result, filter.E_Method, filt); err != nil {
		fmt.Println(err)
	}
	return result
}

// TestFilterEnumValue does not exhaustively test the filter logic (that happens in TestIsExcluded)
// The test cases are there to test the data transformation logic that happens
func TestFilterMethod(t *testing.T) {
	var emptyRPCType = builder.RpcTypeMessage(builder.NewMessage("empty"), false)
	cases := []struct {
		name    string
		input   *builder.MethodBuilder
		terms   *set.Set
		output  bool
		isError bool
	}{
		{
			name:    "Should return `false` if input does not have the extension",
			input:   builder.NewMethod("method", emptyRPCType, emptyRPCType),
			terms:   set.New(),
			output:  false,
			isError: false,
		},
		{
			name:    "Should return `false` if input has extension with term in Included",
			input:   builder.NewMethod("method", emptyRPCType, emptyRPCType).SetOptions(getMethodFilter([]string{}, []string{"foo"})),
			terms:   set.New("foo"),
			output:  false,
			isError: false,
		},
		{
			name:    "Should return `true` if input has extension with term in Excluded",
			input:   builder.NewMethod("method", emptyRPCType, emptyRPCType).SetOptions(getMethodFilter([]string{"foo"}, []string{})),
			terms:   set.New("foo"),
			output:  true,
			isError: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// EnumValue must be part of an enum for filterEnumValue to work
			builder.NewService("service").AddMethod(tc.input)
			if result, err := filterMethod(tc.input, tc.terms); assert.NoError(t, err) {
				assert.Equal(t, tc.output, result)
			}
		})
	}
}

func getServiceFilter(exclude []string, include []string) *dpb.ServiceOptions {
	filt := &filter.ValueFilter{
		Include: include,
		Exclude: exclude,
	}

	result := &dpb.ServiceOptions{}
	if err := proto.SetExtension(result, filter.E_Service, filt); err != nil {
		fmt.Println(err)
	}
	return result
}

// TestFilterEnumValue does not exhaustively test the filter logic (that happens in TestIsExcluded)
// The test cases are there to test the data transformation logic that happens
func TestFilterService(t *testing.T) {
	var emptyRPCType = builder.RpcTypeMessage(builder.NewMessage("empty"), false)
	cases := []struct {
		name             string
		input            *builder.ServiceBuilder
		terms            *set.Set
		expectedChildren []string
		output           bool
		isError          bool
	}{
		{
			name:             "Should return `false` if input does not have the extension",
			input:            builder.NewService("service"),
			terms:            set.New(),
			expectedChildren: []string{},
			output:           false,
			isError:          false,
		},
		{
			name:             "Should return `false` if input has extension with term in Included",
			input:            builder.NewService("service").SetOptions(getServiceFilter([]string{}, []string{"foo"})),
			terms:            set.New("foo"),
			expectedChildren: []string{},
			output:           false,
			isError:          false,
		},
		{
			name:             "Should return `true` if input has extension with term in Excluded",
			input:            builder.NewService("service").SetOptions(getServiceFilter([]string{"foo"}, []string{})),
			terms:            set.New("foo"),
			expectedChildren: []string{},
			output:           true,
			isError:          false,
		},
		{
			name: "Should include all children without the extension",
			input: builder.NewService("service").
				AddMethod(builder.NewMethod("method1", emptyRPCType, emptyRPCType)).
				AddMethod(builder.NewMethod("method2", emptyRPCType, emptyRPCType)),
			terms:            set.New("foo"),
			expectedChildren: []string{"method1", "method2"},
			output:           false,
			isError:          false,
		},
		{
			name: "Should remove all children that have the term in excluded",
			input: builder.NewService("service").
				AddMethod(builder.NewMethod("method1", emptyRPCType, emptyRPCType).SetOptions(getMethodFilter([]string{"foo"}, []string{}))).
				AddMethod(builder.NewMethod("method2", emptyRPCType, emptyRPCType).SetOptions(getMethodFilter([]string{"bar"}, []string{}))),
			terms:            set.New("foo"),
			expectedChildren: []string{"method2"},
			output:           false,
			isError:          false,
		},
		{
			name: "Should include all children that have the term in included",
			input: builder.NewService("service").
				AddMethod(builder.NewMethod("method1", emptyRPCType, emptyRPCType).SetOptions(getMethodFilter([]string{}, []string{"foo"}))).
				AddMethod(builder.NewMethod("method2", emptyRPCType, emptyRPCType).SetOptions(getMethodFilter([]string{}, []string{"bar"}))),
			terms:            set.New("foo"),
			expectedChildren: []string{"method1"},
			output:           false,
			isError:          false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if result, err := filterService(tc.input, tc.terms); assert.NoError(t, err) {
				assert.Equal(t, tc.output, result)
				children := make([]string, len(tc.input.GetChildren()))
				for i, v := range tc.input.GetChildren() {
					children[i] = v.GetName()
				}
				assert.ElementsMatch(t, tc.expectedChildren, children)
			}
		})
	}
}

func getEnumFilter(exclude []string, include []string) *dpb.EnumOptions {
	filt := &filter.ValueFilter{
		Include: include,
		Exclude: exclude,
	}

	result := &dpb.EnumOptions{}
	if err := proto.SetExtension(result, filter.E_Enum, filt); err != nil {
		fmt.Println(err)
	}
	return result
}

// TestFilterEnumValue does not exhaustively test the filter logic (that happens in TestIsExcluded)
// The test cases are there to test the data transformation logic that happens
func TestFilterEnum(t *testing.T) {
	cases := []struct {
		name             string
		input            *builder.EnumBuilder
		terms            *set.Set
		expectedChildren []string
		output           bool
		isError          bool
	}{
		{
			name:             "Should return `false` if input does not have the extension",
			input:            builder.NewEnum("enum"),
			terms:            set.New(),
			expectedChildren: []string{},
			output:           false,
			isError:          false,
		},
		{
			name:             "Should return `false` if input has extension with term in Included",
			input:            builder.NewEnum("enum").SetOptions(getEnumFilter([]string{}, []string{"foo"})),
			terms:            set.New("foo"),
			expectedChildren: []string{},
			output:           false,
			isError:          false,
		},
		{
			name:             "Should return `true` if input has extension with term in Excluded",
			input:            builder.NewEnum("enum").SetOptions(getEnumFilter([]string{"foo"}, []string{})),
			terms:            set.New("foo"),
			expectedChildren: []string{},
			output:           true,
			isError:          false,
		},
		{
			name: "Should include all children without the extension",
			input: builder.NewEnum("enum").
				AddValue(builder.NewEnumValue("enumvalue1")).
				AddValue(builder.NewEnumValue("enumvalue2")),
			terms:            set.New("foo"),
			expectedChildren: []string{"enumvalue1", "enumvalue2"},
			output:           false,
			isError:          false,
		},
		{
			name: "Should remove all children that have the term in excluded",
			input: builder.NewEnum("enum").
				AddValue(builder.NewEnumValue("enumvalue1").SetOptions(getEnumValueFilter([]string{"foo"}, []string{}))).
				AddValue(builder.NewEnumValue("enumvalue2").SetOptions(getEnumValueFilter([]string{"bar"}, []string{}))),
			terms:            set.New("foo"),
			expectedChildren: []string{"enumvalue2"},
			output:           false,
			isError:          false,
		},
		{
			name: "Should include all children that have the term in included",
			input: builder.NewEnum("enum").
				AddValue(builder.NewEnumValue("enumvalue1").SetOptions(getEnumValueFilter([]string{}, []string{"foo"}))).
				AddValue(builder.NewEnumValue("enumvalue2").SetOptions(getEnumValueFilter([]string{}, []string{"bar"}))),
			terms:            set.New("foo"),
			expectedChildren: []string{"enumvalue1"},
			output:           false,
			isError:          false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if result, err := filterEnum(tc.input, tc.terms); assert.NoError(t, err) {
				assert.Equal(t, tc.output, result)
				children := make([]string, len(tc.input.GetChildren()))
				for i, v := range tc.input.GetChildren() {
					children[i] = v.GetName()
				}
				assert.ElementsMatch(t, tc.expectedChildren, children)
			}
		})
	}
}

func getOneOfFilter(exclude []string, include []string) *dpb.OneofOptions {
	filt := &filter.ValueFilter{
		Include: include,
		Exclude: exclude,
	}

	result := &dpb.OneofOptions{}
	if err := proto.SetExtension(result, filter.E_OneOf, filt); err != nil {
		fmt.Println(err)
	}
	return result
}

// TestFilterOneOf does not exhaustively test the filter logic (that happens in TestIsExcluded)
// The test cases are there to test the data transformation logic that happens
func TestFilterOneOf(t *testing.T) {
	cases := []struct {
		name             string
		input            *builder.OneOfBuilder
		terms            *set.Set
		expectedChildren []string
		output           bool
		isError          bool
	}{
		{
			name:             "Should return `false` if input does not have the extension",
			input:            builder.NewOneOf("oneof"),
			terms:            set.New(),
			expectedChildren: []string{},
			output:           false,
			isError:          false,
		},
		{
			name:             "Should return `false` if input has extension with term in Included",
			input:            builder.NewOneOf("oneof").SetOptions(getOneOfFilter([]string{}, []string{"foo"})),
			terms:            set.New("foo"),
			expectedChildren: []string{},
			output:           false,
			isError:          false,
		},
		{
			name:             "Should return `true` if input has extension with term in Excluded",
			input:            builder.NewOneOf("oneof").SetOptions(getOneOfFilter([]string{"foo"}, []string{})),
			terms:            set.New("foo"),
			expectedChildren: []string{},
			output:           true,
			isError:          false,
		},
		{
			name: "Should include all children without the extension",
			input: builder.NewOneOf("oneof").
				AddChoice(builder.NewField("oneofchoice1", builder.FieldTypeString())).
				AddChoice(builder.NewField("oneofchoice2", builder.FieldTypeString())),
			terms:            set.New("foo"),
			expectedChildren: []string{"oneofchoice1", "oneofchoice2"},
			output:           false,
			isError:          false,
		},
		{
			name: "Should remove all children that have the term in excluded",
			input: builder.NewOneOf("oneof").
				AddChoice(builder.NewField("oneofchoice1", builder.FieldTypeString()).SetOptions(getFieldFilter([]string{"foo"}, []string{}))).
				AddChoice(builder.NewField("oneofchoice2", builder.FieldTypeString()).SetOptions(getFieldFilter([]string{"bar"}, []string{}))),
			terms:            set.New("foo"),
			expectedChildren: []string{"oneofchoice2"},
			output:           false,
			isError:          false,
		},
		{
			name: "Should include all children that have the term in included",
			input: builder.NewOneOf("oneof").
				AddChoice(builder.NewField("oneofchoice1", builder.FieldTypeString()).SetOptions(getFieldFilter([]string{}, []string{"foo"}))).
				AddChoice(builder.NewField("oneofchoice2", builder.FieldTypeString()).SetOptions(getFieldFilter([]string{}, []string{"bar"}))),
			terms:            set.New("foo"),
			expectedChildren: []string{"oneofchoice1"},
			output:           false,
			isError:          false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// one_of must be part of a message for .Build() to work
			builder.NewMessage("message").AddOneOf(tc.input)
			if result, err := filterOneOf(tc.input, tc.terms); assert.NoError(t, err) {
				assert.Equal(t, tc.output, result)
				children := make([]string, len(tc.input.GetChildren()))
				for i, v := range tc.input.GetChildren() {
					children[i] = v.GetName()
				}
				assert.ElementsMatch(t, tc.expectedChildren, children)
			}
		})
	}
}

func getMessageFilter(exclude []string, include []string) *dpb.MessageOptions {
	filt := &filter.ValueFilter{
		Include: include,
		Exclude: exclude,
	}

	result := &dpb.MessageOptions{}
	if err := proto.SetExtension(result, filter.E_Message, filt); err != nil {
		fmt.Println(err)
	}
	return result
}

// TestFilterEnumValue does not exhaustively test the filter logic (that happens in TestIsExcluded)
// The test cases are there to test the data transformation logic that happens
func TestFilterMessage(t *testing.T) {
	cases := []struct {
		name             string
		input            *builder.MessageBuilder
		terms            *set.Set
		expectedChildren []string
		output           bool
		isError          bool
	}{
		{
			name:             "Should return `false` if input does not have the extension",
			input:            builder.NewMessage("message"),
			terms:            set.New(),
			expectedChildren: []string{},
			output:           false,
			isError:          false,
		},
		{
			name:             "Should return `false` if input has extension with term in Included",
			input:            builder.NewMessage("message").SetOptions(getMessageFilter([]string{}, []string{"foo"})),
			terms:            set.New("foo"),
			expectedChildren: []string{},
			output:           false,
			isError:          false,
		},
		{
			name:             "Should return `true` if input has extension with term in Excluded",
			input:            builder.NewMessage("message").SetOptions(getMessageFilter([]string{"foo"}, []string{})),
			terms:            set.New("foo"),
			expectedChildren: []string{},
			output:           true,
			isError:          false,
		},
		{
			name: "Should include all children without the extension",
			input: builder.NewMessage("message").
				AddField(builder.NewField("field1", builder.FieldTypeString())).
				AddField(builder.NewField("field2", builder.FieldTypeString())),
			terms:            set.New("foo"),
			expectedChildren: []string{"field1", "field2"},
			output:           false,
			isError:          false,
		},
		{
			name: "Should remove all children that have the term in excluded",
			input: builder.NewMessage("message").
				AddField(builder.NewField("field1", builder.FieldTypeString()).SetOptions(getFieldFilter([]string{"foo"}, []string{}))).
				AddField(builder.NewField("field2", builder.FieldTypeString()).SetOptions(getFieldFilter([]string{"bar"}, []string{}))),
			terms:            set.New("foo"),
			expectedChildren: []string{"field2"},
			output:           false,
			isError:          false,
		},
		{
			name: "Should include all children that have the term in included",
			input: builder.NewMessage("message").
				AddField(builder.NewField("field1", builder.FieldTypeString()).SetOptions(getFieldFilter([]string{}, []string{"foo"}))).
				AddField(builder.NewField("field2", builder.FieldTypeString()).SetOptions(getFieldFilter([]string{}, []string{"bar"}))),
			terms:            set.New("foo"),
			expectedChildren: []string{"field1"},
			output:           false,
			isError:          false,
		},
		{
			name: "Should remove excluded nested message",
			input: builder.NewMessage("message").
				AddNestedMessage(builder.NewMessage("nestedmessage").SetOptions(getMessageFilter([]string{"foo"}, []string{}))).
				AddField(builder.NewField("field1", builder.FieldTypeString()).SetOptions(getFieldFilter([]string{"foo"}, []string{}))).
				AddField(builder.NewField("field2", builder.FieldTypeString()).SetOptions(getFieldFilter([]string{"bar"}, []string{}))),
			terms:            set.New("foo"),
			expectedChildren: []string{"field2"},
			output:           false,
			isError:          false,
		},
		{
			name: "Should remove excluded nested enum",
			input: builder.NewMessage("message").
				AddNestedEnum(builder.NewEnum("nestedenum").SetOptions(getEnumFilter([]string{"foo"}, []string{}))).
				AddField(builder.NewField("field1", builder.FieldTypeString()).SetOptions(getFieldFilter([]string{"foo"}, []string{}))).
				AddField(builder.NewField("field2", builder.FieldTypeString()).SetOptions(getFieldFilter([]string{"bar"}, []string{}))),
			terms:            set.New("foo"),
			expectedChildren: []string{"field2"},
			output:           false,
			isError:          false,
		},
		{
			name: "Should remove excluded nested extension",
			input: builder.NewMessage("message").
				AddNestedExtension(builder.NewExtension("nestedextension", 12345, builder.FieldTypeString(), builder.NewMessage("extendee")).SetOptions(getFieldFilter([]string{"foo"}, []string{}))).
				AddField(builder.NewField("field1", builder.FieldTypeString()).SetOptions(getFieldFilter([]string{"foo"}, []string{}))).
				AddField(builder.NewField("field2", builder.FieldTypeString()).SetOptions(getFieldFilter([]string{"bar"}, []string{}))),
			terms:            set.New("foo"),
			expectedChildren: []string{"field2"},
			output:           false,
			isError:          false,
		},
		{
			name: "Should remove excluded oneof",
			input: builder.NewMessage("message").
				AddOneOf(builder.NewOneOf("oneof").AddChoice(builder.NewField("oneoffield", builder.FieldTypeString())).SetOptions(getOneOfFilter([]string{"foo"}, []string{}))).
				AddField(builder.NewField("field1", builder.FieldTypeString()).SetOptions(getFieldFilter([]string{"foo"}, []string{}))).
				AddField(builder.NewField("field2", builder.FieldTypeString()).SetOptions(getFieldFilter([]string{"bar"}, []string{}))),
			terms:            set.New("foo"),
			expectedChildren: []string{"field2"},
			output:           false,
			isError:          false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if result, err := filterMessage(tc.input, tc.terms); assert.NoError(t, err) {
				assert.Equal(t, tc.output, result)
				children := make([]string, len(tc.input.GetChildren()))
				for i, v := range tc.input.GetChildren() {
					children[i] = v.GetName()
				}
				assert.ElementsMatch(t, tc.expectedChildren, children)
			}
		})
	}
}

func getFileFilter(exclude []string, include []string) *dpb.FileOptions {
	filt := &filter.ValueFilter{
		Include: include,
		Exclude: exclude,
	}

	result := &dpb.FileOptions{}
	if err := proto.SetExtension(result, filter.E_File, filt); err != nil {
		fmt.Println(err)
	}
	return result
}

// TestFilterFile does not exhaustively test the filter logic (that happens in TestIsExcluded)
// The test cases are there to test the data transformation logic that happens
func TestFilterFile(t *testing.T) {
	cases := []struct {
		name             string
		input            *builder.FileBuilder
		terms            *set.Set
		expectedChildren []string
		output           bool
		isError          bool
	}{
		{
			name:             "Should return `false` if input does not have the extension",
			input:            builder.NewFile("file"),
			terms:            set.New(),
			expectedChildren: []string{},
			output:           false,
			isError:          false,
		},
		{
			name:             "Should return `false` if input has extension with term in Included",
			input:            builder.NewFile("file").SetOptions(getFileFilter([]string{}, []string{"foo"})),
			terms:            set.New("foo"),
			expectedChildren: []string{},
			output:           false,
			isError:          false,
		},
		{
			name:             "Should return `true` if input has extension with term in Excluded",
			input:            builder.NewFile("file").SetOptions(getFileFilter([]string{"foo"}, []string{})),
			terms:            set.New("foo"),
			expectedChildren: []string{},
			output:           true,
			isError:          false,
		},
		{
			name: "Should include all children without the extension",
			input: builder.NewFile("file").
				AddMessage(builder.NewMessage("message1")).
				AddMessage(builder.NewMessage("message2")),
			terms:            set.New("foo"),
			expectedChildren: []string{"message1", "message2"},
			output:           false,
			isError:          false,
		},
		{
			name: "Should remove all children that have the term in excluded",
			input: builder.NewFile("file").
				AddMessage(builder.NewMessage("message1").SetOptions(getMessageFilter([]string{"foo"}, []string{}))).
				AddMessage(builder.NewMessage("message2").SetOptions(getMessageFilter([]string{"bar"}, []string{}))),
			terms:            set.New("foo"),
			expectedChildren: []string{"message2"},
			output:           false,
			isError:          false,
		},
		{
			name: "Should include all children that have the term in included",
			input: builder.NewFile("file").
				AddMessage(builder.NewMessage("message1").SetOptions(getMessageFilter([]string{}, []string{"foo"}))).
				AddMessage(builder.NewMessage("message2").SetOptions(getMessageFilter([]string{}, []string{"bar"}))),
			terms:            set.New("foo"),
			expectedChildren: []string{"message1"},
			output:           false,
			isError:          false,
		},
		{
			name: "Should remove excluded message",
			input: builder.NewFile("file").
				AddMessage(builder.NewMessage("message").SetOptions(getMessageFilter([]string{"foo"}, []string{}))).
				AddMessage(builder.NewMessage("message1").SetOptions(getMessageFilter([]string{}, []string{"foo"}))).
				AddMessage(builder.NewMessage("message2").SetOptions(getMessageFilter([]string{}, []string{"bar"}))),
			terms:            set.New("foo"),
			expectedChildren: []string{"message1"},
			output:           false,
			isError:          false,
		},
		{
			name: "Should remove excluded enum",
			input: builder.NewFile("file").
				AddEnum(builder.NewEnum("enum").SetOptions(getEnumFilter([]string{"foo"}, []string{}))).
				AddMessage(builder.NewMessage("message1").SetOptions(getMessageFilter([]string{}, []string{"foo"}))).
				AddMessage(builder.NewMessage("message2").SetOptions(getMessageFilter([]string{}, []string{"bar"}))),
			terms:            set.New("foo"),
			expectedChildren: []string{"message1"},
			output:           false,
			isError:          false,
		},
		{
			name: "Should remove excluded service",
			input: builder.NewFile("file").
				AddService(builder.NewService("service").SetOptions(getServiceFilter([]string{"foo"}, []string{}))).
				AddMessage(builder.NewMessage("message1").SetOptions(getMessageFilter([]string{}, []string{"foo"}))).
				AddMessage(builder.NewMessage("message2").SetOptions(getMessageFilter([]string{}, []string{"bar"}))),
			terms:            set.New("foo"),
			expectedChildren: []string{"message1"},
			output:           false,
			isError:          false,
		},
		{
			name: "Should remove excluded extension",
			input: builder.NewFile("file").
				AddExtension(builder.NewExtension("extension", 12345, builder.FieldTypeString(), builder.NewMessage("extendee")).SetOptions(getFieldFilter([]string{"foo"}, []string{}))).
				AddMessage(builder.NewMessage("message1").SetOptions(getMessageFilter([]string{}, []string{"foo"}))).
				AddMessage(builder.NewMessage("message2").SetOptions(getMessageFilter([]string{}, []string{"bar"}))),
			terms:            set.New("foo"),
			expectedChildren: []string{"message1"},
			output:           false,
			isError:          false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if result, err := filterFile(tc.input, tc.terms); assert.NoError(t, err) {
				assert.Equal(t, tc.output, result)
				children := make([]string, len(tc.input.GetChildren()))
				for i, v := range tc.input.GetChildren() {
					children[i] = v.GetName()
				}
				assert.ElementsMatch(t, tc.expectedChildren, children)
			}
		})
	}
}

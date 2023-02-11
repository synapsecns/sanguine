package configschema

import (
	coerce_value_fmt "fmt"
	nestingmode_string_strconv "strconv"

	coerce_value_cty "github.com/hashicorp/go-cty/cty"
	empty_value_cty "github.com/hashicorp/go-cty/cty"
	implied_type_cty "github.com/hashicorp/go-cty/cty"
	schema_cty "github.com/hashicorp/go-cty/cty"
	coerce_value_convert "github.com/hashicorp/go-cty/cty/convert"
)

func (b *Block) CoerceValue(in coerce_value_cty.Value) (coerce_value_cty.Value, error) {
	var path coerce_value_cty.Path
	return b.coerceValue(in, path)
}

func (b *Block) coerceValue(in coerce_value_cty.Value, path coerce_value_cty.Path) (coerce_value_cty.Value, error) {
	switch {
	case in.IsNull():
		return coerce_value_cty.NullVal(b.ImpliedType()), nil
	case !in.IsKnown():
		return coerce_value_cty.UnknownVal(b.ImpliedType()), nil
	}

	ty := in.Type()
	if !ty.IsObjectType() {
		return coerce_value_cty.UnknownVal(b.ImpliedType()), path.NewErrorf("an object is required")
	}

	for name := range ty.AttributeTypes() {
		if _, defined := b.Attributes[name]; defined {
			continue
		}
		if _, defined := b.BlockTypes[name]; defined {
			continue
		}
		return coerce_value_cty.UnknownVal(b.ImpliedType()), path.NewErrorf("unexpected attribute %q", name)
	}

	attrs := make(map[string]coerce_value_cty.Value)

	for name, attrS := range b.Attributes {
		var val coerce_value_cty.Value
		switch {
		case ty.HasAttribute(name):
			val = in.GetAttr(name)
		case attrS.Computed || attrS.Optional:
			val = coerce_value_cty.NullVal(attrS.Type)
		default:
			return coerce_value_cty.UnknownVal(b.ImpliedType()), path.NewErrorf("attribute %q is required", name)
		}

		val, err := attrS.coerceValue(val, append(path, coerce_value_cty.GetAttrStep{Name: name}))
		if err != nil {
			return coerce_value_cty.UnknownVal(b.ImpliedType()), err
		}

		attrs[name] = val
	}
	for typeName, blockS := range b.BlockTypes {
		switch blockS.Nesting {

		case NestingSingle, NestingGroup:
			switch {
			case ty.HasAttribute(typeName):
				var err error
				val := in.GetAttr(typeName)
				attrs[typeName], err = blockS.coerceValue(val, append(path, coerce_value_cty.GetAttrStep{Name: typeName}))
				if err != nil {
					return coerce_value_cty.UnknownVal(b.ImpliedType()), err
				}
			default:
				attrs[typeName] = blockS.EmptyValue()
			}

		case NestingList:
			switch {
			case ty.HasAttribute(typeName):
				coll := in.GetAttr(typeName)

				switch {
				case coll.IsNull():
					attrs[typeName] = coerce_value_cty.NullVal(coerce_value_cty.List(blockS.ImpliedType()))
					continue
				case !coll.IsKnown():
					attrs[typeName] = coerce_value_cty.UnknownVal(coerce_value_cty.List(blockS.ImpliedType()))
					continue
				}

				if !coll.CanIterateElements() {
					return coerce_value_cty.UnknownVal(b.ImpliedType()), path.NewErrorf("must be a list")
				}
				l := coll.LengthInt()

				if l == 0 {
					attrs[typeName] = coerce_value_cty.ListValEmpty(blockS.ImpliedType())
					continue
				}
				elems := make([]coerce_value_cty.Value, 0, l)
				{
					path = append(path, coerce_value_cty.GetAttrStep{Name: typeName})
					for it := coll.ElementIterator(); it.Next(); {
						var err error
						idx, val := it.Element()
						val, err = blockS.coerceValue(val, append(path, coerce_value_cty.IndexStep{Key: idx}))
						if err != nil {
							return coerce_value_cty.UnknownVal(b.ImpliedType()), err
						}
						elems = append(elems, val)
					}
				}
				attrs[typeName] = coerce_value_cty.ListVal(elems)
			default:
				attrs[typeName] = coerce_value_cty.ListValEmpty(blockS.ImpliedType())
			}

		case NestingSet:
			switch {
			case ty.HasAttribute(typeName):
				coll := in.GetAttr(typeName)

				switch {
				case coll.IsNull():
					attrs[typeName] = coerce_value_cty.NullVal(coerce_value_cty.Set(blockS.ImpliedType()))
					continue
				case !coll.IsKnown():
					attrs[typeName] = coerce_value_cty.UnknownVal(coerce_value_cty.Set(blockS.ImpliedType()))
					continue
				}

				if !coll.CanIterateElements() {
					return coerce_value_cty.UnknownVal(b.ImpliedType()), path.NewErrorf("must be a set")
				}
				l := coll.LengthInt()

				if l == 0 {
					attrs[typeName] = coerce_value_cty.SetValEmpty(blockS.ImpliedType())
					continue
				}
				elems := make([]coerce_value_cty.Value, 0, l)
				{
					path = append(path, coerce_value_cty.GetAttrStep{Name: typeName})
					for it := coll.ElementIterator(); it.Next(); {
						var err error
						idx, val := it.Element()
						val, err = blockS.coerceValue(val, append(path, coerce_value_cty.IndexStep{Key: idx}))
						if err != nil {
							return coerce_value_cty.UnknownVal(b.ImpliedType()), err
						}
						elems = append(elems, val)
					}
				}
				attrs[typeName] = coerce_value_cty.SetVal(elems)
			default:
				attrs[typeName] = coerce_value_cty.SetValEmpty(blockS.ImpliedType())
			}

		case NestingMap:
			switch {
			case ty.HasAttribute(typeName):
				coll := in.GetAttr(typeName)

				switch {
				case coll.IsNull():
					attrs[typeName] = coerce_value_cty.NullVal(coerce_value_cty.Map(blockS.ImpliedType()))
					continue
				case !coll.IsKnown():
					attrs[typeName] = coerce_value_cty.UnknownVal(coerce_value_cty.Map(blockS.ImpliedType()))
					continue
				}

				if !coll.CanIterateElements() {
					return coerce_value_cty.UnknownVal(b.ImpliedType()), path.NewErrorf("must be a map")
				}
				l := coll.LengthInt()
				if l == 0 {
					attrs[typeName] = coerce_value_cty.MapValEmpty(blockS.ImpliedType())
					continue
				}
				elems := make(map[string]coerce_value_cty.Value)
				{
					path = append(path, coerce_value_cty.GetAttrStep{Name: typeName})
					for it := coll.ElementIterator(); it.Next(); {
						var err error
						key, val := it.Element()
						if key.Type() != coerce_value_cty.String || key.IsNull() || !key.IsKnown() {
							return coerce_value_cty.UnknownVal(b.ImpliedType()), path.NewErrorf("must be a map")
						}
						val, err = blockS.coerceValue(val, append(path, coerce_value_cty.IndexStep{Key: key}))
						if err != nil {
							return coerce_value_cty.UnknownVal(b.ImpliedType()), err
						}
						elems[key.AsString()] = val
					}
				}

				useObject := false
				switch {
				case coll.Type().IsObjectType():
					useObject = true
				default:

					ety := coll.Type().ElementType()
					for _, v := range elems {
						if !v.Type().Equals(ety) {
							useObject = true
							break
						}
					}
				}

				if useObject {
					attrs[typeName] = coerce_value_cty.ObjectVal(elems)
				} else {
					attrs[typeName] = coerce_value_cty.MapVal(elems)
				}
			default:
				attrs[typeName] = coerce_value_cty.MapValEmpty(blockS.ImpliedType())
			}

		default:

			panic(coerce_value_fmt.Errorf("unsupported nesting mode %#v", blockS.Nesting))
		}
	}

	return coerce_value_cty.ObjectVal(attrs), nil
}

func (a *Attribute) coerceValue(in coerce_value_cty.Value, path coerce_value_cty.Path) (coerce_value_cty.Value, error) {
	val, err := coerce_value_convert.Convert(in, a.Type)
	if err != nil {
		return coerce_value_cty.UnknownVal(a.Type), path.NewError(err)
	}
	return val, nil
}

func (b *Block) EmptyValue() empty_value_cty.Value {
	vals := make(map[string]empty_value_cty.Value)
	for name, attrS := range b.Attributes {
		vals[name] = attrS.EmptyValue()
	}
	for name, blockS := range b.BlockTypes {
		vals[name] = blockS.EmptyValue()
	}
	return empty_value_cty.ObjectVal(vals)
}

func (a *Attribute) EmptyValue() empty_value_cty.Value {
	return empty_value_cty.NullVal(a.Type)
}

func (b *NestedBlock) EmptyValue() empty_value_cty.Value {
	switch b.Nesting {
	case NestingSingle:
		return empty_value_cty.NullVal(b.Block.ImpliedType())
	case NestingGroup:
		return b.Block.EmptyValue()
	case NestingList:
		if ty := b.Block.ImpliedType(); ty.HasDynamicTypes() {
			return empty_value_cty.EmptyTupleVal
		} else {
			return empty_value_cty.ListValEmpty(ty)
		}
	case NestingMap:
		if ty := b.Block.ImpliedType(); ty.HasDynamicTypes() {
			return empty_value_cty.EmptyObjectVal
		} else {
			return empty_value_cty.MapValEmpty(ty)
		}
	case NestingSet:
		return empty_value_cty.SetValEmpty(b.Block.ImpliedType())
	default:

		return empty_value_cty.NullVal(empty_value_cty.DynamicPseudoType)
	}
}

func (b *Block) ImpliedType() implied_type_cty.Type {
	if b == nil {
		return implied_type_cty.EmptyObject
	}

	atys := make(map[string]implied_type_cty.Type)

	for name, attrS := range b.Attributes {
		atys[name] = attrS.Type
	}

	for name, blockS := range b.BlockTypes {
		if _, exists := atys[name]; exists {
			panic("invalid schema, blocks and attributes cannot have the same name")
		}

		childType := blockS.Block.ImpliedType()

		switch blockS.Nesting {
		case NestingSingle, NestingGroup:
			atys[name] = childType
		case NestingList:

			if childType.HasDynamicTypes() {
				atys[name] = implied_type_cty.DynamicPseudoType
			} else {
				atys[name] = implied_type_cty.List(childType)
			}
		case NestingSet:
			if childType.HasDynamicTypes() {
				panic("can't use cty.DynamicPseudoType inside a block type with NestingSet")
			}
			atys[name] = implied_type_cty.Set(childType)
		case NestingMap:

			if childType.HasDynamicTypes() {
				atys[name] = implied_type_cty.DynamicPseudoType
			} else {
				atys[name] = implied_type_cty.Map(childType)
			}
		default:
			panic("invalid nesting type")
		}
	}

	return implied_type_cty.Object(atys)
}

func _() {

	var x [1]struct{}
	_ = x[nestingModeInvalid-0]
	_ = x[NestingSingle-1]
	_ = x[NestingGroup-2]
	_ = x[NestingList-3]
	_ = x[NestingSet-4]
	_ = x[NestingMap-5]
}

const _NestingMode_name = "nestingModeInvalidNestingSingleNestingGroupNestingListNestingSetNestingMap"

var _NestingMode_index = [...]uint8{0, 18, 31, 43, 54, 64, 74}

func (i NestingMode) String() string {
	if i < 0 || i >= NestingMode(len(_NestingMode_index)-1) {
		return "NestingMode(" + nestingmode_string_strconv.FormatInt(int64(i), 10) + ")"
	}
	return _NestingMode_name[_NestingMode_index[i]:_NestingMode_index[i+1]]
}

type StringKind int

const (
	StringPlain StringKind = iota

	StringMarkdown
)

type Block struct {
	Attributes map[string]*Attribute

	BlockTypes map[string]*NestedBlock

	Description     string
	DescriptionKind StringKind

	Deprecated bool
}

type Attribute struct {
	Type schema_cty.Type

	Description     string
	DescriptionKind StringKind

	Required bool

	Optional bool

	Computed bool

	Sensitive bool

	Deprecated bool
}

type NestedBlock struct {
	Block

	Nesting NestingMode

	MinItems, MaxItems int
}

type NestingMode int

const (
	nestingModeInvalid NestingMode = iota

	NestingSingle

	NestingGroup

	NestingList

	NestingSet

	NestingMap
)

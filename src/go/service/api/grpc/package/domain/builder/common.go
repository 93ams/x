package builder

import (
	"grpc/package/domain/model"
	"grpc/package/domain/model/meta"
)

func (b *EnumBuilder) Meta(m meta.Meta) *EnumBuilder {
	b.T.Meta = m
	return b
}
func (b *EnumBuilder) Comment(c *model.Comment) *EnumBuilder {
	b.T.Comment = c
	return b
}
func (b *EnumBuilder) Comments(c ...*model.Comment) *EnumBuilder {
	b.T.Comments = append(b.T.Comments, c...)
	return b
}
func (b *EnumFieldBuilder) Meta(m meta.Meta) *EnumFieldBuilder {
	b.T.Meta = m
	return b
}
func (b *EnumFieldBuilder) Comment(c *model.Comment) *EnumFieldBuilder {
	b.T.Comment = c
	return b
}
func (b *EnumFieldBuilder) Comments(c ...*model.Comment) *EnumFieldBuilder {
	b.T.Comments = append(b.T.Comments, c...)
	return b
}
func (b *ExtendBuilder) Meta(m meta.Meta) *ExtendBuilder {
	b.T.Meta = m
	return b
}
func (b *ExtendBuilder) Comment(c *model.Comment) *ExtendBuilder {
	b.T.Comment = c
	return b
}
func (b *ExtendBuilder) Comments(c ...*model.Comment) *ExtendBuilder {
	b.T.Comments = append(b.T.Comments, c...)
	return b
}
func (b *ExtensionsBuilder) Meta(m meta.Meta) *ExtensionsBuilder {
	b.T.Meta = m
	return b
}
func (b *ExtensionsBuilder) Comment(c *model.Comment) *ExtensionsBuilder {
	b.T.Comment = c
	return b
}
func (b *ExtensionsBuilder) Comments(c ...*model.Comment) *ExtensionsBuilder {
	b.T.Comments = append(b.T.Comments, c...)
	return b
}
func (b *FieldBuilder) Meta(m meta.Meta) *FieldBuilder {
	b.T.Meta = m
	return b
}
func (b *FieldBuilder) Comment(c *model.Comment) *FieldBuilder {
	b.T.Comment = c
	return b
}
func (b *FieldBuilder) Comments(c ...*model.Comment) *FieldBuilder {
	b.T.Comments = append(b.T.Comments, c...)
	return b
}
func (b *GroupFieldBuilder) Meta(m meta.Meta) *GroupFieldBuilder {
	b.T.Meta = m
	return b
}
func (b *GroupFieldBuilder) Comment(c *model.Comment) *GroupFieldBuilder {
	b.T.Comment = c
	return b
}
func (b *GroupFieldBuilder) Comments(c ...*model.Comment) *GroupFieldBuilder {
	b.T.Comments = append(b.T.Comments, c...)
	return b
}
func (b *ImportBuilder) Meta(m meta.Meta) *ImportBuilder {
	b.T.Meta = m
	return b
}
func (b *ImportBuilder) Comment(c *model.Comment) *ImportBuilder {
	b.T.Comment = c
	return b
}
func (b *ImportBuilder) Comments(c ...*model.Comment) *ImportBuilder {
	b.T.Comments = append(b.T.Comments, c...)
	return b
}
func (b *MapFieldBuilder) Meta(m meta.Meta) *MapFieldBuilder {
	b.T.Meta = m
	return b
}
func (b *MapFieldBuilder) Comment(c *model.Comment) *MapFieldBuilder {
	b.T.Comment = c
	return b
}
func (b *MapFieldBuilder) Comments(c ...*model.Comment) *MapFieldBuilder {
	b.T.Comments = append(b.T.Comments, c...)
	return b
}
func (b *MessageBuilder) Meta(m meta.Meta) *MessageBuilder {
	b.T.Meta = m
	return b
}
func (b *MessageBuilder) Comment(c *model.Comment) *MessageBuilder {
	b.T.Comment = c
	return b
}
func (b *MessageBuilder) Comments(c ...*model.Comment) *MessageBuilder {
	b.T.Comments = append(b.T.Comments, c...)
	return b
}
func (b *OneOfBuilder) Meta(m meta.Meta) *OneOfBuilder {
	b.T.Meta = m
	return b
}
func (b *OneOfBuilder) Comment(c *model.Comment) *OneOfBuilder {
	b.T.Comment = c
	return b
}
func (b *OneOfBuilder) Comments(c ...*model.Comment) *OneOfBuilder {
	b.T.Comments = append(b.T.Comments, c...)
	return b
}
func (b *OneOfFieldBuilder) Meta(m meta.Meta) *OneOfFieldBuilder {
	b.T.Meta = m
	return b
}
func (b *OneOfFieldBuilder) Comment(c *model.Comment) *OneOfFieldBuilder {
	b.T.Comment = c
	return b
}
func (b *OneOfFieldBuilder) Comments(c ...*model.Comment) *OneOfFieldBuilder {
	b.T.Comments = append(b.T.Comments, c...)
	return b
}
func (b *OptionBuilder) Meta(m meta.Meta) *OptionBuilder {
	b.T.Meta = m
	return b
}
func (b *OptionBuilder) Comment(c *model.Comment) *OptionBuilder {
	b.T.Comment = c
	return b
}
func (b *OptionBuilder) Comments(c ...*model.Comment) *OptionBuilder {
	b.T.Comments = append(b.T.Comments, c...)
	return b
}
func (b *PackageBuilder) Meta(m meta.Meta) *PackageBuilder {
	b.T.Meta = m
	return b
}
func (b *PackageBuilder) Comment(c *model.Comment) *PackageBuilder {
	b.T.Comment = c
	return b
}
func (b *PackageBuilder) Comments(c ...*model.Comment) *PackageBuilder {
	b.T.Comments = append(b.T.Comments, c...)
	return b
}
func (b *ReservedBuilder) Meta(m meta.Meta) *ReservedBuilder {
	b.T.Meta = m
	return b
}
func (b *ReservedBuilder) Comment(c *model.Comment) *ReservedBuilder {
	b.T.Comment = c
	return b
}
func (b *ReservedBuilder) Comments(c ...*model.Comment) *ReservedBuilder {
	b.T.Comments = append(b.T.Comments, c...)
	return b
}
func (b *RPCBuilder) Meta(m meta.Meta) *RPCBuilder {
	b.T.Meta = m
	return b
}
func (b *RPCBuilder) Comment(c *model.Comment) *RPCBuilder {
	b.T.Comment = c
	return b
}
func (b *RPCBuilder) Comments(c ...*model.Comment) *RPCBuilder {
	b.T.Comments = append(b.T.Comments, c...)
	return b
}
func (b *ServiceBuilder) Meta(m meta.Meta) *ServiceBuilder {
	b.T.Meta = m
	return b
}
func (b *ServiceBuilder) Comment(c *model.Comment) *ServiceBuilder {
	b.T.Comment = c
	return b
}
func (b *ServiceBuilder) Comments(c ...*model.Comment) *ServiceBuilder {
	b.T.Comments = append(b.T.Comments, c...)
	return b
}
func (b *SyntaxBuilder) Meta(m meta.Meta) *SyntaxBuilder {
	b.T.Meta = m
	return b
}
func (b *SyntaxBuilder) Comment(c *model.Comment) *SyntaxBuilder {
	b.T.Comment = c
	return b
}
func (b *SyntaxBuilder) Comments(c ...*model.Comment) *SyntaxBuilder {
	b.T.Comments = append(b.T.Comments, c...)
	return b
}

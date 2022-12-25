package model

type (
	Visitor interface {
		VisitComment(*Comment)
		VisitEmptyStatement(*EmptyStatement) bool
		VisitEnum(*Enum) bool
		VisitEnumField(*EnumField) bool
		VisitExtend(*Extend) bool
		VisitExtensions(*Extensions) bool
		VisitField(*Field) bool
		VisitGroupField(*GroupField) bool
		VisitImport(*Import) bool
		VisitMapField(*MapField) bool
		VisitMessage(*Message) bool
		VisitOneof(*OneOf) bool
		VisitOneofField(*OneOfField) bool
		VisitOption(*Option) bool
		VisitPackage(*Package) bool
		VisitReserved(*Reserved) bool
		VisitRPC(*RPC) bool
		VisitService(*Service) bool
		VisitSyntax(*Syntax) bool
	}
	Visitee interface{ Accept(v Visitor) }
)

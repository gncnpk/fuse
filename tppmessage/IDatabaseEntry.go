package tppmessage

type IDatabaseEntry interface {
	CreateSchema() error
	TableName() string
}

package jvm

import "go-on-jvm/jvm/constantpool"

type Statement struct {
}

func newStatement() Statement {
	return Statement{}
}

func (s Statement) fillConstantsPool(pool *constantpool.ConstantPool) {

}

type statementCompiler struct {
	Statement
	Pool *constantpool.ConstantPool
}

func newStatementCompiler(statement Statement, pool *constantpool.ConstantPool) *statementCompiler {
	return &statementCompiler{statement, pool}
}

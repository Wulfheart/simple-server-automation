package ssa

import "io"

type StandardTask struct {
	Output io.Writer
}

func (st *StandardTask) SetOutput(writer io.Writer) {
	st.Output = writer
}

type Task interface {
	SetOutput(writer io.Writer)
	Execute() error
}

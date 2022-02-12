package api

import "io"

type handler func(reader io.Reader, writer io.Writer)

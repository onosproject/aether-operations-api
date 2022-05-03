package graphql

import (
	fmt "fmt"
	graphql "github.com/99designs/gqlgen/graphql"
	io "io"
)

type FieldInput = Field
type RpcInput = Rpc
type SvcInput = Svc

func MarshalType(x Type) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = fmt.Fprintf(w, "%q", x.String())
	})
}

func UnmarshalType(v interface{}) (Type, error) {
	code, ok := v.(string)
	if ok {
		return Type(Type_value[code]), nil
	}
	return 0, fmt.Errorf("cannot unmarshal Type enum")
}

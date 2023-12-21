package day20

import "testing"

var config1 = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a
`

var config2 = `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output
`

func TestModuleCreation(t *testing.T) {
}

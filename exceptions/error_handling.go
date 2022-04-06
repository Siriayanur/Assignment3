package exceptions

import (
	"errors"
	"fmt"
)

var errMap = make(map[string]string)

var ErrInvalidNode = errors.New("NODE ERROR ")
var ErrInvalidDependency = errors.New("DEPENDENCY ERROR ")

func CreateErrorStatements() {
	errMap["idExists"] = "cannot create node, ID already exists "
	errMap["idNotExists"] = "node ID doesn't exist "
	errMap["cyclicDependency"] = "cannot create the cyclic dependency between parentID and childID "
}

func InvalidOperation(statement string, errKind error) error {
	return fmt.Errorf("%w :: %s", errKind, errMap[statement])
}

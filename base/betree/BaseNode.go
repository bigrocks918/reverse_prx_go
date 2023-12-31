package betree

import "fmt"

const (
	SUCCESS = iota
	ERROR   = iota
)

const (
	COMPOSITE = "composite"
	DECORATOR = "decorator"
	ACTION    = "action"
	CONDITION = "condition"
)

type (
	BaseNode struct {
		Type string
		Name string
	}

	IBaseNode interface {
		OnEnter(tick int64)
		OnExec(tick int64) bool
		OnExit(tick int64)
		Init()
		GetType() string
		GetName() string
		SetName(string)
	}
)

/**
 * Enter method, override this to use. It is called every time a node is
 * asked to execute, before the tick itself.
 *
 * @method enter
 * @param {Tick} tick A tick instance.
**/
func (b *BaseNode) OnEnter(tick int64) {

}

/**
 * Tick method, override this to use. This method must contain the real
 * execution of node (perform a task, call children, etc.). It is called
 * every time a node is asked to execute.
 *
 * @method tick
 * @param {Tick} tick A tick instance.
**/
func (b *BaseNode) OnExec(tick int64) bool {
	fmt.Println("tick BaseWorker")
	return false
}

/**
 * Exit method, override this to use. Called every time in the end of the
 * execution.
 *
 * @method exit
 * @param {Tick} tick A tick instance.
**/
func (b *BaseNode) OnExit(tick int64) {

}

func (b *BaseNode) Init() {

}

func (b *BaseNode) GetType() string {
	return b.Type
}

func (b *BaseNode) GetName() string {
	return b.Name
}

func (b *BaseNode) SetName(name string) {
	b.Name = name
}

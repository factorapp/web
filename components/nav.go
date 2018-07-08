package components

import (
	"github.com/factorapp/factor/markup"
	"github.com/gowasm/vecty"
)

type Nav struct {
	vecty.Core
	MyProp      string `vecty:"Prop"`
	CurrentPath string
}

func (n *Nav) OnAfterPrint(e *markup.Event) {

}

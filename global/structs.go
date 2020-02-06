package global

var (
	XConfig    map[string]string
	FrontVars  map[string]string //Fronrt End Server Variables
	ServerVars map[string]string //Fronrt End Server Variables
)

// 1. Receivers /////////////////////////////////////////////////////////////////
type ComVar struct {
	Key   string
	Value string
}

type Controller struct {
	// Ctx            *context.Context
	Ctx            Context
	controllerName string
	actionName     string
	GlobalVars     []ComVar
	Data           map[interface{}]interface{}
}

type Context struct {
	Ask    AbangoAsk
	Answer AbangoAnswer
}

type AbangoAsk struct {
	AuthToken  string
	AskName    string
	UniqueId   string
	Body       []byte
	ServerVars []ComVar
}

type AbangoAnswer struct {
	Status string
	Body   []byte
}

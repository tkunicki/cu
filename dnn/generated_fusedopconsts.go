package cudnn

/* Generated by gencudnn. DO NOT EDIT */

// #include <cudnn.h>
import "C"
import (
	"runtime"
	"unsafe"
)

// FusedOpConsts is a representation of cudnnFusedOpsConstParamPack_t.
type FusedOpConsts struct {
	internal C.cudnnFusedOpsConstParamPack_t

	paramLabel FusedOpsConstParamLabel
	param      Memory
}

// NewFusedOpConsts creates a new FusedOpConsts.
func NewFusedOpConsts(paramLabel FusedOpsConstParamLabel, param Memory) (retVal *FusedOpConsts, err error) {
	var internal C.cudnnFusedOpsConstParamPack_t
	if err := result(C.cudnnCreateFusedOpsConstParamPack(&internal)); err != nil {
		return nil, err
	}

	if err := result(C.cudnnSetFusedOpsConstParamPackAttribute(internal, paramLabel.C(), unsafe.Pointer(param.Uintptr()))); err != nil {
		return nil, err
	}

	retVal = &FusedOpConsts{
		internal:   internal,
		paramLabel: paramLabel,
		param:      param,
	}
	runtime.SetFinalizer(retVal, destroyFusedOpConsts)
	return retVal, nil
}

// C returns the internal cgo representation.
func (f *FusedOpConsts) C() C.cudnnFusedOpsConstParamPack_t { return f.internal }

// ParamLabel returns the internal paramLabel.
func (f *FusedOpConsts) ParamLabel() FusedOpsConstParamLabel { return f.paramLabel }

// Param returns the internal param.
func (f *FusedOpConsts) Param() Memory { return f.param }

func destroyFusedOpConsts(obj *FusedOpConsts) { C.cudnnDestroyFusedOpsConstParamPack(obj.internal) }

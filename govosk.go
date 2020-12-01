package govosk

// #cgo CFLAGS: -g -O1 -DFST_NO_DYNAMIC_LINKING -I${SRCDIR}/../kaldi/src
// #cgo CPPFLAGS: -g -O1 -DFST_NO_DYNAMIC_LINKING -I${SRCDIR}/../kaldi/src -I${SRCDIR}/../kaldi/tools/openfst/include
// #cgo CXXFLAGS: -std=c++11
// #cgo LDFLAGS: -L${SRCDIR}/../kaldi/src -lgfortran -lpthread
// #cgo LDFLAGS: ${SRCDIR}/../kaldi/src/online2/kaldi-online2.a
// #cgo LDFLAGS: ${SRCDIR}/../kaldi/src/decoder/kaldi-decoder.a
// #cgo LDFLAGS: ${SRCDIR}/../kaldi/src/ivector/kaldi-ivector.a
// #cgo LDFLAGS: ${SRCDIR}/../kaldi/src/gmm/kaldi-gmm.a
// #cgo LDFLAGS: ${SRCDIR}/../kaldi/src/nnet3/kaldi-nnet3.a
// #cgo LDFLAGS: ${SRCDIR}/../kaldi/src/tree/kaldi-tree.a
// #cgo LDFLAGS: ${SRCDIR}/../kaldi/src/feat/kaldi-feat.a
// #cgo LDFLAGS: ${SRCDIR}/../kaldi/src/lat/kaldi-lat.a
// #cgo LDFLAGS: ${SRCDIR}/../kaldi/src/lm/kaldi-lm.a
// #cgo LDFLAGS: ${SRCDIR}/../kaldi/src/hmm/kaldi-hmm.a
// #cgo LDFLAGS: ${SRCDIR}/../kaldi/src/transform/kaldi-transform.a
// #cgo LDFLAGS: ${SRCDIR}/../kaldi/src/cudamatrix/kaldi-cudamatrix.a
// #cgo LDFLAGS: ${SRCDIR}/../kaldi/src/matrix/kaldi-matrix.a
// #cgo LDFLAGS: ${SRCDIR}/../kaldi/src/fstext/kaldi-fstext.a
// #cgo LDFLAGS: ${SRCDIR}/../kaldi/src/util/kaldi-util.a
// #cgo LDFLAGS: ${SRCDIR}/../kaldi/src/base/kaldi-base.a
// #cgo LDFLAGS: ${SRCDIR}/../kaldi/tools/OpenBLAS/libopenblas.a
// #cgo LDFLAGS: ${SRCDIR}/../kaldi/tools/openfst/lib/libfst.a
// #cgo LDFLAGS: ${SRCDIR}/../kaldi/tools/openfst/lib/libfstngram.a
// #cgo LDFLAGS: -lgfortran -lpthread
// #include <stdlib.h>
// #include "vosk_api.h"
import "C"
import (
	"fmt"
)

func VoskFinalResult(recognizer *VoskRecognizer, buffer []byte) string {
	cbuf := C.CBytes(buffer)
	defer C.free(cbuf)
	final := C.vosk_recognizer_accept_waveform(recognizer.rec, (*C.char)(cbuf), C.int(len(buffer)))
	if final != C.int(0) {
		fmt.Println("final")
	} else {
		fmt.Println("not final")
	}
	result := C.GoString(C.vosk_recognizer_final_result(recognizer.rec))
	fmt.Println(result)
	return result
}

// NewModel creates a new VoskModel instance
func NewModel(modelPath string) (*VoskModel, error) {
	var internal *C.struct_VoskModel
	internal = C.vosk_model_new(C.CString(modelPath))
	model := &VoskModel{model: internal}
	return model, nil
}

// NewRecognizer creates a new VoskRecognizer instance
func NewRecognizer(model *VoskModel) (*VoskRecognizer, error) {
	var internal *C.struct_VoskRecognizer
	internal = C.vosk_recognizer_new(model.model, 16000.0)
	rec := &VoskRecognizer{rec: internal}
	return rec, nil
}

func freeModel(model *VoskModel) {
	C.vosk_model_free(model.model)
}

func freeRecognizer(recognizer *VoskRecognizer) {
	C.vosk_recognizer_free(recognizer.rec)
}

// NewSpkModel creates a new VoskSpkModel instance
func NewSpkModel() {
	// TODO
}

// VoskModel contains a reference to the C VoskModel
type VoskModel struct {
	model *C.struct_VoskModel
}

// VoskSpkModel contains a reference to the C VoskSpkModel
type VoskSpkModel struct {
	spkModel *C.struct_VoskSpkModel
}

// VoskRecognizer contains a reference to the  C VoskRecognizer
type VoskRecognizer struct {
	rec *C.struct_VoskRecognizer
}
